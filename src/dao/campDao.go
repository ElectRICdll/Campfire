package dao

import . "campfire/entity"

type CampDao interface {
	CampInfo(queryMemberID ID, projID ID, campID ID) (Camp, error)

	AddCamp(queryMemberID ID, projID ID, camp Camp) error

	SetCampInfo(queryOwnerID ID, projID ID, camp Camp) error

	DeleteCamp(queryOwnerID, projID ID, campID ID) error

	MemberList(queryMemberID ID, projID ID, campID ID) ([]Member, error)

	MemberInfo(queryMemberID ID, projID ID, campID ID, userID ID) (Member, error)

	AddMember(queryOwnerID ID, projID ID, campID ID, userID ID) error

	DeleteMember(queryOwnerID ID, projID ID, campID ID, userID ID) error

	SetMemberInfo(projID ID, campID ID, member Member) error

	AnnouncementInfo(queryMemberID ID, projID ID, campID ID, annoID ID) (Announcement, error)

	Announcements(queryMemberID ID, projID ID, campID ID) ([]Announcement, error)

	EditAnnouncement(queryOwnerID ID, projID ID, campID ID, anno Announcement) error

	AddAnnouncement(queryOwnerID ID, projID ID, campID ID, anno Announcement) error

	DeleteAnnouncement(queryOwnerID ID, projID ID, campID ID, annoID ID) error
}
