package tg

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
)

type Client struct {
	client *tgbotapi.BotAPI
}

func New(tgClient Config) (*Client, error) {
	client, err := tgbotapi.NewBotAPI(tgClient.Token)
	if err != nil {
		return nil, errors.Wrap(err, "NewBotAPI")
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) SendMessage(text string, userID int64) error {
	_, err := c.client.Send(tgbotapi.NewMessage(userID, text))
	if err != nil {
		return errors.Wrap(err, "client.Send")
	}
	return nil
}

func (c *Client) ListenUpdates(msgModel *messages.Model) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.client.GetUpdatesChan(u)

	log.Println("listening for messages")

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			log.Printf("[%s] %s %s", update.Message.From.UserName, update.Message.Command(), update.Message.CommandArguments())

			err := msgModel.IncomingMessage(messages.Message{
				UserID:    update.Message.From.ID,
				Command:   update.Message.Command(),
				Arguments: update.Message.CommandArguments(),
			})
			if err != nil {
				log.Println("error processing message:", err)
			}
		}
	}
}
