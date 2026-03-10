package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jot",
	Short: "A simple note-taking CLI",
	Long:  `job is a CLI tool to add, read, and list your notes specific to the directory you're working in'.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(readCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(removeCmd)
}
