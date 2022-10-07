package messages

type Client interface {
	SendMessage(msg Message) error
	SetupCurrencyKeyboard(msg *Message)
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
