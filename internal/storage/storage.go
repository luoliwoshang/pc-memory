package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"memory-manager/internal/models"
)

type StorageInterface interface {
	Save(memory *models.Memory) error
	Load() (*models.MemoryStore, error)
	LoadAll() ([]*models.Memory, error)
	Delete(id string) error
	Search(query string) ([]*models.Memory, error)
}

type FileStorage struct {
	dataDir  string
	filename string
}

func NewFileStorage() *FileStorage {
	homeDir, _ := os.UserHomeDir()
	dataDir := filepath.Join(homeDir, ".memory-manager")
	
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		os.MkdirAll(dataDir, 0755)
	}
	
	return &FileStorage{
		dataDir:  dataDir,
		filename: "memories.json",
	}
}

func (fs *FileStorage) getFilePath() string {
	return filepath.Join(fs.dataDir, fs.filename)
}

func (fs *FileStorage) Save(memory *models.Memory) error {
	store, err := fs.Load()
	if err != nil {
		// 如果文件不存在，创建新的store
		store = &models.MemoryStore{
			Memories: make([]models.Memory, 0),
			Version:  "1.0",
		}
	}
	
	// 检查是否已存在相同ID的记忆，如果存在则更新
	found := false
	for i, existing := range store.Memories {
		if existing.ID == memory.ID {
			store.Memories[i] = *memory
			found = true
			break
		}
	}
	
	if !found {
		store.Memories = append(store.Memories, *memory)
	}
	
	return fs.saveStore(store)
}

func (fs *FileStorage) Load() (*models.MemoryStore, error) {
	filePath := fs.getFilePath()
	
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return &models.MemoryStore{
			Memories: make([]models.Memory, 0),
			Version:  "1.0",
		}, nil
	}
	
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read memory file: %w", err)
	}
	
	var store models.MemoryStore
	if err := json.Unmarshal(data, &store); err != nil {
		return nil, fmt.Errorf("failed to parse memory file: %w", err)
	}
	
	return &store, nil
}

func (fs *FileStorage) LoadAll() ([]*models.Memory, error) {
	store, err := fs.Load()
	if err != nil {
		return nil, err
	}
	
	memories := make([]*models.Memory, len(store.Memories))
	for i := range store.Memories {
		memories[i] = &store.Memories[i]
	}
	
	return memories, nil
}

func (fs *FileStorage) Delete(id string) error {
	store, err := fs.Load()
	if err != nil {
		return err
	}
	
	for i, memory := range store.Memories {
		if memory.ID == id {
			store.Memories = append(store.Memories[:i], store.Memories[i+1:]...)
			return fs.saveStore(store)
		}
	}
	
	return fmt.Errorf("memory with ID %s not found", id)
}

func (fs *FileStorage) Search(query string) ([]*models.Memory, error) {
	memories, err := fs.LoadAll()
	if err != nil {
		return nil, err
	}
	
	var results []*models.Memory
	for _, memory := range memories {
		if contains(memory.Content, query) || 
		   contains(memory.Category, query) ||
		   containsInTags(memory.Tags, query) {
			results = append(results, memory)
		}
	}
	
	return results, nil
}

func (fs *FileStorage) saveStore(store *models.MemoryStore) error {
	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal memory store: %w", err)
	}
	
	filePath := fs.getFilePath()
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write memory file: %w", err)
	}
	
	return nil
}

func contains(text, query string) bool {
	return len(query) > 0 && len(text) >= len(query) && 
		   findSubstring(text, query) >= 0
}

func findSubstring(text, substr string) int {
	if len(substr) == 0 {
		return 0
	}
	for i := 0; i <= len(text)-len(substr); i++ {
		if text[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func containsInTags(tags []string, query string) bool {
	for _, tag := range tags {
		if contains(tag, query) {
			return true
		}
	}
	return false
}