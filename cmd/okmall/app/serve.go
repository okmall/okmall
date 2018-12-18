package app

import (
	"context"
	"github.com/alimy/logus"
	_ "github.com/gin-gonic/gin"
	"github.com/okmall/okmall/pkg/config"
	"github.com/spf13/cobra"
)

var AppContext = context.Background()

const (
	certFilePathDefault = "cert.pem" // certificate file default path
	keyFilePathDefault  = "key.pem"  // key file used in https server default path
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start to serve okmall service",
	Long:  "this cmd will start a https server to provide okmall service",
	Run:   serveRun,
}

var (
	muxType     string
	certFile    string
	keyFile     string
	enableHttps bool
)

func init() {
	serveCmd.Flags().StringVarP(&certFile, "cert", "c", certFilePathDefault, "certificate path used in https connect")
	serveCmd.Flags().StringVarP(&keyFile, "key", "k", keyFilePathDefault, "key path used in https connect")
	serveCmd.Flags().BoolVarP(&enableHttps, "https", "s", false, "whether use https serve connect")
}

func serveRun(cmd *cobra.Command, args []string) {
	// TODO
	config.LoadConfig(&AppContext, config.DefaultConfigFilePath)
	logus.Debug("start serve okmall service")
}