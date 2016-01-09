package error

import (
	"encoding/json"
	"fmt"
)

type ErrorReason struct {
	Type   string `json:"type"`
	Reason string `json:"reason"`
}

type AppbaseError struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Err     *ErrorReason `json:"error"`
	Raw     *json.RawMessage
}

func (a *AppbaseError) Error() string {
	return fmt.Sprintf("%d: %s: [%s] %s ", a.Status, a.Message, a.Err.Type, a.Err.Reason)
}
