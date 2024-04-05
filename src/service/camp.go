package service

import (
	"campfire/auth"
	"campfire/dao"
	. "campfire/entity"
	"campfire/util"
	"campfire/ws"
)

type CampService interface {
	PublicCamps(queryID uint, projID uint) ([]BriefCampDTO, error)

	CampInfo(queryID uint, campID uint) (Camp, error)

	CreateCamp(queryID uint, camp Camp, users ...uint) (uint, error)

	EditCampInfo(queryID uint, camp Camp) error

	DisableCamp(queryID uint, campID uint) error

	ExitCamp(queryID uint, campID uint) error

	GiveOwner(queryID uint, campID uint, userID uint) error

	GiveRuler(queryID uint, campID uint, userID uint) error

	DelRuler(queryID uint, campID uint, userID uint) error

	MemberList(queryID uint, campID uint) ([]Member, error)

	MemberInfo(queryID uint, campID uint, userID uint) (Member, error)

	InviteMember(queryID uint, campID uint, userID uint) error

	KickMember(queryID uint, campID uint, userID uint) error

	EditMemberInfo(member Member) error

	Announcements(queryID uint, campID uint) ([]Announcement, error)

	AnnouncementInfo(queryID uint, campID uint, annoID uint) (Announcement, error)

	CreateAnnouncement(queryID uint, anno Announcement) error

	EditAnnouncementInfo(queryID uint, anno Announcement) error

	DeleteAnnouncement(queryID uint, campID uint, annoID uint) error

	SetTitle(queryID uint, campID uint, userID uint, title string) error
}

func NewCampService() CampService {
	return campService{
		mention:   SessionServiceContainer,
		access:    auth.SecurityInstance,
		campQuery: dao.CampDaoContainer,
		projQuery: dao.ProjectDaoContainer,
		userQuery: dao.UserDaoContainer,
	}
}

type campService struct {
	mention   *ws.SessionService
	access    auth.SecurityGuard
	campQuery dao.CampDao
	projQuery dao.ProjectDao
	userQuery dao.UserDao
}

func (c campService) SetTitle(queryID uint, campID uint, userID uint, title string) error {
	if err := c.access.IsUserACampLeader(campID, queryID); err != nil {
		return err
	}
	if err := c.campQuery.SetMemberInfo(Member{
		UserID: userID,
		CampID: campID,
		Title:  title,
	}); err != nil {
		return err
	}
	return nil
}

func (c campService) ExitCamp(queryID uint, campID uint) error {
	if err := c.campQuery.DeleteMember(campID, queryID); err != nil {
		return err
	}
	return nil
}

func (c campService) GiveOwner(queryID uint, campID uint, userID uint) error {
	if err := c.access.IsUserACampLeader(campID, queryID); err != nil {
		return err
	}
	if err := c.campQuery.TransferOwner(campID, userID); err != nil {
		return err
	}
	return nil
}

func (c campService) GiveRuler(queryID uint, campID uint, userID uint) error {
	if err := c.access.IsUserACampLeader(campID, queryID); err != nil {
		return err
	}
	if err := c.campQuery.Promotion(campID, userID); err != nil {
		return err
	}
	return nil
}

func (c campService) DelRuler(queryID uint, campID uint, userID uint) error {
	if err := c.access.IsUserACampLeader(campID, queryID); err != nil {
		return err
	}
	if err := c.campQuery.Demotion(campID, userID); err != nil {
		return err
	}
	return nil
}

func (c campService) MemberList(queryID uint, campID uint) ([]Member, error) {
	if err := c.access.IsUserACampMember(campID, queryID); err != nil {
		return nil, err
	}
	res, err := c.campQuery.MemberList(campID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c campService) MemberInfo(queryID uint, campID uint, userID uint) (Member, error) {
	if err := c.access.IsUserACampMember(campID, queryID); err != nil {
		return Member{}, err
	}
	member, err := c.campQuery.MemberInfo(campID, userID)
	if err != nil {
		return Member{}, err
	}
	return member, err
}

func (c campService) PublicCamps(queryID uint, projID uint) ([]BriefCampDTO, error) {
	if err := c.access.IsUserAProjMember(projID, queryID); err != nil {
		return nil, err
	}
	res, err := c.projQuery.CampsOfProject(projID)
	if err != nil {
		return nil, err
	}
	camps := make([]BriefCampDTO, 0)
	for _, camp := range res {
		camps = append(camps, camp.BriefDTO())
	}
	return camps, nil
}

func (c campService) CampInfo(queryID uint, campID uint) (Camp, error) {
	if err := c.access.IsUserACampMember(campID, queryID); err != nil {
		return Camp{}, err
	}
	res, err := c.campQuery.CampInfo(campID, "Members.User", "Announcements", "MessageRecords")
	if err != nil {
		return Camp{}, err
	}
	return res, nil
}

func (c campService) CreateCamp(queryID uint, camp Camp, usersID ...uint) (uint, error) {
	if err := c.access.IsUserAProjMember(camp.ProjID, queryID); err != nil {
		return 0, err
	}
	for _, userID := range usersID {
		if err := c.access.IsUserAProjMember(camp.ProjID, userID); err != nil {
			return 0, util.NewExternalError("some user have no access.")
		}
	}
	id, err := c.campQuery.AddCamp(queryID, camp, usersID...)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (c campService) EditCampInfo(queryID uint, camp Camp) error {
	if err := c.access.IsUserACampLeader(camp.ID, queryID); err != nil {
		return err
	}
	if err := c.campQuery.SetCampInfo(camp); err != nil {
		return err
	}
	if err := c.mention.NotifyByEvent(&ws.CampInfoChangedEvent{
		Camp: camp,
	}, ws.CampInfoChangedEventType); err != nil {
		return err
	}
	return nil
}

func (c campService) DisableCamp(queryID uint, campID uint) error {
	if err := c.access.IsUserACampLeader(campID, queryID); err != nil {
		return err
	}
	if err := c.campQuery.DeleteCamp(campID); err != nil {
		return err
	}
	if err := c.mention.NotifyByEvent(&ws.CampDisableEvent{
		CampID: campID,
	}, ws.CampDisableEventType); err != nil {
		return err
	}
	return nil
}

func (c campService) InviteMember(queryID uint, campID uint, userID uint) error {
	if err := c.access.IsUserACampLeader(campID, queryID); err != nil {
		return err
	}
	if err := c.access.IsUserACampMember(campID, userID); err == nil {
		return util.NewExternalError("用户已在群聊中")
	}
	if err := c.campQuery.AddMember(Member{UserID: userID, CampID: campID}); err != nil {
		return err
	}
	res, err := c.campQuery.CampInfo(campID)
	if err != nil {
		return err
	}
	if err := c.mention.NotifyByEvent(ws.NewCampInvitationEvent(res.BriefDTO(), userID), ws.CampInvitationEventType); err != nil {
		return err
	}
	return nil
}

func (c campService) KickMember(queryID uint, campID uint, userID uint) error {
	if err := c.access.IsUserACampLeader(campID, queryID); err != nil {
		return err
	}
	if queryID == userID {
		return util.NewExternalError("群主无法踢除自身！")
	}
	if err := c.campQuery.DeleteMember(campID, userID); err != nil {
		return err
	}
	if err := c.mention.NotifyByEvent(&ws.MemberExitedEvent{
		CampID: campID,
	}, ws.MemberExitedEventType); err != nil {
		return err
	}
	return nil
}

func (c campService) EditMemberInfo(member Member) error {
	if err := c.campQuery.SetMemberInfo(member); err != nil {
		return err
	}
	if err := c.mention.NotifyByEvent(&ws.MemberInfoChangedEvent{
		Member: Member{
			UserID: member.UserID,
			Title:  member.Title,
		},
	}, ws.MemberInfoChangedEventType); err != nil {
		return err
	}
	return nil
}

func (c campService) Announcements(queryID uint, campID uint) ([]Announcement, error) {
	if err := c.access.IsUserACampMember(campID, queryID); err != nil {
		return nil, err
	}
	res, err := c.campQuery.Announcements(campID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c campService) AnnouncementInfo(queryID uint, campID uint, annoID uint) (Announcement, error) {
	if err := c.access.IsUserACampMember(campID, queryID); err != nil {
		return Announcement{}, err
	}
	res, err := c.campQuery.AnnouncementInfo(campID, annoID)
	if err != nil {
		return Announcement{}, err
	}
	return res, nil
}

func (c campService) CreateAnnouncement(queryID uint, anno Announcement) error {
	if err := c.access.IsUserACampLeader(anno.CampID, queryID); err != nil {
		return err
	}
	err := c.campQuery.AddAnnouncement(anno)
	if err != nil {
		return err
	}
	if err := c.mention.NotifyByEvent(&ws.NewAnnouncementEvent{
		Announcement: anno,
	}, ws.NewAnnouncementEventType); err != nil {
		return err
	}
	return nil
}

func (c campService) EditAnnouncementInfo(queryID uint, anno Announcement) error {
	if err := c.access.IsUserACampLeader(anno.CampID, queryID); err != nil {
		return err
	}
	err := c.campQuery.EditAnnouncement(anno)
	return err
}

func (c campService) DeleteAnnouncement(queryID uint, campID uint, annoID uint) error {
	if err := c.access.IsUserACampLeader(campID, queryID); err != nil {
		return err
	}
	err := c.campQuery.DeleteAnnouncement(campID, annoID)
	return err
}
