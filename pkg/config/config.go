package config

import (
	"context"
	"github.com/okmall/okmall/models"
	"github.com/unisx/logus"
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
	if !config.Debug {
		logus.InProduction()
	}
}
