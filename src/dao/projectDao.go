package dao

import "campfire/entity"

type ProjectDao interface {
	ProjectInfo(projectId int) (entity.Project, error)

	SetProjectName(projectId int, name string) error

	AddProject(project entity.Project) error

	DeleteProject(projectId int) error

	AddCampsite(campsite entity.Campsite) error

	DeleteCampsite(projectId int, campsiteId int) error

	TaskInfo(projectId int, taskId int) (entity.Task, error)

	TasksInfo(projectId int) ([]entity.Task, error)

	SetTaskInfo(projectId int, taskId int) error

	AddTask(projectId int, task entity.Task) error

	DeleteTask(projectId int, taskId int) error
}
