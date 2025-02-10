package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all aliases",
	Run: func(_ *cobra.Command, _ []string) {
		akaDir := getAkaDir()
		if err := listAliases(akaDir); err != nil {
			fmt.Fprintln(os.Stderr, "Error listing aliases:", err)
			os.Exit(1)
		}
	},
}
