package service

var (
	LoginServiceContainer    = NewLoginService()
	SecurityServiceContainer = NewSecurityService()
	SessionServiceContainer  = NewSessionService()
	UserServiceContainer     = NewUserService()
)
