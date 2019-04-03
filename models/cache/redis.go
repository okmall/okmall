package cache

import (
	"github.com/mediocregopher/radix/v3"
	"github.com/okmall/okmall/pkg/context"
	"github.com/okmall/okmall/pkg/json"
	"github.com/unisx/logus"

	"net/http"
)

// RedisCache Cache interface implement by redis
type RedisCache struct {
	// TODO
}

// EntryTo write cached entry to gin.Context
func (r *RedisCache) EntryTo(key string, c context.Context) bool {
	if content, err := r.entryFrom(key); err == nil {
		c.Status(http.StatusOK)
		header := c.Writer.Header()
		header.Set("Content-Type", "application/json; charset=utf-8")
		c.Writer.Write(content)
		return true
	}
	return false
}

// CacheFrom cache entry
func (r *RedisCache) CacheFrom(key string, entry interface{}) {
	jsonVal, err := json.Marshal(entry)
	if err != nil {
		logus.E("marshl entry", err)
		return
	}
	cache <- &jsonEntry{key: key, val: string(jsonVal)}
}

func (r *RedisCache) entryFrom(key string) ([]byte, error) {
	var entry []byte
	err := rds.Do(radix.Cmd(&entry, "JSON.GET", key, "."))
	if err == nil && len(entry) == 0 {
		err = errNoExistEntry
	}
	return entry, err
}
