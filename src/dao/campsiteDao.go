package dao

import "campfire/entity"

type CampsiteDao interface {
	CampsiteInfo(campsiteId int) (entity.Campsite, error)

	SetCampsiteName(campsiteId int, name string) error

	MemberList(campsiteId int) ([]entity.Member, error)

	MemberInfo(campsiteId int, userId int) (entity.Member, error)

	AddMember(campsiteId int, userId int) error

	DeleteMember(campsiteId int, userId int) error

	SetMemberNickname(campsiteId int, userId int, nickname string) error

	SetMemberTitle(campsiteId int, userId int, title string) error

	AnnouncementInfo(campsiteId int, annoId int) (entity.Announcement, error)

	AnnouncementsInfo(campsiteId int) ([]entity.Announcement, error)

	AddAnnouncement(campsiteId int, anno entity.Announcement) error

	DeleteAnnouncement(campsiteId int, annoId int) error

	AllMessageRecord(campsiteId int) ([]entity.Message, error)

	AllTentMessageRecord(campsiteId int, tentId int) ([]entity.Message, error)

	MessageRecord(campsiteId int, messageId int) (entity.Message, error)

	TentMessageRecord(campsiteId int, tentId int, messageId int) (entity.Message, error)
}
