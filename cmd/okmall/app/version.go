package app

import (
	"fmt"
	"github.com/okmall/okmall/version"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of OkMall",
	Long:  `All software has versions. This is OkMall's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Version)
	},
}
