package cmd

import (
	"fmt"
	"memory/internal/models"
	"memory/internal/storage"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [text]",
	Short: "Add a new memory",
	Long:  `Add a new memory to store important information.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		text := strings.Join(args, " ")
		
		if err := addMemory(text); err != nil {
			fmt.Printf("Error adding memory: %v\n", err)
			return
		}
		
		fmt.Printf("Memory added successfully!\n")
	},
}

func addMemory(text string) error {
	// Initialize storage
	storage, err := storage.NewStorage()
	if err != nil {
		return fmt.Errorf("failed to initialize storage: %w", err)
	}
	
	// Load existing store
	store, err := storage.LoadStore()
	if err != nil {
		return fmt.Errorf("failed to load memory store: %w", err)
	}
	
	// Create new memory
	memory := models.NewMemory(text)
	
	// Add to store
	store.AddMemory(memory)
	
	// Save store
	if err := storage.SaveStore(store); err != nil {
		return fmt.Errorf("failed to save memory: %w", err)
	}
	
	return nil
}

func init() {
	rootCmd.AddCommand(addCmd)
}