package dao

import "campfire/entity"

type CampDao interface {
	CampInfo(queryMemberID int, campID int, userID int) (entity.Camp, error)

	AddCamp(queryMemberID int, projID int, camp entity.Camp) error

	SetCampInfo(queryOwnerID int, camp entity.Camp) error

	DeleteCamp(queryOwnerID, campID int) error

	MemberList(queryMemberID int, campID int) ([]entity.Member, error)

	MemberInfo(queryMemberID int, campID int, userID int) (entity.Member, error)

	AddMember(queryOwnerID int, campID int, userID int) error

	DeleteMember(queryOwnerID int, campID int, userID int) error

	SetMemberInfo(campID int, member entity.Member) error

	AnnouncementInfo(queryMemberID int, campID int, annoID int) (entity.Announcement, error)

	Announcements(queryMemberID int, campID int) ([]entity.Announcement, error)

	EditAnnouncement(queryOwnerID int, campID int, anno entity.Announcement) error

	AddAnnouncement(queryOwnerID int, campID int, anno entity.Announcement) error

	DeleteAnnouncement(queryOwnerID int, campID int, annoID int) error
}
