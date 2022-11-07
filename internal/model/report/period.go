package report

import (
	"fmt"
	"time"
)

func GetWeekPeriod(now time.Time) (beginPeriod, endPeriod time.Time) {
	nowWeekday := now.Weekday()
	if nowWeekday == 0 {
		nowWeekday = 7
	}
	beginPeriod = now.AddDate(0, 0, -int(nowWeekday)+1)
	endPeriod = beginPeriod.AddDate(0, 0, 6)
	return
}

func GetMonthPeriod(now time.Time) (beginPeriod, endPeriod time.Time) {
	currYear, currMonth, _ := now.Date()
	beginPeriod = time.Date(currYear, currMonth, 1, 0, 0, 0, 0, now.Location())
	endPeriod = beginPeriod.AddDate(0, 1, -1)
	return
}

func GetYearPeriod(now time.Time) (beginPeriod, endPeriod time.Time) {
	currYear, _, _ := now.Date()
	beginPeriod = time.Date(currYear, 1, 1, 0, 0, 0, 0, now.Location())
	endPeriod = beginPeriod.AddDate(1, 0, -1)
	return
}

func GetPeriod(period string) (beginPeriod, endPeriod time.Time, err error) {

	err = nil
	tn := time.Now()
	switch period {
	case "week":
		beginPeriod, endPeriod = GetWeekPeriod(tn)
		return
	case "month":
		beginPeriod, endPeriod = GetMonthPeriod(tn)
		return
	case "year":
		beginPeriod, endPeriod = GetYearPeriod(tn)
		return
	}
	return tn, tn, fmt.Errorf("Error in period")
}

func compareAGreaterOrEqualB(a, b time.Time) bool {
	return a.Unix()-b.Unix() >= 0
}

func DeterminePeriod(date time.Time, now time.Time) (period []string, err error) {

	begin, end := GetWeekPeriod(now)
	if compareAGreaterOrEqualB(date, begin) && compareAGreaterOrEqualB(end, date) {
		period = append(period, "week")
	}

	begin, _ = GetMonthPeriod(now)
	if compareAGreaterOrEqualB(date, begin) && compareAGreaterOrEqualB(end, date) {
		period = append(period, "month")
	}

	begin, _ = GetYearPeriod(now)
	if compareAGreaterOrEqualB(date, begin) && compareAGreaterOrEqualB(end, date) {
		period = append(period, "year")
	}

	if period == nil {
		return nil, fmt.Errorf("Error in determine period date %s", date.Format("02.01.2006"))
	}
	return
}
