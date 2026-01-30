package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [pattern]",
	Short: "List aliases (optionally filtered by pattern)",
	Long: `List all aliases or filter them using a glob pattern.

Examples:
  aka list           # List all aliases
  aka list *dev*     # List aliases containing "dev"
  aka list aws*      # List aliases starting with "aws"
  aka list *kube*    # List aliases containing "kube"`,
	Args: cobra.MaximumNArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		akaDir := getAkaDir()
		pattern := ""
		if len(args) > 0 {
			pattern = args[0]
		}
		if err := listAliases(akaDir, pattern); err != nil {
			fmt.Fprintln(os.Stderr, "Error listing aliases:", err)
			os.Exit(1)
		}
	},
}
