package service

type FileService interface {
	FileDetail(filename string)

	UploadFile(filename string)

	DownloadFile(filename string)
}
