package tgServer

import (
	"time"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/storage"
)

type TgServer struct {
	storage       storage.Storage
	dateFormatter DateFormatter
}

type DateFormatter struct {
	format       string
	budgetFormat string
}

func New(storage storage.Storage, config Config) (*TgServer, error) {
	return &TgServer{
		storage: storage,
		dateFormatter: DateFormatter{
			format:       config.FormatDate,
			budgetFormat: config.BudgetFormatDate,
		},
	}, nil
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
