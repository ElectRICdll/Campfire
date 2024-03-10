package entity

import "time"

type Message struct {
	ID      uint `json:"m_id" gorm:"primaryKey;autoIncrement"`
	OwnerID uint `json:"o_id" gorm:"index;not null"`
	CampID  uint `json:"c_id" gorm:"index;not null"`
	ReplyID uint `json:"r_id" gorm:"index;not null"`

	Timestamp time.Time `json:"timestamp"`
	Type      int       `json:"m_type"`
	Content   string    `json:"-"`
}
