package service

import (
	. "campfire/entity"
	"github.com/gin-gonic/gin"
)

type ProjectService interface {
	CreateProject(project Project) error

	ProjectInfo(queryID uint, projectID uint) (BriefProjectDTO, error)

	EditProjectInfo(queryID uint, project Project) error

	DisableProject(queryID uint, projID uint) error

	// UploadProject 暂时搁置
	UploadProject(queryID uint)

	// DownloadProject 暂时搁置
	DownloadProject(queryID uint)

	CreateTask(queryID uint, task Task) error

	Tasks(queryID uint, projID uint) ([]TaskDTO, error)

	TaskInfo(queryID uint, projID uint, taskID uint) (TaskDTO, error)

	EditTaskInfo(queryID uint, projID uint, task Task) error

	DeleteTask(queryID uint, projID uint, taskID uint) error

	// 以下暂时搁置
	FileCatalogue(*gin.Context)

	FileDetail(*gin.Context)

	UploadFile(*gin.Context)

	DownloadFile(*gin.Context)

	DeleteFile(*gin.Context)

	DirectoryDetail(*gin.Context)

	CreateDirectory(*gin.Context)

	DeleteDirectory(*gin.Context)
}
