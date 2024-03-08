package service

import (
	. "campfire/entity"
)

type CampService interface {
	PublicCamps(queryID ID, projID ID) ([]Camp, error)

	CampInfo(queryID ID, projID ID, campID ID) (CampDTO, error)

	CreateCamp(queryID ID, projID ID, camp Camp) error

	EditCampInfo(queryID ID, projID ID, camp Camp) error

	DisableCamp(queryID ID, projID ID, campID ID) error

	MemberList(queryID ID, projID ID, campID ID) ([]Member, error)

	MemberInfo(queryID ID, projID ID, campID ID, userID ID) (Member, error)

	InviteMember(queryID ID, projID ID, campID ID, userID ID) error

	KickMember(queryID ID, projID ID, campID ID, userID ID) error

	EditNickname(projID ID, campID ID, userID ID, nickname string) error

	EditMemberTitle(projID ID, campID ID, userID ID, title string) error

	AnnouncementInfo(queryID ID, projID ID, campID ID, annoID ID) (AnnouncementDTO, error)

	CreateAnnouncement(queryID ID, projID ID, campID ID, anno Announcement) error

	EditAnnouncementInfo(queryID ID, projID ID, campID ID, anno Announcement) error

	DeleteAnnouncement(queryID ID, projID ID, campID ID, annoID ID) error
}

type campService struct {
}
