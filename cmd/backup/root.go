package backup

import (
	"github.com/okmall/okmall/cmd/core"
	"github.com/spf13/cobra"
	"github.com/unisx/logus"
)

const (
	remoteDefault = "127.0.0.1:8013" // default remote address
)

var (
	remoteAddr string
	fileName   string
	inDebug    bool
)

func init() {
	importCmd := &cobra.Command{
		Use:   "import",
		Short: "Import data to remote okmall service",
		Long:  "Import data to remote okmall service",
		Run:   importRun,
	}

	exportCmd := &cobra.Command{
		Use:   "export",
		Short: "Export data from remote okmall service data",
		Long:  "Export data from remote okmall service data",
		Run:   exportRun,
	}

	// Parse flags for importCmd
	importCmd.Flags().StringVarP(&remoteAddr, "remote", "r", remoteDefault, "remote kmlet service listen address")
	importCmd.Flags().StringVarP(&fileName, "file", "f", "", "content file name that will be import to kmlet service")
	importCmd.Flags().BoolVarP(&inDebug, "debug", "d", false, "whether in debug mode")

	// Parse flags for exportCmd
	exportCmd.Flags().StringVarP(&remoteAddr, "remote", "r", remoteDefault, "remote kmlet service listen address")
	exportCmd.Flags().StringVarP(&fileName, "file", "f", "", "content file name that will be export from kmlet service")
	exportCmd.Flags().BoolVarP(&inDebug, "debug", "d", false, "whether in debug mode")

	// Register importCmd/exportCmd as sub-command
	core.Register(importCmd)
	core.Register(exportCmd)
}

func importRun(cmd *cobra.Command, args []string) {
	setup()
	importFrom(fileName, remoteAddr)
}

func exportRun(cmd *cobra.Command, args []string) {
	setup()
	exportTo(fileName, remoteAddr)
}

// setup pre setup before start serve
func setup() {
	if !inDebug {
		logus.InProduction()
	}
}
