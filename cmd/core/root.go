package core

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:  "okmall",
		Long: "okmall service",
	}
)

// Setup set root command name,short-describe, long-describe
// return &cobra.Command to custom other options
func Setup(use, short, long string) *cobra.Command {
	rootCmd.Use = use
	rootCmd.Short = short
	rootCmd.Long = long
	return rootCmd
}

// Register add sub-command
func Register(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

// Execute start application
func Execute() {
	rootCmd.Execute()
}
