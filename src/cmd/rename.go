package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:   "rename <old-name> <new-name>",
	Short: "Rename an existing alias",
	Long:  `Rename an alias by creating a new one with the same command and removing the old one.`,
	Args:  cobra.ExactArgs(2),
	Run: func(_ *cobra.Command, args []string) {
		akaDir := getAkaDir()
		oldName := args[0]
		newName := args[1]
		
		// Check if old alias exists
		oldPath := filepath.Join(akaDir, oldName+".alias")
		if _, err := os.Stat(oldPath); err != nil {
			if os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "Error: alias '%s' does not exist\n", oldName)
				os.Exit(1)
			}
			fmt.Fprintf(os.Stderr, "Error reading alias: %v\n", err)
			os.Exit(1)
		}
		
		// Check if new alias already exists
		newPath := filepath.Join(akaDir, newName+".alias")
		if _, err := os.Stat(newPath); err == nil {
			fmt.Fprintf(os.Stderr, "Error: alias '%s' already exists\n", newName)
			os.Exit(1)
		}
		
		// Read old alias content
		content, err := os.ReadFile(oldPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading alias '%s': %v\n", oldName, err)
			os.Exit(1)
		}
		
		command := strings.TrimSpace(string(content))
		
		// Create new alias
		if err := addAlias(akaDir, newName, command); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating new alias: %v\n", err)
			os.Exit(1)
		}
		
		// Remove old alias
		if err := os.Remove(oldPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error removing old alias: %v\n", err)
			// Try to cleanup the new alias
			os.Remove(newPath)
			os.Exit(1)
		}
		
		fmt.Printf("Alias renamed: %s → %s\n", oldName, newName)
	},
}
