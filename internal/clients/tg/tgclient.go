package tg

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/currancy"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
)

type Client struct {
	client *tgbotapi.BotAPI
}

func New(tgClient Config, parser *currancy.Parser) (*Client, error) {
	currencies := parser.GetAbbreviations()
	initKeyboards(currencies)

	client, err := tgbotapi.NewBotAPI(tgClient.Token)
	if err != nil {
		return nil, errors.Wrap(err, "NewBotAPI")
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) SendMessage(msg messages.Message) error {
	tgMsg := tgbotapi.NewMessage(msg.UserID, msg.Text)
	tgMsg.ReplyMarkup = msg.Keyboard

	_, err := c.client.Send(tgMsg)
	if err != nil {
		return errors.Wrap(err, "client.Send")
	}
	return nil
}

func (c *Client) SetupCurrencyKeyboard(msg *messages.Message) {
	msg.Keyboard = oneTimeCurrencyKeyboard
}

func (c *Client) ListenUpdates(msgModel *messages.Model) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.client.GetUpdatesChan(u)

	log.Println("listening for messages")

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			err := msgModel.IncomingMessage(messages.Message{
				UserID:    update.Message.From.ID,
				Command:   update.Message.Command(),
				Arguments: update.Message.CommandArguments(),
				Text:      update.Message.Text,
			})
			if err != nil {
				log.Println("error processing message:", err)
			}
		}
	}
}
