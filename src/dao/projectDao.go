package dao

import (
	. "campfire/entity"
	"campfire/util"
	"fmt"
	"gorm.io/gorm/clause"
	"time"

	//"fmt"
	//"time"

	//. "campfire/util"

	//"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

type ProjectDao interface {
	ProjectInfo(projID uint, with ...string) (Project, error) //需要鉴权

	SetProjectInfo(project Project) error //需要鉴权

	AddProject(ownerID uint, proj *Project, usersID ...uint) error

	DeleteProject(projID uint) error //需要鉴权

	MemberList(projID uint) ([]ProjectMember, error) //需要鉴权

	MemberInfo(projID uint, userID uint) (ProjectMember, error) //需要鉴权

	SetOwner(projID, userID uint) error

	RemoveTitle(projID, userID uint) error

	AddMember(member ProjectMember) error //需要鉴权

	DeleteMember(projID uint, userID uint) error //需要鉴权

	SetMemberInfo(member ProjectMember) error

	TasksOfProject(projID uint) ([]Task, error) //需要鉴权

	TaskInfo(taskID uint) (Task, error) //需要鉴权

	SetTaskInfo(projID uint, task Task) error //需要鉴权

	AddTask(task Task) (uint, uint, error) //需要鉴权

	DeleteTask(taskID uint) error //需要鉴权

	CampsOfProject(projID uint) ([]Camp, error)
}

func NewProjectDao() ProjectDao {
	return &projectDao{
		db: DBConn(),
	}
}

type projectDao struct {
	db *gorm.DB
}

func (d *projectDao) ProjectInfo(projID uint, with ...string) (Project, error) {
	var project Project
	var db *gorm.DB = d.db
	for _, value := range with {
		db = db.Preload(value)
	}
	var result = db.Model(&project).Where("id = ?", projID).First(&project)
	if result.Error == gorm.ErrRecordNotFound {
		return Project{}, util.NewExternalError("no such data")
	}
	if result.Error != nil {
		return Project{}, result.Error
	}

	return project, nil
}

func (d *projectDao) SetProjectInfo(project Project) error {
	var result = d.db.Updates(&project)
	if result.Error == gorm.ErrRecordNotFound {
		return util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *projectDao) AddProject(ownerID uint, proj *Project, usersID ...uint) error {
	proj.OwnerID = ownerID
	proj.BeginAt = time.Now()

	tran := d.db.Begin()
	if err := tran.Create(proj).Error; err != nil {
		tran.Rollback()
		return err
	}
	if err := tran.Model(proj).Association("Members").Append(&ProjectMember{ProjID: proj.ID, UserID: ownerID}); err != nil {
		tran.Rollback()
		return err
	}
	for _, userID := range usersID {
		if err := tran.Model(proj).Association("Members").Append(&ProjectMember{ProjID: proj.ID, UserID: userID}); err != nil {
			tran.Rollback()
			return err
		}
	}
	proj.Path = fmt.Sprintf("%s/%d-%s.git", util.CONFIG.NativeStorageRootPath, proj.ID, proj.Title)
	if err := tran.Updates(proj).Error; err != nil {
		tran.Rollback()
		return err
	}
	tran.Commit()
	return nil
}

func (d *projectDao) DeleteProject(projID uint) error {
	result := d.db.Select(clause.Associations).Delete(&Project{ID: projID})
	if result.Error == gorm.ErrRecordNotFound {
		return util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *projectDao) MemberList(projID uint) ([]ProjectMember, error) {
	var projmember []ProjectMember
	var result = d.db.Where("proj_id = ?", projID).Find(&projmember)
	if result.Error == gorm.ErrRecordNotFound {
		return projmember, util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return projmember, result.Error
	}
	return projmember, nil
}

func (d *projectDao) MemberInfo(projID uint, userID uint) (ProjectMember, error) {
	var projmember ProjectMember
	var result = d.db.Where("proj_id = ? AND user_id = ?", projID, userID).Find(&projmember)
	if result.Error == gorm.ErrRecordNotFound {
		return projmember, util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return projmember, result.Error
	}
	return projmember, nil
}

func (d *projectDao) SetOwner(projID uint, userID uint) error {
	var project Project
	err := d.db.Where("id = ?", projID).First(&project).Error
	if err == gorm.ErrRecordNotFound {
		return util.NewExternalError("no such data")
	}
	if err != nil {
		return err
	}

	//tran := d.db.Begin()
	var owner = ProjectMember{
		UserID: userID,
	}
	//if err := tran.Model(&project).Association("Owner").Find(&owner); err != nil {
	//	tran.Rollback()
	//	return err
	//}
	//if owner.UserID == userID {
	//	return util.NewExternalError("user has already been owner")
	//}
	//if err := tran.Model(&project).Association("Regulars").Append(&owner); err != nil {
	//	tran.Rollback()
	//	return err
	//}
	//if err := tran.Where("user_id = ? AND proj_id = ?", userID, projID).First(&owner).Error; err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		tran.Rollback()
	//		return util.NewExternalError("no such data")
	//	}
	//	tran.Rollback()
	//	return err
	//}
	if err := d.db.Model(&project).Update("Owner", owner).Error; err != nil {
		return err
	}
	d.db.Commit()
	return nil
}

func (d *projectDao) AddMember(member ProjectMember) error {
	var project Project
	var result = d.db.Where("id = ?", member.ProjID).Find(&project)
	if result.Error == gorm.ErrRecordNotFound {
		return util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return result.Error
	}

	if err := d.db.Model(&project).Association("Regulars").Append(&member); err != nil {
		return err
	}

	return nil
}

func (d *projectDao) DeleteMember(projID uint, userID uint) error {
	var result = d.db.Where("proj_id = ? AND user_id = ?", projID, userID).Delete(&ProjectMember{})
	if result.Error == gorm.ErrRecordNotFound {
		return util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *projectDao) SetMemberInfo(member ProjectMember) error {
	var result = d.db.Model(&member).Where("proj_id = ? and user_id = ?", member.ProjID, member.UserID).Updates(&member)
	if result.Error == gorm.ErrRecordNotFound {
		return util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *projectDao) TasksOfProject(projID uint) ([]Task, error) {
	var task []Task
	var result = d.db.Where("proj_id = ?", projID).Find(&task)
	if result.Error == gorm.ErrRecordNotFound {
		return task, util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return task, result.Error
	}
	return task, nil
}

func (d *projectDao) TaskInfo(taskID uint) (Task, error) {
	var task Task
	var result = d.db.Where("id = ?", taskID).Find(&task)
	if result.Error == gorm.ErrRecordNotFound {
		return task, util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return task, result.Error
	}
	return task, nil
}

func (d *projectDao) SetTaskInfo(projID uint, task Task) error {
	var result = d.db.Where("proj_id = ?", projID).Updates(&task)
	if result.Error == gorm.ErrRecordNotFound {
		return util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *projectDao) AddTask(task Task) (uint, uint, error) {
	var result = d.db.Save(&task)
	if result.Error == gorm.ErrRecordNotFound {
		return 0, 0, util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return 0, 0, result.Error
	}
	return task.ProjID, task.ID, nil
}

func (d *projectDao) DeleteTask(taskID uint) error {
	var result = d.db.Where("id = ?", taskID).Delete(&Task{})
	if result.Error == gorm.ErrRecordNotFound {
		return util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *projectDao) CampsOfProject(projID uint) ([]Camp, error) {
	var camp []Camp
	var result = d.db.Where("proj_id = ?", projID).Find(&camp)
	if result.Error == gorm.ErrRecordNotFound {
		return camp, util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return camp, result.Error
	}
	return camp, nil
}

func (d *projectDao) RemoveTitle(projID, userID uint) error {
	var member *ProjectMember
	if err := d.db.Model(member).Where("proj_id = ? and user_id = ?", projID, userID).Updates(member).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return util.NewExternalError("no such data")
		}
		return err
	}
	return nil
}
