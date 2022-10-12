package storage

import (
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

type NotesDB interface {
	GetNote(id int64, date string) ([]diary.Note, error)
	AddNote(id int64, date string, note diary.Note) error
}

type RatesDB interface {
	GetRate(abbreviation string) (*diary.Valute, error)
	AddRate(valute diary.Valute) error
	SetDefaultCurrency() error
}

type UsersDB interface {
	GetUserAbbValute(userID int64) (string, error)
	AddUserAbbValute(userID int64, abbreviation string) error
}

type Storage interface {
	NotesDB
	UsersDB
	RatesDB
}
