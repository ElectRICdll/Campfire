package service

import (
	"campfire/ws"
)

var (
	LoginServiceContainer   = NewLoginService()
	MessageServiceContainer = NewMessageService()
	SessionServiceContainer = ws.NewSessionService()
	UserServiceContainer    = NewUserService()
	TaskServiceContainer    = NewTaskService()
	ProjectServiceContainer = NewProjectService()
	CampServiceContainer    = NewCampService()
)
