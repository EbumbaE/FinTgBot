package producer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/request"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/pkg/logger"
	"go.uber.org/zap"
)

type Producer struct {
	asyncProducer sarama.AsyncProducer
	kafkaTopic    string
	brokersList   []string
}

func New(cfg Config) *Producer {
	return &Producer{
		kafkaTopic:  cfg.KafkaTopic,
		brokersList: cfg.BrokersList,
	}
}

func (p *Producer) InitProducer(ctx context.Context) (err error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0
	config.Producer.Return.Successes = true

	p.asyncProducer, err = sarama.NewAsyncProducer(p.brokersList, config)
	if err != nil {
		return fmt.Errorf("starting Sarama producer: %w", err)
	}

	return nil
}

func (p *Producer) StartConsumeError() {
	go func() {
		// TODO graceful
		for err := range p.asyncProducer.Errors() {
			logger.Error("failed to write message", zap.Error(err))
		}
	}()
}

func (p *Producer) SendReportRequest(ctx context.Context, r request.ReportRequest) (err error) {

	value, err := json.Marshal(r)
	if err != nil {
		return
	}

	msg := sarama.ProducerMessage{
		Topic:   p.kafkaTopic,
		Key:     sarama.StringEncoder("report"),
		Value:   sarama.StringEncoder(value),
		Headers: []sarama.RecordHeader{{Key: []byte("report"), Value: []byte(value)}},
	}

	p.asyncProducer.Input() <- &msg
	successMsg := <-p.asyncProducer.Successes()
	logger.Info("successful to write message", zap.Int64("offset", successMsg.Offset))

	return
}

func (p *Producer) SendAddNoteInCache(ctx context.Context, r request.AddNoteInCacheRequest) (err error) {

	value, err := json.Marshal(r)
	if err != nil {
		return
	}

	msg := sarama.ProducerMessage{
		Topic:   p.kafkaTopic,
		Key:     sarama.StringEncoder("add_note_in_cache"),
		Value:   sarama.StringEncoder(value),
		Headers: []sarama.RecordHeader{{Key: []byte("add_note_in_cache"), Value: []byte(value)}},
	}

	p.asyncProducer.Input() <- &msg
	successMsg := <-p.asyncProducer.Successes()
	logger.Info("successful to write message", zap.Int64("offset", successMsg.Offset))

	return
}

func (p *Producer) Close() error {
	return p.asyncProducer.Close()
}
