package service

import (
	"campfire/auth"
	"campfire/cache"
	"campfire/dao"
	. "campfire/entity"
	"campfire/ws"
)

type TaskService interface {
	CreateTask(queryID uint, task Task) error

	Tasks(queryID uint, projID uint) ([]TaskDTO, error)

	TaskInfo(queryID uint, projID uint, taskID uint) (TaskDTO, error)

	EditTaskInfo(queryID uint, projID uint, task Task) error

	DeleteTask(queryID uint, projID uint, taskID uint) error
}

func NewTaskService() TaskService {
	return taskService{
		mention: SessionServiceContainer,
		query:   dao.ProjectDaoContainer,
		sec:     auth.SecurityInstance,
	}
}

type taskService struct {
	mention *ws.SessionService
	query   dao.ProjectDao
	sec     auth.SecurityGuard
}

func (p taskService) CreateTask(queryID uint, task Task) error {
	if err := p.sec.IsUserHavingTitle(task.ProjID, queryID); err != nil {
		return err
	}
	if err := p.query.AddTask(task); err != nil {
		return err
	}
	if err := cache.StoreTaskToProject(task.ProjID, task); err != nil {
		return err
	}
	task.StartATimer()
	if err := p.mention.NotifyByEvent(&ws.NewTaskEvent{
		TaskDTO: task.DTO(),
	}, ws.NewTaskEventType); err != nil {
		return err
	}
	return nil
}

func (p taskService) Tasks(queryID uint, projID uint) ([]TaskDTO, error) {
	if err := p.sec.IsUserAProjMember(projID, queryID); err != nil {
		return nil, err
	}
	res, err := p.query.TasksOfProject(projID)
	if err != nil {
		return nil, err
	}
	return TasksDTO(res), nil
}

func (p taskService) TaskInfo(queryID uint, projID uint, taskID uint) (TaskDTO, error) {
	if err := p.sec.IsUserAProjMember(projID, queryID); err != nil {
		return TaskDTO{}, err
	}
	res, err := p.query.TaskInfo(taskID)
	if err != nil {
		return TaskDTO{}, err
	}
	return res.DTO(), nil
}

func (p taskService) EditTaskInfo(queryID uint, projID uint, task Task) error {
	if err := p.sec.IsUserHavingTitle(task.ProjID, queryID); err != nil {
		return err
	}
	if err := cache.EditTaskFromCache(task); err != nil {
		return err
	}
	err := p.query.SetTaskInfo(projID, task)
	if err != nil {
		return err
	}
	return nil
}

func (p taskService) DeleteTask(queryID uint, projID uint, taskID uint) error {
	if err := p.sec.IsUserATaskOwner(projID, taskID, queryID); err != nil {
		return err
	}
	err := p.query.DeleteTask(taskID)
	return err
}
