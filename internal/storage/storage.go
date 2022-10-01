package storage

import "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"

type DiaryDB interface {
	Get(id int64, date string) ([]diary.Note, error)
	Set(id int64, date string, note diary.Note) error
}

type Storage interface {
	DiaryDB
}
