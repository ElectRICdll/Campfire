package util

import (
	"encoding/base64"
	"io/ioutil"
)

func FileToBase64(filePath string) (string, error) {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(fileData), nil
}
