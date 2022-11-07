package tgServer

import (
	"time"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/report"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/servers/middleware"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/storage"
)

type ReportCache interface {
	AddReportInCache(userID int64, period string, addedReport report.ReportFormat) (err error)
	GetReportFromCache(userID int64, period string) (getReport report.ReportFormat, err error)
	AddNoteInCacheReports(userID int64, date time.Time, note diary.Note) error
}

type TgServer struct {
	storage       storage.Storage
	cache         ReportCache
	dateFormatter DateFormatter
	Metrics       *middleware.Metrics
}

type DateFormatter struct {
	format       string
	budgetFormat string
}

func New(storage storage.Storage, cache ReportCache, config Config) (*TgServer, error) {
	return &TgServer{
		storage: storage,
		cache:   cache,
		dateFormatter: DateFormatter{
			format:       config.FormatDate,
			budgetFormat: config.BudgetFormatDate,
		},
	}, nil
}

func (s *TgServer) InitMiddleware() {
	s.Metrics = middleware.NewMetrics()
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
