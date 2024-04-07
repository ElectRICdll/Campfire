package main

import (
	"campfire/api"
	"campfire/auth"
	"campfire/cache"
	"campfire/dao"
	"campfire/entity"
	"campfire/log"
	"campfire/util"

	"github.com/gin-gonic/gin"
)

func registerDependencies(engine *gin.Engine) {
	security := auth.SecurityInstance

	engine.Use(security.CorsMiddleware())

	login := api.NewLoginController()
	engine.POST("/login", login.Login)
	engine.POST("/reg", login.Register)

	session := api.NewSessionController()
	engine.GET("/ws", session.NewSession)

	user := api.NewUserController()
	engine.GET("/user/:user_id", security.AuthMiddleware(), user.UserInfo)
	engine.GET("/user/search", user.FindUsersByName)
	engine.GET("/user/camps/private", security.AuthMiddleware(), user.PrivateCamps)
	engine.GET("/user/camps", security.AuthMiddleware(), user.PublicCamps)
	engine.GET("/user/projects", security.AuthMiddleware(), user.Projects)
	engine.GET("/user/project/:project_id/tasks", security.AuthMiddleware(), user.Tasks)
	engine.POST("/user/edit", security.AuthMiddleware(), user.EditUserInfo)
	engine.POST("/user/edit/p", security.AuthMiddleware(), user.ChangePassword)
	engine.POST("/user/edit/avatar", security.AuthMiddleware(), user.UploadAvatar)
	engine.POST("/user/tasks", security.AuthMiddleware(), user.Tasks)

	proj := api.NewProjectController()
	engine.GET("/project/:project_id", security.AuthMiddleware(), proj.ProjectInfo)
	engine.POST("/project/:project_id/del", security.AuthMiddleware(), proj.DisableProject)
	engine.GET("/project/:project_id/camps", security.AuthMiddleware(), proj.PublicCamps)
	engine.POST("/project/new_proj", security.AuthMiddleware(), proj.CreateProject)
	engine.POST("/project/:project_id/edit", security.AuthMiddleware(), proj.EditProjectInfo)
	engine.POST("/project/:project_id/new_camp", security.AuthMiddleware(), proj.CreateCamp)
	engine.POST("/project/:project_id/invite", security.AuthMiddleware(), proj.InviteMember)
	engine.POST("/project/:project_id/kick", security.AuthMiddleware(), proj.KickMember)

	task := api.NewTaskController()
	engine.POST("/project/:project_id/new_task", security.AuthMiddleware(), task.CreateTask)
	engine.POST("/project/:project_id/tasks/:task_id/edit", security.AuthMiddleware(), task.EditTaskInfo)
	engine.POST("/project/:project_id/tasks/:task_id/del", security.AuthMiddleware(), task.DeleteTask)
	engine.GET("/project/:project_id/tasks/:task_id", security.AuthMiddleware(), task.TaskInfo)
	engine.GET("/project/:project_id/tasks", security.AuthMiddleware(), task.Tasks)

	camp := api.NewCampController()
	engine.GET("/camp/:camp_id", security.AuthMiddleware(), camp.CampInfo)
	engine.GET("/camp/:camp_id/msg", security.AuthMiddleware(), camp.MessageRecord)
	engine.POST("/camp/:camp_id/edit", security.AuthMiddleware(), camp.EditCampInfo)
	engine.POST("/camp/:camp_id/del", security.AuthMiddleware(), camp.DisableCamp)
	engine.POST("/camp/:camp_id/members/invite", security.AuthMiddleware(), camp.InviteMember)
	engine.POST("/camp/:camp_id/members/kick", security.AuthMiddleware(), camp.KickMember)
	engine.POST("/camp/:camp_id/members/edit", security.AuthMiddleware(), camp.EditMyMemberInfo)
	engine.POST("/camp/:camp_id/exit", security.AuthMiddleware(), camp.ExitCamp)
	engine.POST("/camp/:camp_id/promotion", security.AuthMiddleware(), camp.Promotion)
	engine.POST("/camp/:camp_id/demotion", security.AuthMiddleware(), camp.Demotion)
	engine.POST("/camp/:camp_id/own", security.AuthMiddleware(), camp.GiveOwner)
	engine.POST("/camp/:camp_id/title/set", security.AuthMiddleware(), camp.SetTitle)

	git := api.NewGitController()
	engine.GET("/project/:project_id/workplace/:branch/clone", security.AuthMiddleware(), git.Clone)
	engine.GET("/project/:project_id/workplace/:branch/open", security.AuthMiddleware(), git.OpenFile)
	engine.GET("/project/:project_id/workplace/:branch/dir", security.AuthMiddleware(), git.RepoDir)
	engine.POST("/project/:project_id/workplace/:branch/commit", security.AuthMiddleware(), git.Commit)
	engine.POST("/project/:project_id/workplace/:branch/create", security.AuthMiddleware(), git.CreateBranch)
	engine.POST("/project/:project_id/workplace/:branch/rm", security.AuthMiddleware(), git.RemoveBranch)

	engine.GET("/:gitPath/*any", git.GitHTTPBackend)
	engine.POST("/:gitPath/*any", git.GitHTTPBackend)
	engine.PUT("/:gitPath/*any", git.GitHTTPBackend)
	engine.PATCH("/:gitPath/*any", git.GitHTTPBackend)
	engine.DELETE("/:gitPath/*any", git.GitHTTPBackend)
	engine.HEAD("/:gitPath/*any", git.GitHTTPBackend)

	engine.Handle("PROPFIND", "/:gitPath/*any", git.GitHTTPBackend)
	engine.Handle("MKCOL", "/:gitPath/*any", git.GitHTTPBackend)
	engine.Handle("LOCK", "/:gitPath/*any", git.GitHTTPBackend)
	engine.Handle("UNLOCK", "/:gitPath/*any", git.GitHTTPBackend)
	engine.Handle("MOVE", "/:gitPath/*any", git.GitHTTPBackend)

	file := api.NewFileController()
	engine.GET("/user/:user_id/avatar", file.Avatar)
}

func main() {

	db := dao.DB

	if err := db.AutoMigrate(
		&entity.Project{},
		&entity.Release{},
		&entity.User{},
		&entity.Member{},
		&entity.TaskExecutors{},
		&entity.TaskReceivers{},
		&entity.Task{},
		&entity.Camp{},
		&entity.Announcement{},
		&entity.Message{},
		&entity.ProjectMember{},
	); err != nil {
		log.Error(err.Error())
		return
	}

	cache.InitCache()
	cache.InitProjectCache()
	cache.InitCampCache()

	r := gin.Default()

	r.MaxMultipartMemory = 2 << 20
	registerDependencies(r)

	if err := r.Run(":" + util.CONFIG.Port); err != nil {

	}
}
