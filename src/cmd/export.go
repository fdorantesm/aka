package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export [file]",
	Short: "Export aliases to a file or stdout",
	Run: func(cmd *cobra.Command, args []string) {
		akaDir := getAkaDir()
		out := ""
		if len(args) > 0 {
			out = args[0]
		}
		if err := exportAliases(akaDir, out); err != nil {
			fmt.Fprintln(os.Stderr, "Error exporting aliases:", err)
			os.Exit(1)
		}
	},
}
