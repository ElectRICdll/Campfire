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

func NewMessageDao() MessageDao {
	return messageDao{
		DBConn(),
	}
}

type messageDao struct {
	db *gorm.DB
}

func (d messageDao) AddMessageRecord(msg ...Message) error {
	result := DB.Create(&msg)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d messageDao) PullCampMessageRecord(campID uint, beginMessageID uint, msgCount uint) ([]Message, error) {
	var messages []Message

	query := d.db.Where("camp_id = ?", campID).Order("timestamp DESC")

	if beginMessageID != 0 {
		query = query.Where("id < ?", beginMessageID)
	}

	if err := query.Limit(50).Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

func (d messageDao) MessageRecord(campID uint, msgID uint) (Message, error) {
	var message Message
	result := DB.Where("camp_id = ? AND id = ?", campID, msgID).Find(&message)
	if result.Error == gorm.ErrRecordNotFound {
		return message, util.NewExternalError("no record found")
	}
	if result.Error != nil {
		return message, result.Error
	}
	return message, nil
}
