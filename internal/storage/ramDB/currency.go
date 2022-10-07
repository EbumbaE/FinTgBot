package ramDB

import (
	"fmt"
	"sync"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/currency"
)

type CurrencyDatabase struct {
	db    map[string]currency.Valute
	mutex sync.Mutex
}

func (d *Database) GetCurrency(abbreviation string) (currency.Valute, error) {
	d.Currency.mutex.Lock()
	defer d.Currency.mutex.Unlock()
	answer, ok := d.Currency.db[abbreviation]
	if !ok {
		return currency.Valute{}, fmt.Errorf("No such currency")
	}
	return answer, nil
}

func (d *Database) SetCurrency(valute currency.Valute) error {
	d.Currency.mutex.Lock()
	defer d.Currency.mutex.Unlock()
	d.Currency.db[valute.Abbreviation] = valute
	return nil
}
