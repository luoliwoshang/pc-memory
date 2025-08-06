package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"memory-manager/internal/memory"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a memory by ID",
	Long: `Delete a specific memory from the store using its ID.

Examples:
  memory delete 20240806123456-abc123`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		
		service := memory.NewService()
		if err := service.DeleteMemory(id); err != nil {
			fmt.Printf("Error deleting memory: %v\n", err)
			return
		}
		
		fmt.Printf("Memory with ID %s deleted successfully.\n", id)
	},
}