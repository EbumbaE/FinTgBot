package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct {
	AmountRequests        prometheus.Counter
	SummaryResponseTime   prometheus.Summary
	HistogramResponseTime *prometheus.HistogramVec
}

func init() {
	http.Handle("/metrics", promhttp.Handler())
	log.Println(http.ListenAndServe(":2112", nil))
}

func NewMetrics() *Metrics {
	amountRequests := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "ozon",
		Name:      "in_amount_requests_total",
	})
	summaryResponseTime := promauto.NewSummary(prometheus.SummaryOpts{
		Namespace: "ozon",
		Name:      "summary_response_time_seconds",
		Objectives: map[float64]float64{
			0.5:  0.1,
			0.9:  0.01,
			0.99: 0.001,
		},
	})
	histogramResponseTime := promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "ozon",
			Name:      "histogram_response_time_seconds",
			Buckets:   []float64{0.0001, 0.0005, 0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 2},
		},
		[]string{"code"},
	)

	return &Metrics{
		AmountRequests:        amountRequests,
		SummaryResponseTime:   summaryResponseTime,
		HistogramResponseTime: histogramResponseTime,
	}

}

func (m *Middleware) MetricsMiddleware() MiddlewareFunc {
	return func(ctx context.Context, msgModel MessageModel, tgMsg *tgbotapi.Message) {
		startTime := time.Now()
		duration := time.Since(startTime)
		m.Metrics.SummaryResponseTime.Observe(duration.Seconds())
		m.Metrics.AmountRequests.Inc()

		m.wrappedFunc(ctx, msgModel, tgMsg)
	}
}
