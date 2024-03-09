package dao

import . "campfire/entity"

type CampDao interface {
	CampInfo(queryMemberID uint, campID uint) (Camp, error)

	AddCamp(queryMemberID uint, camp Camp) error

	SetCampInfo(queryOwnerID uint, camp Camp) error

	DeleteCamp(queryOwnerID, campID uint) error

	MemberList(queryMemberID uint, campID uint) ([]Member, error)

	MemberInfo(queryMemberID uint, campID uint, userID uint) (Member, error)

	AddMember(queryOwnerID uint, campID uint, userID uint) error

	DeleteMember(queryOwnerID uint, campID uint, userID uint) error

	SetMemberInfo(campID uint, member Member) error

	AnnouncementInfo(queryMemberID uint, campID uint, annoID uint) (Announcement, error)

	Announcements(queryMemberID uint, campID uint) ([]Announcement, error)

	EditAnnouncement(queryOwnerID uint, campID uint, anno Announcement) error

	AddAnnouncement(queryOwnerID uint, campID uint, anno Announcement) error

	DeleteAnnouncement(queryOwnerID uint, campID uint, annoID uint) error
}
