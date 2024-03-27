package gitserver

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

// CreateRepository 创建一个新的 Git 仓库
func CreateRepository(repoPath string) error {
	err := os.MkdirAll(repoPath, 0755)
	if err != nil {
		return fmt.Errorf("无法创建仓库目录：%s", err)
	}

	_, err = git.PlainInit(repoPath, false)
	if err != nil {
		return fmt.Errorf("无法初始化仓库：%s", err)
	}

	return nil
}
