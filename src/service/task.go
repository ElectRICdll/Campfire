package service

import (
	"campfire/auth"
	"campfire/cache"
	"campfire/dao"
	. "campfire/entity"
	"campfire/util"
	"campfire/ws"
)

type TaskService interface {
	CreateTask(queryID uint, task Task) (uint, uint, error)

	Tasks(queryID uint, projID uint) ([]Task, error)

	TaskInfo(queryID uint, projID uint, taskID uint) (Task, error)

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

func (p taskService) CreateTask(queryID uint, task Task) (uint, uint, error) {
	if err := p.sec.IsUserHavingTitle(task.ProjID, queryID); err != nil {
		return 0, 0, err
	}
	if ok := util.ValidateTitle(task.Title); !ok {
		return 0, 0, util.NewExternalError("illegal title format")
	}
	projID, userID, err := p.query.AddTask(task)
	if err != nil {
		return 0, 0, err
	}
	task.StartATimer()
	if err := cache.StoreTaskToProject(task.ProjID, task); err != nil {
		return 0, 0, err
	}
	if err := p.mention.NotifyByEvent(&ws.NewTaskEvent{
		Task: task,
	}, ws.NewTaskEventType); err != nil {
		return 0, 0, err
	}
	return projID, userID, nil
}

func (p taskService) Tasks(queryID uint, projID uint) ([]Task, error) {
	if err := p.sec.IsUserAProjMember(projID, queryID); err != nil {
		return nil, err
	}
	res, err := p.query.TasksOfProject(projID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p taskService) TaskInfo(queryID uint, projID uint, taskID uint) (Task, error) {
	if err := p.sec.IsUserAProjMember(projID, queryID); err != nil {
		return Task{}, err
	}
	res, err := p.query.TaskInfo(taskID)
	if err != nil {
		return Task{}, err
	}
	return res, nil
}

func (p taskService) EditTaskInfo(queryID uint, projID uint, task Task) error {
	if err := p.sec.IsUserHavingTitle(task.ProjID, queryID); err != nil {
		return err
	}
	if len(task.Title) != 0 && !util.ValidateTitle(task.Title) {
		return util.NewExternalError("illegal title format")
	}
	//if err := cache.EditTaskFromCache(task); err != nil {
	//	return err
	//}
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
