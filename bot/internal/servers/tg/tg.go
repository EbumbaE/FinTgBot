package tgServer

import (
	"context"
	"time"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/request"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/servers/middleware"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/storage"
)

type Producer interface {
	SendReportRequest(context.Context, request.ReportRequest) error
	SendAddNoteInCache(context.Context, request.AddNoteInCacheRequest) error
}

type TgServer struct {
	storage       storage.Storage
	dateFormatter DateFormatter
	Metrics       *middleware.TgMetrics
	producer      Producer
}

type DateFormatter struct {
	format       string
	budgetFormat string
}

func New(storage storage.Storage, producer Producer, config Config) (*TgServer, error) {
	return &TgServer{
		storage: storage,
		dateFormatter: DateFormatter{
			format:       config.FormatDate,
			budgetFormat: config.BudgetFormatDate,
		},
		producer: producer,
	}, nil
}

func (s *TgServer) InitMiddleware() {
	s.Metrics = middleware.NewTgMetrics()
}

func (d *DateFormatter) FormatDateTimeToString(date time.Time) string {
	return date.Format(d.format)
}

func (d *DateFormatter) FormatDateStringToTime(date string) (t time.Time, err error) {
	t, err = time.Parse(d.format, date)
	return
}

func (df *DateFormatter) CorrectDate(date string) (string, error) {
	parseDate, err := time.Parse(df.format, date)
	if err != nil {
		return "", err
	}
	return df.FormatDateTimeToString(parseDate), nil
}

func (df *DateFormatter) CorrectMonthYear(date string) (string, error) {
	parseDate, err := time.Parse(df.budgetFormat, date)
	if err != nil {
		return "", err
	}
	return df.FormatDateTimeToString(parseDate), nil
}
