package storage

import (
	"memory/internal/models"
	"os"
	"path/filepath"
)

const (
	configDirName = "memory"
	dataFileName  = "memories.json"
)

// Storage handles file operations for memories
type Storage struct {
	dataPath string
}

// NewStorage creates a new storage instance
func NewStorage() (*Storage, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configDir := filepath.Join(homeDir, ".config", configDirName)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, err
	}

	dataPath := filepath.Join(configDir, dataFileName)
	
	return &Storage{
		dataPath: dataPath,
	}, nil
}

// LoadStore loads the memory store from file, creates new if doesn't exist
func (s *Storage) LoadStore() (*models.MemoryStore, error) {
	store := models.NewMemoryStore()
	
	if _, err := os.Stat(s.dataPath); os.IsNotExist(err) {
		return store, nil
	}
	
	data, err := os.ReadFile(s.dataPath)
	if err != nil {
		return nil, err
	}
	
	if len(data) == 0 {
		return store, nil
	}
	
	err = store.FromJSON(data)
	if err != nil {
		return nil, err
	}
	
	return store, nil
}

// SaveStore saves the memory store to file
func (s *Storage) SaveStore(store *models.MemoryStore) error {
	data, err := store.ToJSON()
	if err != nil {
		return err
	}
	
	return os.WriteFile(s.dataPath, data, 0600)
}