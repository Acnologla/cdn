package cache

import (
	"github.com/Acnologla/cdn/internal/core/domain"
	"github.com/Acnologla/cdn/internal/core/port"
	lru "github.com/hashicorp/golang-lru/v2"
)

type Cache struct {
	cache *lru.Cache[string, *domain.File]
}

func (c *Cache) Get(key string) (*domain.File, bool) {
	return c.cache.Get(key)
}

func (c *Cache) Set(key string, value *domain.File) {
	c.cache.Add(key, value)
}

func (c *Cache) Delete(key string) {
	c.cache.Remove(key)
}

func (c *Cache) Clear() {
	c.cache.Purge()
}

func NewLRUCache(max int) port.Cache {
	l, _ := lru.New[string, *domain.File](max)
	return &Cache{
		cache: l,
	}
}
