package models

import (
	"strings"
	"testing"
	"time"
)

func TestNewMemory(t *testing.T) {
	content := "测试记忆内容"
	memory := NewMemory(content)
	
	if memory.Content != content {
		t.Errorf("Expected content %s, got %s", content, memory.Content)
	}
	
	if memory.Priority != 3 {
		t.Errorf("Expected default priority 3, got %d", memory.Priority)
	}
	
	if memory.ID == "" {
		t.Error("Expected non-empty ID")
	}
	
	if memory.Timestamp.IsZero() {
		t.Error("Expected non-zero timestamp")
	}
	
	if memory.Metadata == nil {
		t.Error("Expected initialized metadata map")
	}
}

func TestGenerateID(t *testing.T) {
	id1 := generateID()
	id2 := generateID()
	
	if id1 == id2 {
		t.Error("Expected unique IDs")
	}
	
	if !strings.Contains(id1, "-") {
		t.Error("Expected ID to contain hyphen separator")
	}
	
	parts := strings.Split(id1, "-")
	if len(parts) != 2 {
		t.Error("Expected ID to have timestamp and random parts")
	}
	
	// 验证时间戳部分格式
	if len(parts[0]) != 14 {
		t.Error("Expected timestamp part to be 14 characters")
	}
	
	// 验证随机字符串长度
	if len(parts[1]) != 6 {
		t.Error("Expected random part to be 6 characters")
	}
}

func TestMemoryJSON(t *testing.T) {
	memory := &Memory{
		ID:        "test-id",
		Content:   "test content",
		Timestamp: time.Date(2024, 8, 6, 12, 0, 0, 0, time.UTC),
		Category:  "测试分类",
		Priority:  4,
		Tags:      []string{"tag1", "tag2"},
		Metadata:  map[string]string{"key": "value"},
	}
	
	jsonData, err := memory.ToJSON()
	if err != nil {
		t.Fatalf("Failed to serialize to JSON: %v", err)
	}
	
	var newMemory Memory
	err = newMemory.FromJSON(jsonData)
	if err != nil {
		t.Fatalf("Failed to deserialize from JSON: %v", err)
	}
	
	if newMemory.ID != memory.ID {
		t.Errorf("Expected ID %s, got %s", memory.ID, newMemory.ID)
	}
	
	if newMemory.Content != memory.Content {
		t.Errorf("Expected content %s, got %s", memory.Content, newMemory.Content)
	}
	
	if newMemory.Category != memory.Category {
		t.Errorf("Expected category %s, got %s", memory.Category, newMemory.Category)
	}
	
	if newMemory.Priority != memory.Priority {
		t.Errorf("Expected priority %d, got %d", memory.Priority, newMemory.Priority)
	}
}