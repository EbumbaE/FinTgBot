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
