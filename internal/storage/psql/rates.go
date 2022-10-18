package psql

import (
	"database/sql"
	"fmt"
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
	const queryInsert = `
		INSERT INTO rates(
			id,
			abbreviation,
			name,
			value
		) VALUES (
			$1, $2, $3, $4
		);
	`
	const queryUpdate = `
		UPDATE rates
		SET updated_at = now(),
			abbreviation = $2,
			name = $3,
			value = $4
		WHERE id = $1; 
	`

	_, err1 := d.Rates.db.Exec(queryInsert,
		currency.ID,
		currency.Abbreviation,
		currency.Name,
		currency.Value,
	)
	if err1 != nil {
		_, err2 := d.Rates.db.Exec(queryUpdate,
			currency.ID,
			currency.Abbreviation,
			currency.Name,
			currency.Value,
		)
		if err2 != nil {
			return fmt.Errorf("Insert and Update rates: %v, %v", err1, err2)
		}
	}

	return nil
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
