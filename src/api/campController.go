package api

import (
	"campfire/entity"
	"campfire/service"
	"campfire/util"
	"github.com/gin-gonic/gin"
)

type CampController interface {
	CampInfo(*gin.Context)

	EditCampInfo(*gin.Context)

	DisableCamp(*gin.Context)

	InviteMember(*gin.Context)

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

/*
EditCampInfo
method: POST
path: /camp/:camp_id/edit
jwt_auth: true
*/
func (p campController) EditCampInfo(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		CID uint `json:"c_id"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	camp := entity.BriefCampDTO{}
	if err := ctx.BindJSON(&camp); err != nil {
		responseError(ctx, err)
		return
	}
	if err := p.campService.EditCampInfo(userID, entity.Camp{
		ID:      camp.ID,
		OwnerID: camp.OwnerID,
		ProjID:  camp.ProjID,
		Name:    camp.Name,
	}); err != nil {
		responseError(ctx, err)
		return
	}
	responseSuccess(ctx)
	return
}

func (p campController) DisableCamp(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		CID uint `json:"c_id"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	if err := p.campService.DisableCamp(userID, uri.CID); err != nil {
		responseError(ctx, err)
		return
	}
	responseSuccess(ctx)
	return
}

func (p campController) InviteMember(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		CID uint `json:"c_id"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	user := entity.UserDTO{}
	if err := ctx.BindJSON(&user); err != nil {
		responseError(ctx, err)
		return
	}
	if err := p.campService.InviteMember(userID, uri.CID, user.ID); err != nil {
		responseError(ctx, err)
		return
	}
	responseSuccess(ctx)
	return
}

func (p campController) KickMember(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		CID uint `json:"c_id"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	user := entity.UserDTO{}
	if err := ctx.BindJSON(&user); err != nil {
		responseError(ctx, err)
		return
	}
	if err := p.campService.KickMember(userID, uri.CID, user.ID); err != nil {
		responseError(ctx, err)
		return
	}
	responseSuccess(ctx)
	return
}

func (p campController) EditMemberInfo(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	member := entity.MemberDTO{}
	uri := struct {
		CID uint `uri:"p_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	if err := ctx.BindJSON(&member); err != nil {
		responseError(ctx, util.ExternalError{Message: "invalid syntax"})
	}
	if err := p.campService.EditMemberTitle(uri.CID, userID, member.Title); err != nil {
		responseError(ctx, err)
		return
	}
	return
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
		responseError(ctx, util.ExternalError{Message: "invalid syntax"})
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
