package middleware

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	AmountCommands              prometheus.Counter
	SummaryCountStatisticTime   prometheus.Summary
	HistogramCountStatisticTime prometheus.Histogram
	SummaryAddNoteInCacheTime   prometheus.Summary
	HistogramAddNoteInCacheTime prometheus.Histogram
}

func NewMetrics() *Metrics {
	amountCommands := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "ozon",
		Name:      "report_amount_commands",
	})

	summaryCountStatisticTime := promauto.NewSummary(prometheus.SummaryOpts{
		Namespace: "ozon",
		Name:      "report_summary_countStatistic_time_seconds",
		Objectives: map[float64]float64{
			0.5:  0.1,
			0.9:  0.01,
			0.99: 0.001,
		},
	})

	histogramCountStatisticTime := promauto.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "ozon",
			Name:      "report_histogram_countStatistic_time_seconds",
			Buckets:   []float64{0.0001, 0.0005, 0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 2},
		},
	)

	summaryAddNoteInCacheTime := promauto.NewSummary(prometheus.SummaryOpts{
		Namespace: "ozon",
		Name:      "report_summary_addNoteInCache_time_seconds",
		Objectives: map[float64]float64{
			0.5:  0.1,
			0.9:  0.01,
			0.99: 0.001,
		},
	})

	histogramAddNoteInCacheTime := promauto.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "ozon",
			Name:      "report_histogram_addNoteInCache_time_seconds",
			Buckets:   []float64{0.0001, 0.0005, 0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 2},
		},
	)

	return &Metrics{
		AmountCommands:              amountCommands,
		SummaryCountStatisticTime:   summaryCountStatisticTime,
		HistogramCountStatisticTime: histogramCountStatisticTime,
		SummaryAddNoteInCacheTime:   summaryAddNoteInCacheTime,
		HistogramAddNoteInCacheTime: histogramAddNoteInCacheTime,
	}
}
