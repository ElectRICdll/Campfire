package entity

import "time"

type Announcement struct {
	ID      ID
	OwnerID ID
	Title   string
	Begin   time.Time
	Content string
}

type AnnouncementDTO struct {
	ID      int       `json:"ID"`
	CampID  int       `json:"c_id"`
	OwnerID int       `json:"o_id"`
	Title   string    `json:"title"`
	Begin   time.Time `json:"begin"`
	Content string    `json:"content"`
}
