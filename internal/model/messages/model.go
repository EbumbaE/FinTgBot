package messages

type Client interface {
	SendMessage(msg Message) error
	SetupCurrencyKeyboard(msg *Message)
}

type Messanger interface {
	IsCurrency(text string) bool
	MessageDefault(msg *Message) (answer string, err error)
	MessageSetReportCurrency(msg *Message) (answer string, err error)
}

type Commander interface {
	CommandStart(msg *Message) (answer string, err error)
	CommandHelp(msg *Message) (answer string, err error)
	CommandSetNote(msg *Message) (answer string, err error)
	CommandGetStatistic(msg *Message) (answer string, err error)
	CommandDefault(msg *Message) (answer string, err error)
	CommandSetBudget(msg *Message) (answer string, err error)
	CommandGetBudget(msg *Message) (answer string, err error)
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
