package gitserver

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

// CloneRepository 克隆远程仓库到指定路径
func CloneRepository(remoteURL, localPath string) error {

	_, err := git.PlainClone(localPath, false, &git.CloneOptions{
		URL:      remoteURL,
		Progress: os.Stdout,
	})
	if err != nil {
		return fmt.Errorf("无法克隆仓库：%s", err)
	}

	return nil
}
