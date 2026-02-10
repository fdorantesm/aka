package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update <alias>",
	Short: "Update an existing alias",
	Long:  `Update the command associated with an alias. You'll be prompted to enter the new command.`,
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		akaDir := getAkaDir()
		aliasName := args[0]
		
		// Check if alias exists
		filePath := filepath.Join(akaDir, aliasName+".alias")
		if _, err := os.Stat(filePath); err != nil {
			if os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "Error: alias '%s' does not exist\n", aliasName)
				os.Exit(1)
			}
			fmt.Fprintf(os.Stderr, "Error reading alias: %v\n", err)
			os.Exit(1)
		}
		
		// Read current command
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading alias: %v\n", err)
			os.Exit(1)
		}
		
		currentCmd := strings.TrimSpace(string(content))
		
		// Show current command
		fmt.Printf("Current command: %s\n", currentCmd)
		fmt.Print("Enter new command: ")
		
		// Read new command from stdin
		reader := bufio.NewReader(os.Stdin)
		newCmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}
		
		newCmd = strings.TrimSpace(newCmd)
		
		// Check if command is empty
		if newCmd == "" {
			fmt.Println("Error: command cannot be empty")
			os.Exit(1)
		}
		
		// Update alias
		if err := addAlias(akaDir, aliasName, newCmd); err != nil {
			fmt.Fprintf(os.Stderr, "Error updating alias: %v\n", err)
			os.Exit(1)
		}
		
		fmt.Printf("Alias '%s' updated successfully\n", aliasName)
	},
}
