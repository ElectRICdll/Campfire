package entity

import "time"

type Message struct {
	ID      uint `json:"msgID" gorm:"primaryKey;autoIncrement"`
	OwnerID uint `json:"ownerID" gorm:"index;not null"`
	CampID  uint `json:"campID" gorm:"index;not null"`
	ReplyID uint `json:"replyID" gorm:"index;not null"`

	Timestamp time.Time `json:"timestamp"`
	Type      int       `json:"mType"`
	Content   string    `json:"-"`
}
