package consumer

import (
	"encoding/json"
	"time"

	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/report"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/request"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/pkg/logger"
	"go.uber.org/zap"
)

type Valute interface {
	GetAbbreviation() string
	GetValue() float64
}

type ReportCache interface {
	AddReportInCache(userID int64, period string, addedReport report.ReportFormat) (err error)
	GetReportFromCache(userID int64, period string) (getReport report.ReportFormat, err error)
	AddNoteInCacheReports(userID int64, date time.Time, note diary.Note) error
}

type ConsumeHandler struct {
}

func (c *ConsumeHandler) Setup(sarama.ConsumerGroupSession) error {
	logger.Info("consumer is setup")
	return nil
}

func (c *ConsumeHandler) Cleanup(sarama.ConsumerGroupSession) error {
	logger.Info("consumer is cleanup")
	return nil
}

func (c *ConsumeHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		logger.Info("consumer is claim msg",
			zap.String("Topic", msg.Topic),
			zap.Int64("Offset", msg.Offset),
			zap.Int32("Partition", msg.Partition),
			zap.ByteString("Key", msg.Key),
			zap.ByteString("Value", msg.Value),
		)

		var reportRequest request.ReportRequest
		if err := json.Unmarshal(msg.Value, reportRequest); err != nil {
			logger.Error("unmarshal report request", zap.Error(err))
		}

		// cacheReport, err := t.cache.GetReportFromCache(msg.UserID, period)
		// if err != nil {
		// 	logger.Info("get cache report", zap.Error(err))
		// } else {
		// 	strReport, err := report.FormatReportToString(&cacheReport, period, userRateCurrency)
		// 	return strReport, err
		// }

		// countReport, err := report.CountStatistic(msg.UserID, period, t.storage, &t.dateFormatter)
		// if err != nil {
		// 	answer = "not done"
		// 	logger.Error("count statistic", zap.Error(err))
		// 	return
		// }
		// answer, err = report.FormatReportToString(&countReport, period, userRateCurrency)
		// if err != nil {
		// 	answer = "not done"
		// 	logger.Error("format statistic to string", zap.Error(err))
		// 	return
		// }

		// if err := t.cache.AddReportInCache(msg.UserID, period, countReport); err != nil {
		// 	logger.Error("add report in cache", zap.Error(err))
		// }

		session.MarkMessage(msg, "")
	}

	return nil
}
