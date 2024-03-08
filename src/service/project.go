package service

import (
	. "campfire/entity"
	"github.com/gin-gonic/gin"
)

type ProjectService interface {
	CreateProject(project Project) error

	ProjectInfo(queryID ID, projectID ID) (BriefProjectDTO, error)

	EditProjectInfo(queryID ID, project Project) error

	DisableProject(queryID ID, projID ID) error

	// UploadProject 暂时搁置
	UploadProject(queryID ID)

	// DownloadProject 暂时搁置
	DownloadProject(queryID ID)

	CreateTask(queryID ID, task Task) error

	Tasks(queryID ID, projID ID) ([]TaskDTO, error)

	TaskInfo(queryID ID, projID ID, taskID ID) (TaskDTO, error)

	EditTaskInfo(queryID ID, projID ID, task Task) error

	DeleteTask(queryID ID, projID ID, taskID ID) error

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
