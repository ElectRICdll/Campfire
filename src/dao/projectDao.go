package dao

import . "campfire/entity"

type ProjectDao interface {
	ProjectInfo(queryMemberID ID, projID ID) (Project, error)

	SetProjectName(queryOwnerID, projID ID, name string) error

	AddProject(queryUserID, proj Project) error

	DeleteProject(queryOwnerID, projID ID) error

	MemberInfo(queryMemberID ID, projID ID, userID ID) (Member, error)

	AddMember(queryOwnerID ID, projID ID, userID ID) error

	DeleteMember(queryOwnerID ID, projID ID, userID ID) error

	SetMemberInfo(campID ID, member Member) error

	TasksOfProject(queryMemberID, projID ID) ([]Task, error)

	TaskInfo(queryMemberID ID, projID ID, taskID ID) (Task, error)

	SetTaskInfo(queryOwnerID ID, projID ID, taskID ID) error

	AddTask(queryProjMemberID, projID ID, task Task) error

	DeleteTask(queryOwnerID, projID ID, taskID ID) error

	CampsOfProject(queryMemberID, projID ID) ([]Camp, error)
}
