package request

import "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"

type ReportRequest struct {
	UserID       int64
	Period       string
	UserCurrency diary.Valute
}
