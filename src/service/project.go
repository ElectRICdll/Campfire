package service

import (
	"campfire/entity"
	"github.com/gin-gonic/gin"
)

type ProjectService interface {
	CreateProject(project entity.Project) error

	ProjectInfo(projectID int) (entity.BriefProjectDTO, error)

	EditProjectInfo(*gin.Context)

	DisableProject(*gin.Context)

	UploadProject(*gin.Context)

	DownloadProject(*gin.Context)

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
