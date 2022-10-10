package tgServer

import (
	"context"
	"log"
	"time"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
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

func setDefaultCurrancy(db storage.Storage) {
	db.SetRate(diary.Valute{
		Abbreviation: "RUB",
		Name:         "Российский рубль",
		Value:        1,
	})
}

func (t *TgServer) InitCurrancies(ctx context.Context, currencies chan diary.Valute) {

	setDefaultCurrancy(t.storage)

	go func() {
		for {
			select {
			case valute := <-currencies:
				if err := t.storage.SetRate(valute); err != nil {
					log.Printf("[Error in set currency %s] %v\n", valute.Abbreviation, err)
				}
			case <-ctx.Done():
				log.Println("accepting currencies from the parser is off")
				return
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
