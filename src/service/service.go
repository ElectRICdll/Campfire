package service

import (
	"campfire/ws"
)

var (
	LoginServiceContainer    = NewLoginService()
	MessageServiceContainer  = NewMessageService()
	SecurityServiceContainer = NewSecurityService()
	SessionServiceContainer  = ws.NewSessionService()
	UserServiceContainer     = NewUserService()
	TaskServiceContainer     = NewTaskService()
	ProjectServiceContainer  = NewProjectService()
	CampServiceContainer     = NewCampService()
)
