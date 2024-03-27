package gitserver

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// GetBranches 获取指定仓库的所有分支
func GetBranches(repoPath string) ([]string, error) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, fmt.Errorf("打开仓库失败：%s", err)
	}

	branches, err := repo.Branches()
	if err != nil {
		return nil, fmt.Errorf("获取分支列表失败：%s", err)
	}

	var branchNames []string
	err = branches.ForEach(func(ref *plumbing.Reference) error {
		branchNames = append(branchNames, ref.Name().Short())
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("遍历分支列表失败：%s", err)
	}

	return branchNames, nil
}
