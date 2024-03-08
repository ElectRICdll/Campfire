package dao

import "campfire/entity"

type MessageDao interface {
	AddMessageRecord(projID int, campID int, msg ...entity.Message) error

	PullCampMessageRecord(projID int, campID int, beginMessageID int, msgCount int) ([]entity.Message, error)

	MessageRecord(projID int, campID int, msgID int) (entity.Message, error)
}
