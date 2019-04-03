package mina

import (
	"github.com/alimy/mir"
	"github.com/okmall/okmall/models/core"
	"net/http"
)

// Collect handle collect info
type Collect struct {
	Group       mir.Group `mir:"mina/v1"`
	list        mir.Get   `mir:"/collect/list/"`
	addOrDelete mir.Post  `mir:"/collect/addor/delete/"`

	*core.Context
}

// List get all collect info
func (collect *Collect) List(c Context) {
	// TODO
	c.String(http.StatusOK, "get collect list")
}

// AddOrDelete add or delete collect info
func (collect *Collect) AddOrDelete(c Context) {
	// TODO
	c.String(http.StatusOK, "add or delete collect")
}
