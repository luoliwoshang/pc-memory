package cmd

import (
	"fmt"
	"memory/internal/storage"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all memories",
	Long:  `Display all stored memories.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := listMemories(); err != nil {
			fmt.Printf("Error listing memories: %v\n", err)
			return
		}
	},
}

func listMemories() error {
	// Initialize storage
	storage, err := storage.NewStorage()
	if err != nil {
		return fmt.Errorf("failed to initialize storage: %w", err)
	}
	
	// Load store
	store, err := storage.LoadStore()
	if err != nil {
		return fmt.Errorf("failed to load memory store: %w", err)
	}
	
	if len(store.Memories) == 0 {
		fmt.Println("No memories found.")
		return nil
	}
	
	fmt.Printf("Found %d memories:\n\n", len(store.Memories))
	
	for i, memory := range store.Memories {
		fmt.Printf("%d. [%s] %s\n", 
			i+1, 
			memory.CreatedAt.Format("2006-01-02 15:04"), 
			memory.OriginalText)
	}
	
	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
}