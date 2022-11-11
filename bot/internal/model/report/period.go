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
