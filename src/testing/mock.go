package msg

type Message struct{}

func (m *Message) Send(email, subject string, body []byte) error {
	return nil
}

// Define an interface that describes the methods you use on Message
type Messager interface {
	Send(email, subject string, body []byte) error
}

// Passes the Messager interface instead of the Message type
func Alert(m Messager, problem []byte) error {
	return m.Send("noc@examplecom", "Critical Error", problem)
}
