package tgServer

import (
	"context"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
)

func (t *TgServer) MessageSetReportCurrency(ctx context.Context, msg *messages.Message) (answer string, err error) {
	if t.Metrics != nil {
		t.Metrics.AmountActionWithCurrency.Inc()
		t.Metrics.AmountMessage.Inc()
	}

	valute, err := t.storage.GetRate(msg.Text)
	if err != nil {
		answer = err.Error()
		return
	}

	err = t.storage.SetUserAbbValute(msg.UserID, valute.Abbreviation)
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

func (t *TgServer) MessageDefault(ctx context.Context, msg *messages.Message) (answer string, err error) {
	if t.Metrics != nil {
		t.Metrics.AmountMessage.Inc()
		t.Metrics.AmountDefaultMsgAndComm.Inc()
	}

	return "What you mean?", nil
}
