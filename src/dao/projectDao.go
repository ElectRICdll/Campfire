package dao

import . "campfire/entity"

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
}

type projectDao struct{}

func (d projectDao) ProjectInfo(queryMemberID uint, projID uint) (Project, error) {
	var projmember ProjectMember
	var project Project
	var result = DB.Where("UserID = ? and ProjID = ?", queryMemberID, projID).Find(&projmember)
	if result.Error != nil {
		return project, result.Error
	}
	if result == nil {
		return project, ExternalError{}
	}

	result = DB.Where("ID = ?", projID).Find(&project)

	if result.Error != nil {
		return project, result.Error
	}
	if result == nil {
		return project, ExternalError{}
	}
	return project, nil
}

func (d projectDao) SetProjectInfo(queryOwnerID uint, project Project) error {
	result := DB.Where("OwnerID = ?", queryOwnerID).Save(&project)
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	return nil
}

func (d projectDao) AddProject(proj Project) error {
	result := DB.Save(&proj)
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	return nil
}

func (d projectDao) DeleteProject(queryOwnerID, projID uint) error {
	result := DB.Where("OwnerID = ? and ID = ?", queryOwnerID, projID).Delete(&Project{})
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	return nil
}

func (d projectDao) MemberList(queryMemberID uint, projID uint) ([]ProjectMember, error) {
	var projmember []ProjectMember
	var result = DB.Where("UserID = ? and ProjID = ?", queryMemberID, projID).Find(&projmember)
	if result.Error != nil {
		return projmember, result.Error
	}
	if result == nil {
		return projmember, ExternalError{}
	}

	result = DB.Where("ProjID = ?", projID).Find(&projmember)
	if result.Error != nil {
		return projmember, result.Error
	}
	if result == nil {
		return projmember, ExternalError{}
	}
	return projmember, nil
}

func (d projectDao) MemberInfo(queryMemberID uint, projID uint, userID uint) (ProjectMember, error) {
	var projmember ProjectMember
	var result = DB.Where("UserID = ? and ProjID = ?", queryMemberID, projID).Find(&projmember)
	if result.Error != nil {
		return projmember, result.Error
	}
	if result == nil {
		return projmember, ExternalError{}
	}
	result = DB.Where("ProjID = ? and UserID = ?", projID, userID).Find(&projmember)
	if result.Error != nil {
		return projmember, result.Error
	}
	if result == nil {
		return projmember, ExternalError{}
	}
	return projmember, nil
}

func (d projectDao) AddMember(queryOwnerID uint, projID uint, userID uint) error {
	var result = DB.Where("OwnerID = ? and ID = ?", queryOwnerID, projID).Find(&Project{})
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	var projmember = ProjectMember{ProjID: projID, UserID: userID}
	result = DB.Save(&projmember)
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	return nil
}

func (d projectDao) DeleteMember(queryOwnerID uint, projID uint, userID uint) error {
	var result = DB.Where("OwnerID = ? and ID = ?", queryOwnerID, projID).Find(&Project{})
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	result = DB.Where("projID = ? and userID = ?", projID, userID).Delete(&ProjectMember{})
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	return nil
}

func (d projectDao) SetMemberInfo(projID uint, member ProjectMember) error {
	var result = DB.Where("projID = ?", projID).Save(&member)
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	return nil
}

func (d projectDao) TasksOfProject(queryMemberID, projID uint) ([]Task, error) {
	var projmember ProjectMember
	var task []Task
	var result = DB.Where("UserID = ? and ProjID = ?", queryMemberID, projID).Find(&projmember)
	if result.Error != nil {
		return task, result.Error
	}
	if result == nil {
		return task, ExternalError{}
	}
	result = DB.Where("ProjID = ?", projID).Find(&task)
	if result.Error != nil {
		return task, result.Error
	}
	if result == nil {
		return task, ExternalError{}
	}
	return task, nil
}

func (d projectDao) TaskInfo(queryMemberID uint, projID uint, taskID uint) (Task, error) {
	var task Task
	var result = DB.Where("UserID = ? and ProjID = ?", queryMemberID, projID).Find(&ProjectMember{})
	if result.Error != nil {
		return task, result.Error
	}
	if result == nil {
		return task, ExternalError{}
	}
	result = DB.Where("ID = ?", taskID).Find(&task)
	if result.Error != nil {
		return task, result.Error
	}
	if result == nil {
		return task, ExternalError{}
	}
	return task, nil
}

func (d projectDao) SetTaskInfo(queryOwnerID uint, projID uint, task Task) error {
	var result = DB.Where("OwnerID = ? and ID = ?", queryOwnerID, projID).Find(&Project{})
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	result = DB.Where("projID = ?", projID).Save(&task)
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	return nil
}

func (d projectDao) AddTask(queryProjMemberID uint, task Task) error {
	var result = DB.Where("OwnerID = ? and ID = ?", queryProjMemberID, task.ProjID).Find(&Project{})
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	result = DB.Save(&task)
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	return nil
}

func (d projectDao) DeleteTask(queryOwnerID, projID uint, taskID uint) error {
	var result = DB.Where("OwnerID = ? and ID = ?", queryOwnerID, projID).Find(&Project{})
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	result = DB.Where("ID = ?", taskID).Delete(&Task{})
	if result.Error != nil {
		return result.Error
	}
	if result == nil {
		return ExternalError{}
	}
	return nil
}

func (d projectDao) CampsOfProject(queryMemberID, projID uint) ([]Camp, error) {
	var projmember ProjectMember
	var camp []Camp
	var result = DB.Where("UserID = ? and ProjID = ?", queryMemberID, projID).Find(&projmember)
	if result.Error != nil {
		return camp, result.Error
	}
	if result == nil {
		return camp, ExternalError{}
	}
	result = DB.Where("ProjID = ?", projID).Find(&camp)
	if result.Error != nil {
		return camp, result.Error
	}
	if result == nil {
		return camp, ExternalError{}
	}
	return camp, nil
}
