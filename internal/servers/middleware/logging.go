package middleware

import (
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/logger"
)

func LoggingMiddleware(msg messages.Message) {
	logger.Info("bibi")
}
