package psql

import (
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	Rates *RatesDB
	Diary *DiaryDB
	Users *UsersDB
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
	return nil
}

func (d *Database) Close() {
	d.Diary.db.Close()
	d.Rates.db.Close()
	d.Users.db.Close()
	log.Println("All db is closed")
}
