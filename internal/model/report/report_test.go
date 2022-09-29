package report

import (
	"testing"
	"time"
)

type Period struct {
	begin, end string
}

type Date struct {
	day, month, year int
}

type Testpair struct {
	values  []Date
	average Period
}

var weekTestsPair = []Testpair{
	{
		[]Date{{26, 9, 2022}, {27, 9, 2022}, {28, 9, 2022}, {29, 9, 2022}, {30, 9, 2022}, {1, 10, 2022}, {2, 10, 2022}},
		Period{"26.09.2022", "02.10.2022"},
	},
}

var monthTestsPair = []Testpair{
	{
		[]Date{{1, 9, 2022}, {10, 9, 2022}, {28, 9, 2022}, {29, 9, 2022}},
		Period{"01.09.2022", "30.09.2022"},
	},
	{
		[]Date{{1, 10, 2022}, {10, 10, 2022}, {28, 10, 2022}, {29, 10, 2022}},
		Period{"01.10.2022", "31.10.2022"},
	},
}

var yearTestsPair = []Testpair{
	{
		[]Date{{26, 1, 2022}, {27, 5, 2022}, {2, 4, 2022}, {20, 12, 2022}},
		Period{"01.01.2022", "31.12.2022"},
	},
	{
		[]Date{{26, 1, 2021}, {27, 5, 2021}, {2, 4, 2021}, {20, 12, 2021}},
		Period{"01.01.2021", "31.12.2021"},
	},
}

func Test_GetWeekPeriod(t *testing.T) {

	for _, pair := range weekTestsPair {
		for _, value := range pair.values {
			testNowTime := time.Date(value.year, time.Month(value.month), value.day, 0, 0, 0, 0, time.Now().Location())
			beginPeriod, endPeriod := getWeekPeriod(testNowTime)

			begin := beginPeriod.Format("02.01.2006")
			end := endPeriod.Format("02.01.2006")
			if begin != pair.average.begin || end != pair.average.end {
				t.Errorf("expected: %s-%s, but get %s-%s", pair.average.begin, pair.average.end, begin, end)
			}
		}
	}
}

func Test_GetMonthPeriod(t *testing.T) {
	for _, pair := range monthTestsPair {
		for _, value := range pair.values {
			testNowTime := time.Date(value.year, time.Month(value.month), value.day, 0, 0, 0, 0, time.Now().Location())
			beginPeriod, endPeriod := getMonthPeriod(testNowTime)

			begin := beginPeriod.Format("02.01.2006")
			end := endPeriod.Format("02.01.2006")
			if begin != pair.average.begin || end != pair.average.end {
				t.Errorf("expected: %s-%s, but get %s-%s", pair.average.begin, pair.average.end, begin, end)
			}
		}
	}
}

func Test_GetYearPeriod(t *testing.T) {
	for _, pair := range yearTestsPair {
		for _, value := range pair.values {
			testNowTime := time.Date(value.year, time.Month(value.month), value.day, 0, 0, 0, 0, time.Now().Location())
			beginPeriod, endPeriod := getYearPeriod(testNowTime)

			begin := beginPeriod.Format("02.01.2006")
			end := endPeriod.Format("02.01.2006")
			if begin != pair.average.begin || end != pair.average.end {
				t.Errorf("expected: %s-%s, but get %s-%s", pair.average.begin, pair.average.end, begin, end)
			}
		}
	}
}
