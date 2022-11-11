package psql

import (
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	Diary *DiaryDB
}

func New(cfg Config) (db *Database, err error) {
	db = &Database{}
	db.Diary, err = NewDiaryDB(cfg.DriverName, cfg.DataSourceName)
	if err != nil {
		return
	}
	return db, nil
}

func (d *Database) CheckHealth() error {
	if err := d.Diary.db.Ping(); err != nil {
		return err
	}
	return nil
}

func (d *Database) Close() error {
	err := d.Diary.db.Close()
	if err != nil {
		return err
	}
	log.Println("All db is closed")
	return nil
}
