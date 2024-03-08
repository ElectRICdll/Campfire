package service

import (
	"campfire/dao"
	"campfire/entity"
	"encoding/json"
)

type MessageHandler func(entity.Message) (json.RawMessage, []entity.User, error)

type MessageService interface {
	MessageRecord(messageId int)

	FindMessageRecordByKeyword(keyword string)

	FindMessageRecordByMember(userId string)

	AllMessageRecord()

	newMessageRecord(message ...entity.Message) error

	PullTentMessageRecord(userId int, campId int, tentId int, beginMessageId int) ([]entity.Message, error)

	PullCampMessageRecord(userId int, campId int, beginMessageId int) ([]entity.Message, error)

	unknownMessageHandler(message entity.Message) (json.RawMessage, []entity.User, error)

	textMessageHandler(message entity.Message) (json.RawMessage, []entity.User, error)

	binaryMessageHandler(message entity.Message) (json.RawMessage, []entity.User, error)

	eventMessageHandler(message entity.Message) (json.RawMessage, []entity.User, error)
}

func NewMessageService() MessageService {
	return messageService{
		query: nil,
	}
}

type messageService struct {
	query dao.CampsiteDao
}

func (s messageService) PullTentMessageRecord(userId int, campId int, tentId int, beginMessageId int) ([]entity.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (s messageService) PullCampMessageRecord(userId int, campId int, beginMessageId int) ([]entity.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (s messageService) newMessageRecord(message ...entity.Message) error {
	//TODO implement me
	panic("implement me")
}

func (s messageService) MessageRecord(messageId int) {
	//TODO implement me
	panic("implement me")
}

func (s messageService) FindMessageRecordByKeyword(keyword string) {
	//TODO implement me
	panic("implement me")
}

func (s messageService) FindMessageRecordByMember(userId string) {
	//TODO implement me
	panic("implement me")
}

func (s messageService) AllMessageRecord() {
	//TODO implement me
	panic("implement me")
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

	members, err := s.query.MemberList(message.CampID)
	if err != nil {
		return nil, nil, err
	}
	users := []entity.User{}
	for _, member := range members {
		users = append(users, *member.User)
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

	members, err := s.query.MemberList(message.CampID)
	if err != nil {
		return nil, nil, err
	}
	users := []entity.User{}
	for _, member := range members {
		users = append(users, *member.User)
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

	members, err := s.query.MemberList(message.CampID)
	if err != nil {
		return nil, nil, err
	}
	users := []entity.User{}
	for _, member := range members {
		users = append(users, *member.User)
	}
	return res, users, nil
}

//func tentTest(message entity.Message, res []byte) (json.RawMessage, []entity.User, error) {
//	if message.TentID != 0 {
//		if value, ok := cache.TestProjects[0].Camps[0].Tents[(entity.ID)(message.TentID)]; ok {
//			return res, []entity.User{*value.Target()}, nil
//		}
//		return nil, nil, entity.ExternalError{
//			Message: "No such private channel.",
//		}
//	}
//}
