package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import [file]",
	Short: "Import aliases from a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		akaDir := getAkaDir()
		if err := importAliases(akaDir, args[0]); err != nil {
			fmt.Fprintln(os.Stderr, "Error importing aliases:", err)
			os.Exit(1)
		}
	},
}
