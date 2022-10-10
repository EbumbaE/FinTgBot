package messages_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	msgmocks "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/mocks/messages"
	dbmocks "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/mocks/storage"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
)

func TestOnStartCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := msgmocks.NewMockClient(ctrl)
	server := msgmocks.NewMockServer(ctrl)

	msg := messages.Message{
		Command: "start",
		UserID:  123,
	}
	sendMsg := messages.Message{
		Text:    "Hello",
		Command: "start",
		UserID:  123,
	}
	server.EXPECT().CommandStart(&msg).Return("Hello", nil)
	client.EXPECT().SendMessage(sendMsg)

	model := messages.New(client, server)
	err := model.IncomingCommand(msg)

	assert.NoError(t, err)
}

func TestOnSetNoteCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := msgmocks.NewMockClient(ctrl)
	server := msgmocks.NewMockServer(ctrl)
	storage := dbmocks.NewMockCurrencyDB(ctrl)

	msg := messages.Message{
		Command:   "setNote",
		Arguments: "29.09.2022 food 453.12",
		UserID:    123,
	}
	sendMsg := messages.Message{
		Text:      "Done",
		Command:   "setNote",
		Arguments: "29.09.2022 food 453.12",
		UserID:    123,
	}

	storage.EXPECT().SetUserAbbValute(msg.UserID, "RUB").Return(nil)
	err := storage.SetUserAbbValute(msg.UserID, "RUB")
	assert.NoError(t, err)

	server.EXPECT().CommandSetNote(&msg).Return("Done", nil)
	client.EXPECT().SendMessage(sendMsg)

	model := messages.New(client, server)
	err = model.IncomingCommand(msg)

	assert.NoError(t, err)
}

func TestOnGetStatisticCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := msgmocks.NewMockClient(ctrl)
	server := msgmocks.NewMockServer(ctrl)
	storage := dbmocks.NewMockCurrencyDB(ctrl)

	msg := messages.Message{
		Command:   "getStatistic",
		Arguments: "week",
		UserID:    123,
	}
	sendMsg := messages.Message{
		Text:      "Statistic for the week in RUB:",
		Command:   "getStatistic",
		Arguments: "week",
		UserID:    123,
	}

	storage.EXPECT().SetUserAbbValute(msg.UserID, "RUB").Return(nil)
	err := storage.SetUserAbbValute(msg.UserID, "RUB")
	assert.NoError(t, err)

	server.EXPECT().CommandGetStatistic(&msg).Return("Statistic for the week in RUB:", nil)
	client.EXPECT().SendMessage(sendMsg)
	model := messages.New(client, server)
	err = model.IncomingCommand(msg)
	assert.NoError(t, err)
}

func TestOnUnknownCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := msgmocks.NewMockClient(ctrl)
	server := msgmocks.NewMockServer(ctrl)

	msg := messages.Message{
		Command: "some text",
		UserID:  123,
	}

	sendMsg := messages.Message{
		Text:    "Unknown command",
		Command: "some text",
		UserID:  123,
	}

	server.EXPECT().CommandDefault(&msg).Return("Unknown command", nil)
	client.EXPECT().SendMessage(sendMsg)

	model := messages.New(client, server)
	err := model.IncomingCommand(msg)

	assert.NoError(t, err)
}

func TestOnRightStatic(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := msgmocks.NewMockClient(ctrl)
	server := msgmocks.NewMockServer(ctrl)
	storage := dbmocks.NewMockCurrencyDB(ctrl)

	model := messages.New(client, server)

	sendSetNoteComand := func(arguments string, userID int64) {
		msg := messages.Message{
			Command:   "setNote",
			Arguments: arguments,
			UserID:    userID,
		}
		sendMsg := messages.Message{
			Text:      "Done",
			Command:   "setNote",
			Arguments: arguments,
			UserID:    userID,
		}
		server.EXPECT().CommandSetNote(&msg).Return("Done", nil)
		client.EXPECT().SendMessage(sendMsg)

		err := model.IncomingCommand(msg)

		if err != nil {
			assert.Error(t, err)
		}
	}

	storage.EXPECT().SetUserAbbValute(int64(123), "RUB").Return(nil)
	err := storage.SetUserAbbValute(int64(123), "RUB")
	assert.NoError(t, err)

	sendSetNoteComand("10.10.2022 food 45.12", 123)
	sendSetNoteComand("09.10.2022 food 100", 123)
	sendSetNoteComand("01.09.2022 food 100", 123)
	sendSetNoteComand("11.10.2022 school 45.01", 123)
	sendSetNoteComand("12.10.2022 school 40.1", 123)

	msg := messages.Message{
		Command:   "getStatistic",
		Arguments: "week",
		UserID:    123,
	}

	sendMsg1 := messages.Message{
		Text:      "Statistic for the week in RUB: \nfood: 45.12\nschool: 85.01",
		Command:   "getStatistic",
		Arguments: "week",
		UserID:    123,
	}
	server.EXPECT().CommandGetStatistic(&msg).Return("Statistic for the week in RUB: \nfood: 45.12\nschool: 85.01", nil)
	client.EXPECT().SendMessage(sendMsg1)

	err = model.IncomingCommand(msg)
	if err != nil {
		assert.Error(t, err)
	}

	sendMsg2 := messages.Message{
		Text:      "Statistic for the week in RUB: \nfood: 145.12\nschool: 85.01",
		Command:   "getStatistic",
		Arguments: "month",
		UserID:    123,
	}
	msg.Arguments = "month"
	server.EXPECT().CommandGetStatistic(&msg).Return("Statistic for the week in RUB: \nfood: 145.12\nschool: 85.01", nil)
	client.EXPECT().SendMessage(sendMsg2)

	err = model.IncomingCommand(msg)
	if err != nil {
		assert.Error(t, err)
	}

	sendMsg3 := messages.Message{
		Text:      "Statistic for the week in RUB: \nfood: 245.12\nschool: 85.01",
		Command:   "getStatistic",
		Arguments: "year",
		UserID:    123,
	}
	msg.Arguments = "year"
	server.EXPECT().CommandGetStatistic(&msg).Return("Statistic for the week in RUB: \nfood: 245.12\nschool: 85.01", nil)
	client.EXPECT().SendMessage(sendMsg3)

	err = model.IncomingCommand(msg)
	if err != nil {
		assert.Error(t, err)
	}

	assert.NoError(t, err)
}
