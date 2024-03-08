package entity

import "time"

type Task struct {
	ID         ID
	CreatorID  ID
	ReceiverID []ID
	Title      string
	Begin      time.Time
	End        time.Time
	Content    string
	Status     Status
	Belong     *Project
}

type TaskDTO struct {
	ID         int       `json:"id"`
	ProjectID  int       `json:"p_id"`
	OwnerID    int       `json:"o_id"`
	ReceiverID int       `json:"r_id"`
	Title      string    `json:"name"`
	Begin      time.Time `json:"begin"`
	End        time.Time `json:"end"`
	Content    string    `json:"content"`
	Status     int       `json:"status"`
}
