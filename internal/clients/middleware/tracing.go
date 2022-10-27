package middleware

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/opentracing/opentracing-go"
)

func TracingMiddleware(handleFunc MiddlewareFunc) MiddlewareFunc {
	return func(ctx context.Context, msgModel MessageModel, tgMsg *tgbotapi.Message) {
		span := opentracing.SpanFromContext(ctx)
		if span != nil {
			span.LogKV("incoming request", "got message", "message", tgMsg.Text)
		}
		defer span.Finish()

		wrappedFunc(ctx, msgModel, tgMsg)
	}
}
