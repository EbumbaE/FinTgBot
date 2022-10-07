package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var oneTimeCurrencyKeyboard = tgbotapi.ReplyKeyboardMarkup{}

func intitOneTimeCurrencyKeyboard(currency []string) {
	rowCount := -1
	lenRow := 3
	lenRowCount := -1
	rows := [][]tgbotapi.KeyboardButton{}

	for _, abb := range currency {
		lenRowCount = (lenRowCount + 1) % lenRow

		if lenRowCount == 0 {
			rowCount++
			rows = append(rows, []tgbotapi.KeyboardButton{})
		}
		rows[rowCount] = append(rows[rowCount], tgbotapi.NewKeyboardButton(abb))
	}

	oneTimeCurrencyKeyboard = tgbotapi.NewOneTimeReplyKeyboard(rows...)
}

func initKeyboards(currency []string) {
	intitOneTimeCurrencyKeyboard(currency)
}
