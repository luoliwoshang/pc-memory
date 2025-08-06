package cmd

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"memory-manager/internal/memory"
	"memory-manager/internal/models"
)

var (
	sortBy   string
	category string
	limit    int
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all memories",
	Long: `List all stored memories with optional filtering and sorting.

Examples:
  memory list
  memory list --sort=priority
  memory list --category="账号信息"
  memory list --limit=10`,
	Run: func(cmd *cobra.Command, args []string) {
		service := memory.NewService()
		memories, err := service.ListMemories()
		if err != nil {
			fmt.Printf("Error loading memories: %v\n", err)
			return
		}
		
		// 过滤
		if category != "" {
			filtered := make([]*models.Memory, 0)
			for _, mem := range memories {
				if mem.Category == category {
					filtered = append(filtered, mem)
				}
			}
			memories = filtered
		}
		
		// 排序
		switch sortBy {
		case "priority":
			sort.Slice(memories, func(i, j int) bool {
				return memories[i].Priority > memories[j].Priority
			})
		case "date":
			sort.Slice(memories, func(i, j int) bool {
				return memories[i].Timestamp.After(memories[j].Timestamp)
			})
		case "category":
			sort.Slice(memories, func(i, j int) bool {
				return memories[i].Category < memories[j].Category
			})
		}
		
		// 限制数量
		if limit > 0 && len(memories) > limit {
			memories = memories[:limit]
		}
		
		if len(memories) == 0 {
			fmt.Println("No memories found.")
			return
		}
		
		fmt.Printf("Found %d memories:\n\n", len(memories))
		
		for i, mem := range memories {
			fmt.Printf("[%d] ID: %s\n", i+1, mem.ID)
			fmt.Printf("    Content: %s\n", truncateString(mem.Content, 80))
			fmt.Printf("    Category: %s | Priority: %d/5\n", mem.Category, mem.Priority)
			if len(mem.Tags) > 0 {
				fmt.Printf("    Tags: %s\n", strings.Join(mem.Tags, ", "))
			}
			fmt.Printf("    Created: %s\n", mem.Timestamp.Format(time.RFC3339))
			
			if summary, exists := mem.Metadata["summary"]; exists {
				fmt.Printf("    Summary: %s\n", truncateString(summary, 60))
			}
			fmt.Println()
		}
	},
}

func init() {
	listCmd.Flags().StringVar(&sortBy, "sort", "", "Sort by: priority, date, category")
	listCmd.Flags().StringVar(&category, "category", "", "Filter by category")
	listCmd.Flags().IntVar(&limit, "limit", 0, "Limit number of results")
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}