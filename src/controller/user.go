package controller

import (
	"campfire/entity"
	"campfire/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	UserInfo(*gin.Context)

	FindUsersByName(*gin.Context)

	EditUserInfo(*gin.Context)

	ChangeEmail(*gin.Context)

	ChangePassword(*gin.Context)

	EmailVerify(*gin.Context)
}

func NewUserController() UserController {
	return &userController{
		service.UserServiceContainer,
	}
}

type userController struct {
	s service.UserService
}

func (c *userController) EmailVerify(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c *userController) ChangeEmail(ctx *gin.Context) {

}

/*
UserInfo
获取某用户的相关信息
method: GET
path: /user/{user_id}
jwt_auth: false
*/
func (c *userController) UserInfo(ctx *gin.Context) {
	user := entity.UserDTO{}
	err := ctx.ShouldBindUri(&user)
	if err != nil {
		responseError(ctx, entity.ExternalError{Message: "invalid syntax"})
		return
	}

	res, err := c.s.UserInfo(user.ID)
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
func (c *userController) FindUsersByName(ctx *gin.Context) {
	name := ctx.Query("username")
	users, err := c.s.FindUsersByName(name)

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
func (c *userController) EditUserInfo(ctx *gin.Context) {
	id := (uint)(ctx.Keys["id"].(float64))
	user := entity.UserDTO{ID: id}
	if err := ctx.BindJSON(&user); err != nil {
		responseError(ctx, entity.ExternalError{Message: "invalid syntax"})
		return
	}
	if err := c.s.EditUserInfo(user); err != nil {
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
path: /user/change/p
jwt_auth: true
*/
func (c *userController) ChangePassword(ctx *gin.Context) {
	id := (uint)(ctx.Keys["id"].(float64))
	p := struct {
		Password string `json:"p"`
	}{}
	if err := ctx.BindJSON(&p); err != nil {
		responseError(ctx, entity.ExternalError{Message: "invalid syntax"})
		return
	}
	if err := c.s.ChangePassword(id, p.Password); err != nil {

	}
}
