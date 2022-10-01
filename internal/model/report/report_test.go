package report

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Period struct {
	begin, end time.Time
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
		Period{time.Date(2022, 9, 26, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 10, 02, 0, 0, 0, 0, time.Now().Location())},
	},
}

var monthTestsPair = []Testpair{
	{
		[]Date{{1, 9, 2022}, {10, 9, 2022}, {28, 9, 2022}, {29, 9, 2022}},
		Period{
			time.Date(2022, 9, 1, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 9, 30, 0, 0, 0, 0, time.Now().Location()),
		},
	},
	{
		[]Date{{1, 10, 2022}, {10, 10, 2022}, {28, 10, 2022}, {29, 10, 2022}},
		Period{time.Date(2022, 10, 1, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 10, 31, 0, 0, 0, 0, time.Now().Location())},
	},
}

var yearTestsPair = []Testpair{
	{
		[]Date{{26, 1, 2022}, {27, 5, 2022}, {2, 4, 2022}, {20, 12, 2022}},
		Period{time.Date(2022, 1, 1, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 12, 31, 0, 0, 0, 0, time.Now().Location())},
	},
	{
		[]Date{{26, 1, 2021}, {27, 5, 2021}, {2, 4, 2021}, {20, 12, 2021}},
		Period{time.Date(2021, 1, 1, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2021, 12, 31, 0, 0, 0, 0, time.Now().Location())},
	},
}

func Test_GetWeekPeriod(t *testing.T) {

	for _, pair := range weekTestsPair {
		for _, value := range pair.values {
			testNowTime := time.Date(value.year, time.Month(value.month), value.day, 0, 0, 0, 0, time.Now().Location())
			beginPeriod, endPeriod := getWeekPeriod(testNowTime)

			assert.Equal(t, beginPeriod, pair.average.begin, "expected: %s, but get %s\n", beginPeriod, pair.average.end)
			assert.Equal(t, endPeriod, pair.average.end, "expected: %s, but get %s\n", endPeriod, pair.average.end)
		}
	}
}

func Test_GetMonthPeriod(t *testing.T) {
	for _, pair := range monthTestsPair {
		for _, value := range pair.values {
			testNowTime := time.Date(value.year, time.Month(value.month), value.day, 0, 0, 0, 0, time.Now().Location())
			beginPeriod, endPeriod := getMonthPeriod(testNowTime)

			assert.Equal(t, beginPeriod, pair.average.begin, "expected: %s, but get %s\n", beginPeriod, pair.average.end)
			assert.Equal(t, endPeriod, pair.average.end, "expected: %s, but get %s\n", endPeriod, pair.average.end)
		}
	}
}

func Test_GetYearPeriod(t *testing.T) {
	for _, pair := range yearTestsPair {
		for _, value := range pair.values {
			testNowTime := time.Date(value.year, time.Month(value.month), value.day, 0, 0, 0, 0, time.Now().Location())
			beginPeriod, endPeriod := getYearPeriod(testNowTime)

			assert.Equal(t, beginPeriod, pair.average.begin, "expected: %s, but get %s\n", beginPeriod, pair.average.end)
			assert.Equal(t, endPeriod, pair.average.end, "expected: %s, but get %s\n", endPeriod, pair.average.end)
		}
	}
}
