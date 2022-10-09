package tgServer

import (
	"fmt"
	"strconv"
	"strings"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/report"
)

func parseArguments(lineArgs string, amount int) ([]string, error) {
	args := strings.Split(lineArgs, " ")
	if len(args) != amount {
		return nil, fmt.Errorf("Error amount")
	}
	return args, nil
}

func (t *TgServer) CommandStart(msg *messages.Message) (answer string, err error) {
	return "hello", nil
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

	userValute, err := t.storage.GetUserValute(msg.UserID)
	if err != nil {
		answer = err.Error()
		return
	}

	answer = "Done"
	note := diary.Note{
		Category: args[1],
		Sum:      sum,
		Valute:   userValute,
	}
	err = t.storage.SetNote(msg.UserID, date, note)
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

	valute, err := t.storage.GetUserValute(msg.UserID)
	if err != nil {
		answer = err.Error()
		return
	}

	answer, err = report.CountStatistic(msg.UserID, period, t.storage, &t.dateFormater, valute)
	if err != nil {
		answer = "not done"
		fmt.Printf("[%d] %s", msg.UserID, err)
		return
	}

	return
}

func (t *TgServer) CommandDefault(msg *messages.Message) (answer string, err error) {
	return "Unknown command", nil
}
