package report_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	dbmocks "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/mocks/storage"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/report"
)

func TestGetWeekStatic(t *testing.T) {
	ctrl := gomock.NewController(t)
	storage := dbmocks.NewMockStorage(ctrl)
	cache := dbmocks.NewMockCache(ctrl)

	tg, err := tgServer.New(storage, cache, tgConfig)
	assert.NoError(t, err)

	msg := messages.Message{
		Command:   "getStatistic",
		Arguments: "week",
		UserID:    123,
	}
	userRate := diary.Valute{
		Abbreviation: "USD",
		Value:        30,
	}
	fReport1 := "Statistic for the week in USD:\nfood: %0.2f\nschool: %0.2f\n"
	fReport2 := "Statistic for the week in USD:\nschool: %0.2f\nfood: %0.2f\n"

	var foodSum, schoolSum float64 = 0, 0
	mapReport := report.ReportFormat{}
	ctx := context.Background()

	storage.EXPECT().GetUserAbbValute(msg.UserID).Return(userRate.Abbreviation, nil)
	storage.EXPECT().GetRate(userRate.Abbreviation).Return(&userRate, nil)

	beginPeriod, endPeriod := report.GetWeekPeriod(time.Now())
	for date := beginPeriod; date != endPeriod.AddDate(0, 0, 1); date = date.AddDate(0, 0, 1) {
		foodSum += notes[0].Sum
		schoolSum += notes[1].Sum
		mapReport[notes[0].Category] += notes[0].Sum
		mapReport[notes[1].Category] += notes[1].Sum
		storage.EXPECT().GetNote(msg.UserID, date.Format("02.01.2006")).Return(notes, nil)
	}
	cache.EXPECT().GetReportFromCache(msg.UserID, msg.Arguments).Return(nil, fmt.Errorf("error for mock"))
	cache.EXPECT().AddReportInCache(msg.UserID, msg.Arguments, mapReport).Return(nil)

	answer, err := tg.CommandGetStatistic(ctx, &msg)
	assert.NoError(t, err)

	foodSum /= userRate.Value
	schoolSum /= userRate.Value
	if answer != fmt.Sprintf(fReport1, foodSum, schoolSum) &&
		answer != fmt.Sprintf(fReport2, schoolSum, foodSum) {
		t.Fatalf("unexpected answer: %s", answer)
	}
}

func TestGetMonthStatic(t *testing.T) {
	ctrl := gomock.NewController(t)
	storage := dbmocks.NewMockStorage(ctrl)
	cache := dbmocks.NewMockCache(ctrl)

	tg, err := tgServer.New(storage, cache, tgConfig)
	assert.NoError(t, err)

	msg := messages.Message{
		Command:   "getStatistic",
		Arguments: "month",
		UserID:    123,
	}
	userRate := diary.Valute{
		Abbreviation: "USD",
		Value:        30,
	}
	fReport1 := "Statistic for the month in USD:\nfood: %0.2f\nschool: %0.2f\n"
	fReport2 := "Statistic for the month in USD:\nschool: %0.2f\nfood: %0.2f\n"

	var foodSum, schoolSum float64 = 0, 0
	mapReport := report.ReportFormat{}
	ctx := context.Background()

	storage.EXPECT().GetUserAbbValute(msg.UserID).Return(userRate.Abbreviation, nil)
	storage.EXPECT().GetRate(userRate.Abbreviation).Return(&userRate, nil)

	beginPeriod, endPeriod := report.GetMonthPeriod(time.Now())
	for date := beginPeriod; date != endPeriod.AddDate(0, 0, 1); date = date.AddDate(0, 0, 1) {
		foodSum += notes[0].Sum
		schoolSum += notes[1].Sum
		mapReport[notes[0].Category] += notes[0].Sum
		mapReport[notes[1].Category] += notes[1].Sum
		storage.EXPECT().GetNote(msg.UserID, date.Format("02.01.2006")).Return(notes, nil)
	}
	cache.EXPECT().GetReportFromCache(msg.UserID, msg.Arguments).Return(nil, fmt.Errorf("error for mock"))
	cache.EXPECT().AddReportInCache(msg.UserID, msg.Arguments, mapReport).Return(nil)

	answer, err := tg.CommandGetStatistic(ctx, &msg)
	assert.NoError(t, err)

	foodSum /= userRate.Value
	schoolSum /= userRate.Value
	if answer != fmt.Sprintf(fReport1, foodSum, schoolSum) &&
		answer != fmt.Sprintf(fReport2, schoolSum, foodSum) {
		t.Fatalf("unexpected answer: %s", answer)
	}
}

func TestGetYearStatic(t *testing.T) {
	ctrl := gomock.NewController(t)
	storage := dbmocks.NewMockStorage(ctrl)
	cache := dbmocks.NewMockCache(ctrl)

	tg, err := tgServer.New(storage, cache, tgConfig)
	assert.NoError(t, err)

	msg := messages.Message{
		Command:   "getStatistic",
		Arguments: "year",
		UserID:    123,
	}
	userRate := diary.Valute{
		Abbreviation: "USD",
		Value:        30,
	}
	fReport1 := "Statistic for the year in USD:\nfood: %0.2f\nschool: %0.2f\n"
	fReport2 := "Statistic for the year in USD:\nschool: %0.2f\nfood: %0.2f\n"

	var foodSum, schoolSum float64 = 0, 0
	mapReport := report.ReportFormat{}
	ctx := context.Background()

	storage.EXPECT().GetUserAbbValute(msg.UserID).Return(userRate.Abbreviation, nil)
	storage.EXPECT().GetRate(userRate.Abbreviation).Return(&userRate, nil)

	beginPeriod, endPeriod := report.GetYearPeriod(time.Now())
	for date := beginPeriod; date != endPeriod.AddDate(0, 0, 1); date = date.AddDate(0, 0, 1) {
		foodSum += notes[0].Sum
		schoolSum += notes[1].Sum
		mapReport[notes[0].Category] += notes[0].Sum
		mapReport[notes[1].Category] += notes[1].Sum
		storage.EXPECT().GetNote(msg.UserID, date.Format("02.01.2006")).Return(notes, nil)
	}
	cache.EXPECT().GetReportFromCache(msg.UserID, msg.Arguments).Return(nil, fmt.Errorf("error for mock"))
	cache.EXPECT().AddReportInCache(msg.UserID, msg.Arguments, mapReport).Return(nil)

	answer, err := tg.CommandGetStatistic(ctx, &msg)

	foodSum /= userRate.Value
	schoolSum /= userRate.Value
	assert.NoError(t, err)
	if answer != fmt.Sprintf(fReport1, foodSum, schoolSum) &&
		answer != fmt.Sprintf(fReport2, schoolSum, foodSum) {
		t.Fatalf("unexpected answer: %s", answer)
	}
}
