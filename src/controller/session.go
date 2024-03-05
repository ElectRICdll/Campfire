package controller

import (
	"campfire/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SessionController interface {
	NewSession(*gin.Context)
}

func NewSessionController() SessionController {
	return &sessionController{
		s: service.NewSessionService(),
	}
}

type sessionController struct {
	s service.SessionService
}

/*
NewSession
创建新的ws会话
method: GET
path: /ws (WebSocket)

	request_body {
		"token": string
	}
*/
func (c *sessionController) NewSession(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		responseBadRequest(ctx, "invalid syntax")
	}

	if err := c.s.NewSession(ctx.Writer, ctx.Request, nil, id); err != nil {
		responseInternalError(ctx, err)
		return
	}
}
