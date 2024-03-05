package main

import (
	"campfire/controller"
	"campfire/service"
	"github.com/gin-gonic/gin"
)

func registerHandlers(engine *gin.Engine) {
	auth := service.NewSecurityService()
	login := controller.NewLoginController()
	engine.POST("/login", login.Login)
	engine.POST("/reg", login.Register)

	user := controller.NewUserController()
	engine.GET("/user", auth.AuthMiddleware(), user.UserInfo)
	engine.GET("/user/search", auth.AuthMiddleware(), user.FindUsersByName)

	session := controller.NewSessionController()
	engine.GET("/ws", auth.AuthMiddleware(), session.NewSession)
}

func main() {
	r := gin.Default()

	registerHandlers(r)

	r.Run()
}
