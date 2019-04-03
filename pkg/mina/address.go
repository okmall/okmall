package mina

import (
	"github.com/alimy/mir"
	"github.com/okmall/okmall/models/core"
	"net/http"
)

// Address handle address info
type Address struct {
	Group   mir.Group `mir:"mina/v1"`
	list    mir.Get   `mir:"/address/list/"`
	detail  mir.Get   `mir:"/address/detail/"`
	save    mir.Post  `mir:"/address/sava/"`
	delete  mir.Post  `mir:"/address/delete/"`
	regions mir.Get   `mir:"/region/list/"`

	*core.Context
}

// List get address list
func (a *Address) List(c Context) {
	// TODO
	c.String(http.StatusOK, "get address list")
}

// Detail get address detail
func (a *Address) Detail(c Context) {
	// TODO
	c.String(http.StatusOK, "get address detail")
}

// Save save address info
func (a *Address) Save(c Context) {
	// TODO
	c.String(http.StatusOK, "save address")
}

// Delete remove address info
func (a *Address) Delete(c Context) {
	// TODO
	c.String(http.StatusOK, "delete address")
}

// Regions get regions info
func (a *Address) Regions(c Context) {
	// TODO
	c.String(http.StatusOK, "get region list")
}
