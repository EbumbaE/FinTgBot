package messages

func (m *Model) IncomingCommand(msg Message) (err error) {
	switch msg.Command {
	case "start":
		msg.Text, err = m.tgServer.CommandStart(&msg)
	case "help":
		msg.Text, err = m.tgServer.CommandHelp(&msg)
	case "setNote":
		msg.Text, err = m.tgServer.CommandSetNote(&msg)
	case "getStatistic":
		msg.Text, err = m.tgServer.CommandGetStatistic(&msg)
	case "setBudget":
		msg.Text, err = m.tgServer.CommandSetBudget(&msg)
	case "getBudget":
		msg.Text, err = m.tgServer.CommandGetBudget(&msg)
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
