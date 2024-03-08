package controller

import (
	"campfire/service"
	"github.com/gin-gonic/gin"
)

// MessageController (暂时弃用)
type MessageController interface {
	MessageRecord(*gin.Context)

	FindMessageRecordByKeyword(*gin.Context)

	FindMessageRecordByMember(*gin.Context)

	AllMessageRecord(*gin.Context)
}

func NewMessageController() MessageController {
	return messageController{
		s: service.MessageServiceContainer,
	}
}

type messageController struct {
	s service.MessageService
}

func (c messageController) MessageRecord(ctx *gin.Context) {
	// TODO
	panic("implement me")
}

func (c messageController) FindMessageRecordByKeyword(ctx *gin.Context) {
	// TODO
	panic("implement me")
}

func (c messageController) FindMessageRecordByMember(ctx *gin.Context) {
	// TODO
	panic("implement me")
}

func (c messageController) AllMessageRecord(ctx *gin.Context) {
	// TODO
	panic("implement me")
}
