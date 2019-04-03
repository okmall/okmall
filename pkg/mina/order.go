package mina

import (
	"github.com/alimy/mir"
	"github.com/okmall/okmall/models/core"
	"net/http"
)

// Order handle order info
type Order struct {
	Group   mir.Group `mir:"mina/v1"`
	submit  mir.Post  `mir:"/order/submit/"`
	list    mir.Get   `mir:"/order/list/"`
	detail  mir.Get   `mir:"/order/detail/"`
	cancel  mir.Post  `mir:"/order/cancel/"`
	express mir.Get   `mir:"/order/express/"`

	*core.Context
}

// List get all order info
func (o *Order) List(c Context) {
	// TODO
	c.String(http.StatusOK, "get order list")
}

// Detail get order detail
func (o *Order) Detail(c Context) {
	// TODO
	c.String(http.StatusOK, "get order detail")
}

// Cancel cancel order
func (o *Order) Cancel(c Context) {
	// TODO
	c.String(http.StatusOK, "cancel order")
}

// Express get order express info
func (o *Order) Express(c Context) {
	// TODO
	c.String(http.StatusOK, "get order express")
}

// Submit commit order
func (o *Order) Submit(c Context) {
	// TODO
	c.String(http.StatusOK, "submit order")
}
