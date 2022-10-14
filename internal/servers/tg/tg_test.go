package tgServer_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	dbmocks "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/mocks/storage"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
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
	beginPeriod := time.Date(2022, 10, 2, 0, 0, 0, 0, time.Now().Location())
	endPeriod := time.Date(2022, 10, 31, 0, 0, 0, 0, time.Now().Location())
	for pDate := beginPeriod; pDate != endPeriod; pDate = pDate.AddDate(0, 0, 1) {
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

	beginPeriod := time.Date(2022, 10, 1, 0, 0, 0, 0, time.Now().Location())
	endPeriod := time.Date(2022, 10, 31, 0, 0, 0, 0, time.Now().Location())
	for pDate := beginPeriod; pDate != endPeriod; pDate = pDate.AddDate(0, 0, 1) {
		storage.EXPECT().GetNote(userID, pDate.Format("02.01.2006")).Return(note, nil)
	}

	answer, err := tg.CheckBudget(userID, budget.Date, 200, 1)

	assert.Equal(t, "Done", answer)

	assert.NoError(t, err)
}
