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

type CampDisableEvent struct {
	Timestamp time.Time
	CampID    uint
}

func (a CampDisableEvent) ScopeID() uint {
	return a.CampID
}

type MemberInfoChangedEvent struct {
	Timestamp time.Time `json:"timestamp"`
	UserID    uint      `json:"user_id"`
	CampID    uint      `json:"camp_id"`
	entity.MemberDTO
}

func (e MemberInfoChangedEvent) ScopeID() uint {
	return e.CampID
}
