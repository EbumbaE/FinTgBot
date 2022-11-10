package consumer

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/pkg/logger"
	"go.uber.org/zap"
)

var (
	KafkaTopic         = "example-topic"
	KafkaConsumerGroup = "example-consumer-group"
	BrokersList        = []string{"localhost:9092"}
	Assignor           = "range"
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
		return fmt.Errorf("starting consumer group: %w", err)
	}

	return nil
}

func (c *Consumer) StartConsumerGroup(ctx context.Context) {
	ch := &ConsumeHandler{}
	//TODO обернуть в спан и логгер

	go func() {
		err := c.consumerGroup.Consume(ctx, []string{c.kafkaTopic}, ch)
		if err != nil {
			logger.Error("consuming via handler", zap.Error(err))
		}
	}()

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
	for message := range claim.Messages() {
		logger.Info("consumer is claim msg",
			zap.String("Topic", message.Topic),
			zap.Int64("Offset", message.Offset),
			zap.Int32("Partition", message.Partition),
			zap.ByteString("Key", message.Key),
			zap.ByteString("Value", message.Value),
		)

		session.MarkMessage(message, "")
	}

	return nil
}
