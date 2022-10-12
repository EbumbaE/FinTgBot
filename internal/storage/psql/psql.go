package psql

import "log"

type Database struct {
	Rates RatesDB
	Notes NotesDB
	Users UsersDB
}

func New(cfg Config) (*Database, error) {
	return &Database{
		Rates: *NewRatesDB(cfg.DriverName, cfg.DataSourceName),
		Notes: *NewNotesDB(cfg.DriverName, cfg.DataSourceName),
		Users: *NewUsersDB(cfg.DriverName, cfg.DataSourceName),
	}, nil
}

func (d *Database) CheckHealth() error {
	if err := d.Notes.db.Ping(); err != nil {
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
	d.Notes.db.Close()
	d.Rates.db.Close()
	d.Users.db.Close()
	log.Println("All db is closed")
}
