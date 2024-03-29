package tgServer

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/EbumbaE/FinTgBot/bot/internal/model/diary"
	"github.com/EbumbaE/FinTgBot/bot/internal/model/messages"
	"github.com/EbumbaE/FinTgBot/bot/internal/model/report"
	"github.com/EbumbaE/FinTgBot/bot/internal/model/request"
	"github.com/EbumbaE/FinTgBot/bot/internal/storage"
	"github.com/EbumbaE/FinTgBot/bot/pkg/logger"
	"go.uber.org/zap"
)

func parseArguments(lineArgs string, amount int) ([]string, error) {
	args := strings.Split(lineArgs, " ")
	if len(args) != amount {
		return nil, fmt.Errorf("Error amount")
	}
	return args, nil
}

func deltaUserValuteToDefault(db storage.Storage, userID int64) (delta float64, err error) {

	userAbbValute, err := db.GetUserAbbValute(userID)
	if err != nil {
		return 1, err
	}

	userRateValute, err := db.GetRate(userAbbValute)
	if err != nil {
		return 1, err
	}

	return userRateValute.Value, nil
}

func (t *TgServer) CheckBudget(userID int64, date string, sum, delta float64) (answer string, err error) {
	budget, err := t.storage.GetMonthlyBudget(userID, date)
	if err != nil {
		budget = &diary.Budget{
			Value: 0,
		}
	}

	if budget.Value == 0 {
		return "Done", nil
	}

	dateBudget, err := t.dateFormatter.CorrectMonthYear(date)
	if err != nil {
		return err.Error(), err
	}
	timeBudget, err := t.dateFormatter.FormatDateStringToTime(dateBudget)
	if err != nil {
		return err.Error(), err
	}

	monthSum, err := report.CountMonthSumInDBCurrency(userID, t.storage, &t.dateFormatter, timeBudget)
	if err != nil {
		return err.Error(), err
	}
	budgetRate, err := t.storage.GetRate(budget.Abbreviation)
	if err != nil {
		return err.Error(), err
	}
	userAbbValute, err := t.storage.GetUserAbbValute(userID)
	if err != nil {
		return err.Error(), err
	}

	differ := budget.Value*budgetRate.Value*delta - (monthSum*delta + sum)
	if differ < 0 {
		return fmt.Sprintf("Over budget by %0.2f %s", -1*differ, userAbbValute), nil
	}

	return "Done", nil
}

func (t *TgServer) CommandSetNote(ctx context.Context, msg *messages.Message) (answer string, err error) {
	if t.Metrics != nil {
		t.Metrics.AmountActionWithNotes.Inc()
		t.Metrics.AmountCommand.Inc()
	}

	args, err := parseArguments(msg.Arguments, 3)
	if err != nil {
		answer = "error in arguments"
		return
	}

	sum, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		answer = "error in sum"
		return
	}

	date, err := t.dateFormatter.CorrectDate(args[0])
	if err != nil {
		answer = "error in date, try this format 27.09.2022"
		return
	}

	delta, err := deltaUserValuteToDefault(t.storage, msg.UserID)
	if err != nil {
		answer = err.Error()
		return
	}

	monthYear := date[3:]
	answer, err = t.CheckBudget(msg.UserID, monthYear, sum, 1.0/delta)
	if err != nil {
		return
	}

	note := diary.Note{
		Category: args[1],
		Sum:      sum * delta,
		Currency: "RUB",
	}
	if err := t.storage.AddNote(msg.UserID, date, note); err != nil {
		answer = err.Error()
		logger.Error("add note: ", zap.Error(err))
	}

	timeNote, err := t.dateFormatter.FormatDateStringToTime(date)
	if err != nil {
		logger.Error("format date string to time in set note", zap.Error(err))
		return
	}

	r := request.AddNoteInCacheRequest{
		UserID:   msg.UserID,
		TimeNote: timeNote,
		Note:     note,
	}

	if err = t.producer.SendAddNoteInCache(ctx, r); err != nil {
		logger.Error("send to add note in cache", zap.Error(err))
		return
	}

	return
}

func (t *TgServer) CommandGetStatistic(ctx context.Context, msg *messages.Message) error {
	if t.Metrics != nil {
		t.Metrics.AmountActionWithStatistic.Inc()
		t.Metrics.AmountCommand.Inc()
	}

	args, err := parseArguments(msg.Arguments, 1)
	if err != nil {
		return err
	}
	period := args[0]

	userAbbCurrency, err := t.storage.GetUserAbbValute(msg.UserID)
	if err != nil {
		return err
	}
	userCurrency, err := t.storage.GetRate(userAbbCurrency)
	if err != nil {
		return err
	}

	r := request.ReportRequest{
		UserID:       msg.UserID,
		Period:       period,
		DateFormat:   t.dateFormatter.format,
		UserCurrency: *userCurrency,
	}

	if err := t.producer.SendReportRequest(ctx, r); err != nil {
		return err
	}

	return nil
}

func (t *TgServer) CommandSetBudget(ctx context.Context, msg *messages.Message) (answer string, err error) {
	if t.Metrics != nil {
		t.Metrics.AmountActionWithBudgets.Inc()
		t.Metrics.AmountCommand.Inc()
	}

	args, err := parseArguments(msg.Arguments, 3)
	if err != nil {
		answer = "error in arguments"
		return
	}

	monthYear := args[0]
	_, err = t.dateFormatter.CorrectMonthYear(monthYear)
	if err != nil {
		answer = "error in date, try this format 09.2022"
		return
	}

	sum, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		answer = "error in sum"
		return
	}

	answer = "Done"
	budget := diary.Budget{
		Value:        sum,
		Abbreviation: args[2],
		Date:         monthYear,
	}
	err = t.storage.AddMonthlyBudget(msg.UserID, budget)
	if err != nil {
		logger.Error("add monthly budget: ", zap.Error(err))
		answer = "error in storage: set budget"
	}
	return
}

func (t *TgServer) CommandGetBudget(ctx context.Context, msg *messages.Message) (answer string, err error) {
	if t.Metrics != nil {
		t.Metrics.AmountActionWithBudgets.Inc()
		t.Metrics.AmountCommand.Inc()
	}

	args, err := parseArguments(msg.Arguments, 1)
	if err != nil {
		answer = "error in arguments"
		return
	}

	monthYear := args[0]
	_, err = t.dateFormatter.CorrectMonthYear(monthYear)
	if err != nil {
		answer = "error in date, try this format 09.2022"
		return
	}

	userAbbCurrency, err := t.storage.GetUserAbbValute(msg.UserID)
	if err != nil {
		answer = err.Error()
		logger.Error("get user currency abbreviation: ", zap.Error(err))
		return
	}
	userRate, err := t.storage.GetRate(userAbbCurrency)
	if err != nil {
		answer = err.Error()
		logger.Error("get rate error: ", zap.Error(err))
		return
	}

	answer, err = report.GetBudgetReport(msg.UserID, t.storage, &t.dateFormatter, userRate, monthYear)
	if err != nil {
		answer = err.Error()
		logger.Error("get budget report: ", zap.Error(err))
		return
	}

	return
}

func (t *TgServer) CommandStart(ctx context.Context, msg *messages.Message) (answer string, err error) {
	if t.Metrics != nil {
		t.Metrics.AmountCommand.Inc()
	}
	if err := t.storage.CheckUser(msg.UserID); err == nil {
		t.Metrics.AmountNewUsers.Inc()
	}

	answer =
		`hello, some commands:
/setNote дата категория сумма
пример: /setNote 27.09.2022 food 453.12
добавляет трату в заданный день по заданной категории, отвечает Done в случае успешной записи

/getStatistic период (week, month или year)
example: /getStatistic week
выводит статистику за заданный период, ответа на команду:
Statistic for the week:
food: 245.12
school: 85.01

/selectCurrency
дает выбор валюты для команд getStatistic, setNote, getMonthlyBudget

/setBudget дата сумма валюта
example: /setBudget 10.2022 254 USD
устанавливает бюджет на месяц

/getBudget дата
example: /getBudget 10.2022
выводит расход за месяц рабочей валюте

/start
выводит информацию о командах`

	return
}

func (t *TgServer) CommandDefault(ctx context.Context, msg *messages.Message) (answer string, err error) {
	if t.Metrics != nil {
		t.Metrics.AmountCommand.Inc()
		t.Metrics.AmountDefaultMsgAndComm.Inc()
	}

	return "Unknown command", nil
}
