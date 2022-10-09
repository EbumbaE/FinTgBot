package storage

import (
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

type DiaryDB interface {
	GetNote(id int64, date string) ([]diary.Note, error)
	SetNote(id int64, date string, note diary.Note) error
}

type CurrencyDB interface {
	GetValute(abbreviation string) (diary.Valute, error)
	SetValute(valute diary.Valute) error
	GetUserValute(userID int64) (diary.Valute, error)
	SetUserValute(userID int64, valute diary.Valute) error
}

type Storage interface {
	DiaryDB
	CurrencyDB
}
