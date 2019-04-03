package mina

import (
	"github.com/alimy/mir"
	"github.com/okmall/okmall/models/core"
	"net/http"
)

// Search handler search info
type Search struct {
	Group        mir.Group `mir:"mina/v1"`
	index        mir.Get   `mir:"/search/index/"`
	result       mir.Get   `mir:"/search/result/"`
	helper       mir.Get   `mir:"/search/helper/"`
	clearHistory mir.Get   `mir:"/search/clearhistory/"`

	*core.Context
}

// Index get search info
func (s *Search) Index(c Context) {
	// TODO
	c.String(http.StatusOK, "get search index data")
}

// Result get search result
func (s *Search) Result(c Context) {
	// TODO
	c.String(http.StatusOK, "get search result data")
}

// Helper get search helper
func (s *Search) Helper(c Context) {
	// TODO
	c.String(http.StatusOK, "get search helper data")
}

// ClearHistory get search clear history
func (s *Search) ClearHistory(c Context) {
	// TODO
	c.String(http.StatusOK, "clear search history")
}
