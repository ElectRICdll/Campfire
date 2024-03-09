package controller

import (
	"campfire/entity"
	"campfire/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ChatController interface {
	PrivateCamps(*gin.Context)

	PublicCamps(*gin.Context)

	CampInfo(*gin.Context)

	MessageRecord(*gin.Context)

	EditCampInfo(*gin.Context)

	Projects(*gin.Context)
}

func NewChatController() ChatController {
	return chatController{}
}

type chatController struct {
	userService    service.UserService
	campService    service.CampService
	messageService service.MessageService
}

/*
PrivateCamps
查看与用户有关联的所有私聊频道
method: GET
path: /user/private_camps
jwt_auth: true
*/
func (c chatController) PrivateCamps(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))

	res, err := c.userService.PrivateCamps(userID)
	responseJSON(ctx, res, err)

	return
}

/*
PublicCamps
查看与用户有关联的全部群聊
method: GET
path: /user/public_camps
jwt_auth: true
*/
func (c chatController) PublicCamps(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))

	res, err := c.userService.PublicCamps(userID)
	responseJSON(ctx, res, err)

	return
}

/*
CampInfo
查看与用户有关联的某个群聊
method: GET
path: /{project_id}/{campsite_id}
jwt_auth: true
*/
func (c chatController) CampInfo(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		CID uint `uri:"c_id" binding:"required"`
	}{}

	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}

	res, err := c.campService.CampInfo(userID, uri.CID)
	responseJSON(ctx, res, err)

	return
}

/*
EditCampInfo
修改群聊信息，只支持基础信息的修改，包括群聊名称，群主转让等。
method: POST
path: /{project_id}/{campsite_id}/edit
jwt_auth: true
*/
func (c chatController) EditCampInfo(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	camp := entity.CampDTO{}
	uri := struct {
		CID uint `uri:"c_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
	}
	if err := ctx.BindJSON(&camp); err != nil {
		responseError(ctx, entity.ExternalError{Message: "invalid syntax"})
	}
	if err := c.campService.EditCampInfo(userID, entity.Camp{
		ID:      uri.CID,
		Name:    camp.Name,
		OwnerID: camp.OwnerID,
	}); err != nil {
		responseError(ctx, err)
		return
	}
	return
}

/*
MessageRecord
更新Campsite的消息记录，数量由配置文件决定
method: GET
path: /{project_id}/{campsite_id}/record
jwt_auth: true
*/
func (c chatController) MessageRecord(ctx *gin.Context) {
	userID := (int)(ctx.Keys["id"].(float64))

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

	res, err := c.messageService.PullMessageRecord(userID, uri.CID, begin)
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
	userID := (uint)(ctx.Keys["id"].(float64))

	res, err := c.userService.Projects(userID)
	responseJSON(ctx, res, err)

	return
}
