package ramDB

import (
	"sync"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/currancy"
)

type CurrancyDatabase struct {
	db    map[string]currancy.Valute
	mutex sync.Mutex
}

func (d *Database) GetCurrancy(abbreviation string) (currancy.Valute, error) {
	d.Currancy.mutex.Lock()
	defer d.Currancy.mutex.Unlock()
	answer := d.Currancy.db[abbreviation]
	return answer, nil
}

func (d *Database) SetCurrancy(valute currancy.Valute) error {
	d.Currancy.mutex.Lock()
	defer d.Currancy.mutex.Unlock()
	d.Currancy.db[valute.Abbreviation] = valute
	return nil
}
