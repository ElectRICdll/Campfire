package service

import (
	"campfire/entity"
	"encoding/json"
)

type MessageHandler func(entity.Message) (json.RawMessage, []entity.User, error)

type MessageService interface {
	unknownMessageHandler(message entity.Message) (json.RawMessage, []entity.User, error)

	textMessageHandler(message entity.Message) (json.RawMessage, []entity.User, error)

	binaryMessageHandler(message entity.Message) (json.RawMessage, []entity.User, error)

	eventMessageHandler(message entity.Message) (json.RawMessage, []entity.User, error)
}

func NewMessageService(callback func(dto entity.Message) error) MessageService {
	return messageService{
		us: UserServiceContainer,
	}
}

type messageService struct {
	us UserService
}

func (s messageService) unknownMessageHandler(message entity.Message) (json.RawMessage, []entity.User, error) {
	return nil, nil, entity.ExternalError{
		Message: "unknown type message."}
}

func (s messageService) textMessageHandler(message entity.Message) (json.RawMessage, []entity.User, error) {
	res, err := json.Marshal(entity.TextMessage{
		Message: message,
		Content: (string)(message.Content),
	})
	if err != nil {
		return nil, nil, entity.ExternalError{
			Message: "No such private channel.",
		}
	}

	users, err := s.us.campsiteList(message.CampsiteID)
	if err != nil {
		return nil, nil, err
	}
	return res, users, nil
}

func (s messageService) binaryMessageHandler(message entity.Message) (json.RawMessage, []entity.User, error) {
	res, err := json.Marshal(entity.BinaryMessage{
		Message: message,
	})
	if err != nil {
		return nil, nil, entity.ExternalError{
			Message: "No such private channel.",
		}
	}

	users, err := s.us.campsiteList(message.CampsiteID)
	if err != nil {
		return nil, nil, err
	}
	return res, users, nil
}

func (s messageService) eventMessageHandler(message entity.Message) (json.RawMessage, []entity.User, error) {
	res, err := json.Marshal(entity.TextMessage{
		Message: message,
		Content: (string)(message.Content),
	})
	if err != nil {
		return nil, nil, entity.ExternalError{
			Message: "No such private channel.",
		}
	}

	users, err := s.us.campsiteList(message.CampsiteID)
	if err != nil {
		return nil, nil, err
	}
	return res, users, nil
}

//func tentTest(message entity.Message, res []byte) (json.RawMessage, []entity.User, error) {
//	if message.TentID != 0 {
//		if value, ok := cache.TestProjects[0].Campsites[0].Tents[(entity.ID)(message.TentID)]; ok {
//			return res, []entity.User{*value.Target()}, nil
//		}
//		return nil, nil, entity.ExternalError{
//			Message: "No such private channel.",
//		}
//	}
//}
