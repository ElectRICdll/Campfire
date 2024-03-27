package api

import (
	"campfire/entity"
	"campfire/service"
	"campfire/util"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(*gin.Context)

	Register(*gin.Context)
}

func NewLoginController() LoginController {
	return &loginController{
		service.LoginServiceContainer,
	}
}

type loginController struct {
	loginService service.LoginService
}

/*
Login
method: POST
path: /login

	request_body: {
		"email": string,
		"p": string
	}

	response_body: {
		"res": string,
		"id": int,
		"token": string,
		"avatar_url": string
	}
*/
func (c *loginController) Login(ctx *gin.Context) {
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := ctx.BindJSON(&body); err != nil {
		responseError(ctx, util.NewExternalError("invalid syntax"))
	}
	res, err := c.loginService.Login(body.Email, body.Password)
	responseJSON(ctx, res, err)
	return
}

/*
Register
method: POST
path: /reg

	request_body: {
		"email": string,
		"username": string,
		"p": string
	}
*/
func (c *loginController) Register(ctx *gin.Context) {
	var newUser struct {
		entity.User
		Password string `json:"password"`
	}
	if err := ctx.BindJSON(&newUser); err != nil {
		responseError(ctx, err)
		return
	}
	res, err := c.loginService.Register(entity.User{
		Email:    newUser.Email,
		Name:     newUser.Name,
		Password: newUser.Password,
	}, newUser.Password)
	resStruct := struct {
		ID uint `json:"id"`
	}{res}
	responseJSON(ctx, resStruct, err)
	return
}
