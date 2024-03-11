package service

import (
	"campfire/dao"
	. "campfire/entity"
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

	EditNickname(campID uint, userID uint, nickname string) error

	EditMemberTitle(campID uint, userID uint, title string) error

	Announcements(queryID uint, campID uint) ([]AnnouncementDTO, error)

	AnnouncementInfo(queryID uint, campID uint, annoID uint) (AnnouncementDTO, error)

	CreateAnnouncement(queryID uint, campID uint, anno Announcement) error

	EditAnnouncementInfo(queryID uint, campID uint, anno Announcement) error

	DeleteAnnouncement(queryID uint, campID uint, annoID uint) error
}

func NewCampService() CampService {
	return campService{
		mention:   SessionServiceContainer,
		campQuery: dao.CampDaoContainer,
		projQuery: dao.ProjectDaoContainer,
	}
}

type campService struct {
	mention   SessionService
	campQuery dao.CampDao
	projQuery dao.ProjectDao
}

func (c campService) MemberList(queryID uint, campID uint) ([]MemberDTO, error) {
	res, err := c.campQuery.MemberList(queryID, campID)
	if err != nil {
		return nil, err
	}
	return MembersDTO(res), nil
}

func (c campService) MemberInfo(queryID uint, campID uint, userID uint) (MemberDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (c campService) PublicCamps(queryID uint, projID uint) ([]BriefCampDTO, error) {
	res, err := c.projQuery.CampsOfProject(queryID, projID)
	if err != nil {
		return nil, err
	}
	camps := []BriefCampDTO{}
	for _, camp := range res {
		camps = append(camps, camp.BriefDTO())
	}
	return camps, nil
}

func (c campService) CampInfo(queryID uint, campID uint) (CampDTO, error) {
	res, err := c.campQuery.CampInfo(queryID, campID)
	if err != nil {
		return CampDTO{}, err
	}
	return res.DTO(), nil
}

func (c campService) CreateCamp(queryID uint, camp Camp) error {
	err := c.campQuery.AddCamp(queryID, camp)
	return err
}

func (c campService) EditCampInfo(queryID uint, camp Camp) error {
	if err := c.campQuery.SetCampInfo(queryID, camp); err != nil {
		return err
	}
	// TODO
	c.mention.notify(Notification{})
	return nil
}

func (c campService) DisableCamp(queryID uint, campID uint) error {
	if err := c.campQuery.DeleteCamp(queryID, campID); err != nil {
		return err
	}
	// TODO
	c.mention.notify(Notification{})
	return nil
}

func (c campService) InviteMember(queryID uint, campID uint, userID uint) error {
	if err := c.campQuery.AddMember(queryID, campID, userID); err != nil {
		return err
	}
	// TODO
	c.mention.notify(Notification{})
	return nil
}

func (c campService) KickMember(queryID uint, campID uint, userID uint) error {
	if err := c.campQuery.DeleteMember(queryID, campID, userID); err != nil {
		return err
	}
	// TODO
	c.mention.notify(Notification{})
	return nil
}

func (c campService) EditNickname(campID uint, userID uint, nickname string) error {
	if err := c.campQuery.SetMemberInfo(campID, Member{
		UserID:   userID,
		Nickname: nickname,
	}); err != nil {
		return err
	}
	// TODO
	c.mention.notify(Notification{})
	return nil
}

func (c campService) EditMemberTitle(campID uint, userID uint, title string) error {
	if err := c.campQuery.SetMemberInfo(campID, Member{
		UserID: userID,
		Title:  title,
	}); err != nil {
		return err
	}
	// TODO
	c.mention.notify(Notification{})
	return nil
}

func (c campService) Announcements(queryID uint, campID uint) ([]AnnouncementDTO, error) {
	res, err := c.campQuery.Announcements(queryID, campID)
	if err != nil {
		return nil, err
	}
	return AnnouncementsDTO(res), nil
}

func (c campService) AnnouncementInfo(queryID uint, campID uint, annoID uint) (AnnouncementDTO, error) {
	res, err := c.campQuery.AnnouncementInfo(queryID, campID, annoID)
	if err != nil {
		return AnnouncementDTO{}, err
	}
	return res.DTO(), nil
}

func (c campService) CreateAnnouncement(queryID uint, campID uint, anno Announcement) error {
	err := c.campQuery.AddAnnouncement(queryID, campID, anno)
	if err != nil {
		return err
	}
	// TODO
	c.mention.notify(Notification{})
	return nil
}

func (c campService) EditAnnouncementInfo(queryID uint, campID uint, anno Announcement) error {
	err := c.campQuery.EditAnnouncement(queryID, campID, anno)
	return err
}

func (c campService) DeleteAnnouncement(queryID uint, campID uint, annoID uint) error {
	err := c.campQuery.DeleteAnnouncement(queryID, campID, annoID)
	return err
}
