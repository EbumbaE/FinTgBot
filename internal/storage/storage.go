package storage

import (
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/currancy"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

type DiaryDB interface {
	GetNote(id int64, date string) ([]diary.Note, error)
	SetNote(id int64, date string, note diary.Note) error
}

type CurrancyDB interface {
	GetCurrancy(abbreviation string) (currancy.Valute, error)
	SetCurrancy(valute currancy.Valute) error
}

type Storage interface {
	DiaryDB
	CurrancyDB
}
