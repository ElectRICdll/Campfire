package api

import (
	"campfire/service"
	"campfire/util"
	"github.com/gin-gonic/gin"
)

type FileController interface {
	Upload(*gin.Context)

	Avatar(*gin.Context)
}

func NewFileController() FileController {
	return fileController{
		userService: service.UserServiceContainer,
	}
}

type fileController struct {
	userService service.UserService
}

func (f fileController) Upload(ctx *gin.Context) {

}

func (f fileController) Avatar(ctx *gin.Context) {
	userID := struct {
		ID uint `uri:"user_id"`
	}{}

	err := ctx.BindUri(&userID)
	if err != nil {
		responseError(ctx, util.NewExternalError("invalid syntax"))
		return
	}

	res, err := f.userService.UserInfo(userID.ID)
	avatar, err := util.FileToBase64(res.AvatarUrl)
	if err != nil {
		avatar = ""
	}
	responseJSON(ctx, avatar, err)
	return
}
