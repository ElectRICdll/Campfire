package controller

import "github.com/gin-gonic/gin"

type ProjectController interface {
	CreateProject(*gin.Context)

	ProjectInfo(*gin.Context)

	EditProjectInfo(*gin.Context)

	DisableProject(*gin.Context)

	UploadProject(*gin.Context)

	DownloadProject(*gin.Context)

	CreateTask(*gin.Context)

	Tasks(*gin.Context)

	TaskInfo(*gin.Context)

	EditTaskInfo(*gin.Context)

	DeleteTask(*gin.Context)

	FileCatalogue(*gin.Context)

	FileDetail(*gin.Context)

	UploadFile(*gin.Context)

	DownloadFile(*gin.Context)

	DeleteFile(*gin.Context)

	DirectoryDetail(*gin.Context)

	CreateDirectory(*gin.Context)

	DeleteDirectory(*gin.Context)
}

func NewProjectController() ProjectController {
	return projectController{}
}

type projectController struct {
}

func (p projectController) CreateProject(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) ProjectInfo(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) EditProjectInfo(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) DisableProject(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) UploadProject(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) DownloadProject(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) CreateTask(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) Tasks(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) TaskInfo(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) EditTaskInfo(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) DeleteTask(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) FileCatalogue(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) FileDetail(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) UploadFile(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) DownloadFile(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) DeleteFile(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) DirectoryDetail(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) CreateDirectory(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p projectController) DeleteDirectory(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}
