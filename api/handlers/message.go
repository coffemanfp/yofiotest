package handlers

import "fmt"

type message struct {
	Error string `json:"error,omitempty"`
}

func newErrorMessage(s string, args ...interface{}) message {
	return message{Error: fmt.Sprintf(s, args...)}
}
