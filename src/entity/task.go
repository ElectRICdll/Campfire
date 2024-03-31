package entity

import (
	"campfire/util"
	"time"
)

const (
	Unknown = iota
	Activated
	Deactivated
)

const (
	NoStatus = iota
	NotStart
	Processing
	Completed
	Expired
)

type Task struct {
	ID      uint `gorm:"primaryKey;autoIncrement" json:"id"`
	ProjID  uint `gorm:"primaryKey;autoIncrement:false" json:"projID"`
	OwnerID uint `gorm:"not null" json:"ownerID"`

	Owner     User            `gorm:"foreignKey:OwnerID"`
	Receivers []TaskReceivers `gorm:"foreignKey:ProjectID,TaskID;references:ProjID,ID;onDelete:CASCADE"`
	Executors []TaskExecutors `gorm:"foreignKey:ProjectID,TaskID;references:ProjID,ID;onDelete:CASCADE"`

	Title       string    `gorm:"not null" json:"title"`
	IsFree      bool      `gorm:"not null" json:"isFree"`
	BeginAt     time.Time `gorm:"not null" json:"begin"`
	EndAt       time.Time `json:"end"`
	Content     string    `json:"content"`
	Status      int       `gorm:"not null" json:"status"`
	*util.Timer `gorm:"-" json:"-"`
}

func (t *Task) StartATimer() {
	t.Timer = &util.Timer{
		Duration: t.EndAt.Sub(t.BeginAt),
	}
	t.Start(t.SetStatus, Deactivated)
}

func (t *Task) SetStatus(code int) {
	t.Status = code
}

type TaskExecutors struct {
	TaskID       uint `gorm:"primaryKey"`
	ProjectID    uint `gorm:"primaryKey"`
	MemberUserID uint `gorm:"primaryKey"`
}

type TaskReceivers struct {
	TaskID       uint `gorm:"primaryKey"`
	ProjectID    uint `gorm:"primaryKey"`
	MemberUserID uint `gorm:"primaryKey"`
}
