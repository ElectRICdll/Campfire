package gitserver

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// CreateTag 创建一个新的标签
func CreateTag(repoPath, tagName, commitHash string) error {
	// 打开现有仓库
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return fmt.Errorf("打开仓库失败：%s", err)
	}

	// 解析提交哈希
	hash := plumbing.NewHash(commitHash)

	// 创建一个新的标签
	tag := plumbing.ReferenceName(fmt.Sprintf("refs/tags/%s", tagName))
	ref := plumbing.NewHashReference(tag, hash)
	if err := repo.Storer.SetReference(ref); err != nil {
		return fmt.Errorf("创建标签失败：%s", err)
	}

	return nil
}
