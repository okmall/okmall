package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile, userLicense string

	rootCmd = &cobra.Command{
		Use:   "okmall",
		Short: "e-shop service agent",
		Long:  `a e-shop service agent`,
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(versionCmd)
}

// Execute executes the root command.
func Execute() {
	rootCmd.Execute()
}