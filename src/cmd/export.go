package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export [file]",
	Short: "Export aliases to a JSON file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		akaDir := getAkaDir()
		aliases := make(map[string]string)
		entries, err := os.ReadDir(akaDir)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading aliases:", err)
			os.Exit(1)
		}
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".alias") {
				aliasName := strings.TrimSuffix(entry.Name(), ".alias")
				content, err := os.ReadFile(filepath.Join(akaDir, entry.Name()))
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error reading alias:", err)
					os.Exit(1)
				}
				aliases[aliasName] = strings.TrimSpace(string(content))
			}
		}
		data, err := json.MarshalIndent(aliases, "", "  ")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error encoding JSON:", err)
			os.Exit(1)
		}
		if err := os.WriteFile(args[0], data, 0644); err != nil {
			fmt.Fprintln(os.Stderr, "Error writing file:", err)
			os.Exit(1)
		}
		fmt.Printf("Aliases exported to %s\n", args[0])
	},
}
