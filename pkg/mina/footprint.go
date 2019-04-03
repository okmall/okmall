package mina

import (
	"github.com/alimy/mir"
	"github.com/okmall/okmall/models/core"
	"net/http"
)

// Footprint handle footprint info
type Footprint struct {
	Group  mir.Group `mir:"mina/v1"`
	list   mir.Get   `mir:"/footprint/list/"`
	delete mir.Post  `mir:"/footprint/delete/"`

	*core.Context
}

// List get all footprint info
func (f *Footprint) List(c Context) {
	// TODO
	c.String(http.StatusOK, "get footprint list")
}

// Delete remove footprint info
func (f *Footprint) Delete(c Context) {
	// TODO
	c.String(http.StatusOK, "delete footprint")
}
