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
	Timestamp    time.Time              `json:"timestamp"`
	TargetID     uint                   `json:"targetID"`
	IsAccepted   int                    `json:"isAccepted"`
	KeepDuration time.Duration          `json:"keepTime"`
	Data         entity.BriefProjectDTO `json:"projectInfo"`
	util.Timer   `json:"-"`
}

func NewProjectInvitationEvent(data entity.BriefProjectDTO, targetID uint) *ProjectInvitationEvent {
	return &ProjectInvitationEvent{
		Timestamp:    time.Now(),
		TargetID:     targetID,
		IsAccepted:   0,
		KeepDuration: util.CONFIG.InvitationKeepDuration,
		Data:         data,
	}
}

func (a *ProjectInvitationEvent) Activate() {
	a.Start(a.Received, Expired)
}

func (a *ProjectInvitationEvent) Received(newStatus int) {
	a.Stop()
	a.IsAccepted = newStatus
}

func (a *ProjectInvitationEvent) ScopeID() uint {
	return a.TargetID
}

type CampInvitationEvent struct {
	Timestamp    time.Time           `json:"timestamp"`
	SourceID     uint                `json:"sourceID"`
	TargetID     uint                `json:"targetID"`
	IsAccepted   int                 `json:"isAccepted"`
	KeepDuration time.Duration       `json:"keepDuration"`
	Data         entity.BriefCampDTO `json:"campInfo"`
	util.Timer   `json:"-"`
}

func NewCampInvitationEvent(data entity.BriefCampDTO, targetID uint) *CampInvitationEvent {
	return &CampInvitationEvent{
		Timestamp:    time.Now(),
		TargetID:     targetID,
		IsAccepted:   0,
		KeepDuration: util.CONFIG.InvitationKeepDuration,
		Data:         data,
	}
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
