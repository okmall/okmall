package mina

import (
	"github.com/alimy/mir"
	"github.com/okmall/okmall/models/core"
	"net/http"
)

// Wx handle wei xin info
type Wx struct {
	Group         mir.Group `mir:"mina/v1"`
	authLoginByWx mir.Post  `mir:"/auth/loginByWeixin/"`
	payPrepayId   mir.Get   `mir:"/pay/prepay/"`

	*core.Context
}

// AuthLoginByWx login by wx
func (w *Wx) AuthLoginByWx(c Context) {
	// TODO
	c.String(http.StatusOK, "auth login by weixin")
}

// PayPrepayId get pay prepay id
func (w *Wx) PayPrepayId(c Context) {
	// TODO
	c.String(http.StatusOK, "get pay prepay id")
}
