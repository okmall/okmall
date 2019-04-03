package models

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/okmall/okmall/models/cache"
	"github.com/okmall/okmall/models/core"
	"github.com/okmall/okmall/models/model"
	"github.com/okmall/okmall/models/mysql"
	"github.com/okmall/okmall/models/sqlite3"
	"github.com/okmall/okmall/pkg/assets"
	"github.com/unisx/logus"
	"os"
)

var (
	config *model.Config
	repo   core.Repository
)

// InitWith initialize models with custom config file
func InitWith(path string) {
	config = &model.Config{}

	// init config
	if err := loadConfig(config); err != nil {
		logus.F("load conifg error", err)
	}
	if path == "" {
		// Empty
	} else if fileIsExist(path) { // set config from custom config file
		customConfig(config, path)
	} else {
		logus.Info("custom config file is not exist so use default config",
			logus.String("path", path))
	}

	// init Repository
	initRepository()
}

// Config return application's config
func Config() *model.Config {
	return config
}

// NewContext return new core.Context instance
func NewContext() *core.Context {
	return &core.Context{
		Cache: &cache.RedisCache{},
		Repo:  repo,
	}
}

// fileIsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func fileIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func initRepository() {
	var err error
	switch config.DB.Type {
	case model.DbMySQL:
		repo, err = mysql.NewRepository(config)
	case model.DbSqlite3:
		repo, err = sqlite3.NewRepository(config)
	default:
		err = fmt.Errorf("unknown database type: %s", config.DB.Type)
	}
	if err != nil {
		logus.F("init repository error", err)
	}
}

func loadConfig(config *model.Config) error {
	data := assets.MustAsset("assets/config/app.toml")
	_, err := toml.Decode(string(data), config)
	return err
}

func customConfig(config *model.Config, path string) {
	customConfig := &model.Config{}

	meta, err := toml.DecodeFile(path, customConfig)
	if err != nil {
		logus.E("decode custom config error", err)
		return
	}
	for _, key := range meta.Keys() {
		if len(key) == 1 { // top section just continue
			continue
		}
		switch key[0] {
		case "application":
			customAppSection(config, customConfig, key)
		case "serve":
			customServeSection(config, customConfig, key)
		case "database":
			customDbSection(config, customConfig, key)
		case "redis":
			customRedisSection(config, customConfig, key)
		}
	}
}

func customAppSection(config *model.Config, custom *model.Config, key toml.Key) {
	switch key[1] {
	case "name":
		config.Application.Name = custom.Application.Name
	case "version":
		config.Application.Version = custom.Application.Version
	case "authors":
		config.Application.Authors = custom.Application.Authors
	case "description":
		config.Application.Description = custom.Application.Description
	case "license":
		config.Application.License = custom.Application.License
	}
}

func customServeSection(config *model.Config, custom *model.Config, key toml.Key) {
	switch key[1] {
	case "schema":
		config.Serve.Schema = custom.Serve.Schema
	case "addr":
		config.Serve.Addr = custom.Serve.Addr
	case "port":
		config.Serve.Port = custom.Serve.Port
	case "certFile":
		config.Serve.CertFile = custom.Serve.CertFile
	case "keyFile":
		config.Serve.KeyFile = custom.Serve.KeyFile
	}
}

func customDbSection(config *model.Config, custom *model.Config, key toml.Key) {
	switch key[1] {
	case "type":
		config.DB.Type = custom.DB.Type
	case "host":
		config.DB.Host = custom.DB.Host
	case "user":
		config.DB.User = custom.DB.User
	case "passwd":
		config.DB.Passwd = custom.DB.Passwd
	case "param":
		config.DB.Param = custom.DB.Param
	case "sslMode":
		config.DB.SslMode = custom.DB.SslMode
	case "path":
		config.DB.Path = custom.DB.Path
	}
}

func customRedisSection(config *model.Config, custom *model.Config, key toml.Key) {
	switch key[1] {
	case "network":
		config.Redis.Network = custom.Redis.Network
	case "addr":
		config.Redis.Addr = custom.Redis.Addr
	}
}
