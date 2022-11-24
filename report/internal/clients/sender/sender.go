package sender

import (
	"context"
	"fmt"
	"time"

	"github.com/EbumbaE/FinTgBot/libs/api"
	"github.com/EbumbaE/FinTgBot/report/internal/clients/middleware"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SenderClient struct {
	target  string
	client  api.SenderClient
	conn    *grpc.ClientConn
	metrics *middleware.Metrics
}

type Message struct {
	UserID int64
	Text   string
}

func New(cfg Config) *SenderClient {
	return &SenderClient{
		target:  cfg.Target,
		metrics: middleware.NewMetrics(),
	}
}

func (s *SenderClient) StartConnection() (err error) {
	s.conn, err = grpc.Dial(s.target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	s.client = api.NewSenderClient(s.conn)

	return
}

func (s *SenderClient) SendMessage(ctx context.Context, msg Message) (err error) {
	startTime := time.Now()
	defer func(startTime time.Time) {
		duration := time.Since(startTime)
		s.metrics.SummarySendMessage.Observe(duration.Seconds())
		s.metrics.HistogramSendMessage.Observe(duration.Seconds())
	}(startTime)

	span, nctx := opentracing.StartSpanFromContext(ctx, "send message")
	if span != nil {
		span.LogKV("send message request", "send message", "user_id", msg.UserID, "text", msg.Text)
		defer span.Finish()
	}

	request := &api.SendMessageRequest{
		UserID: msg.UserID,
		Text:   msg.Text,
	}
	response, err := s.client.SendMessage(nctx, request)
	if err != nil {
		return err
	}
	if response.Status != api.Status_SUCCESS {
		return fmt.Errorf("sender client: response status in send message: %w", err)
	}

	return nil
}

func (s *SenderClient) Close() error {
	if err := s.conn.Close(); err != nil {
		return err
	}

	return nil
}
