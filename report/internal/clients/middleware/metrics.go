package middleware

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	AmountSendMessages   prometheus.Counter
	SummarySendMessage   prometheus.Summary
	HistogramSendMessage prometheus.Histogram
}

func NewMetrics() *Metrics {
	amountSendMessages := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "ozon",
		Name:      "report_send_messages",
	})

	summarySendMessage := promauto.NewSummary(prometheus.SummaryOpts{
		Namespace: "ozon",
		Name:      "report_summary_send_messages_time_seconds",
		Objectives: map[float64]float64{
			0.5:  0.1,
			0.9:  0.01,
			0.99: 0.001,
		},
	})

	histogramSendMessage := promauto.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "ozon",
			Name:      "report_histogram_send_messages_time_seconds",
			Buckets:   []float64{0.0001, 0.0005, 0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 2},
		},
	)

	return &Metrics{
		AmountSendMessages:   amountSendMessages,
		SummarySendMessage:   summarySendMessage,
		HistogramSendMessage: histogramSendMessage,
	}
}
