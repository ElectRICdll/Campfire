package entity

import "time"

type Task struct {
	ID         ID `gorm:"primaryKey"`
	OwnerID    ID
	ProjID     ID
	ReceiverID []ID

	Title   string
	Begin   time.Time
	End     time.Time
	Content string
	Status  int

	Owner User    `gorm:"foreignKey:OwnerID"`
	Proj  Project `gorm:"foreignKey:ProjID"`
}

type TaskDTO struct {
	ID         ID   `json:"id"`
	ProjectID  ID   `json:"p_id"`
	OwnerID    ID   `json:"o_id"`
	ReceiverID []ID `json:"r_id"`

	Title   string    `json:"name"`
	Begin   time.Time `json:"begin"`
	End     time.Time `json:"end"`
	Content string    `json:"content"`
	Status  int       `json:"status"`
}
