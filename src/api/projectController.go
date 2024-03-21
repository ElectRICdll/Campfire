package api

import (
	"campfire/entity"
	"campfire/service"
	"campfire/util"
	"github.com/gin-gonic/gin"
)

type ProjectController interface {
	CreateProject(*gin.Context)

	ProjectInfo(*gin.Context)

	EditProjectInfo(*gin.Context)

	DisableProject(*gin.Context)

	CreateCamp(*gin.Context)

	PublicCamps(*gin.Context)
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
path: /user/{project_id}/new_camp
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
	camp := entity.CampDTO{}
	if err := ctx.BindJSON(&camp); err != nil {
		responseError(ctx, util.ExternalError{Message: "invalid syntax."})
		return
	}

	if err := p.campService.CreateCamp(userID, entity.Camp{
		Name:    camp.Name,
		OwnerID: userID,
		ProjID:  uri.PID,
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
path: /user/{project_id}/camps
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
		responseError(ctx, util.ExternalError{Message: "invalid syntax."})
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
path: /user/{project_id}/edit
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
		responseError(ctx, util.ExternalError{Message: "invalid syntax"})
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
path: /user/{project_id}/del
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
