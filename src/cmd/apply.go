package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Generate shell functions for aliases",
	Run: func(cmd *cobra.Command, args []string) {
		akaDir := getAkaDir()
		if err := applyAliases(akaDir); err != nil {
			fmt.Fprintln(os.Stderr, "Error generating alias definitions:", err)
			os.Exit(1)
		}
	},
}
