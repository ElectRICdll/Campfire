package event

import (
	"time"
)

type Notification struct {
	Timestamp   time.Time `json:"timestamp"`
	EType       int       `json:"e_type"`
	ReceiversID []uint    `json:"-"`
	Event       Event     `json:"event_info"`
}

func (n Notification) GetEventType() int {
	return n.EType
}

func (n Notification) GetEvent() Event {
	return n.Event
}
