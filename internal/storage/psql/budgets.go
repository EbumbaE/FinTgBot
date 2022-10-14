package psql

import (
	"database/sql"
	"log"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
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
		SELECT	user_id,
				date,
				value,
				abbreviation
		FROM budgets
		WHERE user_id = $1 AND date = $2
	`

	var getUserID, getDate, getAbbreviation string
	var getMonthlyBudget float64
	err := d.Budgets.db.QueryRow(query, userID, date).Scan(&getUserID, &getDate, &getMonthlyBudget, &getAbbreviation)
	return &diary.Budget{
		Abbreviation: getAbbreviation,
		Value:        getMonthlyBudget,
		Date:         getDate,
	}, err
}

func (d *Database) AddMonthlyBudget(userID int64, monthlyBudget diary.Budget) error {
	const queryInsert = `
		INSERT INTO budgets (
			created_at,
			user_id,
			date,
			value,
			abbreviation
		) VALUES (
			now(), $1, $2, $3, $4
		);
	`
	const queryUpdate = `
		UPDATE budgets
		SET created_at = now(),
			value = $3,
			abbreviation = $4
		WHERE user_id = $1 AND date = $2; 
	`

	_, err := d.Budgets.db.Exec(queryInsert,
		userID,
		monthlyBudget.Date,
		monthlyBudget.Value,
		monthlyBudget.Abbreviation,
	)
	if err != nil {
		_, err = d.Budgets.db.Exec(queryUpdate,
			userID,
			monthlyBudget.Date,
			monthlyBudget.Value,
			monthlyBudget.Abbreviation,
		)
	}

	return err
}
