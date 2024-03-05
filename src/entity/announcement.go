package entity

import "time"

type Announcement struct {
	ID        ID
	CreatorID ID
	Begin     time.Time
	Content   string
}
