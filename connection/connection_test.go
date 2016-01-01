package connection

import (
	"testing"
)

func TestConnection(t *testing.T) {
	conn, err := NewConnection("https://scalr.api.appbase.io", "dW9DQYdot", "40d5db8b-36c8-41ac-b6e9-d26d7e34ce1e", "testapp2")
	if err != nil {
		t.Error(err)
		return
	}

	err = conn.Ping()
	if err != nil {
		t.Error(err)
		return
	}
}
