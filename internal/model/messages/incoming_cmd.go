package messages

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

func (m *Model) IncomingCommand(ctx context.Context, msg Message) (err error) {
	span, nctx := opentracing.StartSpanFromContext(ctx, "incoming command")
	if span != nil {
		span.LogKV("command", msg.Command)
		defer span.Finish()
	}

	switch msg.Command {
	case "start":
		msg.Text, err = m.tgServer.CommandStart(nctx, &msg)
	case "help":
		msg.Text, err = m.tgServer.CommandHelp(nctx, &msg)
	case "setNote":
		if span != nil {
			span.LogKV("set note", msg.Command, "args", msg.Arguments)
		}
		msg.Text, err = m.tgServer.CommandSetNote(nctx, &msg)
	case "getStatistic":
		if span != nil {
			span.LogKV("get statistic", msg.Command, "args", msg.Arguments)
		}
		msg.Text, err = m.tgServer.CommandGetStatistic(nctx, &msg)
	case "setBudget":
		if span != nil {
			span.LogKV("set budget", msg.Command, "args", msg.Arguments)
		}
		msg.Text, err = m.tgServer.CommandSetBudget(nctx, &msg)
	case "getBudget":
		if span != nil {
			span.LogKV("get budget", msg.Command, "args", msg.Arguments)
		}
		msg.Text, err = m.tgServer.CommandGetBudget(nctx, &msg)
	case "selectCurrency":
		m.tgClient.SetupCurrencyKeyboard(&msg)
		msg.Text = "Setup value:"
	default:
		msg.Text, err = m.tgServer.CommandDefault(nctx, &msg)
	}

	if err != nil {
		msg.Text = err.Error()
	}
	return m.tgClient.SendMessage(msg)
}
