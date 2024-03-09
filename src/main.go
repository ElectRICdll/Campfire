package main

import (
	"campfire/controller"
	"campfire/dao"
	"campfire/entity"
	"campfire/service"
	"github.com/gin-gonic/gin"
)

func registerHandlers(engine *gin.Engine) {
	auth := service.NewSecurityService()
	login := controller.NewLoginController()
	engine.POST("/login", login.Login)
	engine.POST("/reg", login.Register)

	user := controller.NewUserController()
	engine.GET("/user/:id", auth.AuthMiddleware(), user.UserInfo)
	engine.GET("/search", auth.AuthMiddleware(), user.FindUsersByName)

	session := controller.NewSessionController()
	engine.GET("/ws", auth.AuthMiddleware(), session.NewSession)
}

func main() {
	r := gin.Default()

	db := dao.DB

	err2 := db.AutoMigrate(
		&entity.Project{},
		&entity.Member{},
		&entity.User{},
		&entity.Task{},
		&entity.Camp{},
		&entity.Announcement{},
		&entity.Message{},
	)

	if err2 != nil {
		println(err2)
		return
	}

	registerHandlers(r)

	err := r.Run()
	if err != nil {

	}
}
