package psql

import (
	"database/sql"
	"log"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

type RatesDB struct {
	db *sql.DB
}

func NewRatesDB(driverName, dataSourceName string) (*RatesDB, error) {
	rateDB, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Println("rate database open: ", err)
		return nil, err
	}

	return &RatesDB{rateDB}, nil
}

func (d *Database) AddRate(currency diary.Valute) error {
	const query = `
		INSERT INTO rates(
			created_at,
			abbreviation,
			name,
			value
		) VALUES (
			now(), $1, $2, $3
		);
	`

	_, err := d.Rates.db.Exec(query,
		currency.Abbreviation,
		currency.Name,
		currency.Value,
	)

	return err
}

func (d *Database) GetRate(abbreviation string) (*diary.Valute, error) {
	const query = `
		SELECT  created_at,
				abbreviation,
				name,
				value
		FROM rates
		WHERE abbreviation = $1
		ORDER BY created_at DESC
	`

	var getCreatedAt, getAbbreviation, getName string
	var getValue float64
	err := d.Rates.db.QueryRow(query, abbreviation).Scan(&getCreatedAt, &getAbbreviation, &getName, &getValue)

	return &diary.Valute{
		Abbreviation: getAbbreviation,
		Name:         getName,
		Value:        getValue,
	}, err
}

func (d *Database) SetDefaultCurrency() error {
	return d.AddRate(diary.Valute{
		Abbreviation: "RUB",
		Name:         "Российский рубль",
		Value:        1,
	})
}
