package report

import "fmt"

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
