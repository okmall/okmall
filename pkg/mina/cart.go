package mina

import (
	"github.com/alimy/mir"
	"github.com/okmall/okmall/models/core"
	"net/http"
)

// Cart handle cart info
type Cart struct {
	Group    mir.Group `mir:"mina/v1"`
	list     mir.Get   `mir:"/cart/index/"`
	add      mir.Post  `mir:"/cart/add/"`
	update   mir.Post  `mir:"/cart/update/"`
	delete   mir.Post  `mir:"/cart/delete/"`
	checked  mir.Post  `mir:"/cart/checked/"`
	count    mir.Get   `mir:"/cart/goodscount/"`
	checkout mir.Post  `mir:"/cart/checkout/"`

	*core.Context
}

// List get all cart info
func (cart *Cart) List(c Context) {
	// TODO
	c.String(http.StatusOK, "get cart list")
}

// Add add a cart info
func (cart *Cart) Add(c Context) {
	// TODO
	c.String(http.StatusOK, "add cart")
}

// Update update cart info
func (cart *Cart) Update(c Context) {
	// TODO
	c.String(http.StatusOK, "update cart")
}

// Delete remove cart info
func (cart *Cart) Delete(c Context) {
	// TODO
	c.String(http.StatusOK, "delete cart")
}

// Checked check cart info
func (cart *Cart) Checked(c Context) {
	// TODO
	c.String(http.StatusOK, "checked cart")
}

// Count get cart count info
func (cart *Cart) Count(c Context) {
	// TODO
	c.String(http.StatusOK, "get cart goods count")
}

// Checkout checkout cart info
func (cart *Cart) Checkout(c Context) {
	// TODO
	c.String(http.StatusOK, "checkout cart")
}
