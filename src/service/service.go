package service

import "campfire/service/ws-service"

var (
	LoginServiceContainer    = NewLoginService()
	MessageServiceContainer  = NewMessageService()
	SecurityServiceContainer = NewSecurityService()
	SessionServiceContainer  = ws_service.NewSessionService()
	UserServiceContainer     = NewUserService()
	TaskServiceContainer     = NewTaskService()
	ProjectServiceContainer  = NewProjectService()
	CampServiceContainer     = NewCampService()
)
