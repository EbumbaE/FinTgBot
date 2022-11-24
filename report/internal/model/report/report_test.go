package report_test

import (
	"fmt"
	"testing"
	"time"

	dbmocks "github.com/EbumbaE/FinTgBot/report/internal/mocks/storage"
	"github.com/EbumbaE/FinTgBot/report/internal/model/diary"
	"github.com/EbumbaE/FinTgBot/report/internal/model/report"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	notes = []diary.Note{
		{
			Category: "food",
			Sum:      15,
			Currency: "RUB",
		},
		{
			Category: "school",
			Sum:      10,
			Currency: "RUB",
		},
	}
	userRate = diary.Valute{
		Abbreviation: "USD",
		Value:        30,
	}
	UserID int64 = 123

	fReport1 string = "Statistic for the %s in USD:\nfood: %0.2f\nschool: %0.2f\n"
	fReport2 string = "Statistic for the %s in USD:\nschool: %0.2f\nfood: %0.2f\n"
)

func TestGetWeekStatic(t *testing.T) {
	ctrl := gomock.NewController(t)
	storage := dbmocks.NewMockStorage(ctrl)

	var foodSum, schoolSum float64 = 0, 0
	mapReport := report.ReportFormat{}

	beginPeriod, endPeriod := report.GetWeekPeriod(time.Now())
	for date := beginPeriod; date != endPeriod.AddDate(0, 0, 1); date = date.AddDate(0, 0, 1) {
		foodSum += notes[0].Sum
		schoolSum += notes[1].Sum
		mapReport[notes[0].Category] += notes[0].Sum
		mapReport[notes[1].Category] += notes[1].Sum
		storage.EXPECT().GetNote(UserID, date.Format("02.01.2006")).Return(notes, nil)
	}

	period := "week"
	countReport, err := report.CountStatistic(UserID, period, storage, "02.01.2006")
	assert.NoError(t, err)
	answer, err := report.FormatReportToString(&countReport, period, userRate)
	assert.NoError(t, err)

	foodSum /= userRate.Value
	schoolSum /= userRate.Value
	if answer != fmt.Sprintf(fReport1, period, foodSum, schoolSum) &&
		answer != fmt.Sprintf(fReport2, period, schoolSum, foodSum) {
		t.Fatalf("unexpected answer: %s", answer)
	}
}

func TestGetMonthStatic(t *testing.T) {
	ctrl := gomock.NewController(t)
	storage := dbmocks.NewMockStorage(ctrl)

	var foodSum, schoolSum float64 = 0, 0
	mapReport := report.ReportFormat{}

	beginPeriod, endPeriod := report.GetMonthPeriod(time.Now())
	for date := beginPeriod; date != endPeriod.AddDate(0, 0, 1); date = date.AddDate(0, 0, 1) {
		foodSum += notes[0].Sum
		schoolSum += notes[1].Sum
		mapReport[notes[0].Category] += notes[0].Sum
		mapReport[notes[1].Category] += notes[1].Sum
		storage.EXPECT().GetNote(UserID, date.Format("02.01.2006")).Return(notes, nil)
	}

	period := "month"
	countReport, err := report.CountStatistic(UserID, period, storage, "02.01.2006")
	assert.NoError(t, err)
	answer, err := report.FormatReportToString(&countReport, period, userRate)
	assert.NoError(t, err)

	foodSum /= userRate.Value
	schoolSum /= userRate.Value
	if answer != fmt.Sprintf(fReport1, period, foodSum, schoolSum) &&
		answer != fmt.Sprintf(fReport2, period, schoolSum, foodSum) {
		t.Fatalf("unexpected answer: %s", answer)
	}
}

func TestGetYearStatic(t *testing.T) {
	ctrl := gomock.NewController(t)
	storage := dbmocks.NewMockStorage(ctrl)

	var foodSum, schoolSum float64 = 0, 0
	mapReport := report.ReportFormat{}

	beginPeriod, endPeriod := report.GetYearPeriod(time.Now())
	for date := beginPeriod; date != endPeriod.AddDate(0, 0, 1); date = date.AddDate(0, 0, 1) {
		foodSum += notes[0].Sum
		schoolSum += notes[1].Sum
		mapReport[notes[0].Category] += notes[0].Sum
		mapReport[notes[1].Category] += notes[1].Sum
		storage.EXPECT().GetNote(UserID, date.Format("02.01.2006")).Return(notes, nil)
	}

	period := "year"
	countReport, err := report.CountStatistic(UserID, period, storage, "02.01.2006")
	assert.NoError(t, err)
	answer, err := report.FormatReportToString(&countReport, period, userRate)
	assert.NoError(t, err)

	foodSum /= userRate.Value
	schoolSum /= userRate.Value
	if answer != fmt.Sprintf(fReport1, period, foodSum, schoolSum) &&
		answer != fmt.Sprintf(fReport2, period, schoolSum, foodSum) {
		t.Fatalf("unexpected answer: %s", answer)
	}
}
