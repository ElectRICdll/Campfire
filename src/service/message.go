package service

import (
	"campfire/dao"
	"campfire/entity"
	"campfire/util"
	"errors"
)

type MessageHandler func(*entity.Notification) error

type MessageService interface {
	MessageRecord(messageID uint)

	FindMessageRecordByKeyword(keyword string)

	FindMessageRecordByMember(userID string)

	AllMessageRecord()

	newMessageRecord(message ...entity.Message) error

	PullMessageRecord(campID uint, beginMessageID uint) ([]entity.Message, error)

	eventMessageHandler(msg *entity.Notification) error
}

func NewMessageService() MessageService {
	return messageService{
		messageQuery: nil,
	}
}

type messageService struct {
	messageQuery dao.MessageDao
	campQuery    dao.CampDao
	projQuery    dao.ProjectDao
}

func (s messageService) PullMessageRecord(campID uint, beginMessageID uint) ([]entity.Message, error) {
	res, err := s.messageQuery.PullCampMessageRecord(campID, beginMessageID, util.CONFIG.MessageRecordCount)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s messageService) newMessageRecord(message ...entity.Message) error {
	err := s.messageQuery.AddMessageRecord(message...)
	return err
}

func (s messageService) MessageRecord(messageID uint) {
	//TODO implement me
	panic("implement me")
}

func (s messageService) FindMessageRecordByKeyword(keyword string) {
	//TODO implement me
	panic("implement me")
}

func (s messageService) FindMessageRecordByMember(userID string) {
	//TODO implement me
	panic("implement me")
}

func (s messageService) AllMessageRecord() {
	//TODO implement me
	panic("implement me")
}

func (s messageService) eventMessageHandler(msg *entity.Notification) error {
	switch entity.EventTypeIndex[msg.EType].Scope {
	case entity.OnCamp:
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
	case entity.OnProject:
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
	case entity.OnSomeone:

	default:
		return errors.New("unknown scope area")
	}
	//if err := msg.Event.Execute(); err != nil {
	//	return err
	//}
	return nil
}
