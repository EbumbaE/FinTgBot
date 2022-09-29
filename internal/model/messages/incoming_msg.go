package messages

import "fmt"

type MessageSender interface {
	SendMessage(text string, userID int64) error
}

type Commander interface {
	СommandStart(msg *Message) (answer string)
	CommandHelp(msg *Message) (answer string)
	СommandSetNote(msg *Message) (answer string)
	СommandGetStatistic(msg *Message) (answer string)
	СommandDefault(msg *Message) (answer string)
}

type Model struct {
	tgClient MessageSender
	tgServer Commander
}

func New(tgClient MessageSender, tgServer Commander) *Model {
	return &Model{
		tgClient: tgClient,
		tgServer: tgServer,
	}
}

type Message struct {
	Text      string
	Arguments string
	UserID    int64
	Command   string
}

func (m *Model) IncomingMessage(msg Message) error {

	switch msg.Command {
	case "start":
		msg.Text = m.tgServer.СommandStart(&msg)
	case "help":
		msg.Text = m.tgServer.CommandHelp(&msg)
	case "setNote":
		msg.Text = m.tgServer.СommandSetNote(&msg)
	case "getStatistic":
		msg.Text = m.tgServer.СommandGetStatistic(&msg)
	default:
		msg.Text = m.tgServer.СommandDefault(&msg)
	}

	if msg.Text != "" {
		return m.tgClient.SendMessage(msg.Text, msg.UserID)
	}
	return fmt.Errorf("null msg.Text")
}
