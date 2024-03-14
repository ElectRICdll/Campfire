package ws_service

import (
	"campfire/dao"
	"campfire/entity"
	. "campfire/entity/ws-entity"
	"errors"
)

type EventService struct {
	projQuery    dao.ProjectDao
	campQuery    dao.CampDao
	messageQuery dao.MessageDao
}

func (s EventService) HandleEvent(msg *Notification) error {
	_, scope := GetEventByType((EventType)(msg.EType))
	switch scope {
	case OnCamp:
		res, err := s.campQuery.MemberList(1, msg.Event.ScopeID())
		if err != nil {
			return err
		}
		msg.ReceiversID = func(members []entity.Member) []uint {
			res := []uint{}
			for _, member := range members {
				res = append(res, member.UserID)
			}
			return res
		}(res)
	case OnProject:
		res, err := s.projQuery.MemberList(1, msg.Event.ScopeID())
		if err != nil {
			return err
		}
		msg.ReceiversID = func(members []entity.ProjectMember) []uint {
			res := []uint{}
			for _, member := range members {
				res = append(res, member.UserID)
			}
			return res
		}(res)
	case OnSomeone:

	default:
		return errors.New("unknown scope area")
	}

	if event, ok := msg.Event.(InviteEvent); ok {
		event.Activate()
	} else if event, ok := msg.Event.(MessageEvent); ok {
		if err := s.messageQuery.AddMessageRecord(event.ToMessage()); err != nil {
			return err
		}
	}
	return nil
}
