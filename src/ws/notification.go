package ws

import (
	"time"
)

type Notification struct {
	OperatorID  uint      `json:"userID"`
	Timestamp   time.Time `json:"timestamp"`
	EType       int       `json:"eType"`
	ReceiversID []uint    `json:"-"`
	Event       Event     `json:"data"`
}

func (n Notification) GetEventType() int {
	return n.EType
}

func (n Notification) GetEvent() Event {
	return n.Event
}
