package service

import (
	"campfire/dao"
	. "campfire/entity"
	"campfire/util"
	"campfire/ws"
)

type CampService interface {
	PublicCamps(queryID uint, projID uint) ([]BriefCampDTO, error)

	CampInfo(queryID uint, campID uint) (CampDTO, error)

	CreateCamp(queryID uint, camp Camp) error

	EditCampInfo(queryID uint, camp Camp) error

	DisableCamp(queryID uint, campID uint) error

	MemberList(queryID uint, campID uint) ([]MemberDTO, error)

	MemberInfo(queryID uint, campID uint, userID uint) (MemberDTO, error)

	InviteMember(queryID uint, campID uint, userID uint) error

	KickMember(queryID uint, campID uint, userID uint) error

	EditMemberInfo(campID uint, member Member) error

	Announcements(queryID uint, campID uint) ([]AnnouncementDTO, error)

	AnnouncementInfo(queryID uint, campID uint, annoID uint) (AnnouncementDTO, error)

	CreateAnnouncement(queryID uint, anno Announcement) error

	EditAnnouncementInfo(queryID uint, anno Announcement) error

	DeleteAnnouncement(queryID uint, campID uint, annoID uint) error
}

func NewCampService() CampService {
	return campService{
		mention:   SessionServiceContainer,
		access:    SecurityServiceContainer,
		campQuery: dao.CampDaoContainer,
		projQuery: dao.ProjectDaoContainer,
	}
}

type campService struct {
	mention   *ws.SessionService
	access    SecurityService
	campQuery dao.CampDao
	projQuery dao.ProjectDao
}

func (c campService) MemberList(queryID uint, campID uint) ([]MemberDTO, error) {
	if err := c.access.IsUserACampMember(campID, queryID); err != nil {
		return nil, err
	}
	res, err := c.campQuery.MemberList(campID)
	if err != nil {
		return nil, err
	}
	return MembersDTO(res), nil
}

func (c campService) MemberInfo(queryID uint, campID uint, userID uint) (MemberDTO, error) {
	if err := c.access.IsUserACampMember(campID, queryID); err != nil {
		return MemberDTO{}, err
	}
	member, err := c.campQuery.MemberInfo(campID, userID)
	if err != nil {
		return MemberDTO{}, err
	}
	return member.DTO(), err
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

func (c campService) CampInfo(queryID uint, campID uint) (CampDTO, error) {
	if err := c.access.IsUserACampMember(campID, queryID); err != nil {
		return CampDTO{}, err
	}
	res, err := c.campQuery.CampInfo(campID)
	if err != nil {
		return CampDTO{}, err
	}
	return res.DTO(), nil
}

func (c campService) CreateCamp(queryID uint, camp Camp) error {
	if err := c.access.IsUserAProjMember(camp.ProjID, queryID); err != nil {
		return err
	}
	camp.Members = append(camp.Members, Member{
		UserID:   camp.OwnerID,
		IsLeader: true,
	})
	err := c.campQuery.AddCamp(camp)
	return err
}

func (c campService) EditCampInfo(queryID uint, camp Camp) error {
	if err := c.access.IsUserACampLeader(camp.ID, queryID); err != nil {
		return err
	}
	if err := c.campQuery.SetCampInfo(camp); err != nil {
		return err
	}
	if err := c.mention.NotifyByEvent(&ws.CampInfoChangedEvent{
		CampDTO: camp.DTO(),
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
	if err := c.campQuery.AddMember(campID, userID); err != nil {
		return err
	}
	if err := c.mention.NotifyByEvent(&ws.CampInvitationEvent{
		SourceID:     queryID,
		TargetID:     userID,
		IsAccepted:   0,
		KeepDuration: util.CONFIG.InvitationKeepDuration,
		BriefCampDTO: BriefCampDTO{
			ID: campID,
		},
	}, ws.CampInvitationEventType); err != nil {
		return err
	}
	return nil
}

func (c campService) KickMember(queryID uint, campID uint, userID uint) error {
	if err := c.access.IsUserACampLeader(campID, queryID); err != nil {
		return err
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

func (c campService) EditMemberInfo(campID uint, member Member) error {
	if err := c.campQuery.SetMemberInfo(campID, member); err != nil {
		return err
	}
	if err := c.mention.NotifyByEvent(&ws.MemberInfoChangedEvent{
		MemberDTO: MemberDTO{
			UserID: member.UserID,
			Title:  member.Title,
		},
	}, ws.MemberInfoChangedEventType); err != nil {
		return err
	}
	return nil
}

func (c campService) Announcements(queryID uint, campID uint) ([]AnnouncementDTO, error) {
	if err := c.access.IsUserACampMember(campID, queryID); err != nil {
		return nil, err
	}
	res, err := c.campQuery.Announcements(campID)
	if err != nil {
		return nil, err
	}
	return AnnouncementsDTO(res), nil
}

func (c campService) AnnouncementInfo(queryID uint, campID uint, annoID uint) (AnnouncementDTO, error) {
	if err := c.access.IsUserACampMember(campID, queryID); err != nil {
		return AnnouncementDTO{}, err
	}
	res, err := c.campQuery.AnnouncementInfo(campID, annoID)
	if err != nil {
		return AnnouncementDTO{}, err
	}
	return res.DTO(), nil
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
		AnnouncementDTO: anno.DTO(),
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
