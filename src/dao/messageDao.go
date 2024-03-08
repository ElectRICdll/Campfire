package dao

import "campfire/entity"

type MessageDao interface {
	AddMessageRecord(campID int, msg ...entity.Message) error

	AddTentMessageRecord(campID int, tentID int, msg ...entity.Message) error

	PullCampsiteMessageRecord(campID int, beginMessageID int, msgCount int) ([]entity.Message, error)

	PullTentMessageRecord(campID int, tentID int, beginMessageID int, msgCount int) ([]entity.Message, error)

	MessageRecord(campID int, msgID int) (entity.Message, error)

	TentMessageRecord(campID int, tentID int, msgID int) (entity.Message, error)
}
