package messages

import "context"

type Client interface {
	SendMessage(msg Message) error
	SetupCurrencyKeyboard(msg *Message)
}

type Messanger interface {
	IsCurrency(string) bool
	MessageDefault(context.Context, *Message) (string, error)
	MessageSetReportCurrency(context.Context, *Message) (string, error)
}

type Commander interface {
	CommandStart(context.Context, *Message) (string, error)
	CommandSetNote(context.Context, *Message) (string, error)
	CommandGetStatistic(context.Context, *Message) error
	CommandDefault(context.Context, *Message) (string, error)
	CommandSetBudget(context.Context, *Message) (string, error)
	CommandGetBudget(context.Context, *Message) (string, error)
}

type Server interface {
	Messanger
	Commander
}

type Model struct {
	tgClient Client
	tgServer Server
}

type Message struct {
	Text      string
	Arguments string
	UserID    int64
	Command   string
	Keyboard  any
}

func New(tgClient Client, tgServer Server) *Model {
	return &Model{
		tgClient: tgClient,
		tgServer: tgServer,
	}
}
