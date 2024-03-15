package dao

import (
	. "campfire/entity"
	. "campfire/util"

	"gorm.io/gorm"
)

type ProjectDao interface {
	ProjectInfo(queryMemberID uint, projID uint) (Project, error)

	SetProjectInfo(queryOwnerID uint, project Project) error

	AddProject(proj Project) error

	DeleteProject(queryOwnerID, projID uint) error

	MemberList(queryMemberID uint, projID uint) ([]ProjectMember, error)

	MemberInfo(queryMemberID uint, projID uint, userID uint) (ProjectMember, error)

	AddMember(queryOwnerID uint, projID uint, userID uint) error

	DeleteMember(queryOwnerID uint, projID uint, userID uint) error

	SetMemberInfo(campID uint, member ProjectMember) error

	TasksOfProject(queryMemberID, projID uint) ([]Task, error)

	TaskInfo(queryMemberID uint, projID uint, taskID uint) (Task, error)

	SetTaskInfo(queryOwnerID uint, projID uint, task Task) error

	AddTask(queryProjMemberID uint, task Task) error

	DeleteTask(queryOwnerID, projID uint, taskID uint) error

	CampsOfProject(queryMemberID, projID uint) ([]Camp, error)

	IsUserAProjectMember(projID uint, userID uint) (bool, error)
}

type projectDao struct{}

func (d projectDao) IsUserAProjectMember(projID uint, userID uint) (bool, error) {
	panic("wait for implement")
}

func (d projectDao) ProjectInfo(queryMemberID uint, projID uint) (Project, error) {
	var project Project
	var result = DB.Preload("Tasks").Preload("Camps").Preload("Members.User").
		Joins("JOIN project_members ON project_members.project_id = projects.id").
		Joins("JOIN users ON users.id = project_members.user_id").
		Where("projects.id = ? AND users.id = ?", projID, queryMemberID).
		First(&project)
	if result.Error == gorm.ErrRecordNotFound {
		return project, NewExternalError("Access denied.")
	}
	if result.Error != nil {
		return project, result.Error
	}

	return project, nil
}

func (d projectDao) SetProjectInfo(queryOwnerID uint, project Project) error {
	result := DB.Where("OwnerID = ?", queryOwnerID).Save(&project)
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d projectDao) AddProject(proj Project) error {
	result := DB.Save(&proj)
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d projectDao) DeleteProject(queryOwnerID, projID uint) error {
	result := DB.Where("OwnerID = ? AND ID = ?", queryOwnerID, projID).Delete(&Project{})
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d projectDao) MemberList(queryMemberID uint, projID uint) ([]ProjectMember, error) {
	var projmember []ProjectMember
	var result = DB.Where("UserID = ? AND ProjID = ?", queryMemberID, projID).Find(&projmember)
	if result.Error == gorm.ErrRecordNotFound {
		return projmember, ExternalError{}
	}
	if result.Error != nil {
		return projmember, result.Error
	}

	result = DB.Where("ProjID = ?", projID).Find(&projmember)
	if result.Error == gorm.ErrRecordNotFound {
		return projmember, ExternalError{}
	}
	if result.Error != nil {
		return projmember, result.Error
	}
	return projmember, nil
}

func (d projectDao) MemberInfo(queryMemberID uint, projID uint, userID uint) (ProjectMember, error) {
	var projmember ProjectMember
	var result = DB.Where("UserID = ? AND ProjID = ?", queryMemberID, projID).Find(&projmember)
	if result.Error == gorm.ErrRecordNotFound {
		return projmember, ExternalError{}
	}
	if result.Error != nil {
		return projmember, result.Error
	}
	result = DB.Where("ProjID = ? AND UserID = ?", projID, userID).Find(&projmember)
	if result.Error == gorm.ErrRecordNotFound {
		return projmember, ExternalError{}
	}
	if result.Error != nil {
		return projmember, result.Error
	}
	return projmember, nil
}

func (d projectDao) AddMember(queryOwnerID uint, projID uint, userID uint) error {
	var result = DB.Where("OwnerID = ? AND ID = ?", queryOwnerID, projID).Find(&Project{})
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	var projmember = ProjectMember{ProjID: projID, UserID: userID}
	result = DB.Save(&projmember)
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d projectDao) DeleteMember(queryOwnerID uint, projID uint, userID uint) error {
	var result = DB.Where("OwnerID = ? AND ID = ?", queryOwnerID, projID).Find(&Project{})
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	result = DB.Where("projID = ? AND userID = ?", projID, userID).Delete(&ProjectMember{})
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d projectDao) SetMemberInfo(projID uint, member ProjectMember) error {
	var result = DB.Where("projID = ?", projID).Save(&member)
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d projectDao) TasksOfProject(queryMemberID, projID uint) ([]Task, error) {
	var projmember ProjectMember
	var task []Task
	var result = DB.Where("UserID = ? AND ProjID = ?", queryMemberID, projID).Find(&projmember)
	if result.Error == gorm.ErrRecordNotFound {
		return task, ExternalError{}
	}
	if result.Error != nil {
		return task, result.Error
	}
	result = DB.Where("ProjID = ?", projID).Find(&task)
	if result.Error == gorm.ErrRecordNotFound {
		return task, ExternalError{}
	}
	if result.Error != nil {
		return task, result.Error
	}
	return task, nil
}

func (d projectDao) TaskInfo(queryMemberID uint, projID uint, taskID uint) (Task, error) {
	var task Task
	var result = DB.Where("UserID = ? AND ProjID = ?", queryMemberID, projID).Find(&ProjectMember{})
	if result.Error == gorm.ErrRecordNotFound {
		return task, ExternalError{}
	}
	if result.Error != nil {
		return task, result.Error
	}
	result = DB.Where("ID = ?", taskID).Find(&task)
	if result.Error == gorm.ErrRecordNotFound {
		return task, ExternalError{}
	}
	if result.Error != nil {
		return task, result.Error
	}
	return task, nil
}

func (d projectDao) SetTaskInfo(queryOwnerID uint, projID uint, task Task) error {
	var result = DB.Where("OwnerID = ? AND ID = ?", queryOwnerID, projID).Find(&Project{})
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	result = DB.Where("projID = ?", projID).Save(&task)
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d projectDao) AddTask(queryProjMemberID uint, task Task) error {
	var result = DB.Where("OwnerID = ? AND ID = ?", queryProjMemberID, task.ProjID).Find(&Project{})
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	result = DB.Save(&task)
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d projectDao) DeleteTask(queryOwnerID, projID uint, taskID uint) error {
	var result = DB.Where("OwnerID = ? AND ID = ?", queryOwnerID, projID).Find(&Project{})
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	result = DB.Where("ID = ?", taskID).Delete(&Task{})
	if result.Error == gorm.ErrRecordNotFound {
		return ExternalError{}
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d projectDao) CampsOfProject(queryMemberID, projID uint) ([]Camp, error) {
	var projmember ProjectMember
	var camp []Camp
	var result = DB.Where("UserID = ? AND ProjID = ?", queryMemberID, projID).Find(&projmember)
	if result.Error == gorm.ErrRecordNotFound {
		return camp, ExternalError{}
	}
	if result.Error != nil {
		return camp, result.Error
	}
	result = DB.Where("ProjID = ?", projID).Find(&camp)
	if result.Error == gorm.ErrRecordNotFound {
		return camp, ExternalError{}
	}
	if result.Error != nil {
		return camp, result.Error
	}
	return camp, nil
}
