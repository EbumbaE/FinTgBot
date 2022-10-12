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

	date, err := t.correctDate(args[0])
	if err != nil {
		answer = "error in date, try this format 27.09.2022"
		return
	}

	delta, err := deltaUserValuteToDefault(t.storage, msg.UserID)
	if err != nil {
		answer = err.Error()
		return
	}

	answer = "Done"
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

	answer, err = report.CountStatistic(msg.UserID, period, t.storage, &t.dateFormater, userRateCurrency)
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
		`hellow, some commands:
	/setNote date category sum
	example: /setNote 27.09.2022 food 453.12 
	/getStatistic period (week, month or year)
	example: /getStatistic week
	/selectCurrency`
	return
}

func (t *TgServer) CommandDefault(msg *messages.Message) (answer string, err error) {
	return "Unknown command", nil
}
