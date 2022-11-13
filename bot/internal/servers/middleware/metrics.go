package middleware

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type TgMetrics struct {
	AmountNewUsers            prometheus.Counter
	AmountCommand             prometheus.Counter
	AmountMessage             prometheus.Counter
	AmountActionWithNotes     prometheus.Counter
	AmountActionWithBudgets   prometheus.Counter
	AmountActionWithCurrency  prometheus.Counter
	AmountActionWithStatistic prometheus.Counter
	AmountDefaultMsgAndComm   prometheus.Counter
}

func NewTgMetrics() *TgMetrics {
	amountNewUsers := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "ozon",
		Name:      "amount_new_users",
	})
	amountCommand := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "ozon",
		Name:      "amount_command_total",
	})
	amountMessage := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "ozon",
		Name:      "amount_messages_total",
	})
	amountActionWithNotes := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "ozon",
		Name:      "amount_add_get_notes_total",
	})
	amountActionWithBudgets := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "ozon",
		Name:      "amount_add_get_budget_total",
	})
	amountActionWithCurrency := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "ozon",
		Name:      "amount_set_report_currency_total",
	})
	amountActionWithStatistic := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "ozon",
		Name:      "amount_get_statistic_total",
	})
	amountDefaultMsgAndComm := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "ozon",
		Name:      "amount_unknown_msg_command_total",
	})

	return &TgMetrics{
		AmountNewUsers:            amountNewUsers,
		AmountCommand:             amountCommand,
		AmountMessage:             amountMessage,
		AmountActionWithNotes:     amountActionWithNotes,
		AmountActionWithBudgets:   amountActionWithBudgets,
		AmountActionWithCurrency:  amountActionWithCurrency,
		AmountActionWithStatistic: amountActionWithStatistic,
		AmountDefaultMsgAndComm:   amountDefaultMsgAndComm,
	}

}

type SenderMetrics struct {
	AmountSendMessage prometheus.Counter
}

func NewSenderMetrics() *SenderMetrics {
	amountSendMessage := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "ozon",
		Name:      "amount_senderServer_send_message",
	})
	return &SenderMetrics{
		AmountSendMessage: amountSendMessage,
	}
}
