package gitserver

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// GetCommitHistory 获取指定仓库的提交历史
func GetCommitHistory(repoPath string) ([]*object.Commit, error) {

	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, fmt.Errorf("打开仓库失败：%s", err)
	}
	commitIter, err := repo.CommitObjects()
	if err != nil {
		return nil, fmt.Errorf("获取提交历史失败：%s", err)
	}

	var commits []*object.Commit
	err = commitIter.ForEach(func(commit *object.Commit) error {
		commits = append(commits, commit)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("遍历提交历史失败：%s", err)
	}

	return commits, nil
}
