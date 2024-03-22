package ws

import (
	"campfire/entity"
	"campfire/util"
	"time"
)

const (
	UnHandle = iota
	Accepted
	Refused
	Expired
)

type InviteEvent interface {
	Activate()

	Received(int)
}

type ProjectInvitationEvent struct {
	Timestamp  time.Time     `json:"timestamp"`
	TargetID   uint          `json:"targetID"`
	IsAccepted int           `json:"isAccepted"`
	KeepTime   time.Duration `json:"keepTime"`
	util.Timer `json:"-"`
	entity.BriefProjectDTO
}

func (a *ProjectInvitationEvent) Activate() {
	a.Start(a.Received, Expired)
}

func (a *ProjectInvitationEvent) Received(newStatus int) {
	a.Stop()
	a.IsAccepted = newStatus
}

func (a *ProjectInvitationEvent) ScopeID() uint {
	return a.ID
}

type CampInvitationEvent struct {
	Timestamp    time.Time     `json:"timestamp"`
	SourceID     uint          `json:"sourceID"`
	TargetID     uint          `json:"targetID"`
	IsAccepted   int           `json:"isAccepted"`
	KeepDuration time.Duration `json:"keepDuration"`
	util.Timer   `json:"-"`
	entity.BriefCampDTO
}

func (a *CampInvitationEvent) Activate() {
	a.Start(a.Received, Expired)
}

func (a *CampInvitationEvent) Received(newStatus int) {
	a.Stop()
	a.IsAccepted = newStatus
}

func (a *CampInvitationEvent) ScopeID() uint {
	return a.TargetID
}
