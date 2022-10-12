package psql

import (
	"database/sql"
	"log"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

type NotesDB struct {
	db *sql.DB
}

func NewNotesDB(driverName, dataSourceName string) *NotesDB {
	notesDB, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Println("notes database open: ", err)
	}

	return &NotesDB{notesDB}
}

func (d *Database) GetNote(userID int64, date string) (notes []diary.Note, err error) {
	const query = `
		SELECT
			user_id
			date
			note_category,
			note_currency,
			note_sum
		FROM users
		WHERE user_id = $1 AND date = $1
	`
	rows, err := d.Notes.db.Query(query, userID, date)
	if err != nil {
		return nil, err
	}

	var getUserID int64
	var getDate, getNoteCategory, getNoteCurrency string
	var getNoteSum float64

	for rows.Next() {
		if err := rows.Scan(&getUserID, &getDate, &getNoteCategory, &getNoteCurrency, &getNoteSum); err != nil {
			return nil, err
		}
		notes = append(notes, diary.Note{
			Category: getNoteCategory,
			Sum:      getNoteSum,
			Currency: getNoteCurrency,
		})
	}

	return notes, err
}

func (d *Database) AddNote(userID int64, date string, note diary.Note) error {
	const query = `
		insert into notes(
			created_at
			user_id
			date
			note_category,
			note_currency,
			note_sum
		) values (
			now(), $1, $2, $3, $4, $5
		);
	`

	_, err := d.Rates.db.Exec(query,
		userID,
		date,
		note.Category,
		note.Currency,
		note.Sum,
	)

	return err
}
