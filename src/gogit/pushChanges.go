package gitserver

import (
	"fmt"
	"os/exec"
)

// PushChanges 推送更改到远程仓库
func PushChanges(repoPath string) error {
	cmd := exec.Command("git", "-C", repoPath, "push")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("推送更改失败：%s, 输出：%s", err, string(output))
	}
	return nil
}
