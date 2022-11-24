package storage

import (
	"time"

	"github.com/EbumbaE/FinTgBot/report/internal/model/diary"
	"github.com/EbumbaE/FinTgBot/report/internal/model/report"
	"github.com/bradfitz/gomemcache/memcache"
)

type NotesDB interface {
	GetNote(id int64, date string) ([]diary.Note, error)
}

type Storage interface {
	NotesDB
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
