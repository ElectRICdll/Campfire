package service

import "campfire/entity"

type CampService interface {
	CampInfo(userId int, campId int) (entity.CampDTO, error)

	CreateCamp(camp entity.Camp) error

	EditCampInfo(userId int, camp entity.Camp) error

	DisableCampInfo(camp entity.Camp) error

	MemberList(campId int) ([]entity.Member, error)

	EditNickname(campId int, userId int, nickname string) error

	EditMemberTitle(campId int, userId int, title string) error

	InviteMember(campId int, userId int) error

	KickMember(campId int, userId int) error

	AnnouncementInfo(campId int, annoId int) (entity.AnnouncementDTO, error)

	CreateAnnouncement(anno entity.Announcement) error

	EditAnnouncementInfo(anno entity.Announcement) error

	DeleteAnnouncement(annoId int) error
}

type campService struct {
}
