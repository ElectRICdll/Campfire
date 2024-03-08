package dao

import "campfire/entity"

type CampsiteDao interface {
	CampsitesOfUser(userID int) ([]entity.Campsite, error)

	CampsitesOfProject(queryMemberID, projID int) ([]entity.Campsite, error)

	CampsiteInfo(queryMemberID int, campID int, userID int) (entity.Campsite, error)

	AddCampsite(queryMemberID int, projID int, camp entity.Campsite) error

	SetCampsiteInfo(queryOwnerID int, camp entity.Campsite) error

	DeleteCampsite(queryOwnerID, campID int) error

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
