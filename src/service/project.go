package service

import (
	"campfire/dao"
	. "campfire/entity"
	wsentity "campfire/entity/ws-entity"
	ws_service "campfire/service/ws-service"
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
	mention   *ws_service.SessionService
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
	if bl, err := securityService.IsUserAProjMember(securityService{}, projectID, queryID); bl == false {
		return BriefProjectDTO{}, err
	}
	res, err := p.query.ProjectInfo(queryID, projectID)
	return res.BriefDTO(), err
}

func (p projectService) EditProjectInfo(queryID uint, project Project) error {
	if bl, err := securityService.IsUserAProjLeader(securityService{}, project.ID, queryID); bl == false {
		return err
	}
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
	if bl, err := securityService.IsUserAProjLeader(securityService{}, projID, queryID); bl == false {
		return err
	}
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
