package sender

import (
	"context"
	"fmt"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SenderClient struct {
	target string
	client api.SenderClient
	conn   *grpc.ClientConn
}

type Message struct {
	UserID int64
	Text   string
}

func New(cfg Config) *SenderClient {
	return &SenderClient{
		target: cfg.Target,
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
	request := &api.SendMessageRequest{
		UserID: msg.UserID,
		Text:   msg.Text,
	}
	response, err := s.client.SendMessage(ctx, request)
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
