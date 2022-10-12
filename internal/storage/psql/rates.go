package psql

import (
	"database/sql"
	"log"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

type RatesDB struct {
	db *sql.DB
}

func NewRatesDB(driverName, dataSourceName string) *RatesDB {
	rateDB, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Println("rate database open: ", err)
	}

	return &RatesDB{rateDB}
}

func (d *Database) AddRate(currency diary.Valute) error {
	const query = `
		insert into raets(
			created_at,
			abbreviation,
			name,
			value,
			ts
		) values (
			now(), $1, $2, $3, $4
		);
	`

	_, err := d.Rates.db.Exec(query,
		currency.Abbreviation,
		currency.Name,
		currency.Value,
		currency.TimeStep,
	)

	return err
}

func (d *Database) GetRate(abbreviation string) (*diary.Valute, error) {
	const query = `
		select abbreviation
				name,
				value,
				time_step
		from rates
		where abbreviation = $1
	`

	var rate diary.Valute
	err := d.Rates.db.QueryRow(query, abbreviation).Scan(&rate)
	return &rate, err
}

func (d *Database) SetDefaultCurrency() error {
	return d.AddRate(diary.Valute{
		Abbreviation: "RUB",
		Name:         "Российский рубль",
		Value:        1,
	})
}
