package storage

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Box struct {
	Name            string
	ProjID          uint
	ManagerRootPath string
	RootPath        string
	Source          []File
	CommitRecord    [][]*Commit
}

func NewBox(arg string) {
	//client, err := .New(util.CONFIG.OSSEndPoint, util.CONFIG.OSSAccessKeyID, util.CONFIG.OSSSecretKey)
	//if err != nil {
	//	log.Error(err.Error())
	//}
}

func (b *Box) BoxName() string {
	return fmt.Sprintf("%s-%x", b.Name, b.ProjID)
}

func (b *Box) NewDirectory(dirPath string, dirName string) error {
	fullPath := filepath.Join(b.ManagerRootPath, dirPath+dirName)
	err := os.MkdirAll(fullPath, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
	return nil
}

func (b *Box) RemoveDirectory(dirPath string) error {
	fullPath := filepath.Join(b.ManagerRootPath, dirPath)
	err := os.RemoveAll(fullPath)
	if err != nil {
		return fmt.Errorf("failed to delete directory: %v", err)
	}
	return nil
}

func (b *Box) ListDirectory(dirPath string) ([]string, error) {
	fullPath := filepath.Join(b.ManagerRootPath, dirPath)
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

func (b *Box) NewFile(filePath string, content []byte) error {
	fullPath := filepath.Join(b.ManagerRootPath, filePath)
	err := ioutil.WriteFile(fullPath, content, 0644)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	return nil
}

func (b *Box) ReadFile(filePath string) ([]byte, error) {
	fullPath := filepath.Join(b.ManagerRootPath, filePath)
	content, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}
	return content, nil
}

func (b *Box) AddSnippetToFile(opt *OptAdd) error {
	if opt.TargetBox != b {
		return errors.New("illegal call about target box")
	}
	file, err := os.Open(opt.FilePath)
	if err != nil {
		return err
	}
	tempFile, err := b.createTempFile(opt.FilePath)
	if err != nil {
		return err
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	if _, err := io.CopyN(tempFile, file, opt.BeginPos); err != nil {
		return err
	}
	if _, err := tempFile.Write(opt.AddContent); err != nil {
		return err
	}
	_, err = io.Copy(tempFile, file)
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}

func (b *Box) DeleteSnippetToFile(opt *OptDel) error {
	if opt.TargetBox != b {
		return errors.New("illegal call about target box")
	}
	file, err := os.Open(opt.FilePath)
	if err != nil {
		return err
	}
	tempFile, err := b.createTempFile(opt.FilePath)
	if err != nil {
		return err
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	if _, err := io.CopyN(tempFile, file, opt.BeginPos); err != nil {
		return err
	}
	opt.DeletedContent, err = ioutil.ReadAll(io.LimitReader(file, opt.EndPos-opt.BeginPos))
	if err != nil {
		return err
	}
	_, err = io.CopyN(ioutil.Discard, file, opt.EndPos-opt.BeginPos)
	if err != nil && err != io.EOF {
		return err
	}
	_, err = io.Copy(tempFile, file)
	if err != nil && err != io.EOF {
		return err
	}
	err = os.Rename(tempFile.Name(), opt.FilePath)
	if err != nil {
		return err
	}
	return nil
}

func (b *Box) ReplaceSnippetToFile(opt *OptReplace) error {
	if opt.TargetBox != b {
		return errors.New("illegal call about target box")
	}
	file, err := os.Open(opt.FilePath)
	if err != nil {
		return err
	}
	tempFile, err := b.createTempFile(opt.FilePath)
	if err != nil {
		return err
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	if _, err := io.CopyN(tempFile, file, opt.BeginPos); err != nil {
		return err
	}
	opt.ReplaceContent, err = ioutil.ReadAll(io.LimitReader(file, opt.EndPos-opt.BeginPos))
	if err != nil {
		return err
	}
	_, err = io.CopyN(ioutil.Discard, file, opt.EndPos-opt.BeginPos)
	if err != nil && err != io.EOF {
		return err
	}
	if _, err := tempFile.Write(opt.NewContent); err != nil {
		return err
	}
	_, err = io.Copy(tempFile, file)
	if err != nil && err != io.EOF {
		return err
	}
	err = os.Rename(tempFile.Name(), opt.FilePath)
	if err != nil {
		return err
	}
	return nil
}

func (b *Box) DeleteFile(filePath string) error {
	fullPath := filepath.Join(b.ManagerRootPath, filePath)
	err := os.Remove(fullPath)
	if err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}
	return nil
}

func (b *Box) createTempFile(filePath string) (*os.File, error) {
	tmpFile, err := ioutil.TempFile(filePath, "example.*.txt")
	return tmpFile, err
}
