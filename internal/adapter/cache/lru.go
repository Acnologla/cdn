package cache

import (
	"github.com/Acnologla/cdn/internal/core/port"
	lru "github.com/hashicorp/golang-lru/v2"
)

type Cache struct {
	cache *lru.Cache[string, []byte]
}

func (c *Cache) Get(key string) ([]byte, bool) {
	return c.cache.Get(key)
}

func (c *Cache) Set(key string, value []byte) {
	c.cache.Add(key, value)
}

func (c *Cache) Delete(key string) {
	c.cache.Remove(key)
}

func (c *Cache) Clear() {
	c.cache.Purge()
}

func New() port.Cache {
	l, _ := lru.New[string, []byte](512)
	return &Cache{
		cache: l,
	}
}
