package sender

import (
	"context"
	"fmt"
	"net"
	"sync"

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
	port   string
	sender TgSender
	api.UnimplementedSenderServer
}

func New(cfg Config, sender TgSender) *SenderServer {
	return &SenderServer{
		sender: sender,
		port:   cfg.Port,
	}
}

func (s *SenderServer) StartServe(ctx context.Context) (err error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		logger.Error("failed to listen in sender server", zap.Error(err))
	}
	server := grpc.NewServer()
	api.RegisterSenderServer(server, &SenderServer{})

	go func() {
		logger.Info("sender server is begin")
		defer ctx.Value("allDoneWG").(*sync.WaitGroup).Done()

		go func() {
			if err := server.Serve(listener); err != nil {
				logger.Error("failed to serve in sender server", zap.Error(err))
			}
		}()

		<-ctx.Done()
		server.GracefulStop()

		logger.Info("sender server is end")
	}()
	return
}

func (s *SenderServer) SendMessage(ctx context.Context, r *api.SendMessageRequest) (*api.SendMessageResponse, error) {
	msg := messages.Message{}
	if err := s.sender.SendMessage(msg); err != nil {
		return &api.SendMessageResponse{Status: api.Status_FAIL}, err
	}
	return &api.SendMessageResponse{Status: api.Status_SUCCESS}, nil
}
