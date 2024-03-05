package controller

import (
	"campfire/service"
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
	s service.LoginService
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
		"id": string,
		"avatar": url
	}
*/
func (c *loginController) Login(ctx *gin.Context) {
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := ctx.BindJSON(&body); err != nil {
		responseBadRequest(ctx, "invalid syntax")
	}
	res, err := c.s.Login(body.Email, body.Password)
	responseJSON(ctx, res, err)
	return
}

/*
Register
method: POST
path: /reg

	request_body: {
		"email": string,
		"p": string
	}
*/
func (c *loginController) Register(ctx *gin.Context) {
	// TODO
}
