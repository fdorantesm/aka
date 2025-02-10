package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [alias] [command]",
	Short: "Add a new alias",
	Long:  "Add a new alias. When no alias and command are provided, the command runs in interactive mode.",
	Run: func(cmd *cobra.Command, args []string) {
		akaDir := getAkaDir()
		if len(args) == 0 {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Alias name: ")
			name, err := reader.ReadString('\n')
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error reading name:", err)
				os.Exit(1)
			}
			name = strings.TrimSpace(name)
			fmt.Print("Command: ")
			userCmd, err := reader.ReadString('\n')
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error reading command:", err)
				os.Exit(1)
			}
			userCmd = strings.TrimSpace(userCmd)
			if err := addAlias(akaDir, name, userCmd); err != nil {
				fmt.Fprintln(os.Stderr, "Error adding alias:", err)
				os.Exit(1)
			}
			fmt.Printf("Alias '%s' added successfully.\n", name)
		} else if len(args) >= 2 {
			name := args[0]
			userCmd := strings.Join(args[1:], " ")
			if err := addAlias(akaDir, name, userCmd); err != nil {
				fmt.Fprintln(os.Stderr, "Error adding alias:", err)
				os.Exit(1)
			}
			fmt.Printf("Alias '%s' added successfully.\n", name)
		} else {
			cmd.Usage()
			os.Exit(1)
		}
	},
}
