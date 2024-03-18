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

type Task struct {
	ID        uint     `gorm:"primaryKey;autoIncrement"`
	OwnerID   uint     `gorm:"not null"`
	ProjID    uint     `gorm:"not null"`
	Receivers []Member `gorm:"foreignKey:ID"`

	Title       string
	BeginAt     time.Time
	EndAt       time.Time
	Content     string
	Status      int
	*util.Timer `gorm:"-"`
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

type TaskDTO struct {
	ID          uint   `json:"id"`
	OwnerID     uint   `json:"o_id"`
	ProjID      uint   `json:"p_id"`
	ReceiversID []uint `json:"r_id"`

	Title   string    `json:"task_title"`
	BeginAt time.Time `json:"begin_at"`
	EndAt   time.Time `json:"end_at"`
	Content string    `json:"content"`
	Status  int       `json:"status"`
}

func (t *Task) DTO() TaskDTO {
	return TaskDTO{
		ID:      t.ID,
		OwnerID: t.OwnerID,
		ProjID:  t.ProjID,
		ReceiversID: func(members []Member) []uint {
			res := []uint{}
			for _, member := range members {
				res = append(res, member.ID)
			}
			return res
		}(t.Receivers),
		Title:   t.Title,
		BeginAt: t.BeginAt,
		EndAt:   t.EndAt,
		Content: t.Content,
		Status:  t.Status,
	}
}

func TasksDTO(tasks []Task) []TaskDTO {
	res := []TaskDTO{}
	for _, task := range tasks {
		res = append(res, task.DTO())
	}
	return res
}
