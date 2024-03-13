package storage

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type BoxManager interface {
	AddBox(box Box) error

	RemoveBox(boxName string)

	NewDirectory(dirPath string) error

	RemoveDirectory(name string) error

	ListDirectory(name string) ([]string, error)

	NewFile(filepath string, content []byte) error

	ReadFile(filePath string) ([]byte, error)

	DeleteFile(filePath string) error
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
	Boxes    map[string]Box
}

func (n *nativeManager) AddBox(box Box) error {
	n.Boxes[box.BoxName()] = box
	if err := n.NewDirectory(box.RootPath); err != nil {
		return err
	}
	return nil
}

func (n *nativeManager) RemoveBox(boxName string) {
	delete(n.Boxes, boxName)
}

func (n *nativeManager) NewDirectory(dirPath string) error {
	fullPath := filepath.Join(n.RootPath, dirPath)
	err := os.MkdirAll(fullPath, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
	return nil
}

func (nm *nativeManager) RemoveDirectory(dirPath string) error {
	fullPath := filepath.Join(nm.RootPath, dirPath)
	err := os.RemoveAll(fullPath)
	if err != nil {
		return fmt.Errorf("failed to delete directory: %v", err)
	}
	return nil
}

func (n *nativeManager) ListDirectory(dirPath string) ([]string, error) {
	fullPath := filepath.Join(n.RootPath, dirPath)
	files, err := ioutil.ReadDir(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to list directory: %v", err)
	}
	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames, nil
}

func (n *nativeManager) NewFile(filePath string, content []byte) error {
	fullPath := filepath.Join(n.RootPath, filePath)
	err := ioutil.WriteFile(fullPath, content, 0644)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	return nil
}

func (n *nativeManager) ReadFile(filePath string) ([]byte, error) {
	fullPath := filepath.Join(n.RootPath, filePath)
	content, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}
	return content, nil
}

func (n *nativeManager) DeleteFile(filePath string) error {
	fullPath := filepath.Join(n.RootPath, filePath)
	err := os.Remove(fullPath)
	if err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}
	return nil
}
