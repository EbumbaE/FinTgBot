package middleware

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/logger"
	"go.uber.org/zap"
)

func init() {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
	}
	if _, err := cfg.InitGlobalTracer("finbot"); err != nil {
		logger.Fatal("cannot init tracing", zap.Error(err))
	}
}

func (m *Middleware) TracingMiddleware(next MiddlewareFunc) MiddlewareFunc {
	return func(ctx context.Context, msgModel MessageModel, tgMsg *tgbotapi.Message) {
		span, nctx := opentracing.StartSpanFromContext(ctx, "incoming request")
		if span != nil {
			span.LogKV("incoming request", "got message", "message", tgMsg.Text)
			defer span.Finish()
		}

		next(nctx, msgModel, tgMsg)
	}
}
