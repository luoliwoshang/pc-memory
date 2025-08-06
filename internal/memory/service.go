package memory

import (
	"encoding/json"
	"fmt"
	"strings"

	"memory-manager/internal/claude"
	"memory-manager/internal/models"
	"memory-manager/internal/storage"
)

type Service struct {
	storage     storage.StorageInterface
	claudeClient *claude.ClaudeClient
}

type ProcessedMemoryResponse struct {
	Summary  string   `json:"summary"`
	Category string   `json:"category"`
	Priority int      `json:"priority"`
	Tags     []string `json:"tags"`
}

func NewService() *Service {
	return &Service{
		storage:     storage.NewFileStorage(),
		claudeClient: claude.NewClaudeClient(),
	}
}

func (s *Service) StoreMemory(content string) (*models.Memory, error) {
	memory := models.NewMemory(content)
	
	// 如果Claude可用，处理内容
	if s.claudeClient.IsAvailable() {
		processed, err := s.claudeClient.ProcessMemory(content)
		if err == nil {
			if err := s.parseProcessedResponse(processed, memory); err != nil {
				// 如果解析失败，继续使用原始内容，但记录错误
				fmt.Printf("Warning: Failed to parse Claude response: %v\n", err)
			}
			memory.ProcessedBy = "claude"
		} else {
			fmt.Printf("Warning: Claude processing failed: %v\n", err)
			memory.ProcessedBy = "none"
		}
	} else {
		// Claude不可用时的基本处理
		s.basicProcessing(memory)
		memory.ProcessedBy = "basic"
	}
	
	if err := s.storage.Save(memory); err != nil {
		return nil, fmt.Errorf("failed to save memory: %w", err)
	}
	
	return memory, nil
}

func (s *Service) parseProcessedResponse(response string, memory *models.Memory) error {
	// 尝试从响应中提取JSON部分
	response = strings.TrimSpace(response)
	
	// 寻找JSON块
	jsonStart := strings.Index(response, "{")
	jsonEnd := strings.LastIndex(response, "}")
	
	if jsonStart == -1 || jsonEnd == -1 {
		return fmt.Errorf("no JSON found in response")
	}
	
	jsonStr := response[jsonStart : jsonEnd+1]
	
	var processed ProcessedMemoryResponse
	if err := json.Unmarshal([]byte(jsonStr), &processed); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}
	
	// 更新memory对象
	if processed.Summary != "" {
		memory.Metadata["summary"] = processed.Summary
	}
	if processed.Category != "" {
		memory.Category = processed.Category
	}
	if processed.Priority > 0 && processed.Priority <= 5 {
		memory.Priority = processed.Priority
	}
	if len(processed.Tags) > 0 {
		memory.Tags = processed.Tags
	}
	
	return nil
}

func (s *Service) basicProcessing(memory *models.Memory) {
	content := strings.ToLower(memory.Content)
	
	// 基本分类逻辑
	if strings.Contains(content, "password") || strings.Contains(content, "密码") ||
	   strings.Contains(content, "account") || strings.Contains(content, "账号") {
		memory.Category = "账号信息"
		memory.Priority = 5
		memory.Tags = []string{"账号", "密码", "重要"}
	} else if strings.Contains(content, "task") || strings.Contains(content, "任务") ||
			  strings.Contains(content, "todo") || strings.Contains(content, "工作") {
		memory.Category = "工作任务"
		memory.Priority = 4
		memory.Tags = []string{"任务", "工作"}
	} else if strings.Contains(content, "note") || strings.Contains(content, "笔记") ||
			  strings.Contains(content, "remember") || strings.Contains(content, "记住") {
		memory.Category = "个人笔记"
		memory.Priority = 3
		memory.Tags = []string{"笔记", "记忆"}
	} else {
		memory.Category = "其他"
		memory.Priority = 2
		memory.Tags = []string{"杂项"}
	}
}

func (s *Service) ListMemories() ([]*models.Memory, error) {
	return s.storage.LoadAll()
}

func (s *Service) SearchMemories(query string) ([]*models.Memory, error) {
	return s.storage.Search(query)
}

func (s *Service) DeleteMemory(id string) error {
	return s.storage.Delete(id)
}

func (s *Service) GetMemoryStats() (map[string]interface{}, error) {
	memories, err := s.storage.LoadAll()
	if err != nil {
		return nil, err
	}
	
	stats := map[string]interface{}{
		"total_memories": len(memories),
		"categories":     make(map[string]int),
		"priorities":     make(map[int]int),
	}
	
	categories := stats["categories"].(map[string]int)
	priorities := stats["priorities"].(map[int]int)
	
	for _, memory := range memories {
		categories[memory.Category]++
		priorities[memory.Priority]++
	}
	
	return stats, nil
}