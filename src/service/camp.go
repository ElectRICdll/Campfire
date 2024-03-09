package service

import (
	. "campfire/entity"
)

type CampService interface {
	PublicCamps(queryID uint, projID uint) ([]Camp, error)

	CampInfo(queryID uint, projID uint, campID uint) (CampDTO, error)

	CreateCamp(queryID uint, projID uint, camp Camp) error

	EditCampInfo(queryID uint, projID uint, camp Camp) error

	DisableCamp(queryID uint, projID uint, campID uint) error

	MemberList(queryID uint, projID uint, campID uint) ([]Member, error)

	MemberInfo(queryID uint, projID uint, campID uint, userID uint) (Member, error)

	InviteMember(queryID uint, projID uint, campID uint, userID uint) error

	KickMember(queryID uint, projID uint, campID uint, userID uint) error

	EditNickname(projID uint, campID uint, userID uint, nickname string) error

	EditMemberTitle(projID uint, campID uint, userID uint, title string) error

	AnnouncementInfo(queryID uint, projID uint, campID uint, annoID uint) (AnnouncementDTO, error)

	CreateAnnouncement(queryID uint, projID uint, campID uint, anno Announcement) error

	EditAnnouncementInfo(queryID uint, projID uint, campID uint, anno Announcement) error

	DeleteAnnouncement(queryID uint, projID uint, campID uint, annoID uint) error
}

type campService struct {
}
