package ramDB

import (
	"sync"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

type NoteDatabase struct {
	db    map[int64]map[string][]diary.Note
	mutex sync.Mutex
}

func (d *Database) GetNote(id int64, date string) ([]diary.Note, error) {
	d.Note.mutex.Lock()
	defer d.Note.mutex.Unlock()
	answer := d.Note.db[id][date]
	return answer, nil
}

func (d *Database) SetNote(id int64, date string, note diary.Note) error {
	d.Note.mutex.Lock()
	defer d.Note.mutex.Unlock()

	if d.Note.db[id] == nil {
		d.Note.db[id] = map[string][]diary.Note{}
	}
	if d.Note.db[id][date] == nil {
		d.Note.db[id][date] = []diary.Note{}
	}
	d.Note.db[id][date] = append(d.Note.db[id][date], note)

	return nil
}
