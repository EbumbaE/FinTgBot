package metrics

import (
	"context"
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/logger"
	"go.uber.org/zap"
)

func InitServer(ctx context.Context) {

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	server := &http.Server{Addr: ":8080", Handler: mux}
	go func() {
		defer ctx.Value("allDoneWG").(*sync.WaitGroup).Done()
		logger.Info("metrics server begin")

		go func() {
			if err := server.ListenAndServe(); err != nil {
				logger.Error("metrics server listen and serve: ", zap.Error(err))
			}
		}()

		<-ctx.Done()
		if err := server.Shutdown(ctx); err != nil {
			logger.Error("metrics server shutdown: ", zap.Error(err))
		} else {
			logger.Info("metrics server end")
		}
	}()
}
