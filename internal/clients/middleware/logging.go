package middleware

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/logger"
	"go.uber.org/zap"
)

func (m *Middleware) LoggingMiddleware(next MiddlewareFunc) MiddlewareFunc {
	return func(ctx context.Context, msgModel MessageModel, tgMsg *tgbotapi.Message) {
		logger.Info("incoming request: ", zap.Int64("userid", tgMsg.From.ID), zap.String("text", tgMsg.Text))
		next(ctx, msgModel, tgMsg)
	}
}
