package messages

type Messanger interface {
	IsCurrency(text string) bool
	MessageDefault(msg *Message) (answer string, err error)
	MessageSetCurrency(msg *Message) (answer string, err error)
}

func (m *Model) IncomingMessage(msg Message) error {

	var err error = nil

	if isCurrency := m.tgServer.IsCurrency(msg.Text); isCurrency {
		msg.Text, err = m.tgServer.MessageSetCurrency(&msg)
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
