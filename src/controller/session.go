package controller

import (
	"campfire/service"
	"github.com/gin-gonic/gin"
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
	id := (int)(ctx.Keys["id"].(float64))
	if err := c.s.NewSession(ctx.Writer, ctx.Request, nil, id); err != nil {
		responseError(ctx, err)
		return
	}
}
