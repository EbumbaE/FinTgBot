package tgServer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type FormatterAnswer struct {
	answer string
	err    error
}

type FormatterTestsPair struct {
	value   string
	average FormatterAnswer
}

var correctDateTestsPair = []FormatterTestsPair{
	{value: "02.10.2022", average: FormatterAnswer{"02.10.2022", nil}},
	{value: "24.12.2022", average: FormatterAnswer{"24.12.2022", nil}},
	{value: "02/10/2022", average: FormatterAnswer{"", fmt.Errorf("wrong format date")}},
	{value: "aabb", average: FormatterAnswer{"", fmt.Errorf("not date")}},
	{value: "10.2022", average: FormatterAnswer{"", fmt.Errorf("not date")}},
}
var correctMonthYearTestsPair = []FormatterTestsPair{
	{value: "10.2022", average: FormatterAnswer{"01.10.2022", nil}},
	{value: "01.2020", average: FormatterAnswer{"01.01.2020", nil}},
	{value: "24.2022", average: FormatterAnswer{"", fmt.Errorf("month > 12")}},
	{value: "0.2022", average: FormatterAnswer{"", fmt.Errorf("month = 0")}},
}

func TestCorrectDate(t *testing.T) {
	df := DateFormatter{
		format: "02.01.2006",
	}

	for _, pair := range correctDateTestsPair {
		result, err := df.CorrectDate(pair.value)
		if err != pair.average.err {
			assert.Error(t, err)
		}
		assert.Equal(t, result, pair.average.answer, "expected: %s, but get %s", result, pair.average.answer)
	}
}

func TestCorrectMonthYear(t *testing.T) {
	df := DateFormatter{
		format:       "02.01.2006",
		budgetFormat: "01.2006",
	}

	for _, pair := range correctMonthYearTestsPair {
		result, err := df.CorrectMonthYear(pair.value)
		if err != pair.average.err {
			assert.Error(t, err)
		}
		assert.Equal(t, result, pair.average.answer, "expected: %s, but get %s", result, pair.average.answer)
	}
}
