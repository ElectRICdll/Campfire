package service

import (
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
	}
}

type taskService struct {
	mention *ws.SessionService
	query   dao.ProjectDao
}

func (p taskService) CreateTask(queryID uint, task Task) error {
	task.StartATimer()
	// TODO
	//cache.StoreTaskInCache()
	err := p.query.AddTask(queryID, task)
	if err != nil {
		return err
	}
	if err := p.mention.NotifyByEvent(&ws.NewTaskEvent{
		TaskDTO: task.DTO(),
	}, ws.NewTaskEventType); err != nil {
		return err
	}
	return nil
}

func (p taskService) Tasks(queryID uint, projID uint) ([]TaskDTO, error) {
	res, err := p.query.TasksOfProject(queryID, projID)
	if err != nil {
		return nil, err
	}
	return TasksDTO(res), nil
}

func (p taskService) TaskInfo(queryID uint, projID uint, taskID uint) (TaskDTO, error) {
	res, err := p.query.TaskInfo(queryID, projID, taskID)
	if err != nil {
		return TaskDTO{}, err
	}
	return res.DTO(), nil
}

func (p taskService) EditTaskInfo(queryID uint, projID uint, task Task) error {
	err := p.query.SetTaskInfo(queryID, projID, task)
	if err != nil {
		return err
	}
	return nil
}

func (p taskService) DeleteTask(queryID uint, projID uint, taskID uint) error {
	err := p.query.DeleteTask(queryID, projID, taskID)
	return err
}
