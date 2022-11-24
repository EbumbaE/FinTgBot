package middleware

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/EbumbaE/FinTgBot/bot/internal/model/messages"
	"github.com/EbumbaE/FinTgBot/bot/pkg/logger"
	"go.uber.org/zap"
)

func DetermineRequest() MiddlewareFunc {
	return func(ctx context.Context, msgModel MessageModel, tgMsg *tgbotapi.Message) {
		if tgMsg.IsCommand() {
			err := msgModel.IncomingCommand(ctx, messages.Message{
				UserID:    tgMsg.From.ID,
				Command:   tgMsg.Command(),
				Arguments: tgMsg.CommandArguments(),
			})
			if err != nil {
				logger.Error("incoming command: ", zap.Error(err))
			}
		} else {
			err := msgModel.IncomingMessage(ctx, messages.Message{
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
