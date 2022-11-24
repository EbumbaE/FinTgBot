package consumer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/EbumbaE/FinTgBot/report/internal/clients/sender"
	"github.com/EbumbaE/FinTgBot/report/internal/model/report"
	"github.com/EbumbaE/FinTgBot/report/internal/model/request"
	"github.com/EbumbaE/FinTgBot/report/internal/servers/middleware"
	"github.com/EbumbaE/FinTgBot/report/internal/storage"
	"github.com/EbumbaE/FinTgBot/report/pkg/logger"
	"github.com/Shopify/sarama"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type Valute interface {
	GetAbbreviation() string
	GetValue() float64
}

type ConsumeHandler struct {
	metrics *middleware.Metrics
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
	startTime := time.Now()
	defer func(startTime time.Time) {
		duration := time.Since(startTime)
		ch.metrics.SummaryCountStatisticTime.Observe(duration.Seconds())
		ch.metrics.HistogramCountStatisticTime.Observe(duration.Seconds())
	}(startTime)

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

func (ch *ConsumeHandler) incomingAddNoteInCacheRequest(ctx context.Context, msgValue []byte) error {
	startTime := time.Now()
	defer func(startTime time.Time) {
		duration := time.Since(startTime)
		ch.metrics.SummaryAddNoteInCacheTime.Observe(duration.Seconds())
		ch.metrics.HistogramAddNoteInCacheTime.Observe(duration.Seconds())
	}(startTime)

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
		ch.metrics.AmountCommands.Inc()

		switch string(msg.Key) {
		case "report":
			span, nctx := opentracing.StartSpanFromContext(ctx, "incoming request")
			if span != nil {
				span.LogKV("incoming report request", "got message", "value", msg.Value)
				defer span.Finish()
			}
			err = ch.incomingReportRequest(nctx, msg.Value)
		case "add_note_in_cache":
			span, nctx := opentracing.StartSpanFromContext(ctx, "incoming request")
			if span != nil {
				span.LogKV("incoming add note in cache request", "got message", "value", msg.Value)
				defer span.Finish()
			}
			err = ch.incomingAddNoteInCacheRequest(nctx, msg.Value)
		}

		if err != nil {
			logger.Error("incoming request in consume claim", zap.Error(err))
			return
		}

		session.MarkMessage(msg, "")
	}

	return nil
}
