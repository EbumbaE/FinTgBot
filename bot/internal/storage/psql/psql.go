package psql

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	Rates   *RatesDB
	Diary   *DiaryDB
	Users   *UsersDB
	Budgets *BudgetsDB
}

func New(cfg Config) (db *Database, err error) {
	db = &Database{}
	db.Rates, err = NewRatesDB(cfg.DriverName, cfg.DataSourceName)
	if err != nil {
		return
	}
	db.Diary, err = NewDiaryDB(cfg.DriverName, cfg.DataSourceName)
	if err != nil {
		return
	}
	db.Users, err = NewUsersDB(cfg.DriverName, cfg.DataSourceName)
	if err != nil {
		return
	}
	db.Budgets, err = NewBudgetsDB(cfg.DriverName, cfg.DataSourceName)
	if err != nil {
		return
	}
	return db, nil
}

func (d *Database) CheckHealth() error {
	if err := d.Diary.db.Ping(); err != nil {
		return err
	}
	if err := d.Rates.db.Ping(); err != nil {
		return err
	}
	if err := d.Users.db.Ping(); err != nil {
		return err
	}
	if err := d.Budgets.db.Ping(); err != nil {
		return err
	}
	return nil
}

func (d *Database) Close() error {
	err1 := d.Diary.db.Close()
	err2 := d.Rates.db.Close()
	err3 := d.Users.db.Close()
	err4 := d.Budgets.db.Close()

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return fmt.Errorf("Error in close db: %w %w %w %w", err1, err2, err3, err4)
	}

	log.Println("All db is closed")
	return nil
}
