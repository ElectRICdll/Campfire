package entity

import (
	"time"
)

type AnnouncementDTO struct {
	ID         int       `json:"ID"`
	CampsiteID int       `json:"cs_id"`
	CreatorID  int       `json:"c_id"`
	Begin      time.Time `json:"begin"`
	Content    string    `json:"content"`
}

type LoginDTO struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
	WS    string `json:"ws"`
}

type TaskDTO struct {
	ID         int       `json:"id"`
	ProjectID  int       `json:"p_id"`
	CreatorID  int       `json:"c_id"`
	ReceiverID int       `json:"r_id"`
	Begin      time.Time `json:"begin"`
	End        time.Time `json:"end"`
	Content    string    `json:"content"`
	Status     int       `json:"status"`
}

type UserDTO struct {
	Email     string `json:"email,omitempty"`
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Signature string `json:"signature,omitempty"`
	Status    int    `json:"status,omitempty"`
}
