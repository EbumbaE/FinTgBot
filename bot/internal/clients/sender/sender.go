package sender

import (
	"context"
	"fmt"
	"log"
	"net"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/api"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type TgSender interface {
	SendMessage(msg messages.Message) error
}

type SenderServer struct {
	target string
	sender TgSender
	api.UnimplementedSenderServer
}

func New(cfg Config, sender TgSender) *SenderServer {
	return &SenderServer{
		sender: sender,
		target: cfg.Target,
	}
}

func (s *SenderServer) Init() (err error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.target))
	if err != nil {
		logger.Error("failed to listen in sender server", zap.Error(err))
	}
	server := grpc.NewServer()
	api.RegisterSenderServer(server, &SenderServer{})

	logger.Info("sender server is begin")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return
}

func (s *SenderServer) SendMessage(ctx context.Context, r *api.SendMessageRequest) (*api.SendMessageResponse, error) {
	msg := messages.Message{}
	if err := s.sender.SendMessage(msg); err != nil {
		return &api.SendMessageResponse{}, err
	}
	return &api.SendMessageResponse{}, nil
}
