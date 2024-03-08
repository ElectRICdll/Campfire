package service

var (
	LoginServiceContainer    = NewLoginService()
	MessageServiceContainer  = NewMessageService()
	SecurityServiceContainer = NewSecurityService()
	SessionServiceContainer  = NewSessionService()
	UserServiceContainer     = NewUserService()
)
