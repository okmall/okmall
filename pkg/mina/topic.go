package mina

import (
	"github.com/alimy/mir"
	"github.com/okmall/okmall/models/core"
	"net/http"
)

// Topic handle topic info
type Topic struct {
	Group   mir.Group `mir:"mina/v1"`
	list    mir.Get   `mir:"/topic/list/"`
	detail  mir.Get   `mir:"/topic/detail/"`
	related mir.Get   `mir:"/topic/related/"`

	*core.Context
}

// List get all topic
func (t *Topic) List(c Context) {
	// TODO
	c.String(http.StatusOK, "get topic list")
}

// Detail get topic detail
func (t *Topic) Detail(c Context) {
	// TODO
	c.String(http.StatusOK, "get topic detail")
}

// Related get topic related info
func (t *Topic) Related(c Context) {
	// TODO
	c.String(http.StatusOK, "get related topic")
}
