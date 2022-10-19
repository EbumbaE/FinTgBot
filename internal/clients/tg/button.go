package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Keyboards struct {
	oneTimeCurrencyKeyboard tgbotapi.ReplyKeyboardMarkup
}

func (k *Keyboards) GetCurrencyKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return k.oneTimeCurrencyKeyboard
}

func newOneTimeCurrencyKeyboard(currency []string) tgbotapi.ReplyKeyboardMarkup {
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

	return tgbotapi.NewOneTimeReplyKeyboard(rows...)
}

func NewKeyboards(currency []string) *Keyboards {
	return &Keyboards{
		oneTimeCurrencyKeyboard: newOneTimeCurrencyKeyboard(currency),
	}
}
