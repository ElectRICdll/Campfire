package gitserver

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

// CommitChanges 提交更改到指定仓库
func CommitChanges(repoPath, commitMessage string) error {

	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return fmt.Errorf("打开仓库失败：%s", err)
	}

	wt, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("获取工作树失败：%s", err)
	}

	_, err = wt.Add(".")
	if err != nil {
		return fmt.Errorf("添加文件到暂存区失败：%s", err)
	}

	_, err = wt.Commit(commitMessage, &git.CommitOptions{})
	if err != nil {
		return fmt.Errorf("提交更改失败：%s", err)
	}

	return nil
}
