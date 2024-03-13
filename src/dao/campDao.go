package dao

import (
	. "campfire/entity"
	. "campfire/util"
)

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
type campDao struct{}

func (d campDao) CampInfo(queryMemberID uint, campID uint) (Camp, error) {
	var member Member
	var camp Camp
	var result = DB.Where("UserID = ? and CampID = ?", queryMemberID, campID).Find(&member)
	if result.Error != nil {
		return camp, result.Error
	}
	if result == nil {
		return camp, ExternalError{}
	}

	result = DB.Where("ID = ?", campID).Find(&camp)

	if result.Error != nil {
		return camp, result.Error
	}
	if result == nil {
		return camp, ExternalError{}
	}
	return camp, nil
}

func (d campDao) SetCampInfo(queryOwnerID uint, camp Camp) error {
	result := DB.Where("OwnerID = ?", queryOwnerID).Save(&camp)
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	return nil
}

func (d campDao) AddCamp(queryMemberID uint, camp Camp) error {
	var result = DB.Where("UserID = ? and ID = ?", queryMemberID, camp.ProjID).Find(&Project{})
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	result = DB.Save(&camp)
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	return nil
}

func (d campDao) DeleteCamp(queryOwnerID, campID uint) error {
	result := DB.Where("OwnerID = ? and ID = ?", queryOwnerID, campID).Delete(&Camp{})
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	return nil
}

func (d campDao) MemberList(queryMemberID uint, campID uint) ([]Member, error) {
	var member []Member
	var result = DB.Where("UserID = ? and CampID = ?", queryMemberID, campID).Find(&member)
	if result.Error != nil {
		return member, result.Error
	}
	if result == nil {
		return member, ExternalError{}
	}

	result = DB.Where("CampID = ?", campID).Find(&member)
	if result.Error != nil {
		return member, result.Error
	}
	if result == nil {
		return member, ExternalError{}
	}
	return member, nil
}

func (d campDao) MemberInfo(queryMemberID uint, campID uint, userID uint) (Member, error) {
	var member Member
	var result = DB.Where("UserID = ? and CampID = ?", queryMemberID, campID).Find(&member)
	if result.Error != nil {
		return member, result.Error
	}
	if result == nil {
		return member, ExternalError{}
	}
	result = DB.Where("CampID = ? and UserID = ?", campID, userID).Find(&member)
	if result.Error != nil {
		return member, result.Error
	}
	if result == nil {
		return member, ExternalError{}
	}
	return member, nil
}

func (d campDao) AddMember(queryOwnerID uint, campID uint, userID uint) error {
	var result = DB.Where("OwnerID = ? and ID = ?", queryOwnerID, campID).Find(&Camp{})
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	var member = Member{CampID: campID, UserID: userID}
	result = DB.Save(&member)
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	return nil
}

func (d campDao) DeleteMember(queryOwnerID uint, campID uint, userID uint) error {
	var result = DB.Where("OwnerID = ? and ID = ?", queryOwnerID, campID).Find(&Camp{})
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	result = DB.Where("CampID = ? and userID = ?", campID, userID).Delete(&Member{})
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	return nil
}

func (d campDao) SetMemberInfo(campID uint, member Member) error {
	var result = DB.Where("campID = ?", campID).Save(&member)
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	return nil
}
