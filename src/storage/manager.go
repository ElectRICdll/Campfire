package storage

import (
	"errors"
)

type BoxManager interface {
	FindBox(projID uint)

	AddBox(box Box) error

	RemoveBox(projID uint)

	Push(projID uint, commits []*Commit)

	// Pull TODO
	Pull(projID uint) [][]*Commit

	Clone(projID uint)

	Merge()
}

func NewBoxManager(arg string) (BoxManager, error) {
	if arg == "native" {
		return &nativeManager{}, nil
	} else if arg == "oss" {
		return nil, nil
	}
	return nil, errors.New("invalid arg in creating manager")
}

type ossManager struct {
}

type nativeManager struct {
	RootPath string
	Boxes    map[uint]Box
}

func (n *nativeManager) FindBox(projID uint) {
	//TODO implement me
	panic("implement me")
}

func (n *nativeManager) Push(projID uint, commits []*Commit) {
	for _, commit := range commits {
		for _, opt := range commit.Opts {
			if err := opt.Execute(); err != nil {
				// TODO
			}
		}
	}
}

func (n *nativeManager) Pull(projID uint) [][]*Commit {
	//TODO implement me
	panic("implement me")
}

func (n *nativeManager) Clone(projID uint) {
	//TODO implement me
	panic("implement me")
}

func (n *nativeManager) Merge() {
	//TODO implement me
	panic("implement me")
}

func (n *nativeManager) AddBox(box Box) error {
	n.Boxes[box.ProjID] = box
	// TODO
	if err := box.NewDirectory(box.RootPath, box.BoxName()); err != nil {
		return err
	}
	return nil
}

func (n *nativeManager) RemoveBox(projID uint) {
	delete(n.Boxes, projID)
}
