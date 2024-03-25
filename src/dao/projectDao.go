package dao

import (
	. "campfire/entity"
	"campfire/util"

	//"fmt"
	//"time"

	//. "campfire/util"

	//"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

type ProjectDao interface {
	ProjectInfo(projID uint) (Project, error) //需要鉴权

	SetProjectInfo(project Project) error //需要鉴权

	AddProject(proj Project) error

	DeleteProject(projID uint) error //需要鉴权

	MemberList(projID uint) ([]ProjectMember, error) //需要鉴权

	MemberInfo(projID uint, userID uint) (ProjectMember, error) //需要鉴权

	AddMember(projID uint, userID uint) error //需要鉴权

	DeleteMember(projID uint, userID uint) error //需要鉴权

	SetMemberInfo(campID uint, member ProjectMember) error

	TasksOfProject(projID uint) ([]Task, error) //需要鉴权

	TaskInfo(taskID uint) (Task, error) //需要鉴权

	SetTaskInfo(projID uint, task Task) error //需要鉴权

	AddTask(task Task) error //需要鉴权

	DeleteTask(taskID uint) error //需要鉴权

	CampsOfProject(projID uint) ([]Camp, error)

	IsUserAProjectMember(projID uint, userID uint) (bool, error)
}

func NewProjectDao() ProjectDao {
	return projectDao{}
}

type projectDao struct{}

func (d projectDao) IsUserAProjectMember(projID uint, userID uint) (bool, error) {
	// if project, found := ProjectCache.Get(fmt.Sprintf("%d", projID)); found {
	// 	for _, value := range project.(Project).Members {
	// 		if value.UserID == userID {
	// 			return true, nil
	// 		}
	// 	}
	// }

	// project, err := projectDao.ProjectInfo(projectDao{}, userID, projID)
	// if project.ID != 0 {
	// 	ProjectCache.Set(fmt.Sprintf("%d", projID), &project, cache.DefaultExpiration)
	// 	for _, value := range project.Members {
	// 		if value.UserID == userID {
	// 			return true, nil
	// 		}
	// 	}
	// }
	// return false, err
	return false, DB.Error
}

func (d projectDao) ProjectInfo(projID uint) (Project, error) {

	var project Project
	var result = DB.Where("id = ?",projID).Find(&project)
	if result.Error == gorm.ErrRecordNotFound {
		return project, util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return project, result.Error
	}

	return project, nil
}

func (d projectDao) SetProjectInfo(project Project) error {

	var result = DB.Updates(&project)
	if result.Error == gorm.ErrRecordNotFound {
		return util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d projectDao) AddProject(proj Project) error {
	result := DB.Save(&proj)
	if result.Error == gorm.ErrRecordNotFound {
		return util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d projectDao) DeleteProject(projID uint) error {
	result := DB.Where("id = ?", projID).Delete(&Project{})
	if result.Error == gorm.ErrRecordNotFound {
		return util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d projectDao) MemberList(projID uint) ([]ProjectMember, error) {
	var projmember []ProjectMember
	var result = DB.Where("proj_id = ?", projID).Find(&projmember)
	if result.Error == gorm.ErrRecordNotFound {
		return projmember, util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return projmember, result.Error
	}
	return projmember, nil
}

func (d projectDao) MemberInfo(projID uint, userID uint) (ProjectMember, error) {
	var projmember ProjectMember
	var result = DB.Where("proj_id = ? AND user_id = ?", projID, userID).Find(&projmember)
	if result.Error == gorm.ErrRecordNotFound {
		return projmember, util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return projmember, result.Error
	}
	return projmember, nil
}

func (d projectDao) AddMember(projID uint, userID uint) error {
	var result = DB.Where("id = ?", projID).Find(&Project{})
	var projmember = ProjectMember{ProjID: projID, UserID: userID}
	result = DB.Save(&projmember)
	if result.Error == gorm.ErrRecordNotFound {
		return util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d projectDao) DeleteMember(projID uint, userID uint) error {
	var result = DB.Where("proj_id = ? AND user_id = ?", projID, userID).Delete(&ProjectMember{})
	if result.Error == gorm.ErrRecordNotFound {
		return util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d projectDao) SetMemberInfo(projID uint, member ProjectMember) error {
	var result = DB.Where("proj_id = ?", projID).Updates(&member)
	if result.Error == gorm.ErrRecordNotFound {
		return util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d projectDao) TasksOfProject(projID uint) ([]Task, error) {
	var task []Task
	var result = DB.Where("proj_id = ?", projID).Find(&task)
	if result.Error == gorm.ErrRecordNotFound {
		return task, util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return task, result.Error
	}
	return task, nil
}

func (d projectDao) TaskInfo(taskID uint) (Task, error) {
	var task Task
	var result = DB.Where("id = ?", taskID).Find(&task)
	if result.Error == gorm.ErrRecordNotFound {
		return task, util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return task, result.Error
	}
	return task, nil
}

func (d projectDao) SetTaskInfo(projID uint, task Task) error {
	var result = DB.Where("proj_id = ?", projID).Updates(&task)
	if result.Error == gorm.ErrRecordNotFound {
		return util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d projectDao) AddTask(task Task) error {
	var result = DB.Save(&task)
	if result.Error == gorm.ErrRecordNotFound {
		return util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d projectDao) DeleteTask(taskID uint) error {
	var result = DB.Where("id = ?", taskID).Delete(&Task{})
	if result.Error == gorm.ErrRecordNotFound {
		return util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d projectDao) CampsOfProject(projID uint) ([]Camp, error) {
	var camp []Camp
	var result = DB.Where("proj_id = ?", projID).Find(&camp)
	if result.Error == gorm.ErrRecordNotFound {
		return camp, util.NewExternalError("No such data")
	}
	if result.Error != nil {
		return camp, result.Error
	}
	return camp, nil
}
