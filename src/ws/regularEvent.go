package ws

import (
	"campfire/entity"
	"time"
)

type ProjectInfoChangedEvent struct {
	Timestamp time.Time `json:"timestamp"`
	entity.Project
}

func (e ProjectInfoChangedEvent) ScopeID() uint {
	return e.ID
}

type NewTaskEvent struct {
	Timestamp time.Time `json:"timestamp"`
	entity.Task
}

func (t NewTaskEvent) ScopeID() uint {
	return t.Task.ProjID
}

type NewAnnouncementEvent struct {
	Timestamp time.Time `json:"timestamp"`
	entity.Announcement
}

func (a NewAnnouncementEvent) ScopeID() uint {
	return a.CampID
}

type RequestMessageRecordEvent struct {
	CampID  uint `json:"campID"`
	BeginAt uint `json:"begin"`
}

func (a RequestMessageRecordEvent) ScopeID() uint {
	return 0
}

type CampInfoChangedEvent struct {
	Timestamp time.Time `json:"timestamp"`
	entity.Camp
}

func (e CampInfoChangedEvent) ScopeID() uint {
	return e.ID
}

type CampDisableEvent struct {
	Timestamp time.Time `json:"timestamp"`
	CampID    uint
}

func (a CampDisableEvent) ScopeID() uint {
	return a.CampID
}

type MemberInfoChangedEvent struct {
	Timestamp time.Time `json:"timestamp"`
	entity.Member
}

func (e MemberInfoChangedEvent) ScopeID() uint {
	return e.CampID
}

type MemberExitedEvent struct {
	Timestamp time.Time `json:"timestamp"`
	UserID    uint      `json:"userID"`
	CampID    uint      `json:"campID"`
}

func (e MemberExitedEvent) ScopeID() uint {
	return e.CampID
}
