package entity

import "time"

type Message struct {
	ID      uint `json:"-" gorm:"primaryKey;autoIncrement"`
	OwnerID uint `json:"ownerID"`
	CampID  uint `json:"campID"`
	ReplyID uint `json:"replyID"`

	Timestamp time.Time `json:"timestamp" gorm:"not null"`
	Type      int       `json:"eType" gorm:"not null"`
	Content   string    `json:"content"`
}
