package tgServer

import (
	"fmt"
	"time"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/currency"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/storage"
)

type TgServer struct {
	storage      storage.Storage
	dateFormater DateFormater
}

type DateFormater struct {
	format string
}

func New(storage storage.Storage, config Config) (*TgServer, error) {
	return &TgServer{
		storage:      storage,
		dateFormater: DateFormater{format: config.FormatDate},
	}, nil
}

func (t *TgServer) InitCurrancies(currancies chan currency.Valute) {
	go func() {
		for valute := range currancies {
			if err := t.storage.SetCurrency(valute); err != nil {
				fmt.Printf("Error in set currency: %s", valute.Abbreviation)
			}
		}
	}()
}

func (d *DateFormater) FormatDate(date time.Time) string {
	return date.Format(d.format)
}

func (t *TgServer) correctDate(date string) (string, error) {
	parseDate, err := time.Parse(t.dateFormater.format, date)
	if err != nil {
		return "", err
	}
	return t.dateFormater.FormatDate(parseDate), nil
}
