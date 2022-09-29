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

func (t *TgServer) СommandStart(msg *messages.Message) (answer string) {
	answer = "hello"
	return
}

func (t *TgServer) CommandHelp(msg *messages.Message) (answer string) {
	answer =
		`hellow, some commands:
	/setNote date category sum
	example: /setNote 27.09.2022 food 453.12 
	/getStatistic period (week, month or year)
	example: /getStatistic week`
	return
}

func (t *TgServer) СommandSetNote(msg *messages.Message) (answer string) {

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

	answer = "Done"
	note := diary.Note{
		Category: args[1],
		Sum:      sum,
	}
	t.storage.Set(msg.UserID, date, note)
	return
}

func (t *TgServer) СommandGetStatistic(msg *messages.Message) (answer string) {
	args, err := parseArguments(msg.Arguments, 1)
	if err != nil {
		answer = "error in arguments"
		return
	}

	answer, err = report.CountStatistic(msg.UserID, args[0], t.storage, &t.dateFormater)
	if err != nil {
		answer = "not done"
		fmt.Printf("[%d] %s", msg.UserID, err)
		return
	}

	return
}

func (t *TgServer) СommandDefault(msg *messages.Message) (answer string) {
	answer = "What you mean?"
	return
}
