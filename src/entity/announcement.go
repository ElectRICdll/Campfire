package entity

import "time"

type Announcement struct {
	OwnerID uint      `gorm:"primaryKey;autoIncrement:false" json:"campID"`
	CampID  uint      `gorm:"primaryKey;autoIncrement:false" json:"ownerID"`
	Title   string    `gorm:"not null" json:"title"`
	BeginAt time.Time `gorm:"not null" json:"begin"`
	Content string    `gorm:"not null" json:"content"`
}
