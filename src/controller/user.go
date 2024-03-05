package controller

import (
	"campfire/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController interface {
	UserInfo(ctx *gin.Context)

	FindUsersByName(ctx *gin.Context)
}

func NewUserController() UserController {
	return &userController{
		service.UserServiceContainer,
	}
}

type userController struct {
	s service.UserService
}

/*
UserInfo
获取某用户的相关信息
method: GET
path: /user
query: id
*/
func (c *userController) UserInfo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"e": err.Error(),
		})
		return
	}

	res, err := c.s.UserInfo(id)
	responseJSON(ctx, res, err)

	return
}

/*
FindUsersByName
查找用户
method: GET
path: /user/search
query: name
*/
func (c *userController) FindUsersByName(ctx *gin.Context) {
	name := ctx.Query("name")
	users, err := c.s.FindUsersByName(name)

	responseJSON(ctx, users, err)
	return
}

/*
ChangeUserInfo
用于更改用户信息（除密码外）
method: POST
path: /user/commit_change
*/
//func (c *userController) ChangeUserInfo(ctx *gin.Context) {
//	id := ctx.PostForm("id")
//	signature := ctx.PostForm("signature")
//	user := entity.User{}
//}

func (c *userController) ChangePassword(ctx *gin.Context) {

}
