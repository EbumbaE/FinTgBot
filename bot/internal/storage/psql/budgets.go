package psql

import (
	"database/sql"
	"log"

	"github.com/EbumbaE/FinTgBot/bot/internal/model/diary"
)

type BudgetsDB struct {
	db *sql.DB
}

func NewBudgetsDB(driverName, dataSourceName string) (*BudgetsDB, error) {
	budgetsDB, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Println("rate database open: ", err)
		return nil, err
	}

	return &BudgetsDB{budgetsDB}, nil
}

func (d *Database) GetMonthlyBudget(userID int64, date string) (*diary.Budget, error) {
	const query = `
		SELECT	value,
				abbreviation
		FROM budgets
		WHERE user_id = $1 AND date = $2
	`

	var getAbbreviation string
	var getMonthlyBudget float64
	err := d.Budgets.db.QueryRow(query, userID, date).Scan(&getMonthlyBudget, &getAbbreviation)
	return &diary.Budget{
		Abbreviation: getAbbreviation,
		Value:        getMonthlyBudget,
		Date:         date,
	}, err
}

func (d *Database) AddMonthlyBudget(userID int64, monthlyBudget diary.Budget) error {
	const query = `
		INSERT INTO budgets (
			user_id,
			date,
			value,
			abbreviation
		) VALUES (
			$1, $2, $3, $4
		)
		ON CONFLICT (user_id, date) DO UPDATE
		SET updated_at = now(),
			value = $3,
			abbreviation = $4
	`

	_, err := d.Budgets.db.Exec(query,
		userID,
		monthlyBudget.Date,
		monthlyBudget.Value,
		monthlyBudget.Abbreviation,
	)

	return err
}
