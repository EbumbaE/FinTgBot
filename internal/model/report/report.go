package report

import (
	"fmt"
	"time"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

type Storage interface {
	Get(id int64, date string) []diary.Note
}

type Formater interface {
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

func getPeriod(period string) (beginPeriod, endPeriod time.Time, ok bool) {

	tn := time.Now()
	switch period {
	case "week":
		beginPeriod, endPeriod = getWeekPeriod(tn)
		ok = true
		return
	case "month":
		beginPeriod, endPeriod = getMonthPeriod(tn)
		ok = true
		return
	case "year":
		beginPeriod, endPeriod = getYearPeriod(tn)
		ok = true
		return
	}
	return tn, tn, false
}

func addNoteToTotalSum(notes []diary.Note, totalSum *map[string]float64) {
	for _, note := range notes {
		(*totalSum)[note.Category] += note.Sum
	}
	return
}

func CountStatistic(userID int64, period string, db Storage, formater Formater) (answer string, err error) {

	answer = fmt.Sprintf("Statistic for the %s:\n", period)

	beginPeriod, endPeriod, ok := getPeriod(period)
	if !ok {
		return "", fmt.Errorf("error in period")
	}

	totalSum := map[string]float64{}
	for date := beginPeriod; date != endPeriod; date = date.AddDate(0, 0, 1) {
		notes := db.Get(userID, formater.FormatDate((date)))
		addNoteToTotalSum(notes, &totalSum)
	}

	for category, sum := range totalSum {
		answer += fmt.Sprintf("%s: %.2f\n", category, sum)
	}

	return answer, nil
}
