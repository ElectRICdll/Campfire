package dao

import . "campfire/entity"

type MessageDao interface {
	AddMessageRecord(projID uint, campID uint, msg ...Message) error

	PullCampMessageRecord(projID uint, campID uint, beginMessageID uint, msgCount uint) ([]Message, error)

	MessageRecord(projID uint, campID uint, msgID uint) (Message, error)
}
