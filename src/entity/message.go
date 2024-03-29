package entity

import "time"

type Message struct {
	OwnerID uint `json:"ownerID" gorm:"primaryKey;autoIncrement:false"`
	CampID  uint `json:"campID" gorm:"primaryKey;autoIncrement:false"`
	ReplyID uint `json:"replyID"`

	Timestamp time.Time `json:"timestamp" gorm:"not null"`
	Type      int       `json:"eType" gorm:"not null"`
	Content   string    `json:"content"`
}
