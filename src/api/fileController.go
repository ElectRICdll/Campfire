package api

import (
	"campfire/service"
	"campfire/util"
	"github.com/gin-gonic/gin"
	"io"
	"os"
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
	avatar, err := os.Open(res.AvatarUrl)
	if err != nil {
		responseError(ctx, util.NewExternalError("no avatar found"))
	}
	ctx.Header("Content-Type", "application/octet-stream")

	_, err = io.Copy(ctx.Writer, avatar)
	if err != nil {
		ctx.AbortWithStatusJSON(500, err)
		return
	}
	return
}
