package claude

import (
	"fmt"
	"os/exec"
	"strings"
)

type ClaudeClient struct {
	executable string
}

func NewClaudeClient() *ClaudeClient {
	return &ClaudeClient{
		executable: "claude",
	}
}

func (c *ClaudeClient) ProcessMemory(content string) (string, error) {
	prompt := fmt.Sprintf(`请分析以下内容并提供一个简洁的记忆摘要，包括：
1. 关键信息提取
2. 分类建议（如：账号信息、工作任务、个人笔记等）
3. 重要度评级（1-5）
4. 建议的标签

内容：%s

请用JSON格式回复，包含以下字段：
- summary: 简洁摘要
- category: 分类
- priority: 重要度(1-5)
- tags: 标签数组`, content)

	cmd := exec.Command(c.executable, "-p", prompt)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute claude command: %w", err)
	}

	return strings.TrimSpace(string(output)), nil
}

func (c *ClaudeClient) IsAvailable() bool {
	_, err := exec.LookPath(c.executable)
	return err == nil
}