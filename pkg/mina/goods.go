package mina

import (
	"github.com/alimy/mir"
	"github.com/okmall/okmall/models/core"
	"net/http"
)

// Goods handle goods info
type Goods struct {
	Group       mir.Group `mir:"mina/v1"`
	list        mir.Get   `mir:"/goods/list/"`
	count       mir.Get   `mir:"/goods/count/"`
	category    mir.Get   `mir:"/goods/category/"`
	detail      mir.Get   `mir:"/goods/detail/"`
	listNew     mir.Get   `mir:"/goods/new/"`
	listHot     mir.Get   `mir:"/goods/hot/"`
	listRelated mir.Get   `mir:"/goods/related/"`

	*core.Context
}

// List get all goods
func (g *Goods) List(c Context) {
	// TODO
	c.String(http.StatusOK, "get goods list")
}

// Count get goods count
func (g *Goods) Count(c Context) {
	// TODO
	c.String(http.StatusOK, "get goods count")
}

// Category get goods category
func (g *Goods) Category(c Context) {
	// TODO
	c.String(http.StatusOK, "get goods category")
}

// Detail get goods detail
func (g *Goods) Detail(c Context) {
	// TODO
	c.String(http.StatusOK, "get goods detail")
}

// ListNew get new goods
func (g *Goods) ListNew(c Context) {
	// TODO
	c.String(http.StatusOK, "get new goods list")
}

// ListHot get hot goods
func (g *Goods) ListHot(c Context) {
	// TODO
	c.String(http.StatusOK, "get hot goods list")
}

// ListRelated get related goods
func (g *Goods) ListRelated(c Context) {
	// TODO
	c.String(http.StatusOK, "get related goods list")
}
