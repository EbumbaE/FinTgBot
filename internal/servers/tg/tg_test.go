package tgServer_test

import (
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
	tgConfig := tgServer.Config{
		FormatDate:       "02.01.2006",
		BudgetFormatDate: "01.2006",
	}
	tg, err := tgServer.New(storage, tgConfig)
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
	tgConfig := tgServer.Config{
		FormatDate:       "02.01.2006",
		BudgetFormatDate: "01.2006",
	}
	tg, err := tgServer.New(storage, tgConfig)
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
	tgConfig := tgServer.Config{
		FormatDate:       "02.01.2006",
		BudgetFormatDate: "01.2006",
	}
	tg, err := tgServer.New(storage, tgConfig)
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
	tgConfig := tgServer.Config{
		FormatDate:       "02.01.2006",
		BudgetFormatDate: "01.2006",
	}
	tg, err := tgServer.New(storage, tgConfig)
	assert.NoError(t, err)

	date := "15.10.2022"
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

	answer, err := tg.CommandSetNote(&msg)
	assert.Equal(t, answer, "Done")
	assert.NoError(t, err)
}

func TestCommandGetBudget(t *testing.T) {
	ctrl := gomock.NewController(t)
	storage := dbmocks.NewMockStorage(ctrl)
	tgConfig := tgServer.Config{
		FormatDate:       "02.01.2006",
		BudgetFormatDate: "01.2006",
	}
	tg, err := tgServer.New(storage, tgConfig)
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

	answer, err := tg.CommandGetBudget(&msg)
	assert.Equal(t, "Budget for the month in USD:\n12.40/40.00 USD", answer)
	assert.NoError(t, err)

}

func TestGetStatic(t *testing.T) {
	ctrl := gomock.NewController(t)
	storage := dbmocks.NewMockStorage(ctrl)
	tgConfig := tgServer.Config{
		FormatDate:       "02.01.2006",
		BudgetFormatDate: "01.2006",
	}
	tg, err := tgServer.New(storage, tgConfig)
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
	notes := []diary.Note{
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

	storage.EXPECT().GetUserAbbValute(msg.UserID).Return(userRate.Abbreviation, nil)
	storage.EXPECT().GetRate(userRate.Abbreviation).Return(&userRate, nil)

	tn := time.Now()
	beginPeriod, endPeriod := report.GetWeekPeriod(tn)
	for date := beginPeriod; date != endPeriod.AddDate(0, 0, 1); date = date.AddDate(0, 0, 1) {
		storage.EXPECT().GetNote(msg.UserID, date.Format("02.01.2006")).Return(notes, nil)
	}

	answer, err := tg.CommandGetStatistic(&msg)
	assert.NoError(t, err)
	if answer != "Statistic for the week in USD:\nfood: 3.50\nschool: 2.33\n" &&
		answer != "Statistic for the week in USD:\nschool: 2.33\nfood: 3.50\n" {
		t.Fatalf("unexpected answer: %s", answer)
	}

	msg.Arguments = "month"
	storage.EXPECT().GetUserAbbValute(msg.UserID).Return(userRate.Abbreviation, nil)
	storage.EXPECT().GetRate(userRate.Abbreviation).Return(&userRate, nil)

	beginPeriod, endPeriod = report.GetMonthPeriod(tn)
	for date := beginPeriod; date != endPeriod.AddDate(0, 0, 1); date = date.AddDate(0, 0, 1) {
		storage.EXPECT().GetNote(msg.UserID, date.Format("02.01.2006")).Return(notes, nil)
	}

	answer, err = tg.CommandGetStatistic(&msg)
	assert.NoError(t, err)
	if answer != "Statistic for the month in USD:\nfood: 15.50\nschool: 10.33\n" &&
		answer != "Statistic for the month in USD:\nschool: 10.33\nfood: 15.50\n" {
		t.Fatalf("unexpected answer: %s", answer)
	}

	msg.Arguments = "year"
	storage.EXPECT().GetUserAbbValute(msg.UserID).Return(userRate.Abbreviation, nil)
	storage.EXPECT().GetRate(userRate.Abbreviation).Return(&userRate, nil)

	beginPeriod, endPeriod = report.GetYearPeriod(tn)
	for date := beginPeriod; date != endPeriod.AddDate(0, 0, 1); date = date.AddDate(0, 0, 1) {
		storage.EXPECT().GetNote(msg.UserID, date.Format("02.01.2006")).Return(notes, nil)
	}

	answer, err = tg.CommandGetStatistic(&msg)
	assert.NoError(t, err)
	if answer != "Statistic for the year in USD:\nfood: 182.50\nschool: 121.67\n" &&
		answer != "Statistic for the month in USD:\nschool: 121.67\nfood: 182.50\n" {
		t.Fatalf("unexpected answer: %s", answer)
	}
}
