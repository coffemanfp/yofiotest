package handlers

import "fmt"

// message is a helper struct to return message responses.
type message struct {
	Error string `json:"error,omitempty"`
}

func newErrorMessage(s string, args ...interface{}) message {
	return message{Error: fmt.Sprintf(s, args...)}
}
