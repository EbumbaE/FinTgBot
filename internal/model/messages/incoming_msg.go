package messages

type Client interface {
	SendMessage(msg Message) error
	SetupCurrencyKeyboard(msg *Message)
}

type Commander interface {
	CommandStart(msg *Message) (answer string, err error)
	CommandHelp(msg *Message) (answer string, err error)
	CommandSetNote(msg *Message) (answer string, err error)
	CommandGetStatistic(msg *Message) (answer string, err error)
	CommandDefault(msg *Message) (answer string, err error)
}

type Model struct {
	tgClient Client
	tgServer Commander
}

func New(tgClient Client, tgServer Commander) *Model {
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
	Keyboard  any
}

func (m *Model) IncomingMessage(msg Message) error {

	var err error = nil
	switch msg.Command {
	case "start":
		msg.Text, err = m.tgServer.CommandStart(&msg)
	case "help":
		msg.Text, err = m.tgServer.CommandHelp(&msg)
	case "setNote":
		msg.Text, err = m.tgServer.CommandSetNote(&msg)
	case "getStatistic":
		msg.Text, err = m.tgServer.CommandGetStatistic(&msg)
	case "selectCurrency":
		m.tgClient.SetupCurrencyKeyboard(&msg)
		msg.Text = "Setup value"
	default:
		msg.Text, err = m.tgServer.CommandDefault(&msg)
	}

	if err != nil {
		msg.Text = err.Error()
	}
	return m.tgClient.SendMessage(msg)
}
