package gitserver

import (
	"fmt"
	"os"
)

// DeleteRepository 删除指定路径下的 Git 仓库
func DeleteRepository(repoPath string) error {

	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		return fmt.Errorf("仓库不存在：%s", repoPath)
	}

	if err := os.RemoveAll(repoPath); err != nil {
		return fmt.Errorf("无法删除仓库：%s", err)
	}

	return nil
}
