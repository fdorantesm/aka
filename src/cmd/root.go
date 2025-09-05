package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "aka",
	Short: "A simple alias manager",
	Long:  "A simple alias manager that lets you add, list, and apply command aliases.",
}

func Execute() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(applyCmd)
	rootCmd.AddCommand(exportCmd)
	rootCmd.AddCommand(importCmd)
	rootCmd.AddCommand(versionCmd)
	cobra.CheckErr(rootCmd.Execute())
}
