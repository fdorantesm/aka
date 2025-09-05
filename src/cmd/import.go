package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import [file]",
	Short: "Import aliases from a JSON file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading file:", err)
			os.Exit(1)
		}
		var aliases map[string]string
		if err := json.Unmarshal(data, &aliases); err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing JSON:", err)
			os.Exit(1)
		}
		akaDir := getAkaDir()
		for name, command := range aliases {
			if err := addAlias(akaDir, name, command); err != nil {
				fmt.Fprintln(os.Stderr, "Error writing alias:", err)
				os.Exit(1)
			}
		}
		fmt.Println("Aliases imported successfully.")
	},
}
