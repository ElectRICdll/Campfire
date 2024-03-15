package storage

type File struct {
	Path string `json:"path"`
	Name string `json:"name"`
	Data []byte `json:"data"`
}
