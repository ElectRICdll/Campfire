package main

import (
	"campfire/api"
	"campfire/dao"
	"campfire/entity"
	"campfire/service"
	"github.com/gin-gonic/gin"
)

func registerDependencies(engine *gin.Engine) {
	auth := service.NewSecurityService()
	login := api.NewLoginController()
	engine.POST("/login", login.Login)
	engine.POST("/reg", login.Register)

	session := api.NewSessionController()
	engine.GET("/ws", auth.AuthMiddleware(), session.NewSession)

	user := api.NewUserController()
	engine.GET("/user/:user_id", auth.AuthMiddleware(), user.UserInfo)
	engine.GET("/user/search", user.FindUsersByName)
	engine.GET("/user/camps/private", user.PrivateCamps)
	engine.GET("/user/camps", user.PublicCamps)
	engine.GET("/user/projects", user.Projects)
	engine.GET("/user/:project_id/tasks", user.Tasks)
	engine.POST("/user/edit", user.EditUserInfo)
	engine.POST("/user/edit/p", user.ChangePassword)

	proj := api.NewProjectController()
	engine.GET("/:project_id", proj.ProjectInfo)
	engine.GET("/:project_id/:task_id", proj.TaskInfo)
	engine.GET("/:project_id/tasks", proj.Tasks)
	engine.GET("/:project_id/camps", proj.PublicCamps)
	engine.POST("/user/new_proj", proj.CreateProject)
	engine.POST("/user/:project_id/edit", proj.EditProjectInfo)
	engine.POST("/user/:project_id/del", proj.DisableProject)
	engine.POST("/user/:project_id/new_task", proj.CreateTask)
	engine.POST("/user/:project_id/:task_id/edit", proj.EditTaskInfo)
	engine.POST("/user/:project_id/:task_id/del", proj.DeleteTask)

	camp := api.NewCampController()
	engine.GET("/user/camps/:camp_id", camp.CampInfo)
	engine.POST("/user/:camp_id/edit", camp.EditCampInfo)
	engine.POST("/user/:camp_id/del", camp.DisableCamp)
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
		&entity.ProjectMember{},
	)

	if err2 != nil {
		println(err2)
		return
	}

	registerDependencies(r)

	err := r.Run()
	if err != nil {

	}
}
