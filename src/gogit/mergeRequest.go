package gitserver

import (
	"fmt"
	"os/exec"
)

// MergeRequest 合并指定分支到主分支
func MergeRequest(repoPath, branchName string) error {
	// 拉取最新更改
	if err := execCommand("git", []string{"-C", repoPath, "pull"}); err != nil {
		return fmt.Errorf("拉取最新更改失败：%s", err)
	}

	// 合并指定分支到主分支
	if err := execCommand("git", []string{"-C", repoPath, "merge", "origin/" + branchName}); err != nil {
		return fmt.Errorf("合并请求失败：%s", err)
	}

	// 推送合并后的更改到远程仓库
	if err := execCommand("git", []string{"-C", repoPath, "push"}); err != nil {
		return fmt.Errorf("推送更改失败：%s", err)
	}

	return nil
}

// execCommand 执行命令
func execCommand(command string, args []string) error {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("执行命令失败：%s, 输出：%s", err, string(output))
	}
	return nil
}
