package report

import (
	"time"
)

type Formatter interface {
	FormatDateTimeToString(date time.Time) string
	FormatDateStringToTime(date string) (time.Time, error)
	CorrectMonthYear(date string) (string, error)
}

type Valute interface {
	GetAbbreviation() string
	GetValue() float64
}

type Note interface {
	GetCategory() string
	GetSum() float64
}

type Storage interface {
	GetRate(abbreviation string) (Valute, error)
	GetNote(id int64, date string) ([]Note, error)
}

type ReportFormat map[string]float64

func CountStatistic(userID int64, period string, db Storage, formatter Formatter) (report ReportFormat, err error) {

	beginPeriod, endPeriod, err := GetPeriod(period)
	if err != nil {
		return nil, err
	}

	report = map[string]float64{} //check
	for date := beginPeriod; date != endPeriod.AddDate(0, 0, 1); date = date.AddDate(0, 0, 1) {
		notes, err := db.GetNote(userID, formatter.FormatDateTimeToString(date))
		if err != nil {
			return nil, err
		}
		for _, note := range notes {
			report[note.GetCategory()] += note.GetSum()
		}
	}

	return
}
