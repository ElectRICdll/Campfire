package dao

import (
	. "campfire/entity"
	"campfire/util"

	"gorm.io/gorm"
)

type MessageDao interface {
	AddMessageRecord(msg ...Message) error

	PullCampMessageRecord(campID uint, beginMessageID uint, msgCount uint) ([]Message, error)

	MessageRecord(campID uint, msgID uint) (Message, error)
}

type messageDao struct{}

func (d messageDao) AddMessageRecord(msg ...Message) error {
	result := DB.Create(&msg)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (d messageDao) PullCampMessageRecord(campID uint, beginMessageID uint, msgCount uint) ([]Message, error) {
	var message []Message
	result := DB.Where("campID = ? and ID >= ? and ID <= ?", campID, beginMessageID, beginMessageID-msgCount+1).Find(&message)
	if result.Error == gorm.ErrRecordNotFound {
		return message, util.NewExternalError("no record found")
	}
	if result.Error != nil {
		return message, result.Error
	}
	return message, nil
}
func (d messageDao) MessageRecord(campID uint, msgID uint) (Message, error) {
	var message Message
	result := DB.Where("campID = ? and ID = ?", campID, msgID).Find(&message)
	if result.Error == gorm.ErrRecordNotFound {
		return message, util.NewExternalError("no record found")
	}
	if result.Error != nil {
		return message, result.Error
	}
	return message, nil
}
