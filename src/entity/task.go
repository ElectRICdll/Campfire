package entity

import "time"

type Task struct {
	ID        uint     `gorm:"primaryKey;autoIncrement"`
	OwnerID   uint     `gorm:"not null"`
	ProjID    uint     `gorm:"not null"`
	Receivers []Member `gorm:"foreignKey:ID"`

	Title   string
	Begin   time.Time
	End     time.Time
	Content string
	Status  int
}

type TaskDTO struct {
	ID         uint   `json:"id"`
	OwnerID    uint   `json:"o_id"`
	ProjID     uint   `json:"p_id"`
	ReceiverID []uint `json:"r_id"`

	Title   string    `json:"name"`
	Begin   time.Time `json:"begin"`
	End     time.Time `json:"end"`
	Content string    `json:"content"`
	Status  int       `json:"status"`
}

func (t Task) DTO() TaskDTO {
	return TaskDTO{
		ID:      t.ID,
		OwnerID: t.OwnerID,
		ProjID:  t.ProjID,
		ReceiverID: func(members []Member) []uint {
			res := []uint{}
			for _, member := range members {
				res = append(res, member.ID)
			}
			return res
		}(t.Receivers),
		Title:   t.Title,
		Begin:   t.Begin,
		End:     t.End,
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
