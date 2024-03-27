package util

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CreateZip(filePath string) ([]byte, error) {
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)
	defer zipWriter.Close()

	// 将仓库目录下的所有文件添加到 zip 文件中
	if err := filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(filePath, path)
		if err != nil {
			return err
		}
		if strings.HasSuffix(relPath, ".git") {
			return nil
		}

		header.Name = relPath

		// 文件夹处理
		if info.IsDir() {
			header.Name += "/"
		}

		// 写入zip文件
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// 文件内容写入
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
