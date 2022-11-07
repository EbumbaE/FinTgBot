package storage

import (
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/report"
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
	CheckUser(userID int64) error
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

type ReportCache interface {
	AddReportInCache(userID int64, period string, addedReport report.ReportFormat) (err error)
	GetReportFromCache(userID int64, period string) (getReport report.ReportFormat, err error)
	AddNoteInCacheReports(userID int64, date time.Time, note diary.Note) error
}

type Cache interface {
	Get(key string) (*memcache.Item, error)
	Add(item *memcache.Item) error
	Delete(key string) error
	Ping() error

	ReportCache
}
