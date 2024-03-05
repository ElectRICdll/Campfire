package entity

import "time"

type Task struct {
	ID         ID
	CreatorID  ID
	ReceiverID []ID
	Begin      time.Time
	End        time.Time
	Content    string
	Status     Status
}
