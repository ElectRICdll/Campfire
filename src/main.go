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
	engine.GET("/user/camps/private", auth.AuthMiddleware(), user.PrivateCamps)
	engine.GET("/user/camps", auth.AuthMiddleware(), user.PublicCamps)
	engine.GET("/user/projects", auth.AuthMiddleware(), user.Projects)
	engine.GET("/user/project/:project_id/tasks", auth.AuthMiddleware(), user.Tasks)
	engine.POST("/user/edit", auth.AuthMiddleware(), user.EditUserInfo)
	engine.POST("/user/edit/p", auth.AuthMiddleware(), user.ChangePassword)

	proj := api.NewProjectController()
	engine.GET("/project/:project_id", proj.ProjectInfo)
	engine.GET("/project/:project_id/:task_id", auth.AuthMiddleware(), proj.TaskInfo)
	engine.GET("/project/:project_id/tasks", auth.AuthMiddleware(), proj.Tasks)
	engine.GET("/project/:project_id/camps", auth.AuthMiddleware(), proj.PublicCamps)
	engine.POST("/project/new_proj", auth.AuthMiddleware(), proj.CreateProject)
	engine.POST("/project/:project_id/edit", auth.AuthMiddleware(), proj.EditProjectInfo)
	engine.POST("/project/:project_id/del", auth.AuthMiddleware(), proj.DisableProject)
	engine.POST("/project/:project_id/new_task", auth.AuthMiddleware(), proj.CreateTask)
	engine.POST("/project/:project_id/:task_id/edit", auth.AuthMiddleware(), proj.EditTaskInfo)
	engine.POST("/project/:project_id/:task_id/del", auth.AuthMiddleware(), proj.DeleteTask)
	engine.POST("/project/:project_id/new_camp", auth.AuthMiddleware(), proj.CreateCamp)

	camp := api.NewCampController()
	engine.GET("/camp/:camp_id", camp.CampInfo)
	engine.POST("/camp/:camp_id/edit", camp.EditCampInfo)
	engine.POST("/camp/:camp_id/del", camp.DisableCamp)
	engine.POST("/camp/:camp_id/add", camp.AddMember)
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
