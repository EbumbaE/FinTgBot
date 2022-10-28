package middleware

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/logger"
	"go.uber.org/zap"
)

func DetermineRequest() MiddlewareFunc {
	return func(ctx context.Context, msgModel MessageModel, tgMsg *tgbotapi.Message) {
		if tgMsg.IsCommand() {
			err := msgModel.IncomingCommand(messages.Message{
				UserID:    tgMsg.From.ID,
				Command:   tgMsg.Command(),
				Arguments: tgMsg.CommandArguments(),
			})
			if err != nil {
				logger.Error("incoming command: ", zap.Error(err))
			}
		} else {
			err := msgModel.IncomingMessage(messages.Message{
				UserID: tgMsg.From.ID,
				Text:   tgMsg.Text,
			})
			if err != nil {
				logger.Error("incoming message: ", zap.Error(err))
			}
		}
	}
}

func (m *Middleware) IncomingRequest(ctx context.Context, msgModel MessageModel, tgMsg *tgbotapi.Message) {
	m.wrappedFunc(ctx, msgModel, tgMsg)
}
