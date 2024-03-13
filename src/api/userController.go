package api

import (
	"campfire/entity"
	"campfire/service"
	"campfire/util"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	UserInfo(*gin.Context)

	FindUsersByName(*gin.Context)

	EditUserInfo(*gin.Context)

	ChangeEmail(*gin.Context)

	ChangePassword(*gin.Context)

	EmailVerify(*gin.Context)

	PrivateCamps(*gin.Context)

	PublicCamps(*gin.Context)

	Tasks(*gin.Context)

	Projects(*gin.Context)
}

func NewUserController() UserController {
	return userController{
		userService: service.UserServiceContainer,
		campService: nil,
	}
}

type userController struct {
	userService service.UserService
	campService service.CampService
}

func (c userController) EmailVerify(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c userController) ChangeEmail(ctx *gin.Context) {

}

/*
UserInfo
获取某用户的相关信息
method: GET
path: /user/{user_id}
jwt_auth: false
*/
func (c userController) UserInfo(ctx *gin.Context) {
	user := entity.UserDTO{}
	err := ctx.ShouldBindUri(&user)
	if err != nil {
		responseError(ctx, util.ExternalError{Message: "invalid syntax"})
		return
	}

	res, err := c.userService.UserInfo(user.ID)
	responseJSON(ctx, res, err)

	return
}

/*
FindUsersByName
查找用户
method: GET
path: /user/search
jwt_auth: false
query: username
*/
func (c userController) FindUsersByName(ctx *gin.Context) {
	name := ctx.Query("username")
	users, err := c.userService.FindUsersByName(name)

	responseJSON(ctx, users, err)
	return
}

/*
EditUserInfo
用于更改用户信息（除密码外）
method: POST
path: /user/edit
jwt_auth: true
*/
func (c userController) EditUserInfo(ctx *gin.Context) {
	id := (uint)(ctx.Keys["id"].(float64))
	user := entity.UserDTO{ID: id}
	if err := ctx.BindJSON(&user); err != nil {
		responseError(ctx, util.ExternalError{Message: "invalid syntax"})
		return
	}
	if err := c.userService.EditUserInfo(user); err != nil {
		responseError(ctx, err)
		return
	}
	responseSuccess(ctx)
	return
}

/*
ChangePassword
用于更改密码
method: POST
path: /user/edit/p
jwt_auth: true
*/
func (c userController) ChangePassword(ctx *gin.Context) {
	id := (uint)(ctx.Keys["id"].(float64))
	p := struct {
		Password string `json:"p"`
	}{}
	if err := ctx.BindJSON(&p); err != nil {
		responseError(ctx, util.ExternalError{Message: "invalid syntax"})
		return
	}
	if err := c.userService.ChangePassword(id, p.Password); err != nil {
		responseError(ctx, err)
		return
	}
	responseSuccess(ctx)
	return
}

/*
Tasks
查看与用户有关联的所有任务
method: GET
path: /user/tasks/
jwt_auth: true
*/
func (c userController) Tasks(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))

	res, err := c.userService.Tasks(userID)
	responseJSON(ctx, res, err)
	return
}

/*
PrivateCamps
查看与用户有关联的所有私聊频道
method: GET
path: /user/camps/private
jwt_auth: true
*/
func (c userController) PrivateCamps(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))

	res, err := c.userService.PrivateCamps(userID)
	responseJSON(ctx, res, err)

	return
}

/*
PublicCamps
查看与用户有关联的全部群聊
method: GET
path: /user/camps
jwt_auth: true
*/
func (c userController) PublicCamps(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))

	res, err := c.userService.PublicCamps(userID)
	responseJSON(ctx, res, err)

	return
}

/*
CampInfo
查看与用户有关联的某个群聊
method: GET
path: /user/{camp_id}
jwt_auth: true
*/
func (c userController) CampInfo(ctx *gin.Context) {
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
func (c userController) EditCampInfo(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	camp := entity.CampDTO{}
	uri := struct {
		CID uint `uri:"c_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
	}
	if err := ctx.BindJSON(&camp); err != nil {
		responseError(ctx, util.ExternalError{Message: "invalid syntax"})
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
Projects
查看与用户有关联的全部项目
method: GET
path: /user/projects
jwt_auth: true
*/
func (c userController) Projects(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))

	res, err := c.userService.Projects(userID)
	responseJSON(ctx, res, err)

	return
}
