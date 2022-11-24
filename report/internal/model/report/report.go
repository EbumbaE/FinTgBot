package report

import "github.com/EbumbaE/FinTgBot/report/internal/model/diary"

type Valute interface {
	GetAbbreviation() string
	GetValue() float64
}

type Storage interface {
	GetNote(id int64, date string) ([]diary.Note, error)
}

type ReportFormat map[string]float64

func CountStatistic(userID int64, period string, db Storage, dateFormat string) (report ReportFormat, err error) {

	beginPeriod, endPeriod, err := GetPeriod(period)
	if err != nil {
		return nil, err
	}

	report = map[string]float64{}
	for date := beginPeriod; date != endPeriod.AddDate(0, 0, 1); date = date.AddDate(0, 0, 1) {
		notes, err := db.GetNote(userID, date.Format(dateFormat))
		if err != nil {
			return nil, err
		}
		for _, note := range notes {
			report[note.Category] += note.Sum
		}
	}

	return
}
