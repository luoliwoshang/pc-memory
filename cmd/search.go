package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"memory-manager/internal/memory"
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search memories",
	Long: `Search through stored memories by content, category, or tags.

Examples:
  memory search "密码"
  memory search "工作任务"
  memory search "GitHub"`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.Join(args, " ")
		
		service := memory.NewService()
		results, err := service.SearchMemories(query)
		if err != nil {
			fmt.Printf("Error searching memories: %v\n", err)
			return
		}
		
		if len(results) == 0 {
			fmt.Printf("No memories found for query: %s\n", query)
			return
		}
		
		fmt.Printf("Found %d memories for query '%s':\n\n", len(results), query)
		
		for i, mem := range results {
			fmt.Printf("[%d] ID: %s\n", i+1, mem.ID)
			fmt.Printf("    Content: %s\n", highlightQuery(mem.Content, query))
			fmt.Printf("    Category: %s | Priority: %d/5\n", mem.Category, mem.Priority)
			if len(mem.Tags) > 0 {
				fmt.Printf("    Tags: %s\n", strings.Join(mem.Tags, ", "))
			}
			fmt.Printf("    Created: %s\n", mem.Timestamp.Format(time.RFC3339))
			
			if summary, exists := mem.Metadata["summary"]; exists {
				fmt.Printf("    Summary: %s\n", summary)
			}
			fmt.Println()
		}
	},
}

func highlightQuery(text, query string) string {
	// 简单高亮实现（在实际CLI中可以使用颜色）
	return strings.ReplaceAll(text, query, fmt.Sprintf("[%s]", query))
}