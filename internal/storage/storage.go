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
	SetUserAbbValute(userID int64, abbreviation string) error
}

type BudgetsDB interface {
	GetMonthlyBudget(userID int64, date string) (*diary.Budget, error)
	AddMonthlyBudget(userID int64, monthlyBudget diary.Budget) error
}

type Storage interface {
	NotesDB
	UsersDB
	RatesDB
	BudgetsDB
}
