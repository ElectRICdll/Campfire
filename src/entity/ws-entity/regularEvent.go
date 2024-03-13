package wsentity

import (
	"campfire/entity"
	"time"
)

type NewTaskEvent struct {
	Timestamp time.Time `json:"timestamp"`
	entity.TaskDTO
}

func (t NewTaskEvent) ScopeID() uint {
	return t.TaskDTO.ProjID
}

type NewAnnouncementEvent struct {
	Timestamp time.Time `json:"timestamp"`
	entity.AnnouncementDTO
}

func (a NewAnnouncementEvent) ScopeID() uint {
	return a.CampID
}

type RequestMessageRecordEvent struct {
	CampID  uint `json:"c_id"`
	BeginAt uint `json:"begin_at"`
}

func (a RequestMessageRecordEvent) ScopeID() uint {
	return 0
}
