package storage

type Operation interface {
	Execute() error

	Undo() error
}

type Commit struct {
	Description string
	Opts        []Operation
}

type OptAdd struct {
	TargetBox  *Box
	FilePath   string
	BeginPos   int64
	AddContent []byte
}

func (o *OptAdd) Execute() error {
	return o.TargetBox.AddSnippetToFile(o)
}

func (o *OptAdd) Undo() error {
	return o.TargetBox.DeleteSnippetToFile(&OptDel{
		TargetBox:      o.TargetBox,
		FilePath:       o.FilePath,
		BeginPos:       o.BeginPos,
		EndPos:         o.BeginPos + (int64)(len(o.AddContent)),
		DeletedContent: o.AddContent,
	})
}

type OptDel struct {
	TargetBox      *Box
	FilePath       string
	BeginPos       int64
	EndPos         int64
	DeletedContent []byte
}

func (o *OptDel) Execute() error {
	return o.TargetBox.DeleteSnippetToFile(o)
}

func (o *OptDel) Undo() error {
	return o.TargetBox.AddSnippetToFile(&OptAdd{
		TargetBox:  o.TargetBox,
		FilePath:   o.FilePath,
		BeginPos:   o.BeginPos,
		AddContent: o.DeletedContent,
	})
}

type OptReplace struct {
	TargetBox      *Box
	FilePath       string
	BeginPos       int64
	EndPos         int64
	NewContent     []byte
	ReplaceContent []byte
}

func (o *OptReplace) Execute() error {
	return o.TargetBox.ReplaceSnippetToFile(o)
}

func (o *OptReplace) Undo() error {
	return o.TargetBox.ReplaceSnippetToFile(&OptReplace{
		TargetBox:      o.TargetBox,
		FilePath:       o.FilePath,
		BeginPos:       o.BeginPos,
		EndPos:         o.BeginPos + (int64)(len(o.NewContent)),
		NewContent:     o.ReplaceContent,
		ReplaceContent: o.NewContent,
	})
}

type OptDirAdd struct {
	TargetBox *Box
	FilePath  string
	FileName  string
}

func (o *OptDirAdd) Execute() error {
	return o.TargetBox.NewDirectory(o.FilePath, o.FileName)
}

func (o *OptDirAdd) Undo() error {
	return nil
}

type OptDirDel struct {
	TargetBox *Box
	FilePath  string
	FileName  string
}

func (o *OptDirDel) Execute() error {
	return o.TargetBox.RemoveDirectory(o.FilePath + o.FileName)
}

func (o *OptDirDel) Undo() error {
	return nil
}
