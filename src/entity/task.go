package entity

import "time"

type Task struct {
	ID       uint     `gorm:"primaryKey;autoIncrement"`
	OwnerID  uint     `gorm:"not null"`
	ProjID   uint     `gorm:"not null"`
	Receiver []Member `gorm:"foreignKey:ID"`

	Title   string
	Begin   time.Time
	End     time.Time
	Content string
	Status  int
}

type TaskDTO struct {
	ID         uint   `json:"id"`
	ProjectID  uint   `json:"p_id"`
	OwnerID    uint   `json:"o_id"`
	ReceiverID []uint `json:"r_id"`

	Title   string    `json:"name"`
	Begin   time.Time `json:"begin"`
	End     time.Time `json:"end"`
	Content string    `json:"content"`
	Status  int       `json:"status"`
}
