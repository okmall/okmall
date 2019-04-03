package mina

import (
	"github.com/alimy/mir"
	"github.com/okmall/okmall/models/core"
	"net/http"
)

// Brand handle brand info
type Brand struct {
	Group  mir.Group `mir:"mina/v1"`
	list   mir.Get   `mir:"/brand/list/"`
	detail mir.Get   `mir:"/brand/detail/"`

	*core.Context
}

// List get brand info
func (b *Brand) List(c Context) {
	// TODO
	c.String(http.StatusOK, "get brand list")
}

// Detail get brand detail info
func (b *Brand) Detail(c Context) {
	// TODO
	c.String(http.StatusOK, "get brand detail")
}
