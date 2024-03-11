package api

import (
	"campfire/entity"
	"campfire/service"
	"github.com/gin-gonic/gin"
)

type CampController interface {
	CampInfo(*gin.Context)

	EditCampInfo(*gin.Context)

	DisableCamp(*gin.Context)

	AddMember(*gin.Context)

	KickMember(*gin.Context)

	EditMemberInfo(*gin.Context)
}

func NewCampController() CampController {
	return campController{
		campService: nil,
	}
}

type campController struct {
	campService service.CampService
}

func (p campController) EditCampInfo(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p campController) DisableCamp(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p campController) AddMember(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p campController) KickMember(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p campController) EditMemberInfo(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

/*
CampInfo
群聊信息
method: GET
path: /user/{camp_id}
jwt_auth: true
*/
func (p campController) CampInfo(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID uint `uri:"p_id" binding:"required"`
		CID uint `uri:"c_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	res, err := p.campService.CampInfo(userID, uri.CID)
	responseJSON(ctx, res, err)
	return
}

/*
EditCamp
群聊编辑接口
method: GET
path: /user/{camp_id}/edit
jwt_auth: true
*/
func (p campController) EditCamp(ctx *gin.Context) {
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
	if err := p.campService.EditCampInfo(userID, entity.Camp{
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
path: /user/{camp_id}/del
jwt_auth: true
*/
func (p projectController) DisableCamp(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID uint `uri:"p_id" binding:"required"`
		CID uint `uri:"c_id" binding:"required"`
	}{}
	if err := p.campService.DisableCamp(userID, uri.CID); err != nil {
		responseError(ctx, err)
		return
	}
	return
}