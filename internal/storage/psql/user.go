package psql

import (
	"database/sql"
	"log"
)

type UsersDB struct {
	db *sql.DB
}

func NewUsersDB(driverName, dataSourceName string) *UsersDB {
	usersDB, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Println("users database open: ", err)
	}

	return &UsersDB{usersDB}
}

func (d *Database) GetUserAbbValute(userID int64) (abbrevation string, err error) {
	const query = `
		select  user_id
				abbreviation
		from users
		where user_id = $1
	`

	err = d.Users.db.QueryRow(query, userID).Scan(&abbrevation)
	return abbrevation, err
}

func (d *Database) AddUserAbbValute(userID int64, abbreviation string) error {
	const query = `
		insert into users(
			created_at
			user_id
			abbreviation
		) values (
			now(), $1, $2
		);
	`

	_, err := d.Rates.db.Exec(query,
		userID,
		abbreviation,
	)

	return err
}
