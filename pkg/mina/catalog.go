package mina

import (
	"github.com/alimy/mir"
	"github.com/okmall/okmall/models/core"
	"net/http"
)

// Catalog handle catalog info
type Catalog struct {
	Group   mir.Group `mir:"mina/v1"`
	list    mir.Get   `mir:"/catalog/index/"`
	current mir.Get   `mir:"/catalog/current/"`

	*core.Context
}

// List get all catalog info
func (catalog *Catalog) List(c Context) {
	// TODO
	c.String(http.StatusOK, "get catalog list")
}

// Current get current catalog info
func (catalog *Catalog) Current(c Context) {
	// TODO
	c.String(http.StatusOK, "get current catalog")
}
