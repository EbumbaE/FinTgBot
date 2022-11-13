package producer

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	AmountProducerErrors               prometheus.Counter
	AmountReportRequest                prometheus.Counter
	SummaryReportRequestTime           prometheus.Summary
	HistogramReportRequestTime         prometheus.Histogram
	AmountAddNoteInCacheRequest        prometheus.Counter
	SummaryAddNoteInCacheRequestTime   prometheus.Summary
	HistogramAddNoteInCacheRequestTime prometheus.Histogram
}

func NewMetrics() *Metrics {
	amountProducerErrors := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "ozon",
		Name:      "amount_producer_errors_total",
	})
	amountReportRequest := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "ozon",
		Name:      "amount_report_request_total",
	})
	summaryReportRequestTime := promauto.NewSummary(prometheus.SummaryOpts{
		Namespace: "ozon",
		Name:      "summary_report_request_time_seconds",
		Objectives: map[float64]float64{
			0.5:  0.1,
			0.9:  0.01,
			0.99: 0.001,
		},
	})
	histogramReportRequestTime := promauto.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "ozon",
			Name:      "histogram_report_request_time_seconds",
			Buckets:   []float64{0.0001, 0.0005, 0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 2},
		},
	)
	amountAddNoteInCacheRequest := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "ozon",
		Name:      "amount_addNoteInCache_request_total",
	})
	summaryAddNoteInCacheRequestTime := promauto.NewSummary(prometheus.SummaryOpts{
		Namespace: "ozon",
		Name:      "summary_addNoteInCache_request_time_seconds",
		Objectives: map[float64]float64{
			0.5:  0.1,
			0.9:  0.01,
			0.99: 0.001,
		},
	})
	histogramAddNoteInCacheRequestTime := promauto.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "ozon",
			Name:      "histogram_addNoteInCache_request_time_seconds",
			Buckets:   []float64{0.0001, 0.0005, 0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 2},
		},
	)
	return &Metrics{
		AmountProducerErrors:               amountProducerErrors,
		AmountReportRequest:                amountReportRequest,
		SummaryReportRequestTime:           summaryReportRequestTime,
		HistogramReportRequestTime:         histogramReportRequestTime,
		AmountAddNoteInCacheRequest:        amountAddNoteInCacheRequest,
		SummaryAddNoteInCacheRequestTime:   summaryAddNoteInCacheRequestTime,
		HistogramAddNoteInCacheRequestTime: histogramAddNoteInCacheRequestTime,
	}
}
