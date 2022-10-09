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

func (d *Database) GetRate(abbreviation string) (diary.Valute, error) {
	answer, ok := d.Currency.valutes.Load(abbreviation)
	if !ok {
		return diary.Valute{}, fmt.Errorf("No such currency")
	}
	return answer.(diary.Valute), nil
}

func (d *Database) SetRate(valute diary.Valute) error {
	d.Currency.valutes.Store(valute.Abbreviation, valute)
	return nil
}

func (d *Database) GetUserAbbValute(userID int64) (string, error) {
	answer, ok := d.Currency.userValute.Load(userID)
	if !ok {
		return "", fmt.Errorf("Not have user valute")
	}
	return answer.(string), nil
}

func (d *Database) SetUserAbbValute(userID int64, abbreviation string) error {
	d.Currency.userValute.Store(userID, abbreviation)
	return nil
}
