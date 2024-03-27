package service

import (
	"campfire/auth"
	"campfire/dao"
	. "campfire/entity"
	"campfire/util"
	"campfire/ws"
	"time"
)

type ProjectService interface {
	CreateProject(userID uint, project Project, users ...uint) (uint, error)

	ProjectInfo(queryID, projectID uint) (Project, error)

	EditProjectInfo(queryID uint, project Project) error

	DisableProject(queryID, projID uint) error

	GiveOwner(queryID, projID, userID uint) error

	GiveTitle(queryID, projID, userID uint, title string) error

	RemoveTitle(queryID, projID, userID uint) error

	InviteMember(queryID, projID, userID uint) error

	KickMember(queryID, projID, userID uint) error
}

func NewProjectService() ProjectService {
	return projectService{
		query:     dao.ProjectDaoContainer,
		userQuery: dao.UserDaoContainer,
		mention:   SessionServiceContainer,
		sec:       auth.SecurityInstance,
	}
}

type projectService struct {
	query     dao.ProjectDao
	userQuery dao.UserDao
	mention   *ws.SessionService
	sec       auth.SecurityGuard
}

func (p projectService) CreateProject(userID uint, project Project, usersID ...uint) (uint, error) {
	if ok := util.ValidateTitle(project.Title); !ok {
		return 0, util.NewExternalError("Illegal title format")
	}
	project.BeginAt = time.Now()
	res, err := p.query.AddProject(userID, project, usersID...)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (p projectService) ProjectInfo(queryID uint, projectID uint) (Project, error) {
	if err := p.sec.IsUserAProjMember(projectID, queryID); err != nil {
		return Project{}, err
	}
	res, err := p.query.ProjectInfo(
		projectID,
		"Branches",
		"Releases",
		"Owner",
		"Members.User",
		"Camps",
		"Tasks",
	)
	return res, err
}

func (p projectService) EditProjectInfo(queryID uint, project Project) error {
	if len(project.Title) != 0 && !util.ValidateTitle(project.Title) {
		return util.NewExternalError("Illegal title format")
	}
	if err := p.sec.IsUserAProjLeader(project.ID, queryID); err != nil {
		return err
	}
	err := p.query.SetProjectInfo(project)
	if err != nil {
		return err
	}
	if err := p.mention.NotifyByEvent(&ws.ProjectInfoChangedEvent{
		Project: project,
	}, ws.ProjectInfoChangedEventType); err != nil {
		return err
	}
	return nil
}

func (p projectService) DisableProject(queryID uint, projID uint) error {
	if err := p.sec.IsUserAProjLeader(projID, queryID); err != nil {
		return err
	}
	err := p.query.DeleteProject(projID)
	return err
}

func (p projectService) InviteMember(queryID, projID, userID uint) error {
	if err := p.sec.IsUserAProjLeader(projID, queryID); err != nil {
		return err
	}
	if err := p.sec.IsUserAProjMember(projID, userID); err != nil {
		return util.NewExternalError("user has already been in project")
	}
	if err := p.query.AddMember(ProjectMember{UserID: userID, ProjID: projID}); err != nil {
		return err
	}
	res, err := p.query.ProjectInfo(projID)
	if err != nil {
		return err
	}
	if err := p.mention.NotifyByEvent(ws.NewProjectInvitationEvent(res.BriefDTO(), userID), ws.ProjectInvitationEventType); err != nil {
		return err
	}
	return nil
}

func (p projectService) KickMember(queryID, projID, userID uint) error {
	if err := p.sec.IsUserAProjLeader(projID, queryID); err != nil {
		return err
	}
	if err := p.query.DeleteMember(userID, projID); err != nil {
		return err
	}
	res, err := p.query.ProjectInfo(projID)
	if err != nil {
		return err
	}
	if err := p.mention.NotifyByEvent(&ws.ProjectInfoChangedEvent{
		Project: res,
	}, ws.ProjectInfoChangedEventType); err != nil {
		return err
	}
	return nil
}

func (p projectService) GiveOwner(queryID, projID, userID uint) error {
	if err := p.sec.IsUserAProjLeader(projID, queryID); err != nil {
		return err
	}
	if err := p.query.SetOwner(projID, userID); err != nil {
		return err
	}
	return nil
}

func (p projectService) GiveTitle(queryID, projID, userID uint, title string) error {
	if err := p.sec.IsUserAProjLeader(projID, queryID); err != nil {
		return err
	}
	if err := p.query.SetMemberInfo(ProjectMember{
		UserID: userID,
		ProjID: projID,
		Title:  title,
	}); err != nil {
		return err
	}
	return nil
}

func (p projectService) RemoveTitle(queryID, projID, userID uint) error {
	if err := p.sec.IsUserAProjLeader(projID, queryID); err != nil {
		return err
	}
	if err := p.query.RemoveTitle(projID, userID); err != nil {
		return err
	}
	return nil
}
