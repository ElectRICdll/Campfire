package ws_service

import (
	"campfire/entity/ws-entity"
	"time"
)

type Notification struct {
	Timestamp   time.Time      `json:"timestamp"`
	EType       int            `json:"e_type"`
	ReceiversID []uint         `json:"-"`
	Event       wsentity.Event `json:"event_info"`
}

func (n Notification) GetEventType() int {
	return n.EType
}

func (n Notification) GetEvent() wsentity.Event {
	return n.Event
}
