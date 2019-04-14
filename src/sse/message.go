package sse

import (
	"bytes"
	"fmt"
)

// Message represents a event source message.
type Message struct {
	ID    string `json:"id"`
	Event string `json:"event"`
	Data  string `json:"data"`
	Retry int    `json:"retry"`
}

func SimpleMessage(data string) *Message {
	return NewMessage("", data, "")
}

func NewMessage(id string, data string, event string) *Message {
	return &Message{
		id,
		event,
		data,
		0,
	}
}

func (m *Message) String() string {
	var buffer bytes.Buffer

	if len(m.ID) > 0 {
		buffer.WriteString(fmt.Sprintf("id: %s\n", m.ID))
	}

	if m.Retry > 0 {
		buffer.WriteString(fmt.Sprintf("retry: %d\n", m.Retry))
	}

	if len(m.Event) > 0 {
		buffer.WriteString(fmt.Sprintf("event: %s\n", m.Event))
	}

	if len(m.Data) > 0 {
		buffer.WriteString(fmt.Sprintf("data: %s\n", m.Data))
	}

	buffer.WriteString("\n")

	return buffer.String()
}
