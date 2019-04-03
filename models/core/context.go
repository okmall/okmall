package core

import (
	"github.com/gin-gonic/gin"
	"github.com/okmall/okmall/models/cache"
	"net/http"
)

// Context core context for gin handler
type Context struct {
	cache.Cache

	Repo Repository
}

// Retrieve write response content to c
func (ctx *Context) Retrieve(c *gin.Context, cacheKey string, action func() (interface{}, error)) {
	if ok := ctx.Cache.EntryTo(cacheKey, c); !ok {
		if data, err := action(); err == nil {
			c.JSON(http.StatusOK, data)
			ctx.CacheFrom(cacheKey, data)
		} else {
			c.JSON(http.StatusInternalServerError, data)
		}
	}
}
