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
