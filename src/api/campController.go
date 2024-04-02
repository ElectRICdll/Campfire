package api

import (
	"campfire/auth"
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

	ExitCamp(*gin.Context)

	GiveOwner(*gin.Context)

	Promotion(*gin.Context)

	Demotion(*gin.Context)

	EditMyMemberInfo(*gin.Context)

	SetTitle(*gin.Context)

	MessageRecord(*gin.Context)
}

func NewCampController() CampController {
	return campController{
		campService:    service.CampServiceContainer,
		messageService: service.MessageServiceContainer,
	}
}

type campController struct {
	campService    service.CampService
	messageService service.MessageService
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
		CID uint `uri:"camp_id"`
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
		ID:     camp.ID,
		ProjID: camp.ProjID,
		Name:   camp.Name,
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
		CID uint `uri:"camp_id"`
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
		CID uint `uri:"camp_id"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	user := entity.User{}
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
		CID uint `uri:"camp_id"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	user := entity.User{}
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

func (p campController) EditMyMemberInfo(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	member := entity.Member{}
	uri := struct {
		CID uint `uri:"camp_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	if err := ctx.BindJSON(&member); err != nil {
		responseError(ctx, util.NewExternalError("invalid syntax"))
	}
	if err := p.campService.EditMemberInfo(entity.Member{
		UserID: userID,
		CampID: member.CampID,

		Nickname: member.Nickname,
	}); err != nil {
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
		CID uint `uri:"camp_id" binding:"required"`
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
	proj := entity.Camp{}
	uri := struct {
		PID uint `uri:"project_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	if err := ctx.BindJSON(&proj); err != nil {
		responseError(ctx, util.NewExternalError("invalid syntax"))
	}
	if err := p.campService.EditCampInfo(userID, entity.Camp{
		ID:   uri.PID,
		Name: proj.Name,
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
		CID uint `uri:"camp_id" binding:"required"`
	}{}
	if err := p.campService.DisableCamp(userID, uri.CID); err != nil {
		responseError(ctx, err)
		return
	}
	return
}

func (p campController) SetTitle(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		CID uint `uri:"camp_id" binding:"required"`
	}{}
	body := struct {
		UserID uint   `uri:"userID"`
		Title  string `uri:"title"`
	}{}
	if err := ctx.BindJSON(&body); err != nil {
		responseError(ctx, err)
		return
	}
	if err := p.campService.SetTitle(userID, uri.CID, body.UserID, body.Title); err != nil {
		responseError(ctx, err)
		return
	}
	return
}

func (p campController) MessageRecord(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		CID uint `uri:"camp_id" binding:"required"`
	}{}
	beginID := struct {
		BeginMessageID uint `uri:"beginMessageID"`
	}{}
	if err := ctx.BindJSON(&beginID); err != nil {
		responseError(ctx, err)
		return
	}
	if err := auth.SecurityInstance.IsUserACampMember(uri.CID, userID); err != nil {
		responseError(ctx, util.NewExternalError("access denied"))
	}
	res, err := p.messageService.PullMessageRecord(uri.CID, beginID.BeginMessageID)
	resStruct := struct {
		Msgs []entity.Message `json:"msgs"`
	}{res}
	responseJSON(ctx, resStruct, err)
	return
}

func (p campController) GiveOwner(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		CID uint `uri:"camp_id" binding:"required"`
	}{}
	body := struct {
		UserID uint `json:"userID"`
	}{}
	if err := ctx.BindJSON(&body); err != nil {
		responseError(ctx, err)
		return
	}
	if err := p.campService.GiveOwner(userID, uri.CID, body.UserID); err != nil {
		responseError(ctx, err)
		return
	}
	return
}

func (p campController) Promotion(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		CID uint `uri:"camp_id" binding:"required"`
	}{}
	body := struct {
		UserID uint `json:"userID"`
	}{}
	if err := ctx.BindJSON(&body); err != nil {
		responseError(ctx, err)
		return
	}
	if err := p.campService.GiveRuler(userID, uri.CID, body.UserID); err != nil {
		responseError(ctx, err)
		return
	}
	return
}

func (p campController) Demotion(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		CID uint `uri:"camp_id" binding:"required"`
	}{}
	body := struct {
		UserID uint `json:"userID"`
	}{}
	if err := ctx.BindJSON(&body); err != nil {
		responseError(ctx, err)
		return
	}
	if err := p.campService.DelRuler(userID, uri.CID, body.UserID); err != nil {
		responseError(ctx, err)
		return
	}
	return
}

func (p campController) ExitCamp(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		CID uint `uri:"camp_id" binding:"required"`
	}{}
	if err := p.campService.ExitCamp(userID, uri.CID); err != nil {
		responseError(ctx, err)
		return
	}
	return
}
