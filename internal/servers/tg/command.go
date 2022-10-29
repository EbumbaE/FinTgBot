package tgServer

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/report"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/storage"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/logger"
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
	t.Metrics.AmountActionWithNotes.Inc()
	t.Metrics.AmountCommand.Inc()

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
	err = t.storage.AddNote(msg.UserID, date, note)
	if err != nil {
		answer = err.Error()
		logger.Error("add note: ", zap.Error(err))
	}
	return
}

func (t *TgServer) CommandGetStatistic(ctx context.Context, msg *messages.Message) (answer string, err error) {
	t.Metrics.AmountActionWithStatistic.Inc()
	t.Metrics.AmountCommand.Inc()

	args, err := parseArguments(msg.Arguments, 1)
	if err != nil {
		answer = err.Error()
		return
	}
	period := args[0]

	userAbbCurrency, err := t.storage.GetUserAbbValute(msg.UserID)
	if err != nil {
		answer = err.Error()
		return
	}
	userRateCurrency, err := t.storage.GetRate(userAbbCurrency)
	if err != nil {
		answer = err.Error()
		return
	}

	answer, err = report.CountStatistic(msg.UserID, period, t.storage, &t.dateFormatter, userRateCurrency)
	if err != nil {
		answer = "not done"
		logger.Error("count statistic: ", zap.Error(err))
		return
	}

	return
}

func (t *TgServer) CommandSetBudget(ctx context.Context, msg *messages.Message) (answer string, err error) {
	t.Metrics.AmountActionWithBudgets.Inc()
	t.Metrics.AmountCommand.Inc()

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
	t.Metrics.AmountActionWithBudgets.Inc()
	t.Metrics.AmountCommand.Inc()

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
	t.Metrics.AmountCommand.Inc()
	return "Hello", nil
}

func (t *TgServer) CommandHelp(ctx context.Context, msg *messages.Message) (answer string, err error) {
	t.Metrics.AmountCommand.Inc()

	answer =
		`hello, some commands:
	/setNote date category sum
	example: /setNote 27.09.2022 food 453.12 
	
	/getStatistic period (week, month or year)
	example: /getStatistic week
	
	/setMonthlyBudget sum currency
	example: /setMonthlyBudget 245 USD
	
	/getMonthlyBudget month.year
	example: /getMonthlyBudget 09.2022
	
	/selectReportCurrency
	`
	return
}

func (t *TgServer) CommandDefault(ctx context.Context, msg *messages.Message) (answer string, err error) {
	t.Metrics.AmountCommand.Inc()
	t.Metrics.AmountDefaultMsgAndComm.Inc()
	return "Unknown command", nil
}
