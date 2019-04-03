package model

import (
	"fmt"
	"net/url"
	"strings"
)

// Config application config model
type Config struct {
	Application Application
	Serve       Serve
	DB          Database `toml:"database"`
	Redis       Redis
}

// Application indicate application section config
type Application struct {
	Name        string
	Version     string
	Authors     []string
	Description string
	License     string
}

// Serve indicate serve section config
type Serve struct {
	Schema   string
	Addr     string
	Port     uint16
	CertFile string
	KeyFile  string
}

// Database indicate database section config
type Database struct {
	Type    string
	Host    string
	Name    string
	User    string
	Passwd  string
	Param   string
	SslMode string
	Path    string
}

// Redis indicate redis section config
type Redis struct {
	Network string
	Addr    string
}

// ServeAddr return serve address
func (c *Config) ServeAddr() string {
	return fmt.Sprintf("%s:%d", c.Serve.Addr, c.Serve.Port)
}

// Dsn return database type and DSN(Database Source Name)
func (c *Config) Dsn() (string, string) {
	switch c.DB.Type {
	case DbMySQL:
		return DbMySQL, c.DB.mysqlDsn()
	case DbPostgreSQL:
		return DbPostgreSQL, c.DB.postgreDsn()
	case DbSqlite3:
		return DbSqlite3, c.DB.Path
	default:
		return DbNone, ""
	}
}

func (d *Database) mysqlDsn() string {
	var param = "?"
	if d.Param != "" {
		param = param + d.Param + "&"
	}
	var dsn string
	if d.Host[0] == '/' { // use unix socket
		dsn = fmt.Sprintf("%s:%s@unix(%s)/%s%scharset=utf8mb4&parseTime=true",
			d.User, d.Passwd, d.Host, d.Name, param)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s%scharset=utf8mb4&parseTime=true",
			d.User, d.Passwd, d.Host, d.Name, param)
	}
	return dsn
}

func (d *Database) postgreDsn() string {
	host, port := "127.0.0.1", "5432"
	if strings.Contains(d.Host, ":") && !strings.HasSuffix(d.Host, "]") {
		idx := strings.LastIndex(d.Host, ":")
		host = d.Host[:idx]
		port = d.Host[idx+1:]
	} else if len(d.Host) > 0 {
		host = d.Host
	}
	var param = "?"
	if d.Param != "" {
		param = param + d.Param + "&"
	}
	var dsn string
	if host[0] == '/' { // use unix socket
		dsn = fmt.Sprintf("postgres://%s:%s@:%s/%s%ssslmode=%s&host=%s",
			url.QueryEscape(d.User), url.QueryEscape(d.Passwd), port, d.Name, param, d.SslMode, host)
	} else {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s%ssslmode=%s",
			url.QueryEscape(d.User), url.QueryEscape(d.Passwd), host, port, d.Name, param, d.SslMode)
	}
	return dsn
}
