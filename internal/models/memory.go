package models

import (
	"encoding/json"
	"time"
)

type Memory struct {
	ID          string            `json:"id"`
	Content     string            `json:"content"`
	ProcessedBy string            `json:"processed_by"`
	Timestamp   time.Time         `json:"timestamp"`
	Tags        []string          `json:"tags,omitempty"`
	Category    string            `json:"category,omitempty"`
	Priority    int               `json:"priority,omitempty"` // 1-5, 5 being highest
	Metadata    map[string]string `json:"metadata,omitempty"`
}

type MemoryStore struct {
	Memories []Memory `json:"memories"`
	Version  string   `json:"version"`
	LastSync time.Time `json:"last_sync"`
}

func NewMemory(content string) *Memory {
	return &Memory{
		ID:        generateID(),
		Content:   content,
		Timestamp: time.Now(),
		Priority:  3, // default medium priority
		Metadata:  make(map[string]string),
	}
}

func generateID() string {
	return time.Now().Format("20060102150405") + "-" + randomString(6)
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[time.Now().UnixNano()%int64(len(letters))]
	}
	return string(b)
}

func (m *Memory) ToJSON() ([]byte, error) {
	return json.MarshalIndent(m, "", "  ")
}

func (m *Memory) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}