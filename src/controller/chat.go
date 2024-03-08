package controller

import (
	"campfire/entity"
	"campfire/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ChatController interface {
	Tents(*gin.Context)

	TentInfo(*gin.Context)

	TentMessageRecord(*gin.Context)

	Campsites(*gin.Context)

	CampsiteInfo(*gin.Context)

	CampsiteMessageRecord(*gin.Context)

	EditCampsiteInfo(*gin.Context)

	Projects(*gin.Context)
}

func NewChatController() ChatController {
	return chatController{}
}

type chatController struct {
	userService    service.UserService
	campService    service.CampsiteService
	messageService service.MessageService
}

/*
EditCampsiteInfo
修改群聊信息，只支持基础信息的修改，包括群聊名称，群主转让等。
method: POST
path: /campsites/{campsite_id}/edit
jwt_auth: true
*/
func (c chatController) EditCampsiteInfo(ctx *gin.Context) {
	userId := (int)(ctx.Keys["id"].(float64))
	camp := entity.CampDTO{}
	if err := ctx.BindJSON(&camp); err != nil {
		responseError(ctx, entity.ExternalError{Message: "invalid syntax"})
	}
	if err := c.campService.EditCampsiteInfo(userId, entity.Camp{
		Name:     camp.Name,
		LeaderId: camp.LeaderId,
	}); err != nil {
		responseError(ctx, err)
		return
	}
	return
}

/*
Tents
查看与用户有关联的所有私聊频道
method: GET
path: /user/tents
jwt_auth: true
*/
func (c chatController) Tents(ctx *gin.Context) {
	userId := (int)(ctx.Keys["id"].(float64))

	res, err := c.userService.Tents(userId)
	responseJSON(ctx, res, err)

	return
}

/*
TentInfo
查看与用户有关联的某个私聊频道
method: GET
path: /user/{project_id}/{tent_id}
jwt_auth: true
*/
func (c chatController) TentInfo(ctx *gin.Context) {
	userId := (int)(ctx.Keys["id"].(float64))

	uri := struct {
		PID int `uri:"p_id" binding:"required"`
		TID int `uri:"t_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}

	res, err := c.campService.TentInfo(userId, uri.PID, uri.TID)
	responseJSON(ctx, res, err)

	return
}

/*
TentMessageRecord
更新Tent的消息记录，数量由配置文件决定
method: GET
path: /user/{project_id}/{tent_id}/record
jwt_auth: true
*/
func (c chatController) TentMessageRecord(ctx *gin.Context) {
	userId := (int)(ctx.Keys["id"].(float64))

	begin, err := strconv.Atoi(ctx.Query("begin_at"))
	if err != nil {
		responseError(ctx, err)
		return
	}

	uri := struct {
		PID int `uri:"p_id" binding:"required"`
		TID int `uri:"t_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
	}

	res, err := c.messageService.PullTentMessageRecord(userId, uri.PID, uri.TID, begin)
	responseJSON(ctx, res, err)

	return
}

/*
Campsites
查看与用户有关联的全部群聊
method: GET
path: /user/campsites
jwt_auth: true
*/
func (c chatController) Campsites(ctx *gin.Context) {
	userId := (int)(ctx.Keys["id"].(float64))

	res, err := c.userService.Campsites(userId)
	responseJSON(ctx, res, err)

	return
}

/*
CampsiteInfo
查看与用户有关联的某个群聊
method: GET
path: /user/{campsite_id}
jwt_auth: true
*/
func (c chatController) CampsiteInfo(ctx *gin.Context) {
	userId := (int)(ctx.Keys["id"].(float64))
	uri := struct {
		CID int `uri:"c_id" binding:"required"`
	}{}

	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}

	res, err := c.campService.CampsiteInfo(userId, uri.CID)
	responseJSON(ctx, res, err)

	return
}

/*
CampsiteMessageRecord
更新Campsite的消息记录，数量由配置文件决定
method: GET
path: /user/{campsite_id}/record
jwt_auth: true
*/
func (c chatController) CampsiteMessageRecord(ctx *gin.Context) {
	userId := (int)(ctx.Keys["id"].(float64))

	begin, err := strconv.Atoi(ctx.Query("begin_at"))
	if err != nil {
		responseError(ctx, err)
		return
	}

	uri := struct {
		CID int `uri:"c_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
	}

	res, err := c.messageService.PullCampsiteMessageRecord(userId, uri.CID, begin)
	responseJSON(ctx, res, err)

	return
}

/*
Projects
查看与用户有关联的全部项目
method: GET
path: /user/projects
jwt_auth: true
*/
func (c chatController) Projects(ctx *gin.Context) {
	userId := (int)(ctx.Keys["id"].(float64))

	res, err := c.userService.Projects(userId)
	responseJSON(ctx, res, err)

	return
}
