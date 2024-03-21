package dao

import (
	. "campfire/entity"
	. "campfire/util"

	"gorm.io/gorm"
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

	IsUserACampMember(campID uint, userID uint) (bool, error)
}

func NewCampDao() CampDao {
	return campDao{}
}

type campDao struct{}

func (d campDao) IsUserACampMember(campID uint, userID uint) (bool, error) {
	panic("wait for implement")
}

func (d campDao) CampInfo(campID uint) (Camp, error) {
	var camp Camp
	var result = DB.Where("id = ?", campID).Find(&camp)

	if result.Error == gorm.ErrRecordNotFound {
		return camp, ExternalError{}
	}
	if result.Error != nil {
		return camp, result.Error
	}
	return camp, nil
}

func (d campDao) SetCampInfo(camp Camp) error {
	result := DB.Updates(&camp)
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d campDao) AddCamp(camp Camp) error {
	var result = DB.Save(&camp)
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	return nil
}

func (d campDao) DeleteCamp(campID uint) error {
	result := DB.Where("id = ?", campID).Delete(&Camp{})
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d campDao) MemberList(campID uint) ([]Member, error) {
	var member []Member

	var result = DB.Where("camp_id = ?", campID).Find(&member)
	if result.Error == gorm.ErrRecordNotFound {
		return member, ExternalError{}
	}
	if result.Error != nil {
		return member, result.Error
	}
	return member, nil
}

func (d campDao) MemberInfo(campID uint, userID uint) (Member, error) {
	var member Member
	var result = DB.Where("camp_id = ? AND user_id = ?", campID, userID).Find(&member)
	if result.Error == gorm.ErrRecordNotFound {
		return member, ExternalError{}
	}
	if result.Error != nil {
		return member, result.Error
	}
	return member, nil
}

func (d campDao) AddMember(campID uint, userID uint) error {
	var member = Member{CampID: campID, UserID: userID}
	var result = DB.Save(&member)
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d campDao) DeleteMember(campID uint, userID uint) error {
	var result = DB.Where("camp_id = ? AND user_id = ?", campID, userID).Delete(&Member{})
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d campDao) SetMemberInfo(campID uint, member Member) error {
	var result = DB.Where("camp_id = ?", campID).Updates(&member)
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d campDao) AnnouncementInfo(campID uint, annoID uint) (Announcement, error) {
	var announcement Announcement
	var result = DB.Where("camp_id = ? AND id = ?", campID, annoID).Save(&announcement)
	if result.Error == gorm.ErrRecordNotFound {
		return announcement, ExternalError{}
	}
	if result.Error != nil {
		return announcement, result.Error
	}
	return announcement, nil
}

func (d campDao) Announcements(campID uint) ([]Announcement, error) {
	var announcement []Announcement
	var result = DB.Where("camp_id = ?", campID).Save(&announcement)
	if result.Error == gorm.ErrRecordNotFound {
		return announcement, ExternalError{}
	}
	if result.Error != nil {
		return announcement, result.Error
	}
	return announcement, nil
}
func (d campDao) EditAnnouncement(anno Announcement) error {
	var result = DB.Updates(&anno)
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (d campDao) AddAnnouncement(anno Announcement) error {
	var result = DB.Save(&anno)
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (d campDao) DeleteAnnouncement(campID uint, annoID uint) error {
	var result = DB.Where("camp_id = ? AND id = ?", campID, annoID).Delete(&Announcement{})
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
