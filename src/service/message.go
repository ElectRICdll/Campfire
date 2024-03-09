package service

import (
	"campfire/dao"
	"campfire/entity"
	"campfire/util"
	"encoding/json"
)

type MessageHandler func(entity.Message) (json.RawMessage, error)

type MessageService interface {
	MessageRecord(messageID uint)

	FindMessageRecordByKeyword(keyword string)

	FindMessageRecordByMember(userID string)

	AllMessageRecord()

	newMessageRecord(message ...entity.Message) error

	PullMessageRecord(campID uint, beginMessageID uint) ([]entity.Message, error)

	unknownMessageHandler(message entity.Message) (json.RawMessage, error)

	textMessageHandler(message entity.Message) (json.RawMessage, error)

	binaryMessageHandler(message entity.Message) (json.RawMessage, error)

	eventMessageHandler(message entity.Message) (json.RawMessage, error)
}

func NewMessageService() MessageService {
	return messageService{
		query: nil,
	}
}

type messageService struct {
	query dao.MessageDao
}

func (s messageService) PullMessageRecord(campID uint, beginMessageID uint) ([]entity.Message, error) {
	res, err := s.query.PullCampMessageRecord(campID, beginMessageID, util.CONFIG.MessageRecordCount)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s messageService) newMessageRecord(message ...entity.Message) error {
	err := s.query.AddMessageRecord(message...)
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

func (s messageService) unknownMessageHandler(message entity.Message) (json.RawMessage, error) {
	return nil, entity.ExternalError{Message: "unknown type message."}
}

func (s messageService) textMessageHandler(message entity.Message) (json.RawMessage, error) {
	res, err := json.Marshal(entity.TextMessage{
		Message: message,
		Content: (string)(message.Content),
	})
	if err != nil {
		return nil, entity.ExternalError{
			Message: "No such private channel.",
		}
	}

	if err := s.newMessageRecord(message); err != nil {
		return nil, err
	}

	return res, nil
}

func (s messageService) binaryMessageHandler(message entity.Message) (json.RawMessage, error) {
	res, err := json.Marshal(entity.TextMessage{
		Message: message,
		Content: (string)(message.Content),
	})
	if err != nil {
		return nil, entity.ExternalError{
			Message: "invalid syntax",
		}
	}

	if err := s.newMessageRecord(message); err != nil {
		return nil, err
	}

	return res, nil
}

func (s messageService) eventMessageHandler(message entity.Message) (json.RawMessage, error) {
	res, err := json.Marshal(entity.TextMessage{
		Message: message,
		Content: (string)(message.Content),
	})
	if err != nil {
		return nil, entity.ExternalError{
			Message: "No such private channel.",
		}
	}

	if err := s.newMessageRecord(message); err != nil {
		return nil, err
	}

	return res, nil
}
