package messages

func (m *Model) IncomingMessage(msg Message) error {

	var err error = nil

	if isCurrency := m.tgServer.IsCurrency(msg.Text); isCurrency {
		msg.Text, err = m.tgServer.MessageSetReportCurrency(&msg)
	} else {
		switch msg.Text {
		default:
			msg.Text, err = m.tgServer.MessageDefault(&msg)
		}
	}
	if err != nil {
		msg.Text = err.Error()
	}
	return m.tgClient.SendMessage(msg)
}
