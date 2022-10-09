package storage

import (
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

type DiaryDB interface {
	GetNote(id int64, date string) ([]diary.Note, error)
	SetNote(id int64, date string, note diary.Note) error
}

type CurrencyDB interface {
	GetRate(abbreviation string) (diary.Valute, error)
	SetRate(valute diary.Valute) error
	GetUserAbbValute(userID int64) (string, error)
	SetUserAbbValute(userID int64, abbreviation string) error
}

type Storage interface {
	DiaryDB
	CurrencyDB
}
