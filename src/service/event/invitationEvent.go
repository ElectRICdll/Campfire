package event

import (
	"campfire/entity"
	"time"
)

type InvitationEvent interface {
	Event

	Accept()

	Refuse()
}

type ProjectInvitationEvent struct {
	Timestamp time.Time `json:"timestamp"`
	TargetID  uint      `json:"target_id"`
	entity.BriefProjectDTO
}

func (a ProjectInvitationEvent) Process() func() error {
	//TODO implement me
	panic("implement me")
}

func (a ProjectInvitationEvent) ScopeID() uint {
	return 0
}

func (a ProjectInvitationEvent) Accept() {}

func (a ProjectInvitationEvent) Refuse() {}

type CampInvitationEvent struct {
	Timestamp time.Time `json:"timestamp"`
	TargetID  uint      `json:"target_id"`
	entity.BriefCampDTO
}

func (a CampInvitationEvent) Process() func() error {
	//TODO implement me
	panic("implement me")
}

func (a CampInvitationEvent) ScopeID() uint {
	return a.TargetID
}

func (a CampInvitationEvent) Accept() {}

func (a CampInvitationEvent) Refuse() {}
