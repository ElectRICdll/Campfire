package api

import "github.com/gin-gonic/gin"

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
