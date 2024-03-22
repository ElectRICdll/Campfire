package service

import (
	"campfire/dao"
	. "campfire/entity"
	"campfire/ws"
	"time"
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
}

func NewProjectService() ProjectService {
	return projectService{
		query:     dao.ProjectDaoContainer,
		userQuery: dao.UserDaoContainer,
		mention:   SessionServiceContainer,
	}
}

type projectService struct {
	query     dao.ProjectDao
	userQuery dao.UserDao
	mention   *ws.SessionService
}

func (p projectService) CreateProject(project Project) error {
	project.Members = append(project.Members, ProjectMember{
		UserID:    project.OwnerID,
		IsCreator: true,
	})
	project.BeginAt = time.Now()
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
	if err := p.mention.NotifyByEvent(&ws.ProjectInfoChangedEvent{
		ProjectDTO: project.DTO(),
	}, ws.ProjectInfoChangedEventType); err != nil {
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
