package middleware

import (
	"context"

	"github.com/EbumbaE/FinTgBot/bot/internal/model/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Middleware struct {
	Metrics     *Metrics
	wrappedFunc MiddlewareFunc
}

type MessageModel interface {
	IncomingCommand(context.Context, messages.Message) error
	IncomingMessage(context.Context, messages.Message) error
}

type MiddlewareFunc func(ctx context.Context, msgModel MessageModel, tgMsg *tgbotapi.Message)

func NewMiddleware(Metrics *Metrics) (m *Middleware) {
	m = &Middleware{
		Metrics: Metrics,
	}
	m.wrappedFunc = DetermineRequest()
	m.wrappedFunc = m.LoggingMiddleware(m.wrappedFunc)
	m.wrappedFunc = m.MetricsMiddleware(m.wrappedFunc)
	m.wrappedFunc = m.TracingMiddleware(m.wrappedFunc)

	return
}
