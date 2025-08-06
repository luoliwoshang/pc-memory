package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"memory-manager/internal/memory"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show memory statistics",
	Long: `Display statistics about stored memories including total count,
categories breakdown, and priority distribution.

Examples:
  memory stats`,
	Run: func(cmd *cobra.Command, args []string) {
		service := memory.NewService()
		stats, err := service.GetMemoryStats()
		if err != nil {
			fmt.Printf("Error getting statistics: %v\n", err)
			return
		}
		
		fmt.Println("Memory Statistics:")
		fmt.Println("==================")
		fmt.Printf("Total memories: %d\n\n", stats["total_memories"])
		
		fmt.Println("Categories:")
		categories := stats["categories"].(map[string]int)
		for category, count := range categories {
			fmt.Printf("  %s: %d\n", category, count)
		}
		
		fmt.Println("\nPriorities:")
		priorities := stats["priorities"].(map[int]int)
		for priority := 5; priority >= 1; priority-- {
			if count, exists := priorities[priority]; exists {
				fmt.Printf("  Priority %d: %d memories\n", priority, count)
			}
		}
	},
}