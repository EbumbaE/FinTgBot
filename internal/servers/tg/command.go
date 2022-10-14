package tgServer

import (
	"fmt"
	"strconv"
	"strings"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/report"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/storage"
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

func (t *TgServer) checkBudget(userID int64, date string, sum, delta float64) (answer string, err error) {
	budget, err := t.storage.GetMonthlyBudget(userID, date)
	if err != nil {
		budget = &diary.Budget{
			Value: 0,
		}
	}

	if budget.Value == 0 {
		return "Done", nil
	} else {

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
		if differ >= 0 {
			answer = "Done"
		} else {
			answer = fmt.Sprintf("Over budget by %0.2f %s", -1*differ, userAbbValute)
		}
	}
	return answer, err
}

func (t *TgServer) CommandSetNote(msg *messages.Message) (answer string, err error) {

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
	answer, err = t.checkBudget(msg.UserID, monthYear, sum, 1.0/delta)
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
		answer = "error in storage: set note"
	}
	return
}

func (t *TgServer) CommandGetStatistic(msg *messages.Message) (answer string, err error) {
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
		fmt.Printf("[%d] %s", msg.UserID, err)
		return
	}

	return
}

func (t *TgServer) CommandSetBudget(msg *messages.Message) (answer string, err error) {
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
		answer = "error in storage: set budget"
	}
	return
}

func (t *TgServer) CommandGetBudget(msg *messages.Message) (answer string, err error) {
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
		return
	}
	userRate, err := t.storage.GetRate(userAbbCurrency)
	if err != nil {
		answer = err.Error()
		return
	}

	answer, err = report.GetBudgetReport(msg.UserID, t.storage, &t.dateFormatter, userRate, monthYear)
	if err != nil {
		answer = "not done"
		fmt.Printf("[%d] %s", msg.UserID, err)
		return
	}

	return
}

func (t *TgServer) CommandStart(msg *messages.Message) (answer string, err error) {
	return "Hello", nil
}

func (t *TgServer) CommandHelp(msg *messages.Message) (answer string, err error) {
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

func (t *TgServer) CommandDefault(msg *messages.Message) (answer string, err error) {
	return "Unknown command", nil
}
