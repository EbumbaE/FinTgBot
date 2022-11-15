package psql

import (
	"database/sql"
	"log"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

type DiaryDB struct {
	db *sql.DB
}

func NewDiaryDB(driverName, dataSourceName string) (*DiaryDB, error) {
	diaryDB, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Println("notes database open: ", err)
		return nil, err
	}

	return &DiaryDB{diaryDB}, nil
}

func (d *Database) GetNote(userID int64, date string) (notes []diary.Note, err error) {
	const query = `
		SELECT
			note_category,
			note_currency,
			note_sum
		FROM diary
		WHERE user_id = $1 AND date = $2
	`
	rows, err := d.Diary.db.Query(query, userID, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var getNoteCategory, getNoteCurrency string
	var getNoteSum float64

	for rows.Next() {
		if err := rows.Scan(&getNoteCategory, &getNoteCurrency, &getNoteSum); err != nil {
			return nil, err
		}
		notes = append(notes, diary.Note{
			Category: getNoteCategory,
			Sum:      getNoteSum,
			Currency: getNoteCurrency,
		})
	}

	return notes, rows.Err()
}

func (d *Database) AddNote(userID int64, date string, note diary.Note) error {
	const query = `
		INSERT INTO diary(
			user_id,
			date,
			note_category,
			note_currency,
			note_sum
		) VALUES (
			$1, $2, $3, $4, $5
		);
	`

	_, err := d.Diary.db.Exec(query,
		userID,
		date,
		note.Category,
		note.Currency,
		note.Sum,
	)

	return err
}
