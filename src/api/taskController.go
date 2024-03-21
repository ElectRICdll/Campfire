package api

import (
	"campfire/entity"
	"campfire/service"
	"campfire/util"
	"github.com/gin-gonic/gin"
)

type TaskController interface {
	CreateTask(*gin.Context)

	Tasks(*gin.Context)

	TaskInfo(*gin.Context)

	EditTaskInfo(*gin.Context)

	DeleteTask(*gin.Context)
}

func NewTaskController() TaskController {
	return taskController{
		service: service.TaskServiceContainer,
	}
}

type taskController struct {
	service service.TaskService
}

/*
CreateTask
创建新任务接口
method: POST
path: /{project_id}/new_task
jwt_auth: true
*/
func (p taskController) CreateTask(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID uint `uri:"p_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	task := entity.TaskDTO{}
	if err := ctx.BindJSON(&task); err != nil {
		responseError(ctx, util.ExternalError{Message: "invalid syntax."})
	}

	if err := p.service.CreateTask(userID, entity.Task{
		Title:   task.Title,
		OwnerID: userID,
		//ReceiversID: task.ReceiversID,
		Content: task.Content,
		BeginAt: task.BeginAt,
		EndAt:   task.EndAt,
	}); err != nil {
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
func (p taskController) Tasks(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID uint `uri:"p_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	res, err := p.service.Tasks(userID, uri.PID)
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
func (p taskController) TaskInfo(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID uint `uri:"p_id" binding:"required"`
		TID uint `uri:"t_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	res, err := p.service.TaskInfo(userID, uri.PID, uri.TID)
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
func (p taskController) EditTaskInfo(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID uint `uri:"p_id" binding:"required"`
		TID uint `uri:"t_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	task := entity.TaskDTO{}
	if err := ctx.BindJSON(&task); err != nil {
		responseError(ctx, util.ExternalError{Message: "invalid syntax."})
	}

	if err := p.service.EditTaskInfo(userID, uri.PID, entity.Task{
		ID:      uri.TID,
		Title:   task.Title,
		OwnerID: userID,
		//ReceiversID: task.ReceiversID,
		Content: task.Content,
		BeginAt: task.BeginAt,
		EndAt:   task.EndAt,
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
func (p taskController) DeleteTask(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID uint `uri:"p_id" binding:"required"`
		TID uint `uri:"t_id" binding:"required"`
	}{}
	if err := p.service.DeleteTask(userID, uri.PID, uri.TID); err != nil {
		responseError(ctx, err)
		return
	}
	return
}
