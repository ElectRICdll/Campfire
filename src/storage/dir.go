package storage

type Dir struct {
	Path      string `json:"path"`
	Catalogue []File
}
