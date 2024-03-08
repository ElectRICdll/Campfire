package dao

import "campfire/entity"

type MessageDao interface {
	AddMessageRecord(campID int, msg ...entity.Message) error

	PullCampMessageRecord(campID int, beginMessageID int, msgCount int) ([]entity.Message, error)

	MessageRecord(campID int, msgID int) (entity.Message, error)
}
