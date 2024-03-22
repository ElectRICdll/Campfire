package api

import (
	"campfire/service"
	"campfire/ws"
	"github.com/gin-gonic/gin"
)

type SessionController interface {
	NewSession(*gin.Context)
}

func NewSessionController() SessionController {
	return &sessionController{
		s: service.SessionServiceContainer,
	}
}

type sessionController struct {
	s *ws.SessionService
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
	if err := c.s.NewSession(ctx.Writer, ctx.Request, nil); err != nil {
		responseError(ctx, err)
		return
	}
}
