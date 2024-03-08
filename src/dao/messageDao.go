package dao

import . "campfire/entity"

type MessageDao interface {
	AddMessageRecord(projID ID, campID ID, msg ...Message) error

	PullCampMessageRecord(projID ID, campID ID, beginMessageID ID, msgCount ID) ([]Message, error)

	MessageRecord(projID ID, campID ID, msgID ID) (Message, error)
}
