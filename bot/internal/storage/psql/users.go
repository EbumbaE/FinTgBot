package psql

import (
	"database/sql"
	"log"
)

type UsersDB struct {
	db *sql.DB
}

func NewUsersDB(driverName, dataSourceName string) (*UsersDB, error) {
	usersDB, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Println("users database open: ", err)
		return nil, err
	}

	return &UsersDB{usersDB}, nil
}

func (d *Database) SetUserAbbValute(userID int64, abbreviation string) (err error) {
	const query = `
		INSERT INTO users (
			created_at,
			id,
			report_abbreviation
		) VALUES (
			now(), $1, $2
		)
		ON CONFLICT (id) DO UPDATE
		SET updated_at = now(),
			report_abbreviation = $2;
	`

	_, err = d.Users.db.Exec(query,
		userID,
		abbreviation,
	)
	return
}

func (d *Database) GetUserAbbValute(userID int64) (abbreviation string, err error) {
	const query = `
		SELECT report_abbreviation
		FROM users
		WHERE id = $1
	`
	err = d.Users.db.QueryRow(query, userID).Scan(&abbreviation)
	return
}

func (d *Database) CheckUser(userID int64) (err error) {
	const query = `
		INSERT INTO users (
			created_at,
			id
		) VALUES (
			now(), $1
		)	
	`
	_, err = d.Users.db.Exec(query, userID)
	return
}
