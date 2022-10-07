package ramDB

import (
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/currancy"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

type Database struct {
	Currancy CurrancyDatabase
	Note     NoteDatabase
}

func New() (*Database, error) {
	return &Database{
		Note: NoteDatabase{
			db: map[int64]map[string][]diary.Note{},
		},
		Currancy: CurrancyDatabase{
			db: map[string]currancy.Valute{},
		},
	}, nil
}
