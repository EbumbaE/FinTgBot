package messages

type Commander interface {
	CommandStart(msg *Message) (answer string, err error)
	CommandHelp(msg *Message) (answer string, err error)
	CommandSetNote(msg *Message) (answer string, err error)
	CommandGetStatistic(msg *Message) (answer string, err error)
	CommandDefault(msg *Message) (answer string, err error)
}

func (m *Model) IncomingCommand(msg Message) error {
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
		msg.Text = "Setup value:"
	default:
		msg.Text, err = m.tgServer.CommandDefault(&msg)
	}

	if err != nil {
		msg.Text = err.Error()
	}
	return m.tgClient.SendMessage(msg)
}
