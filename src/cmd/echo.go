package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var echoCmd = &cobra.Command{
	Use:   "echo <alias>",
	Short: "Echo the command for an alias",
	Long:  `Print the command associated with an alias. Useful for copying or inspecting alias definitions.`,
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		akaDir := getAkaDir()
		aliasName := args[0]
		
		filePath := filepath.Join(akaDir, aliasName+".alias")
		if _, err := os.Stat(filePath); err != nil {
			if os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "Error: alias '%s' does not exist\n", aliasName)
				os.Exit(1)
			}
			fmt.Fprintf(os.Stderr, "Error reading alias: %v\n", err)
			os.Exit(1)
		}
		
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading alias: %v\n", err)
			os.Exit(1)
		}
		
		cmd := strings.TrimSpace(string(content))
		fmt.Println(cmd)
	},
}
