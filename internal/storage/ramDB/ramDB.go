package ramDB

import (
	"sync"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

type Database struct {
	db    map[int64]map[string][]diary.Note
	mutex sync.Mutex
}

func New() (*Database, error) {
	return &Database{
		db: map[int64]map[string][]diary.Note{},
	}, nil
}

func (d *Database) Get(id int64, date string) ([]diary.Note, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	answer := d.db[id][date]
	return answer, nil
}

func (d *Database) Set(id int64, date string, note diary.Note) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.db[id] == nil {
		d.db[id] = map[string][]diary.Note{}
	}
	if d.db[id][date] == nil {
		d.db[id][date] = []diary.Note{}
	}
	d.db[id][date] = append(d.db[id][date], note)

	return nil
}
