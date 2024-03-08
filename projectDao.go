package dao

import "campfire/entity"

type ProjectDao interface {
	ProjectInfo(queryMemberID int, projID int) (entity.Project, error)

	SetProjectName(queryOwnerID, projID int, name string) error

	AddProject(queryUserID, proj entity.Project) error

	DeleteProject(queryOwnerID, projID int) error

	MemberInfo(queryMemberID int, projID int, userID int) (entity.Member, error)

	AddMember(queryOwnerID int, projID int, userID int) error

	DeleteMember(queryOwnerID int, projID int, userID int) error

	SetMemberInfo(campID int, member entity.Member) error

	TasksOfProject(queryMemberID, projID int) ([]entity.Task, error)

	TasksOfUser(userID int) ([]entity.Task, error)

	TaskInfo(queryMemberID int, projID int, taskID int) (entity.Task, error)

	SetTaskInfo(queryOwnerID int, projID int, taskID int) error

	AddTask(queryProjMemberID, projID int, task entity.Task) error

	DeleteTask(queryOwnerID, projID int, taskID int) error

	CampsOfProject(queryMemberID, projID int) ([]entity.Camp, error)
}
