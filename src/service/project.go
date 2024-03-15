package service

import (
	"campfire/dao"
	. "campfire/entity"
	wsentity "campfire/entity/ws-entity"
	"campfire/service/ws-service"
	"campfire/storage"
)

type ProjectService interface {
	CreateProject(project Project) error

	ProjectInfo(queryID uint, projectID uint) (BriefProjectDTO, error)

	EditProjectInfo(queryID uint, project Project) error

	DisableProject(queryID uint, projID uint) error

	// UploadProject 暂时搁置
	UploadProject(queryID uint)

	// DownloadProject 暂时搁置
	DownloadProject(queryID uint)

	CreateTask(queryID uint, task Task) error

	Tasks(queryID uint, projID uint) ([]TaskDTO, error)

	TaskInfo(queryID uint, projID uint, taskID uint) (TaskDTO, error)

	EditTaskInfo(queryID uint, projID uint, task Task) error

	DeleteTask(queryID uint, projID uint, taskID uint) error

	FileDetail(queryID uint, projID uint, dirPath string) (storage.File, error)

	UploadFile(queryID uint, projID uint, files ...storage.File)

	DownloadFile(queryID uint, projID uint) ([]storage.File, error)

	DeleteFile(queryID uint, projID uint, filePaths ...string) error

	DirectoryDetail(queryID uint, projID uint, dirPath string) (storage.Dir, error)

	CreateDirectory(queryID uint, projID uint, dirPath string) error

	DeleteDirectory(queryID uint, projID uint, dirPath string) error
}

func NewProjectService() ProjectService {
	return projectService{
		query:   dao.ProjectDaoContainer,
		mention: SessionServiceContainer,
	}
}

type projectService struct {
	query   dao.ProjectDao
	mention *ws_service.SessionService
}

func (p projectService) FileDetail(queryID uint, projID uint, dirPath string) (storage.File, error) {
	//TODO implement me
	panic("implement me")
}

func (p projectService) UploadFile(queryID uint, projID uint, files ...storage.File) {
	//TODO implement me
	panic("implement me")
}

func (p projectService) DownloadFile(queryID uint, projID uint) ([]storage.File, error) {
	//TODO implement me
	panic("implement me")
}

func (p projectService) DeleteFile(queryID uint, projID uint, filePaths ...string) error {
	//TODO implement me
	panic("implement me")
}

func (p projectService) DirectoryDetail(queryID uint, projID uint, dirPath string) (storage.Dir, error) {
	//TODO implement me
	panic("implement me")
}

func (p projectService) CreateDirectory(queryID uint, projID uint, dirPath string) error {
	//TODO implement me
	panic("implement me")
}

func (p projectService) DeleteDirectory(queryID uint, projID uint, dirPath string) error {
	//TODO implement me
	panic("implement me")
}

func (p projectService) CreateProject(project Project) error {
	err := p.query.AddProject(project)
	return err
}

func (p projectService) ProjectInfo(queryID uint, projectID uint) (BriefProjectDTO, error) {
	res, err := p.query.ProjectInfo(queryID, projectID)
	return res.BriefDTO(), err
}

func (p projectService) EditProjectInfo(queryID uint, project Project) error {
	err := p.query.SetProjectInfo(queryID, project)
	if err != nil {
		return err
	}
	if err := p.mention.NotifyByEvent(&wsentity.ProjectInfoChangedEvent{
		ProjectDTO: project.DTO(),
	}, wsentity.ProjectInfoChangedEventType); err != nil {
		return err
	}
	return nil
}

func (p projectService) DisableProject(queryID uint, projID uint) error {
	err := p.query.DeleteProject(queryID, projID)
	return err
}

func (p projectService) UploadProject(queryID uint) {
	//TODO implement me
	panic("implement me")
}

func (p projectService) DownloadProject(queryID uint) {
	//TODO implement me
	panic("implement me")
}

func (p projectService) CreateTask(queryID uint, task Task) error {
	err := p.query.AddTask(queryID, task)
	if err != nil {
		return err
	}
	if err := p.mention.NotifyByEvent(&wsentity.NewTaskEvent{
		TaskDTO: task.DTO(),
	}, wsentity.NewTaskEventType); err != nil {
		return err
	}
	return nil
}

func (p projectService) Tasks(queryID uint, projID uint) ([]TaskDTO, error) {
	res, err := p.query.TasksOfProject(queryID, projID)
	if err != nil {
		return nil, err
	}
	return TasksDTO(res), nil
}

func (p projectService) TaskInfo(queryID uint, projID uint, taskID uint) (TaskDTO, error) {
	res, err := p.query.TaskInfo(queryID, projID, taskID)
	if err != nil {
		return TaskDTO{}, err
	}
	return res.DTO(), nil
}

func (p projectService) EditTaskInfo(queryID uint, projID uint, task Task) error {
	err := p.query.SetTaskInfo(queryID, projID, task)
	if err != nil {
		return err
	}
	return nil
}

func (p projectService) DeleteTask(queryID uint, projID uint, taskID uint) error {
	err := p.query.DeleteTask(queryID, projID, taskID)
	return err
}
