package api

import (
	"campfire/service"
	"campfire/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strconv"
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
	w := ctx.Writer
	w.Header().Set("Content-Type", "multipart/form-data")

	res, err := f.userService.UserInfo(userID.ID)
	avatar, err := os.Open(res.AvatarUrl)
	if err != nil {
		responseError(ctx, util.NewExternalError("no avatar found"))
	}
	defer avatar.Close()

	avatarInfo, err := avatar.Stat()
	if err != nil {
		responseError(ctx, err)
		return
	}

	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", res.AvatarUrl))
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Length", strconv.FormatInt(avatarInfo.Size(), 10))
	ctx.Header("X-Content-Type-Options", "nosniff")

	buffer := make([]byte, 1024*1024)
	for {
		n, err := avatar.Read(buffer)
		if err != nil && err != io.EOF {
			responseError(ctx, err)
			return
		}

		if n == 0 {
			break
		}

		ctx.Writer.Write(buffer[:n])
		ctx.Writer.Flush()
	}
	return
}
