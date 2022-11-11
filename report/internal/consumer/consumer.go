package consumer

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/pkg/logger"
	"go.uber.org/zap"
)

type Consumer struct {
	consumerGroup      sarama.ConsumerGroup
	kafkaTopic         string
	kafkaConsumerGroup string
	brokersList        []string
}

func New(cfg Config) *Consumer {
	return &Consumer{
		kafkaTopic:         cfg.KafkaTopic,
		kafkaConsumerGroup: cfg.KafkaConsumerGroup,
		brokersList:        cfg.BrokersList,
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
	ch := &ConsumeHandler{}
	//TODO обернуть в спан и логгер

	go func() {
		// TODO graceful
		logger.Info("consumer begin")
		err := c.consumerGroup.Consume(ctx, []string{c.kafkaTopic}, ch)
		if err != nil {
			logger.Error("consuming via handler", zap.Error(err))
		}
	}()

}
