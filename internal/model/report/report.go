package report

import (
	"fmt"
	"time"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

type Storage interface {
	Get(id int64, date string) ([]diary.Note, error)
}

type Formatter interface {
	FormatDate(date time.Time) string
}

func getWeekPeriod(now time.Time) (beginPeriod, endPeriod time.Time) {
	nowWeekday := now.Weekday()
	if nowWeekday == 0 {
		nowWeekday = 7
	}
	beginPeriod = now.AddDate(0, 0, -int(nowWeekday)+1)
	endPeriod = beginPeriod.AddDate(0, 0, 6)
	return
}

func getMonthPeriod(now time.Time) (beginPeriod, endPeriod time.Time) {
	currYear, currMonth, _ := now.Date()
	beginPeriod = time.Date(currYear, currMonth, 1, 0, 0, 0, 0, now.Location())
	endPeriod = beginPeriod.AddDate(0, 1, -1)
	return
}

func getYearPeriod(now time.Time) (beginPeriod, endPeriod time.Time) {
	currYear, _, _ := now.Date()
	beginPeriod = time.Date(currYear, 1, 1, 0, 0, 0, 0, now.Location())
	endPeriod = beginPeriod.AddDate(1, 0, -1)
	return
}

func getPeriod(period string) (beginPeriod, endPeriod time.Time, err error) {

	err = nil
	tn := time.Now()
	switch period {
	case "week":
		beginPeriod, endPeriod = getWeekPeriod(tn)
		return
	case "month":
		beginPeriod, endPeriod = getMonthPeriod(tn)
		return
	case "year":
		beginPeriod, endPeriod = getYearPeriod(tn)
		return
	}
	return tn, tn, fmt.Errorf("Error in period")
}

func CountStatistic(userID int64, period string, db Storage, formatter Formatter) (answer string, err error) {

	answer = fmt.Sprintf("Statistic for the %s:\n", period)

	beginPeriod, endPeriod, err := getPeriod(period)
	if err != nil {
		return "", err
	}

	totalSum := map[string]float64{}
	for date := beginPeriod; date != endPeriod; date = date.AddDate(0, 0, 1) {
		notes, err := db.Get(userID, formatter.FormatDate(date))
		if err != nil {
			return "Error in storage: get note", err
		}
		for _, note := range notes {
			totalSum[note.Category] += note.Sum
		}
	}

	for category, sum := range totalSum {
		answer += fmt.Sprintf("%s: %.2f\n", category, sum)
	}

	return answer, nil
}
