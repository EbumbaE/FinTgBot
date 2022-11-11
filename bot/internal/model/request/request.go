package request

import (
	"time"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

type ReportRequest struct {
	UserID       int64
	Period       string
	UserCurrency diary.Valute
}

type AddNoteInCacheRequest struct {
	UserID   int64
	TimeNote time.Time
	Note     diary.Note
}
