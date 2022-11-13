package consumer

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/clients/sender"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/report"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/storage"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/pkg/logger"
	"go.uber.org/zap"
)

type MessageSender interface {
	SendMessage(ctx context.Context, msg sender.Message) (err error)
}

type ReportCache interface {
	AddReportInCache(userID int64, period string, addedReport report.ReportFormat) (err error)
	GetReportFromCache(userID int64, period string) (getReport report.ReportFormat, err error)
	AddNoteInCacheReports(userID int64, date time.Time, note diary.Note) error
}

type Consumer struct {
	consumerGroup      sarama.ConsumerGroup
	kafkaTopic         string
	kafkaConsumerGroup string
	brokersList        []string
	storage            storage.Storage
	cache              ReportCache
	sender             MessageSender
}

func New(cfg Config, storage storage.Storage, cache ReportCache, sender MessageSender) *Consumer {
	return &Consumer{
		kafkaTopic:         cfg.KafkaTopic,
		kafkaConsumerGroup: cfg.KafkaConsumerGroup,
		brokersList:        cfg.BrokersList,
		storage:            storage,
		cache:              cache,
		sender:             sender,
	}
}

func (c *Consumer) InitConsumerGroup(ctx context.Context) (err error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRange}

	c.consumerGroup, err = sarama.NewConsumerGroup(c.brokersList, c.kafkaConsumerGroup, config)
	if err != nil {
		return fmt.Errorf("new consumer group: %w", err)
	}

	return nil
}

func (c *Consumer) StartConsumerGroup(ctx context.Context) {
	go func() {
		defer ctx.Value("allDoneWG").(*sync.WaitGroup).Done()

		go func() {
			ch := &ConsumeHandler{
				storage: c.storage,
				cache:   c.cache,
				sender:  c.sender,
				ctx:     ctx,
			}
			logger.Info("consumer is begin")

			if err := c.consumerGroup.Consume(ctx, []string{c.kafkaTopic}, ch); err != nil {
				logger.Error("consuming via handler", zap.Error(err))
			}
		}()

		<-ctx.Done()
		if err := c.consumerGroup.Close(); err != nil {
			logger.Error("consume group close", zap.Error(err))
		}

		logger.Info("consumer is end")
	}()

	return
}
