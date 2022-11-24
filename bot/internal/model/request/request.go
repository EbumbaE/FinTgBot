package request

import (
	"time"

	"github.com/EbumbaE/FinTgBot/bot/internal/model/diary"
)

type ReportRequest struct {
	UserID       int64
	Period       string
	DateFormat   string
	UserCurrency diary.Valute
}

type AddNoteInCacheRequest struct {
	UserID   int64
	TimeNote time.Time
	Note     diary.Note
}
