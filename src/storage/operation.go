package storage

type Commit struct {
	Description string
	Opts        []Operation
}

type Operation struct {
	OType   int
	BeginAt int
	EndAt   int
	Content string
}

// OType
const (
	Unknown = iota
	Add
	Change
	Delete
)
