package connection

import (
	"testing"
)

func TestConnection(t *testing.T) {
	conn, err := NewConnection("https://scalr.api.appbase.io", "QEVrcElba", "5c13d943-a5d1-4b05-92f3-42707d49fcbb", "es2test1")
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
