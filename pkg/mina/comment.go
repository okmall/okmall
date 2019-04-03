package mina

import (
	"github.com/alimy/mir"
	"github.com/okmall/okmall/models/core"
	"net/http"
)

// Comment handle comment info
type Comment struct {
	Group mir.Group `mir:"mina/v1"`
	list  mir.Get   `mir:"/comment/list/"`
	count mir.Get   `mir:"/comment/count/"`
	post  mir.Post  `mir:"/comment/post/"`

	*core.Context
}

// List get comment info
func (comment *Comment) List(c Context) {
	// TODO
	c.String(http.StatusOK, "get comment list")
}

// Count get comment count
func (comment *Comment) Count(c Context) {
	// TODO
	c.String(http.StatusOK, "get comment count")
}

// Post add a comment
func (comment *Comment) Post(c Context) {
	// TODO
	c.String(http.StatusOK, "post comment")
}
