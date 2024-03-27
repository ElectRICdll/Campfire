package storage

type File struct {
	Path        string `json:"path"`
	Name        string `json:"name"`
	IsDirectory bool   `json:"isDirectory"`
}
