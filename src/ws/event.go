package ws

import (
	"campfire/dao"
	"campfire/entity"
	"errors"
)

type Event interface {
	ScopeID() uint
}

func NewEventService() EventService {
	return EventService{
		dao.ProjectDaoContainer,
		dao.CampDaoContainer,
		dao.MessageDaoContainer,
	}
}

type EventService struct {
	projQuery    dao.ProjectDao
	campQuery    dao.CampDao
	messageQuery dao.MessageDao
}

func (s EventService) HandleEvent(msg *Notification) error {
	scope := ScopeByType(msg.EType)
	switch scope {
	case OnCamp:
		res, err := s.campQuery.MemberList(msg.Event.ScopeID())
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
		res, err := s.projQuery.MemberList(msg.Event.ScopeID())
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
		msg.ReceiversID = append(msg.ReceiversID, msg.Event.ScopeID())
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
