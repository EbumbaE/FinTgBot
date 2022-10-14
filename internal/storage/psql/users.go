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

func (d *Database) SetUserAbbValute(userID int64, abbreviation string) error {
	const queryInsert = `
		INSERT INTO users (
			created_at,
			user_id,
			report_abbreviation
		) VALUES (
			now(), $1, $2
		);
	`
	const queryUpdate = `
		UPDATE users
		SET created_at = now(),
			report_abbreviation = $2
		WHERE user_id = $1; 
	`

	_, err := d.Users.db.Exec(queryInsert,
		userID,
		abbreviation,
	)
	if err != nil {
		_, err = d.Users.db.Exec(queryUpdate,
			userID,
			abbreviation,
		)
		if err != nil {
			return err
		}
	}

	return err
}

func (d *Database) GetUserAbbValute(userID int64) (string, error) {
	const query = `
		SELECT  user_id,
				report_abbreviation
		FROM users
		WHERE user_id = $1
	`

	var getUserID, getAbbreviation string
	err := d.Users.db.QueryRow(query, userID).Scan(&getUserID, &getAbbreviation)
	return getAbbreviation, err
}
