package ramDB

import (
	"fmt"
	"sync"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

type CurrencyDatabase struct {
	valutes    sync.Map
	userValute sync.Map
}

func (d *Database) GetValute(abbreviation string) (diary.Valute, error) {
	answer, ok := d.Currency.valutes.Load(abbreviation)
	if !ok {
		return diary.Valute{}, fmt.Errorf("No such currency")
	}
	return answer.(diary.Valute), nil
}

func (d *Database) SetValute(valute diary.Valute) error {
	d.Currency.valutes.Store(valute.Abbreviation, valute)
	return nil
}

func (d *Database) GetUserValute(userID int64) (diary.Valute, error) {
	answer, ok := d.Currency.userValute.Load(userID)
	if !ok {
		return diary.Valute{}, fmt.Errorf("Need to select currency")
	}
	return answer.(diary.Valute), nil
}

func (d *Database) SetUserValute(userID int64, valute diary.Valute) error {
	d.Currency.userValute.Store(userID, valute)
	return nil
}
