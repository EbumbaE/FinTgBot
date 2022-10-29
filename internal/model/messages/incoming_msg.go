package messages

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

func (m *Model) IncomingMessage(ctx context.Context, msg Message) (err error) {
	span, nctx := opentracing.StartSpanFromContext(ctx, "incoming message")
	if span != nil {
		span.LogKV("message", msg.Command, "text", msg.Text)
		defer span.Finish()
	}

	if isCurrency := m.tgServer.IsCurrency(msg.Text); isCurrency {
		msg.Text, err = m.tgServer.MessageSetReportCurrency(nctx, &msg)
	} else {
		switch msg.Text {
		default:
			msg.Text, err = m.tgServer.MessageDefault(nctx, &msg)
		}
	}
	if err != nil {
		msg.Text = err.Error()
	}
	return m.tgClient.SendMessage(msg)
}
