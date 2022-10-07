package ramDB

import (
	"sync"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

type Database struct {
	Currency CurrencyDatabase
	Note     NoteDatabase
}

func New() (*Database, error) {
	return &Database{
		Note: NoteDatabase{
			db: map[int64]map[string][]diary.Note{},
		},
		Currency: CurrencyDatabase{
			valutes:    sync.Map{},
			userValute: sync.Map{},
		},
	}, nil
}
