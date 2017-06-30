package msg

import "testing"

type MockMessage struct {
	email, subject string
	body           []byte
}

// The MockMessage implements Mesagager
func (m *MockMessage) Send(email, subject string, body []byte) error {
	m.email = email
	m.subject = subject
	m.body = body
	return nil
}

func TestAlert(t *testing.T) {
	msgr := new(MockMessage) // create a new MockMesage
	body := []byte("Critical Error")

	Alert(msgr, body) // Run Alert method with the mock

	if msgr.subject != "Critical Error" {
		t.Errorf("Expectd 'Critical Error', Got '%s' ", msgr.subject)
	}

}
