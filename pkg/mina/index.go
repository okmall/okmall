package mina

import (
	"github.com/alimy/mir"
	"github.com/okmall/okmall/models/core"
)

// Index handle index info
type Index struct {
	Group mir.Group `mir:"mina/v1"`
	list  mir.Get   `mir:"/index/index/"`

	*core.Context
}

// List get index info
func (i *Index) List(c Context) {
	i.Retrieve(c, core.RdsMainPage, i.Repo.GetMainPage)
}
