package api

import (
	"campfire/entity"
	"campfire/service"
	"campfire/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

type UserController interface {
	UserInfo(*gin.Context)

	FindUsersByName(*gin.Context)

	EditUserInfo(*gin.Context)

	UploadAvatar(*gin.Context)

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
	//TODO implement me
	panic("implement me")
}

/*
UserInfo
获取某用户的相关信息
method: GET
path: /user/{user_id}
jwt_auth: false
*/
func (c userController) UserInfo(ctx *gin.Context) {
	userID := struct {
		ID uint `uri:"user_id"`
	}{}

	err := ctx.BindUri(&userID)
	if err != nil {
		responseError(ctx, util.NewExternalError("invalid syntax"))
		return
	}

	res, err := c.userService.UserInfo(userID.ID)
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
	user := entity.User{ID: id}
	if err := ctx.BindJSON(&user); err != nil {
		responseError(ctx, util.NewExternalError("invalid syntax"))
		return
	}
	if err := c.userService.EditUserInfo(user); err != nil {
		responseError(ctx, err)
		return
	}
	responseSuccess(ctx)
	return
}

func (c userController) UploadAvatar(ctx *gin.Context) {
	id := (uint)(ctx.Keys["id"].(float64))
	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		responseError(ctx, err)
		return
	}
	defer file.Close()

	path := fmt.Sprintf("%s/avatar-%d", util.CONFIG.AvatarCacheRootPath, id)

	out, err := os.Create(path)
	if err != nil {
		responseError(ctx, util.NewExternalError("图片解析失败"))
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		responseError(ctx, err)
		return
	}

	if err := c.userService.EditUserInfo(entity.User{ID: id, AvatarUrl: path}); err != nil {
		responseError(ctx, err)
	}

	responseSuccess(ctx)
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
		responseError(ctx, util.NewExternalError("invalid syntax"))
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
		CID uint `uri:"camp_id" binding:"required"`
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
	camp := entity.Camp{}
	uri := struct {
		CID uint `uri:"camp_id" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
	}
	if err := ctx.BindJSON(&camp); err != nil {
		responseError(ctx, util.NewExternalError("invalid syntax"))
	}
	if err := c.campService.EditCampInfo(userID, entity.Camp{
		ID:   uri.CID,
		Name: camp.Name,
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
