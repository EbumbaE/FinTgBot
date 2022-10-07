package storage

import (
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/currency"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

type DiaryDB interface {
	GetNote(id int64, date string) ([]diary.Note, error)
	SetNote(id int64, date string, note diary.Note) error
}

type CurrencyDB interface {
	GetCurrency(abbreviation string) (currency.Valute, error)
	SetCurrency(valute currency.Valute) error
}

type Storage interface {
	DiaryDB
	CurrencyDB
}
