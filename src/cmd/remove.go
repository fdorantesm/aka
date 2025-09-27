package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [alias]",
	Short: "Remove an existing alias",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		akaDir := getAkaDir()
		aliasName := args[0]
		if err := removeAlias(akaDir, aliasName); err != nil {
			fmt.Fprintln(os.Stderr, "Error removing alias:", err)
			os.Exit(1)
		}
		fmt.Printf("Alias '%s' removed successfully.\n", aliasName)
	},
}
