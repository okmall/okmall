package config

import (
	"context"
	"github.com/alimy/logus"
	"github.com/okmall/okmall/models"
	"go.uber.org/zap"
)

const (
	DefaultConfigFilePath = "config/app.toml"

	KeyConfig int = iota
)

var ConfigFilePath = DefaultConfigFilePath

func LoadConfig(ctx *context.Context, configFilePath string) {
	// TODO
	config := &models.Config{Debug: true}
	*ctx = context.WithValue(*ctx, KeyConfig, config)
	loadConfig(config)

}

func loadConfig(config *models.Config) {
	if config.Debug {
		logus.InDevelopment(zap.Development(), zap.AddCaller(), zap.AddCallerSkip(1))
	} else {
		logus.InProduction()
	}
}
