package dao

import . "campfire/entity"

type MessageDao interface {
	AddMessageRecord(campID uint, msg ...Message) error

	PullCampMessageRecord(campID uint, beginMessageID uint, msgCount uint) ([]Message, error)

	MessageRecord(campID uint, msgID uint) (Message, error)
}
