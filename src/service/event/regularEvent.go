package event

import (
	"campfire/entity"
	"time"
)

type NewTaskEvent struct {
	Timestamp time.Time `json:"timestamp"`
	entity.TaskDTO
}

func (t NewTaskEvent) Process() func() error {

}

func (t NewTaskEvent) ScopeID() uint {
	return t.TaskDTO.ProjID
}

type AnnouncementEvent struct {
	Timestamp time.Time `json:"timestamp"`
	entity.AnnouncementDTO
}

func (a AnnouncementEvent) Process() func() error {
	//TODO implement me
	panic("implement me")
}

func (a AnnouncementEvent) ScopeID() uint {
	return a.AnnouncementDTO.CampID
}

type RequestMessageRecordEvent struct {
	CampID  uint `json:"c_id"`
	BeginAt uint `json:"begin_at"`
}

func (a RequestMessageRecordEvent) Publish() func() error {
	//TODO implement me
	panic("implement me")
}

func (a RequestMessageRecordEvent) ScopeID() uint {
	return 0
}
