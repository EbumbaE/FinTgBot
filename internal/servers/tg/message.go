package tgServer

import "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"

func (t *TgServer) MessageSetCurrency(msg *messages.Message) (answer string, err error) {

	valute, err := t.storage.GetRate(msg.Text)
	if err != nil {
		answer = err.Error()
		return
	}

	err = t.storage.AddUserAbbValute(msg.UserID, valute.Abbreviation)
	if err != nil {
		return err.Error(), err
	}
	return "Done", nil
}

func (t *TgServer) IsCurrency(text string) bool {
	_, err := t.storage.GetRate(text)
	if err != nil {
		return false
	}
	return true
}

func (t *TgServer) MessageDefault(msg *messages.Message) (answer string, err error) {
	return "What you mean?", nil
}
