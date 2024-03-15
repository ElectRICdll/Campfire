package wsentity

import (
	"campfire/entity"
	"time"
)

const (
	UnHandle = iota
	Accepted
	Refused
)

type InviteEvent interface {
	Activate(int)
}

type ProjectInvitationEvent struct {
	Timestamp  time.Time     `json:"timestamp"`
	TargetID   uint          `json:"target_id"`
	IsAccepted int           `json:"is_accepted"`
	KeepTime   time.Duration `json:"keep_time"`
	entity.BriefProjectDTO
}

func (a *ProjectInvitationEvent) Activate(newStatus int) {
	a.IsAccepted = newStatus
}

func (a *ProjectInvitationEvent) ScopeID() uint {
	return a.ID
}

type CampInvitationEvent struct {
	Timestamp    time.Time     `json:"timestamp"`
	SourceID     uint          `json:"source_id"`
	TargetID     uint          `json:"target_id"`
	IsAccepted   int           `json:"is_accepted"`
	KeepDuration time.Duration `json:"keep_time"`
	entity.BriefCampDTO
}

func (a *CampInvitationEvent) Activate(newStatus int) {
	a.IsAccepted = newStatus
}

func (a *CampInvitationEvent) ScopeID() uint {
	return a.TargetID
}

type InvitationAcceptEvent struct {
}
