package cmd

import (
	"context"
	"github.com/alimy/logus"
	_ "github.com/gin-gonic/gin"
	"github.com/okmall/okmall/pkg/config"
)

var AppContext = context.Background()

func Run() {
	config.LoadConfig(&AppContext, config.ConfigFilePath)
	logus.Info("okmall start...")
	// TODO
}
