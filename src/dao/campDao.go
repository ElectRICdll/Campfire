package dao

import . "campfire/entity"

type CampDao interface {
	CampInfo(queryMemberID uint, projID uint, campID uint) (Camp, error)

	AddCamp(queryMemberID uint, projID uint, camp Camp) error

	SetCampInfo(queryOwnerID uint, projID uint, camp Camp) error

	DeleteCamp(queryOwnerID, projID uint, campID uint) error

	MemberList(queryMemberID uint, projID uint, campID uint) ([]Member, error)

	MemberInfo(queryMemberID uint, projID uint, campID uint, userID uint) (Member, error)

	AddMember(queryOwnerID uint, projID uint, campID uint, userID uint) error

	DeleteMember(queryOwnerID uint, projID uint, campID uint, userID uint) error

	SetMemberInfo(projID uint, campID uint, member Member) error

	AnnouncementInfo(queryMemberID uint, projID uint, campID uint, annoID uint) (Announcement, error)

	Announcements(queryMemberID uint, projID uint, campID uint) ([]Announcement, error)

	EditAnnouncement(queryOwnerID uint, projID uint, campID uint, anno Announcement) error

	AddAnnouncement(queryOwnerID uint, projID uint, campID uint, anno Announcement) error

	DeleteAnnouncement(queryOwnerID uint, projID uint, campID uint, annoID uint) error
}
