package service

import (
	"campfire/dao"
	"campfire/entity"
	"campfire/util"
	"campfire/ws"
)

type MessageHandler func(*ws.Notification) error

type MessageService interface {
	MessageRecord(messageID uint)

	FindMessageRecordByKeyword(keyword string)

	FindMessageRecordByMember(userID string)

	AllMessageRecord()

	NewMessageRecord(message ...entity.Message) error

	PullMessageRecord(campID uint, beginMessageID uint) ([]entity.Message, error)
}

func NewMessageService() MessageService {
	return messageService{
		messageQuery: dao.MessageDaoContainer,
		campQuery:    dao.CampDaoContainer,
		projQuery:    dao.ProjectDaoContainer,
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

func (s messageService) NewMessageRecord(message ...entity.Message) error {
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
