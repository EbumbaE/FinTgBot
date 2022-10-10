package messages_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mocks "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/mocks/messages"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
)

func TestMessageDefault(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := mocks.NewMockClient(ctrl)
	server := mocks.NewMockServer(ctrl)
	model := messages.New(client, server)

	msg := messages.Message{
		Text:   "Random text",
		UserID: 123,
	}

	sendMsg := messages.Message{
		Text:   "What you mean?",
		UserID: 123,
	}

	server.EXPECT().IsCurrency(msg.Text).Return(false)
	server.EXPECT().MessageDefault(&msg).Return("What you mean?", nil)
	client.EXPECT().SendMessage(sendMsg)

	err := model.IncomingMessage(msg)
	assert.NoError(t, err)
}

func TestIsCurrencyAndSetCurrency(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := mocks.NewMockClient(ctrl)
	server := mocks.NewMockServer(ctrl)
	model := messages.New(client, server)

	rightCurr := func(currency string) {
		msg := messages.Message{
			Text:   currency,
			UserID: 123,
		}
		sendMsg := messages.Message{
			Text:   "Done",
			UserID: 123,
		}

		server.EXPECT().IsCurrency(msg.Text).Return(true)
		server.EXPECT().MessageSetCurrency(&msg).Return("Done", nil)
		client.EXPECT().SendMessage(sendMsg)

		err := model.IncomingMessage(msg)
		assert.NoError(t, err)
	}

	notCurr := func(currency string) {
		msg := messages.Message{
			Text:   currency,
			UserID: 123,
		}
		sendMsg := messages.Message{
			Text:   "What you mean?",
			UserID: 123,
		}

		server.EXPECT().IsCurrency(msg.Text).Return(false)
		server.EXPECT().MessageDefault(&msg).Return(sendMsg.Text, nil)
		client.EXPECT().SendMessage(sendMsg)

		err := model.IncomingMessage(msg)
		assert.NoError(t, err)
	}

	rightCurr("USD")
	rightCurr("RUB")
	notCurr("RND")
}
