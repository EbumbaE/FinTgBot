package ramDB

import (
	"fmt"
	"sync"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/currency"
)

type CurrencyDatabase struct {
	valutes    sync.Map
	userValute sync.Map
}

func (d *Database) GetValute(abbreviation string) (currency.Valute, error) {
	answer, ok := d.Currency.valutes.Load(abbreviation)
	if !ok {
		return currency.Valute{}, fmt.Errorf("No such currency")
	}
	return answer.(currency.Valute), nil
}

func (d *Database) SetValute(valute currency.Valute) error {
	d.Currency.valutes.Store(valute.Abbreviation, valute)
	return nil
}

func (d *Database) GetUserAbbValute(userID int64) (string, error) {
	answer, ok := d.Currency.userValute.Load(userID)
	if !ok {
		return "", fmt.Errorf("No such user")
	}
	return answer.(string), nil
}

func (d *Database) SetUserAbbValute(userID int64, abbreviation string) error {
	d.Currency.userValute.Store(userID, abbreviation)
	return nil
}
