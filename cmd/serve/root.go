package serve

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/okmall/okmall/cmd/core"
	"github.com/okmall/okmall/models"
	"github.com/okmall/okmall/models/cache"
	"github.com/okmall/okmall/models/model"
	"github.com/spf13/cobra"
	"github.com/unisx/logus"
	"net/http"
	"os"
	"time"
)

var (
	address     string
	port        uint16
	certFile    string
	keyFile     string
	configFile  string
	enableHttps bool
	inDebug     bool
)

func init() {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Start to mirMusic service",
		Long:  "This cmd will start a https server to provide mirMusic service",
		Run:   serveRun,
	}

	// Parse flags for serveCmd
	serveCmd.Flags().StringVarP(&address, "addr", "a", "", "service listen address")
	serveCmd.Flags().Uint16VarP(&port, "port", "p", 0, "port for listen")
	serveCmd.Flags().StringVarP(&certFile, "cert", "e", "", "certificate path used in https connect")
	serveCmd.Flags().StringVarP(&keyFile, "key", "k", "", "key path used in https connect")
	serveCmd.Flags().StringVarP(&configFile, "config", "c", "", "custom config file path used to init application")
	serveCmd.Flags().BoolVarP(&enableHttps, "https", "s", false, "whether use https serve connect")
	serveCmd.Flags().BoolVarP(&inDebug, "debug", "d", false, "whether in debug mode")

	// Register serveCmd as sub-command
	core.Register(serveCmd)
}

func serveRun(cmd *cobra.Command, args []string) {
	// initial some logic
	setup()

	// application config
	config := models.Config()

	// Start Cache service
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		cache.Done()
	}()
	cache.Start(ctx, config)

	// Setup http.Server
	server := &http.Server{
		Handler: newGin(),
		Addr:    config.ServeAddr(),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start http.Server
	var err error
	logus.Info("listen and serve",
		logus.String("schema", config.Serve.Schema),
		logus.String("address", config.Serve.Addr),
		logus.Uint16("port", config.Serve.Port))
	switch config.Serve.Schema {
	case model.SvHttps:
		err = server.ListenAndServeTLS(config.Serve.CertFile, config.Serve.KeyFile)
	case model.SvHttp:
		err = server.ListenAndServe()
	}
	if err != nil {
		logus.E("listen and serve error", err)
	}
}

// setup pre setup before start serve
func setup() {
	if !inDebug {
		logus.InProduction()
		gin.SetMode(gin.ReleaseMode)
	}

	// load config from environment
	envConfig()

	// init models first
	models.InitWith(configFile)

	// override certFile/keyFile from cmd/env config
	config := models.Config()
	if address != "" {
		config.Serve.Addr = address
	}
	if port != 0 {
		config.Serve.Port = port
	}
	if certFile != "" {
		config.Serve.KeyFile = keyFile
	}
	if keyFile != "" {
		config.Serve.CertFile = certFile
	}
	if enableHttps {
		config.Serve.Schema = model.SvHttps
	}

	logus.Debug("config detail", logus.Any("config", config))
}

func envConfig() {
	if configFile == "" {
		configFile = os.Getenv(model.EnvConfigFile)
	}
	if certFile == "" {
		certFile = os.Getenv(model.EnvCertFile)
	}
	if keyFile == "" {
		keyFile = os.Getenv(model.EnvKeyFile)
	}
}
