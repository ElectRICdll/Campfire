package wsentity

import (
	"campfire/entity"
	"time"
)

type InviteEvent interface {
	Activate()
}

type ProjectInvitationEvent struct {
	Timestamp  time.Time     `json:"timestamp"`
	TargetID   uint          `json:"target_id"`
	IsAccepted int           `json:"is_accepted"`
	KeepTime   time.Duration `json:"keep_time"`
	entity.BriefProjectDTO
}

func (a *ProjectInvitationEvent) Activate() {

}

func (a *ProjectInvitationEvent) ScopeID() uint {
	return a.ID
}

type CampInvitationEvent struct {
	Timestamp  time.Time     `json:"timestamp"`
	TargetID   uint          `json:"target_id"`
	IsAccepted int           `json:"is_accepted"`
	KeepTime   time.Duration `json:"keep_time"`
	entity.BriefCampDTO
}

func (a *CampInvitationEvent) Activate() {

}

func (a *CampInvitationEvent) ScopeID() uint {
	return a.TargetID
}

type InvitationAcceptEvent struct {
}
