package middleware

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
)

type Middleware struct {
	Metrics     *Metrics
	wrappedFunc MiddlewareFunc
}

type MessageModel interface {
	IncomingCommand(messages.Message) error
	IncomingMessage(messages.Message) error
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
