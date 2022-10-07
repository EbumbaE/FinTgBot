package messages_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mocks "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/mocks/messages"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
)

func Test_OnStartCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := mocks.NewMockClient(ctrl)
	server := mocks.NewMockServer(ctrl)
	msg := messages.Message{
		Command: "start",
		UserID:  123,
	}
	sendMsg := messages.Message{
		Text:   "hello",
		UserID: 123,
	}
	commander.EXPECT().CommandStart(&msg).Return("hello", nil)
	client.EXPECT().SendMessage(sendMsg)

	model := messages.New(client, commander)
	err := model.IncomingMessage(msg)

	assert.NoError(t, err)
}

func Test_OnSetNoteCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockClient(ctrl)
	commander := mocks.NewMockCommander(ctrl)
	msg := messages.Message{
		Command:   "setNote",
		Arguments: "29.09.2022 food 453.12",
		UserID:    123,
	}
	sendMsg := messages.Message{
		Text:   "Done",
		UserID: 123,
	}
	commander.EXPECT().CommandSetNote(&msg).Return("Done", nil)
	sender.EXPECT().SendMessage(sendMsg)

	model := messages.New(sender, commander)
	err := model.IncomingMessage(msg)

	assert.NoError(t, err)
}

func Test_OnGetStatisticCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockClient(ctrl)
	commander := mocks.NewMockCommander(ctrl)
	msg := messages.Message{
		Command:   "getStatistic",
		Arguments: "week",
		UserID:    123,
	}
	sendMsg := messages.Message{
		Text:   "Statistic for the week",
		UserID: 123,
	}
	commander.EXPECT().CommandGetStatistic(&msg).Return("Statistic for the week", nil)
	sender.EXPECT().SendMessage(sendMsg)

	model := messages.New(sender, commander)
	err := model.IncomingMessage(msg)

	assert.NoError(t, err)
}

func Test_OnUnknownCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockClient(ctrl)
	commander := mocks.NewMockCommander(ctrl)

	msg := messages.Message{
		Text:   "some text",
		UserID: 123,
	}

	sendMsg := messages.Message{
		Text:   "What you mean?",
		UserID: 123,
	}

	commander.EXPECT().CommandDefault(&msg).Return("What you mean?", nil)
	sender.EXPECT().SendMessage(sendMsg)

	model := messages.New(sender, commander)
	err := model.IncomingMessage(msg)

	assert.NoError(t, err)
}

func Test_OnRightStatic(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockClient(ctrl)
	commander := mocks.NewMockCommander(ctrl)
	model := messages.New(sender, commander)

	sendSetNoteComand := func(arguments string, userID int64) {
		msg := messages.Message{
			Command:   "setNote",
			Arguments: arguments,
			UserID:    userID,
		}
		sendMsg := messages.Message{
			Text:   "Done",
			UserID: userID,
		}
		commander.EXPECT().CommandSetNote(&msg).Return("Done", nil)
		sender.EXPECT().SendMessage(sendMsg)

		err := model.IncomingMessage(msg)

		if err != nil {
			assert.Error(t, err)
		}
	}

	sendSetNoteComand("29.09.2022 food 45.12", 123)
	sendSetNoteComand("01.09.2022 food 100", 123)
	sendSetNoteComand("01.08.2022 food 100", 123)
	sendSetNoteComand("28.09.2022 school 45.01", 123)
	sendSetNoteComand("26.09.2022 school 40.1", 123)

	msg := messages.Message{
		Command:   "getStatistic",
		Arguments: "week",
		UserID:    123,
	}

	sendMsg1 := messages.Message{
		Text:   "Statistic for the week: \nfood: 45.12\nschool: 85.01",
		UserID: 123,
	}
	commander.EXPECT().CommandGetStatistic(&msg).Return("Statistic for the week: \nfood: 45.12\nschool: 85.01", nil)
	sender.EXPECT().SendMessage(sendMsg1)

	err := model.IncomingMessage(msg)
	if err != nil {
		assert.Error(t, err)
	}

	sendMsg2 := messages.Message{
		Text:   "Statistic for the week: \nfood: 145.12\nschool: 85.01",
		UserID: 123,
	}
	msg.Arguments = "month"
	commander.EXPECT().CommandGetStatistic(&msg).Return("Statistic for the week: \nfood: 145.12\nschool: 85.01", nil)
	sender.EXPECT().SendMessage(sendMsg2)

	err = model.IncomingMessage(msg)
	if err != nil {
		assert.Error(t, err)
	}

	sendMsg3 := messages.Message{
		Text:   "Statistic for the week: \nfood: 245.12\nschool: 85.01",
		UserID: 123,
	}
	msg.Arguments = "year"
	commander.EXPECT().CommandGetStatistic(&msg).Return("Statistic for the week: \nfood: 245.12\nschool: 85.01", nil)
	sender.EXPECT().SendMessage(sendMsg3)

	err = model.IncomingMessage(msg)
	if err != nil {
		assert.Error(t, err)
	}

	assert.NoError(t, err)
}
