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
		Password string `json:"p"`
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
	newUser := struct {
		entity.UserDTO
		Password string `json:"p"`
	}{}
	if err := ctx.BindJSON(&newUser); err != nil {
		responseError(ctx, err)
		return
	}
	if err := c.loginService.Register(newUser.UserDTO, newUser.Password); err != nil {
		responseError(ctx, err)
		return
	}

	responseSuccess(ctx)
	return
}
