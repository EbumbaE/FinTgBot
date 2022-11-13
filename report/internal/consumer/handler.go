package consumer

import (
	"context"
	"encoding/json"

	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/clients/sender"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/report"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/request"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/storage"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/pkg/logger"
	"go.uber.org/zap"
)

type Valute interface {
	GetAbbreviation() string
	GetValue() float64
}

type ConsumeHandler struct {
	storage storage.Storage
	cache   ReportCache
	sender  MessageSender
	ctx     context.Context
}

func (c *ConsumeHandler) Setup(sarama.ConsumerGroupSession) error {
	logger.Info("consumer is setup")
	return nil
}

func (c *ConsumeHandler) Cleanup(sarama.ConsumerGroupSession) error {
	logger.Info("consumer is cleanup")
	return nil
}

func (ch *ConsumeHandler) countReport(reportRequest *request.ReportRequest) (*report.ReportFormat, error) {

	cacheReport, err := ch.cache.GetReportFromCache(reportRequest.UserID, reportRequest.Period)
	if err == nil {
		return &cacheReport, nil
	}

	countReport, err := report.CountStatistic(reportRequest.UserID, reportRequest.Period, ch.storage, reportRequest.DateFormat)
	if err != nil {
		logger.Error("count statistic", zap.Error(err))
		return nil, err
	}

	if err := ch.cache.AddReportInCache(reportRequest.UserID, reportRequest.Period, countReport); err != nil {
		logger.Error("add report in cache", zap.Error(err))
	}

	return &countReport, nil
}

func (ch *ConsumeHandler) sendReport(ctx context.Context, r *report.ReportFormat, request *request.ReportRequest) error {
	strReport, err := report.FormatReportToString(r, request.Period, request.UserCurrency)
	if err != nil {
		return err
	}

	msg := sender.Message{
		UserID: request.UserID,
		Text:   strReport,
	}

	if err := ch.sender.SendMessage(ctx, msg); err != nil {
		return err
	}

	return nil
}

func (ch *ConsumeHandler) incomingReportRequest(ctx context.Context, msgValue []byte) error {
	var reportRequest request.ReportRequest
	if err := json.Unmarshal(msgValue, &reportRequest); err != nil {
		logger.Error("unmarshal report request", zap.Error(err))
		return err
	}

	report, err := ch.countReport(&reportRequest)
	if err != nil {
		return err
	}
	if err := ch.sendReport(ctx, report, &reportRequest); err != nil {
		return err
	}

	return nil
}

func (ch *ConsumeHandler) addNoteInCache(addNoteRequest request.AddNoteInCacheRequest) error {
	return ch.cache.AddNoteInCacheReports(addNoteRequest.UserID, addNoteRequest.TimeNote, addNoteRequest.Note)
}

func (ch *ConsumeHandler) incomingAddNoteInCacheRequest(msgValue []byte) error {
	var addNoteRequest request.AddNoteInCacheRequest
	if err := json.Unmarshal(msgValue, &addNoteRequest); err != nil {
		logger.Error("unmarshal request in addNoteInCache", zap.Error(err))
		return err
	}

	return ch.addNoteInCache(addNoteRequest)
}

func (ch *ConsumeHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) (err error) {
	ctx := ch.ctx

	for msg := range claim.Messages() {
		logger.Info("consumer is claim msg",
			zap.String("Topic", msg.Topic),
			zap.Int64("Offset", msg.Offset),
			zap.Int32("Partition", msg.Partition),
			zap.ByteString("Key", msg.Key),
			zap.ByteString("Value", msg.Value),
		)

		switch string(msg.Key) {
		case "report":
			err = ch.incomingReportRequest(ctx, msg.Value)
		case "add_note_in_cache":
			err = ch.incomingAddNoteInCacheRequest(msg.Value)
		}

		if err != nil {
			logger.Error("incoming request in consume claim", zap.Error(err))
			return
		}

		session.MarkMessage(msg, "")
	}

	return nil
}
