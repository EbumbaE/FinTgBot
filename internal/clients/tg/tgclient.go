package tg

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
)

type Parser interface {
	GetAbbreviations() []string
}

type Client struct {
	client *tgbotapi.BotAPI
}

func New(tgClient Config, parser Parser) (*Client, error) {
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

func (c *Client) ListenUpdates(ctx context.Context, msgModel *messages.Model) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.client.GetUpdatesChan(u)

	log.Println("listening for messages")

	go func() {
		for {
			select {
			case update := <-updates:
				if update.Message != nil {

					log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

					if update.Message.IsCommand() {
						err := msgModel.IncomingCommand(messages.Message{
							UserID:    update.Message.From.ID,
							Command:   update.Message.Command(),
							Arguments: update.Message.CommandArguments(),
						})
						if err != nil {
							log.Println("error in incomming command:", err)
						}
					} else {
						err := msgModel.IncomingMessage(messages.Message{
							UserID: update.Message.From.ID,
							Text:   update.Message.Text,
						})
						if err != nil {
							log.Println("error in incomming message:", err)
						}
					}
				}
			case <-ctx.Done():
				defer updates.Clear()
				defer log.Println("listening messages is off")
				return
			}
		}
	}()
}
