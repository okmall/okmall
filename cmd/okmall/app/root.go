package app

import (
	"github.com/alimy/logus"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	// Used for flags.
	cfgFile, userLicense string

	rootCmd = &cobra.Command{
		Use:   "okmall",
		Short: "okmall service agent",
		Long:  `okmall service agent`,
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(versionCmd)
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logus.Panic("execute failure", zap.Error(err))
	}
}