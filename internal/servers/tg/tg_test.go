package tgServer_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	dbmocks "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/mocks/storage"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/report"
	tgServer "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/servers/tg"
)

func TestOverCheckBudget(t *testing.T) {
	ctrl := gomock.NewController(t)
	storage := dbmocks.NewMockStorage(ctrl)
	cache := dbmocks.NewMockCache(ctrl)
	tgConfig := tgServer.Config{
		FormatDate:       "02.01.2006",
		BudgetFormatDate: "01.2006",
	}
	tg, err := tgServer.New(storage, cache, tgConfig)
	assert.NoError(t, err)

	var userID int64 = 123
	date := "01.10.2022"
	budget := diary.Budget{
		Value:        200,
		Abbreviation: "RUB",
		Date:         "10.2022",
	}
	budgetRate := diary.Valute{
		Abbreviation: "RUB",
		Value:        1,
	}
	note := []diary.Note{
		{
			Category: "food",
			Sum:      100,
			Currency: "RUB",
		},
	}

	storage.EXPECT().GetMonthlyBudget(userID, budget.Date).Return(&budget, nil)
	storage.EXPECT().GetRate(budget.Abbreviation).Return(&budgetRate, nil)
	storage.EXPECT().GetUserAbbValute(userID).Return("RUB", nil)

	storage.EXPECT().GetNote(userID, date).Return(note, nil)
	tn := time.Date(2022, 10, 15, 0, 0, 0, 0, time.Now().Location())
	beginPeriod, endPeriod := report.GetMonthPeriod(tn)
	for pDate := beginPeriod.AddDate(0, 0, 1); pDate != endPeriod.AddDate(0, 0, 1); pDate = pDate.AddDate(0, 0, 1) {
		storage.EXPECT().GetNote(userID, pDate.Format("02.01.2006")).Return(nil, nil)
	}

	answer, err := tg.CheckBudget(userID, budget.Date, 200, 1)
	assert.Equal(t, "Over budget by 100.00 RUB", answer)
	assert.NoError(t, err)
}

func TestNullCheckBudget(t *testing.T) {
	ctrl := gomock.NewController(t)
	storage := dbmocks.NewMockStorage(ctrl)
	cache := dbmocks.NewMockCache(ctrl)
	tgConfig := tgServer.Config{
		FormatDate:       "02.01.2006",
		BudgetFormatDate: "01.2006",
	}
	tg, err := tgServer.New(storage, cache, tgConfig)
	assert.NoError(t, err)

	var userID int64 = 123
	budget := diary.Budget{
		Value:        0,
		Abbreviation: "RUB",
		Date:         "10.2022",
	}

	storage.EXPECT().GetMonthlyBudget(userID, budget.Date).Return(&budget, nil)

	answer, err := tg.CheckBudget(userID, budget.Date, 200, 1)
	assert.Equal(t, "Done", answer)
	assert.NoError(t, err)
}

func TestDoneCheckBudget(t *testing.T) {
	ctrl := gomock.NewController(t)
	storage := dbmocks.NewMockStorage(ctrl)
	cache := dbmocks.NewMockCache(ctrl)

	tgConfig := tgServer.Config{
		FormatDate:       "02.01.2006",
		BudgetFormatDate: "01.2006",
	}
	tg, err := tgServer.New(storage, cache, tgConfig)
	assert.NoError(t, err)

	var userID int64 = 123
	budget := diary.Budget{
		Value:        1000,
		Abbreviation: "RUB",
		Date:         "10.2022",
	}
	budgetRate := diary.Valute{
		Abbreviation: "RUB",
		Value:        1,
	}
	note := []diary.Note{
		{
			Category: "food",
			Sum:      1,
			Currency: "RUB",
		},
	}

	storage.EXPECT().GetMonthlyBudget(userID, budget.Date).Return(&budget, nil)
	storage.EXPECT().GetRate(budget.Abbreviation).Return(&budgetRate, nil)
	storage.EXPECT().GetUserAbbValute(userID).Return("RUB", nil)

	tn := time.Date(2022, 10, 15, 0, 0, 0, 0, time.Now().Location())
	beginPeriod, endPeriod := report.GetMonthPeriod(tn)
	for pDate := beginPeriod; pDate != endPeriod.AddDate(0, 0, 1); pDate = pDate.AddDate(0, 0, 1) {
		storage.EXPECT().GetNote(userID, pDate.Format("02.01.2006")).Return(note, nil)
	}

	answer, err := tg.CheckBudget(userID, budget.Date, 200, 1)
	assert.Equal(t, "Done", answer)
	assert.NoError(t, err)
}

func TestSetNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	storage := dbmocks.NewMockStorage(ctrl)
	cache := dbmocks.NewMockCache(ctrl)

	tgConfig := tgServer.Config{
		FormatDate:       "02.01.2006",
		BudgetFormatDate: "01.2006",
	}
	tg, err := tgServer.New(storage, cache, tgConfig)
	assert.NoError(t, err)

	date := "15.10.2022"
	timeNote, err := time.Parse("02.01.2006", date)
	assert.NoError(t, err)

	msg := messages.Message{
		UserID:    123,
		Arguments: "15.10.2022 food 112",
		Command:   "setNote",
	}
	userRate := diary.Valute{
		Abbreviation: "USD",
		Value:        30,
	}
	budget := diary.Budget{
		Value: 0,
		Date:  "10.2022",
	}
	note := diary.Note{
		Category: "food",
		Sum:      112 * userRate.Value,
		Currency: "RUB",
	}

	storage.EXPECT().GetUserAbbValute(msg.UserID).Return(userRate.Abbreviation, nil)
	storage.EXPECT().GetRate(userRate.Abbreviation).Return(&userRate, nil)

	storage.EXPECT().GetMonthlyBudget(msg.UserID, budget.Date).Return(nil, fmt.Errorf("this user haven't budgets"))

	storage.EXPECT().AddNote(msg.UserID, date, note).Return(nil)
	cache.EXPECT().AddNoteInCacheReports(msg.UserID, timeNote, note).Return(nil)

	ctx := context.Background()
	answer, err := tg.CommandSetNote(ctx, &msg)
	assert.Equal(t, answer, "Done")
	assert.NoError(t, err)
}

func TestCommandGetBudget(t *testing.T) {
	ctrl := gomock.NewController(t)
	storage := dbmocks.NewMockStorage(ctrl)
	cache := dbmocks.NewMockCache(ctrl)

	tgConfig := tgServer.Config{
		FormatDate:       "02.01.2006",
		BudgetFormatDate: "01.2006",
	}
	tg, err := tgServer.New(storage, cache, tgConfig)
	assert.NoError(t, err)

	msg := messages.Message{
		UserID:    123,
		Arguments: "10.2022",
		Command:   "getBudget",
	}
	userRate := diary.Valute{
		Abbreviation: "USD",
		Value:        30,
	}
	budget := diary.Budget{
		Value:        1200,
		Date:         "10.2022",
		Abbreviation: "RUB",
	}
	budgetRate := diary.Valute{
		Abbreviation: "RUB",
		Value:        1,
	}
	notes := []diary.Note{
		{
			Category: "food",
			Sum:      12,
			Currency: "RUB",
		},
	}

	storage.EXPECT().GetUserAbbValute(msg.UserID).Return(userRate.Abbreviation, nil)
	storage.EXPECT().GetRate(userRate.Abbreviation).Return(&userRate, nil)

	storage.EXPECT().GetMonthlyBudget(msg.UserID, budget.Date).Return(&budget, nil)
	storage.EXPECT().GetRate(budgetRate.Abbreviation).Return(&budgetRate, nil)

	tn := time.Date(2022, 10, 15, 0, 0, 0, 0, time.Now().Location())
	beginPeriod, endPeriod := report.GetMonthPeriod(tn)
	for date := beginPeriod; date != endPeriod.AddDate(0, 0, 1); date = date.AddDate(0, 0, 1) {
		storage.EXPECT().GetNote(msg.UserID, date.Format("02.01.2006")).Return(notes, nil)
	}

	ctx := context.Background()
	answer, err := tg.CommandGetBudget(ctx, &msg)
	assert.Equal(t, "Budget for the month in USD:\n12.40/40.00 USD", answer)
	assert.NoError(t, err)

}

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
	tgConfig = tgServer.Config{
		FormatDate:       "02.01.2006",
		BudgetFormatDate: "01.2006",
	}
	userRate = diary.Valute{
		Abbreviation: "USD",
		Value:        30,
	}
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
