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
	sender := mocks.NewMockMessageSender(ctrl)
	commander := mocks.NewMockCommander(ctrl)
	msg := messages.Message{
		Command: "start",
		UserID:  123,
	}
	commander.EXPECT().СommandStart(&msg).Return("hello")
	sender.EXPECT().SendMessage("hello", int64(123))

	model := messages.New(sender, commander)
	err := model.IncomingMessage(msg)

	assert.NoError(t, err)
}

func Test_OnSetNoteCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	commander := mocks.NewMockCommander(ctrl)
	msg := messages.Message{
		Command:   "setNote",
		Arguments: "29.09.2022 food 453.12",
		UserID:    123,
	}
	commander.EXPECT().СommandSetNote(&msg).Return("Done")
	sender.EXPECT().SendMessage("Done", int64(123))

	model := messages.New(sender, commander)
	err := model.IncomingMessage(msg)

	assert.NoError(t, err)
}

func Test_OnGetStatisticCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	commander := mocks.NewMockCommander(ctrl)
	msg := messages.Message{
		Command:   "getStatistic",
		Arguments: "week",
		UserID:    123,
	}
	commander.EXPECT().СommandGetStatistic(&msg).Return("Statistic for the week")
	sender.EXPECT().SendMessage("Statistic for the week", int64(123))

	model := messages.New(sender, commander)
	err := model.IncomingMessage(msg)

	assert.NoError(t, err)
}

func Test_OnUnknownCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	commander := mocks.NewMockCommander(ctrl)

	msg := messages.Message{
		Text:   "some text",
		UserID: 123,
	}

	commander.EXPECT().СommandDefault(&msg).Return("What you mean?")
	sender.EXPECT().SendMessage("What you mean?", int64(123))

	model := messages.New(sender, commander)
	err := model.IncomingMessage(msg)

	assert.NoError(t, err)
}

func Test_OnRightStatic(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	commander := mocks.NewMockCommander(ctrl)
	model := messages.New(sender, commander)

	sendSetNoteComand := func(arguments string, userID int64) {
		msg := messages.Message{
			Command:   "setNote",
			Arguments: arguments,
			UserID:    userID,
		}
		commander.EXPECT().СommandSetNote(&msg).Return("Done")
		sender.EXPECT().SendMessage("Done", int64(123))

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
	commander.EXPECT().СommandGetStatistic(&msg).Return("Statistic for the week: \nfood: 45.12\nschool: 85.01")
	sender.EXPECT().SendMessage("Statistic for the week: \nfood: 45.12\nschool: 85.01", int64(123))

	err := model.IncomingMessage(msg)
	if err != nil {
		assert.Error(t, err)
	}

	msg.Arguments = "month"
	commander.EXPECT().СommandGetStatistic(&msg).Return("Statistic for the week: \nfood: 145.12\nschool: 85.01")
	sender.EXPECT().SendMessage("Statistic for the week: \nfood: 145.12\nschool: 85.01", int64(123))

	err = model.IncomingMessage(msg)
	if err != nil {
		assert.Error(t, err)
	}

	msg.Arguments = "year"
	commander.EXPECT().СommandGetStatistic(&msg).Return("Statistic for the week: \nfood: 245.12\nschool: 85.01")
	sender.EXPECT().SendMessage("Statistic for the week: \nfood: 245.12\nschool: 85.01", int64(123))

	err = model.IncomingMessage(msg)
	if err != nil {
		assert.Error(t, err)
	}

	assert.NoError(t, err)
}
