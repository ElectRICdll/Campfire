package dao

import "campfire/entity"

type CampDao interface {
	CampInfo(queryMemberID int, projID int, campID int) (entity.Camp, error)

	AddCamp(queryMemberID int, projID int, camp entity.Camp) error

	SetCampInfo(queryOwnerID int, projID int, camp entity.Camp) error

	DeleteCamp(queryOwnerID, projID int, campID int) error

	MemberList(queryMemberID int, projID int, campID int) ([]entity.Member, error)

	MemberInfo(queryMemberID int, projID int, campID int, userID int) (entity.Member, error)

	AddMember(queryOwnerID int, projID int, campID int, userID int) error

	DeleteMember(queryOwnerID int, projID int, campID int, userID int) error

	SetMemberInfo(projID int, campID int, member entity.Member) error

	AnnouncementInfo(queryMemberID int, projID int, campID int, annoID int) (entity.Announcement, error)

	Announcements(queryMemberID int, projID int, campID int) ([]entity.Announcement, error)

	EditAnnouncement(queryOwnerID int, projID int, campID int, anno entity.Announcement) error

	AddAnnouncement(queryOwnerID int, projID int, campID int, anno entity.Announcement) error

	DeleteAnnouncement(queryOwnerID int, projID int, campID int, annoID int) error
}
