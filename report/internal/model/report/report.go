package report

import (
	"fmt"
	"time"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
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

type Budget interface {
	GetAbbreviation() string
	GetSum() float64
}

type Note interface {
	GetCategory() string
	GetSum() float64
}

type Storage interface {
	GetRate(abbreviation string) (Valute, error)
	GetNote(id int64, date string) ([]Note, error)
	GetMonthlyBudget(userID int64, date string) (Budget, error)
}

type ReportCache interface {
	AddReportInCache(userID int64, period string, addedReport ReportFormat) (err error)
	GetReportFromCache(userID int64, period string) (getReport ReportFormat, err error)
	AddNoteInCacheReports(userID int64, date time.Time, note diary.Note) error
}

type ReportFormat map[string]float64

type Report struct {
	cache   ReportCache
	storage Storage
}

func New() {

}

func addReportHeader(period, currencyAbb string) string {
	return fmt.Sprintf("Statistic for the %s in %s:\n", period, currencyAbb)
}

func addCategory(category string, sum float64) string {
	return fmt.Sprintf("%s: %.2f\n", category, sum)
}

func FormatReportToString(report *ReportFormat, period string, convertCurrency Valute) (answer string, err error) {
	currencyAbb := convertCurrency.GetAbbreviation()
	answer = addReportHeader(period, currencyAbb)

	delta := 1.0 / convertCurrency.GetValue()

	for category, sum := range *report {
		answer += addCategory(category, sum*delta)
	}

	return
}

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
			totalSum += note.GetSum()
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
