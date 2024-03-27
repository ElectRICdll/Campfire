package dao

import (
	. "campfire/entity"
	. "campfire/util"
	"gorm.io/gorm"
)

type CampDao interface {
	CampInfo(campID uint) (Camp, error)

	AddCamp(ownerID uint, camp Camp, userID ...uint) (uint, error)

	SetCampInfo(camp Camp) error

	DeleteCamp(campID uint) error

	MemberList(campID uint) ([]Member, error)

	MemberInfo(campID uint, userID uint) (Member, error)

	AddMember(member Member) error

	DeleteMember(campID uint, userID uint) error

	SetMemberInfo(member Member) error

	Promotion(campID uint, memberID uint) error

	Demotion(campID uint, memberID uint) error

	TransferOwner(campID uint, memberID uint) error

	AnnouncementInfo(campID uint, annoID uint) (Announcement, error)

	Announcements(campID uint) ([]Announcement, error)

	EditAnnouncement(anno Announcement) error

	AddAnnouncement(anno Announcement) error

	DeleteAnnouncement(campID uint, annoID uint) error
}

func NewCampDao() CampDao {
	return &campDao{
		db: DBConn(),
	}
}

type campDao struct {
	db *gorm.DB
}

func (d *campDao) CampInfo(campID uint) (Camp, error) {
	var camp []Camp
	var result = d.db.Where("id = ?", campID).Find(&camp)

	if len(camp) == 0 {
		return Camp{}, NewExternalError("no such data")
	}
	if result.Error != nil {
		return Camp{}, result.Error
	}
	return camp[0], nil
}

func (d *campDao) SetCampInfo(camp Camp) error {
	result := d.db.Updates(Camp{Name: camp.Name})
	if result.Error == gorm.ErrRecordNotFound {
		return NewExternalError("no such data")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *campDao) AddCamp(ownerID uint, camp Camp, usersID ...uint) (uint, error) {
	tran := d.db.Begin()
	var result = tran.Create(&camp)
	if result.Error != nil {
		tran.Rollback()
		return 0, result.Error
	}
	if err := tran.Model(&camp).Update("Owner", &Member{CampID: camp.ID, UserID: ownerID}).Error; err != nil {
		tran.Rollback()
		return 0, err
	}
	for _, userID := range usersID {
		if err := tran.Model(&camp).Association("Regulars").Append(&Member{CampID: camp.ID, UserID: userID}); err != nil {
			tran.Rollback()
			return 0, err
		}
	}
	tran.Commit()
	return camp.ID, nil
}

func (d *campDao) DeleteCamp(campID uint) error {
	result := d.db.Where("id = ?", campID).Delete(&Camp{})
	if result.Error == gorm.ErrRecordNotFound {
		return NewExternalError("no such data")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *campDao) MemberList(campID uint) ([]Member, error) {
	var member []Member

	var result = d.db.Where("camp_id = ?", campID).Find(&member)
	if len(member) == 0 {
		return nil, NewExternalError("no such data")
	}
	if result.Error != nil {
		return member, result.Error
	}
	return member, nil
}

func (d *campDao) MemberInfo(campID uint, userID uint) (Member, error) {
	var camp Camp
	if err := d.db.First(&camp, "id = ?", campID).Error; err != nil {
		return Member{}, NewExternalError("no such data")
	}

	var member *Member
	if camp.Owner.UserID == userID {
		member = &camp.Owner
	} else {
		for _, ruler := range camp.Rulers {
			if ruler.UserID == userID {
				member = &ruler
				break
			}
		}
		if member == nil {
			for _, regular := range camp.Regulars {
				if regular.UserID == userID {
					member = &regular
					break
				}
			}
		}
	}
	if member == nil {
		return Member{}, NewExternalError("member not found")
	}

	return *member, nil
}

func (d *campDao) AddMember(member Member) error {
	var camp Camp
	var result = d.db.Where("id = ?", member.CampID).Find(&camp)
	if result.Error == gorm.ErrRecordNotFound {
		return NewExternalError("No such data")
	}
	if result.Error != nil {
		return result.Error
	}

	if err := d.db.Model(&camp).Association("Regulars").Append(&member); err != nil {
		return err
	}

	return nil
}

func (d *campDao) DeleteMember(campID uint, userID uint) error {
	var member Member
	var result = d.db.Where("camp_id = ? AND user_id = ?", campID, userID).Delete(&member)
	if result.Error == gorm.ErrRecordNotFound {
		return NewExternalError("no such data")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *campDao) SetMemberInfo(member Member) error {
	err := d.db.Updates(&member).Error
	if err == gorm.ErrRecordNotFound {
		return NewExternalError("no such data")
	}
	if err != nil {
		return err
	}

	return nil
}

func (d *campDao) AnnouncementInfo(campID uint, annoID uint) (Announcement, error) {
	var announcement Announcement
	var result = d.db.Where("camp_id = ? AND id = ?", campID, annoID).Save(&announcement)
	if result.Error == gorm.ErrRecordNotFound {
		return announcement, ExternalError{}
	}
	if result.Error != nil {
		return announcement, result.Error
	}
	return announcement, nil
}

func (d *campDao) Announcements(campID uint) ([]Announcement, error) {
	var announcement []Announcement
	var result = d.db.Where("camp_id = ?", campID).Save(&announcement)
	if result.Error == gorm.ErrRecordNotFound {
		return announcement, ExternalError{}
	}
	if result.Error != nil {
		return announcement, result.Error
	}
	return announcement, nil
}
func (d *campDao) EditAnnouncement(anno Announcement) error {
	var result = d.db.Updates(&anno)
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (d *campDao) AddAnnouncement(anno Announcement) error {
	var result = d.db.Save(&anno)
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (d *campDao) DeleteAnnouncement(campID uint, annoID uint) error {
	var result = d.db.Where("camp_id = ? AND id = ?", campID, annoID).Delete(&Announcement{})
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *campDao) Promotion(campID uint, memberID uint) error {
	var camp Camp
	if err := d.db.First(&camp, "id = ?", campID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return NewExternalError("no such data")
		}
		return err
	}

	var member ProjectMember
	if err := d.db.Where("user_id = ? AND camp_id = ?", memberID, campID).First(&member).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return NewExternalError("no such data")
		}
		return err
	}

	for _, ruler := range camp.Rulers {
		if ruler.UserID == member.UserID {
			return NewExternalError("member has already been ruler")
		}
	}

	tran := d.db.Begin()
	if err := tran.Model(&camp).Association("Rulers").Append(&member); err != nil {
		tran.Rollback()
		return err
	}

	if err := tran.Model(&camp).Association("Regulars").Delete(&member); err != nil && err != gorm.ErrRecordNotFound {
		tran.Rollback()
		return err
	}

	tran.Commit()
	return nil
}

func (d *campDao) Demotion(campID uint, memberID uint) error {
	var camp Camp
	if err := d.db.First(&camp, "id = ?", campID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return NewExternalError("no such data")
		}
		return err
	}

	var member ProjectMember
	if err := d.db.Where("user_id = ? AND camp_id = ?", memberID, campID).First(&member).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return NewExternalError("no such data")
		}
		return err
	}

	var isTitled bool
	for _, ruler := range camp.Rulers {
		if ruler.UserID == member.UserID {
			isTitled = true
			break
		}
	}
	if !isTitled {
		return NewExternalError("member has already been regulars")
	}

	tran := d.db.Begin()
	if err := tran.Model(&camp).Association("Rulers").Delete(&member); err != nil {
		tran.Rollback()
		return err
	}

	if err := tran.Model(&camp).Association("Regulars").Append(&member); err != nil && err != gorm.ErrRecordNotFound {
		tran.Rollback()
		return err
	}

	tran.Commit()
	return nil
}

func (d *campDao) TransferOwner(campID uint, memberID uint) error {
	var camp Camp
	err := d.db.Where("id = ?", campID).First(&camp).Error
	if err == gorm.ErrRecordNotFound {
		return NewExternalError("no such data")
	}
	if err != nil {
		return err
	}

	tran := d.db.Begin()
	var newOwner Member
	if err := tran.Model(&camp).Association("Owner").Find(&newOwner); err != nil {
		tran.Rollback()
		return err
	}
	if err := tran.Model(&camp).Association("Regulars").Append(&newOwner); err != nil {
		tran.Rollback()
		return err
	}
	if err := tran.Where("user_id = ? AND camp_id = ?", memberID, campID).First(&newOwner).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			tran.Rollback()
			return NewExternalError("no such data")
		}
		tran.Rollback()
		return err
	}

	if err := tran.Model(&camp).Update("Owner", newOwner).Error; err != nil {
		tran.Rollback()
		return err
	}
	tran.Commit()
	return nil
}
