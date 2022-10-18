package psql

import (
	"database/sql"
	"fmt"
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
			updated_at,
			user_id,
			report_abbreviation
		) VALUES (
			now(), $1, $2
		);
	`
	const queryUpdate = `
		UPDATE users
		SET updated_at = now(),
			report_abbreviation = $2
		WHERE user_id = $1; 
	`

	_, err1 := d.Users.db.Exec(queryInsert,
		userID,
		abbreviation,
	)
	if err1 != nil {
		_, err2 := d.Users.db.Exec(queryUpdate,
			userID,
			abbreviation,
		)
		if err2 != nil {
			return fmt.Errorf("Insert and Update users: %v, %v", err1, err2)
		}
	}

	return nil
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
