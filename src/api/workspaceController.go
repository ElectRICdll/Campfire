package api

import (
	"campfire/service"
	"github.com/gin-gonic/gin"
)

type WorkspaceController interface {
	UploadProject(*gin.Context)

	DownloadProject(*gin.Context)

	FileCatalogue(*gin.Context)

	FileDetail(*gin.Context)

	UploadFile(*gin.Context)

	DownloadFile(*gin.Context)

	DeleteFile(*gin.Context)

	DirectoryDetail(*gin.Context)

	CreateDirectory(*gin.Context)

	DeleteDirectory(*gin.Context)
}

func NewWorkSpaceController() WorkspaceController {
	return workSpaceController{}
}

type workSpaceController struct {
	workService service.ProjectService
}

func (w workSpaceController) UploadProject(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w workSpaceController) DownloadProject(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w workSpaceController) FileCatalogue(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w workSpaceController) FileDetail(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w workSpaceController) UploadFile(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w workSpaceController) DownloadFile(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w workSpaceController) DeleteFile(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w workSpaceController) DirectoryDetail(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w workSpaceController) CreateDirectory(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w workSpaceController) DeleteDirectory(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}
