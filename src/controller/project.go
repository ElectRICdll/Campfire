package controller

import (
	"campfire/entity"
	"campfire/service"
	"github.com/gin-gonic/gin"
)

type ProjectController interface {
	CreateProject(*gin.Context)

	ProjectInfo(*gin.Context)

	EditProjectInfo(*gin.Context)

	DisableProject(*gin.Context)

	UploadProject(*gin.Context)

	DownloadProject(*gin.Context)

	CreateCamp(*gin.Context)

	PublicCamps(*gin.Context)

	CampInfo(*gin.Context)

	EditCamp(*gin.Context)

	DisableCamp(*gin.Context)

	CreateTask(*gin.Context)

	Tasks(*gin.Context)

	TaskInfo(*gin.Context)

	EditTaskInfo(*gin.Context)

	DeleteTask(*gin.Context)

	FileCatalogue(*gin.Context)

	FileDetail(*gin.Context)

	UploadFile(*gin.Context)

	DownloadFile(*gin.Context)

	DeleteFile(*gin.Context)

	DirectoryDetail(*gin.Context)

	CreateDirectory(*gin.Context)

	DeleteDirectory(*gin.Context)
}

func NewProjectController() ProjectController {
	return projectController{}
}

type projectController struct {
	projService service.ProjectService
	campService service.CampService
}

/*
CreateCamp
创建新群聊接口
method: POST
path: /{project_id}/new_camp
jwt_auth: true
*/
func (p projectController) CreateCamp(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID uint `uri:"p_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	if err := p.projService.DisableProject(userID, uri.PID); err != nil {
		responseError(ctx, err)
		return
	}
	task := entity.CampDTO{}
	if err := ctx.BindJSON(&task); err != nil {
		responseError(ctx, entity.ExternalError{Message: "invalid syntax."})
		return
	}

	if err := p.campService.CreateCamp(userID, uri.PID, entity.Camp{
		Name:    task.Name,
		OwnerID: userID,
	},
	); err != nil {
		responseError(ctx, err)
		return
	}
	return
}

/*
PublicCamps
项目中群聊列表接口
method: GET
path: /{project_id}/camps
jwt_auth: true
*/
func (p projectController) PublicCamps(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID uint `uri:"p_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	res, err := p.campService.PublicCamps(userID, uri.PID)
	responseJSON(ctx, res, err)
	return
}

/*
CampInfo
群聊信息
method: GET
path: /{project_id}/{camp_id}
jwt_auth: true
*/
func (p projectController) CampInfo(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID uint `uri:"p_id" binding:"required"`
		CID uint `uri:"p_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	res, err := p.campService.CampInfo(userID, uri.PID, uri.CID)
	responseJSON(ctx, res, err)
	return
}

/*
EditCamp
群聊编辑接口
method: GET
path: /{project_id}/{camp_id}/edit
jwt_auth: true
*/
func (p projectController) EditCamp(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	proj := entity.CampDTO{}
	uri := struct {
		PID uint `uri:"p_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	if err := ctx.BindJSON(&proj); err != nil {
		responseError(ctx, entity.ExternalError{Message: "invalid syntax"})
	}
	if err := p.campService.EditCampInfo(userID, uri.PID, entity.Camp{
		ID:      uri.PID,
		Name:    proj.Name,
		OwnerID: proj.OwnerID,
	}); err != nil {
		responseError(ctx, err)
		return
	}
	return
}

/*
DisableCamp
群聊编辑接口
method: GET
path: /{project_id}/{camp_id}/del
jwt_auth: true
*/
func (p projectController) DisableCamp(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID uint `uri:"p_id" binding:"required"`
		CID uint `uri:"p_id" binding:"required"`
	}{}
	if err := p.campService.DisableCamp(userID, uri.PID, uri.CID); err != nil {
		responseError(ctx, err)
		return
	}
	return
}

/*
CreateProject
创建新项目接口
method: POST
path: /user/new_proj
jwt_auth: true
*/
func (p projectController) CreateProject(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))

	proj := entity.ProjectDTO{}
	if err := ctx.BindJSON(&proj); err != nil {
		responseError(ctx, entity.ExternalError{Message: "invalid syntax."})
		return
	}

	if err := p.projService.CreateProject(
		entity.Project{
			Title:       proj.Title,
			OwnerID:     userID,
			Description: proj.Description,
			Members:     nil,
			Camps:       nil,
			Tasks:       nil,
			FUrl:        "",
		},
	); err != nil {
		responseError(ctx, err)
	}
}

/*
ProjectInfo
项目信息接口
method: GET
path: /{project_id}
jwt_auth: true
*/
func (p projectController) ProjectInfo(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID uint `uri:"p_id" binding:"required"`
	}{}

	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}

	res, err := p.projService.ProjectInfo(userID, uri.PID)
	responseJSON(ctx, res, err)

	return
}

/*
EditProjectInfo
编辑项目接口
method: POST
path: /{project_id}/edit
jwt_auth: true
*/
func (p projectController) EditProjectInfo(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	proj := entity.ProjectDTO{}
	uri := struct {
		PID uint `uri:"p_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	if err := ctx.BindJSON(&proj); err != nil {
		responseError(ctx, entity.ExternalError{Message: "invalid syntax"})
	}
	if err := p.projService.EditProjectInfo(userID, entity.Project{
		ID:          uri.PID,
		Title:       proj.Title,
		OwnerID:     proj.OwnerID,
		Description: proj.Description,
	}); err != nil {
		responseError(ctx, err)
		return
	}
	return
}

/*
DisableProject
删除项目接口
method: POST
path: /{project_id}/del
jwt_auth: true
*/
func (p projectController) DisableProject(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID uint `uri:"p_id" binding:"required"`
	}{}
	if err := p.projService.DisableProject(userID, uri.PID); err != nil {
		responseError(ctx, err)
		return
	}
	return
}

func (p projectController) UploadProject(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) DownloadProject(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

/*
CreateTask
创建新任务接口
method: POST
path: /{project_id}/new
jwt_auth: true
*/
func (p projectController) CreateTask(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID uint `uri:"p_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	if err := p.projService.DisableProject(userID, uri.PID); err != nil {
		responseError(ctx, err)
		return
	}
	task := entity.TaskDTO{}
	if err := ctx.BindJSON(&task); err != nil {
		responseError(ctx, entity.ExternalError{Message: "invalid syntax."})
	}

	if err := p.projService.CreateTask(userID, entity.Task{
		Title:   task.Title,
		OwnerID: userID,
		//ReceiverID: task.ReceiverID,
		Content: task.Content,
		Begin:   task.Begin,
		End:     task.End,
	},
	); err != nil {
		responseError(ctx, err)
		return
	}
	return
}

/*
Tasks
获取任务接口
method: GET
path: /{project_id}/tasks
jwt_auth: true
*/
func (p projectController) Tasks(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID uint `uri:"p_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	res, err := p.projService.Tasks(userID, uri.PID)
	responseJSON(ctx, res, err)
	return
}

/*
TaskInfo
获取任务接口
method: GET
path: /{project_id}/{task_id}
jwt_auth: true
*/
func (p projectController) TaskInfo(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID uint `uri:"p_id" binding:"required"`
		TID uint `uri:"t_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	res, err := p.projService.TaskInfo(userID, uri.PID, uri.TID)
	responseJSON(ctx, res, err)
	return
}

/*
EditTaskInfo
获取任务接口
method: GET
path: /{project_id}/{task_id}/edit
jwt_auth: true
*/
func (p projectController) EditTaskInfo(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID uint `uri:"p_id" binding:"required"`
		TID uint `uri:"t_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	if err := p.projService.DisableProject(userID, uri.PID); err != nil {
		responseError(ctx, err)
		return
	}
	task := entity.TaskDTO{}
	if err := ctx.BindJSON(&task); err != nil {
		responseError(ctx, entity.ExternalError{Message: "invalid syntax."})
	}

	if err := p.projService.EditTaskInfo(userID, uri.PID, entity.Task{
		ID:      uri.TID,
		Title:   task.Title,
		OwnerID: userID,
		//ReceiverID: task.ReceiverID,
		Content: task.Content,
		Begin:   task.Begin,
		End:     task.End,
	}); err != nil {
		responseError(ctx, err)
		return
	}
	return
}

/*
DeleteTask
获取任务接口
method: POST
path: /{project_id}/{task_id}/del
jwt_auth: true
*/
func (p projectController) DeleteTask(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID uint `uri:"p_id" binding:"required"`
		TID uint `uri:"t_id" binding:"required"`
	}{}
	if err := p.projService.DeleteTask(userID, uri.PID, uri.TID); err != nil {
		responseError(ctx, err)
		return
	}
	return
}

func (p projectController) FileCatalogue(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) FileDetail(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) UploadFile(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) DownloadFile(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) DeleteFile(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) DirectoryDetail(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) CreateDirectory(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) DeleteDirectory(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}
