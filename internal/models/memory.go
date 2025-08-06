package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Memory represents a stored memory item
type Memory struct {
	ID           string    `json:"id"`
	OriginalText string    `json:"original"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// MemoryStore represents the collection of all memories
type MemoryStore struct {
	Memories []Memory `json:"memories"`
	Metadata Metadata `json:"metadata"`
}

// Metadata contains store-level information
type Metadata struct {
	Version    string    `json:"version"`
	TotalCount int       `json:"total_count"`
	LastUpdate time.Time `json:"last_update"`
}

// NewMemory creates a new memory with generated ID and timestamps
func NewMemory(text string) *Memory {
	now := time.Now()
	return &Memory{
		ID:           uuid.New().String(),
		OriginalText: text,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

// NewMemoryStore creates a new empty memory store
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		Memories: []Memory{},
		Metadata: Metadata{
			Version:    "1.0.0",
			TotalCount: 0,
			LastUpdate: time.Now(),
		},
	}
}

// AddMemory adds a new memory to the store
func (ms *MemoryStore) AddMemory(memory *Memory) {
	ms.Memories = append(ms.Memories, *memory)
	ms.Metadata.TotalCount = len(ms.Memories)
	ms.Metadata.LastUpdate = time.Now()
}

// ToJSON converts the memory store to JSON
func (ms *MemoryStore) ToJSON() ([]byte, error) {
	return json.MarshalIndent(ms, "", "  ")
}

// FromJSON loads the memory store from JSON
func (ms *MemoryStore) FromJSON(data []byte) error {
	return json.Unmarshal(data, ms)
}