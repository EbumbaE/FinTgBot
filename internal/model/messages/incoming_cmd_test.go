package messages_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	msgmocks "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/mocks/messages"
	dbmocks "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/mocks/storage"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
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
	storage := dbmocks.NewMockStorage(ctrl)

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
	storage.SetUserAbbValute(msg.UserID, "RUB")

	server.EXPECT().CommandSetNote(&msg).Return("Done", nil)
	client.EXPECT().SendMessage(sendMsg)

	model := messages.New(client, server)
	err := model.IncomingCommand(msg)

	assert.NoError(t, err)
}

func TestOnOverBudgetSetNoteCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := msgmocks.NewMockClient(ctrl)
	server := msgmocks.NewMockServer(ctrl)
	storage := dbmocks.NewMockStorage(ctrl)

	msg := messages.Message{
		Command:   "setNote",
		Arguments: "29.09.2022 food 453.12",
		UserID:    123,
	}
	sendMsg := messages.Message{
		Text:      "Over budget by 0.12 RUB",
		Command:   "setNote",
		Arguments: "29.09.2022 food 453.12",
		UserID:    123,
	}

	monthlyBudget := diary.Budget{
		Value:        453,
		Abbreviation: "RUB",
		Date:         "09.2022",
	}

	storage.EXPECT().AddMonthlyBudget(msg.UserID, monthlyBudget).Return(nil)
	storage.AddMonthlyBudget(msg.UserID, monthlyBudget)

	server.EXPECT().CommandSetNote(&msg).Return(sendMsg.Text, nil)
	client.EXPECT().SendMessage(sendMsg)

	model := messages.New(client, server)
	err := model.IncomingCommand(msg)

	assert.NoError(t, err)
}

func TestOnGetStatisticCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := msgmocks.NewMockClient(ctrl)
	server := msgmocks.NewMockServer(ctrl)
	storage := dbmocks.NewMockStorage(ctrl)

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
