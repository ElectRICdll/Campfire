package api

import (
	"campfire/service"
	"campfire/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
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

	res, err := f.userService.UserInfo(userID.ID)
	if err != nil {
		responseError(ctx, util.NewExternalError("no avatar found"))
		return
	}

	avatar, err := os.Open(res.AvatarUrl)
	if err != nil {
		responseError(ctx, util.NewExternalError("avatar not found"))
		return
	}
	defer avatar.Close()

	buffer := make([]byte, 512)
	_, err = avatar.Read(buffer)
	if err != nil {
		responseError(ctx, err)
		return
	}

	mimeType := http.DetectContentType(buffer)
	if mimeType == "application/octet-stream" {
		mimeType = "image/jpeg"
	}

	avatarInfo, err := avatar.Stat()
	if err != nil {
		responseError(ctx, err)
		return
	}

	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", res.AvatarUrl))
	ctx.Header("Content-Type", mimeType)
	ctx.Header("Content-Length", strconv.FormatInt(avatarInfo.Size(), 10))
	ctx.Writer.WriteHeader(http.StatusOK) // 写入响应头

	// 将文件头重新写入响应体
	_, err = ctx.Writer.Write(buffer)
	if err != nil {
		responseError(ctx, err)
		return
	}

	// 写入剩余的文件内容
	_, err = io.Copy(ctx.Writer, avatar)
	if err != nil {
		responseError(ctx, err)
		return
	}
}
