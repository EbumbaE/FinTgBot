package report

import (
	"fmt"
	"time"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

type Storage interface {
	GetRate(abbreviation string) (*diary.Valute, error)
	GetNote(id int64, date string) ([]diary.Note, error)
	GetMonthlyBudget(userID int64, date string) (*diary.Budget, error)
}

type Formatter interface {
	FormatDateTimeToString(date time.Time) string
	FormatDateStringToTime(date string) (time.Time, error)
	CorrectMonthYear(date string) (string, error)
}

type Valute interface {
	GetAbbreviation() string
	GetValue() float64
}

type Budget interface {
	GetAbbreviation() string
	GetSum() float64
}

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

func getPeriod(period string) (beginPeriod, endPeriod time.Time, err error) {

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

func addReportHeader(period, currencyAbb string) string {
	return fmt.Sprintf("Statistic for the %s in %s:\n", period, currencyAbb)
}

func addCategory(category string, sum float64) string {
	return fmt.Sprintf("%s: %.2f\n", category, sum)
}

func CountStatistic(userID int64, period string, db Storage, formatter Formatter, currency Valute) (answer string, err error) {

	currencyAbb := currency.GetAbbreviation()
	answer = addReportHeader(period, currencyAbb)

	delta := 1.0 / currency.GetValue()

	beginPeriod, endPeriod, err := getPeriod(period)
	if err != nil {
		return "", err
	}

	totalCategorySum := map[string]float64{}
	for date := beginPeriod; date != endPeriod.AddDate(0, 0, 1); date = date.AddDate(0, 0, 1) {
		notes, err := db.GetNote(userID, formatter.FormatDateTimeToString(date))
		if err != nil {
			return "Error in storage: get note", err
		}
		for _, note := range notes {
			totalCategorySum[note.Category] += note.Sum
		}
	}

	for category, sum := range totalCategorySum {
		answer += addCategory(category, sum*delta)
	}

	return answer, nil
}

func addBudgetHeader(period, currencyAbb string) string {
	return fmt.Sprintf("Budget for the %s in %s:\n", period, currencyAbb)
}

func addBudget(totalSum, budgetSum float64, currencyAbb string) string {
	return fmt.Sprintf("%.2f/%.2f %s", totalSum, budgetSum, currencyAbb)
}

func CountMonthSumInDBCurrency(userID int64, db Storage, formatter Formatter, timeBudget time.Time) (float64, error) {

	beginPeriod, endPeriod := GetMonthPeriod(timeBudget)

	var totalSum float64 = 0
	for date := beginPeriod; date != endPeriod.AddDate(0, 0, 1); date = date.AddDate(0, 0, 1) {
		notes, err := db.GetNote(userID, formatter.FormatDateTimeToString(date))
		if err != nil {
			return 0, err
		}
		for _, note := range notes {
			totalSum += note.Sum
		}
	}
	return totalSum, nil
}

func GetBudgetReport(userID int64, db Storage, formatter Formatter, userRate Valute, monthYear string) (answer string, err error) {

	answer = addBudgetHeader("month", userRate.GetAbbreviation())

	date, err := formatter.CorrectMonthYear(monthYear)
	if err != nil {
		return err.Error(), err
	}
	timeBudget, err := formatter.FormatDateStringToTime(date)
	if err != nil {
		return err.Error(), err
	}

	budget, err := db.GetMonthlyBudget(userID, monthYear)
	if err != nil {
		return err.Error(), err
	}

	budgetRate, err := db.GetRate(budget.GetAbbreviation())
	if err != nil {
		return err.Error(), err
	}

	totalSum, err := CountMonthSumInDBCurrency(userID, db, formatter, timeBudget)
	if err != nil {
		return err.Error(), err
	}

	delta := budgetRate.GetValue() / userRate.GetValue()
	answer += addBudget(totalSum/userRate.GetValue(), budget.GetSum()*delta, userRate.GetAbbreviation())

	return answer, err
}
