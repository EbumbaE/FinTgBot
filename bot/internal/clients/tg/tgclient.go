package tg

import (
	"context"
	"sync"

	"github.com/EbumbaE/FinTgBot/bot/internal/clients/middleware"
	"github.com/EbumbaE/FinTgBot/bot/internal/model/messages"
	"github.com/EbumbaE/FinTgBot/bot/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Parser interface {
	GetAbbreviations() []string
}

type Client struct {
	client     *tgbotapi.BotAPI
	Keyboards  *Keyboards
	Metrics    *middleware.Metrics
	Middleware *middleware.Middleware
}

func (c *Client) InitMiddleware() {
	metrics := middleware.NewMetrics()
	c.Metrics = metrics
	c.Middleware = middleware.NewMiddleware(metrics)
}

func New(tgClient Config, parser Parser) (*Client, error) {
	currencies := parser.GetAbbreviations()

	client, err := tgbotapi.NewBotAPI(tgClient.Token)
	if err != nil {
		return nil, errors.Wrap(err, "NewBotAPI")
	}

	return &Client{
		client:    client,
		Keyboards: NewKeyboards(currencies),
	}, nil
}

func (c *Client) SendMessage(msg messages.Message) error {
	if msg.Text == "" {
		logger.Info("try to send empty message")
		return nil
	}

	logger.Info("tgclient: send message", zap.Int64("userid", msg.UserID), zap.String("text", msg.Text))

	tgMsg := tgbotapi.NewMessage(msg.UserID, msg.Text)
	tgMsg.ReplyMarkup = msg.Keyboard

	_, err := c.client.Send(tgMsg)
	if err != nil {
		return errors.Wrap(err, "client.Send")
	}
	return nil
}

func (c *Client) SetupCurrencyKeyboard(msg *messages.Message) {
	msg.Keyboard = c.Keyboards.GetCurrencyKeyboard()
}

func (c *Client) ListenUpdates(ctx context.Context, msgModel *messages.Model) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.client.GetUpdatesChan(u)

	logger.Info("listening messages begin")

	go func() {
		for {
			select {
			case update := <-updates:
				if update.Message != nil {
					c.Middleware.IncomingRequest(ctx, msgModel, update.Message)
				}
			case <-ctx.Done():
				defer ctx.Value("allDoneWG").(*sync.WaitGroup).Done()
				logger.Info("Listening messages end")
				return
			}
		}
	}()
}
