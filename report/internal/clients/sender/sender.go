package sender

import (
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Sender struct {
	target string
	client api.SenderClient
	conn   *grpc.ClientConn
}

func New(cfg Config) *Sender {
	return &Sender{
		target: cfg.Target,
	}
}

func (s *Sender) Init() (err error) {
	s.conn, err = grpc.Dial(s.target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	s.client = api.NewSenderClient(s.conn)

	return
}

func (s *Sender) Close() (err error) {
	if err = s.conn.Close(); err != nil {
		return
	}

	return
}
