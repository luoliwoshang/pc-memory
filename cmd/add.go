package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"memory-manager/internal/memory"
)

var addCmd = &cobra.Command{
	Use:   "add [content]",
	Short: "Add a new memory",
	Long: `Add a new memory to the store. The content will be processed by Claude AI
if available to provide intelligent categorization and tagging.

Examples:
  memory add "我的GitHub账号密码是abc123"
  memory add "明天需要完成项目报告"
  memory add "记住要买牛奶和面包"`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		content := strings.Join(args, " ")
		
		service := memory.NewService()
		savedMemory, err := service.StoreMemory(content)
		if err != nil {
			fmt.Printf("Error storing memory: %v\n", err)
			return
		}
		
		fmt.Printf("Memory saved successfully!\n")
		fmt.Printf("ID: %s\n", savedMemory.ID)
		fmt.Printf("Content: %s\n", savedMemory.Content)
		fmt.Printf("Category: %s\n", savedMemory.Category)
		fmt.Printf("Priority: %d/5\n", savedMemory.Priority)
		if len(savedMemory.Tags) > 0 {
			fmt.Printf("Tags: %s\n", strings.Join(savedMemory.Tags, ", "))
		}
		fmt.Printf("Processed by: %s\n", savedMemory.ProcessedBy)
		
		if summary, exists := savedMemory.Metadata["summary"]; exists {
			fmt.Printf("AI Summary: %s\n", summary)
		}
	},
}