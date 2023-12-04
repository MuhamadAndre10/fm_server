package cache

import "github.com/allegro/bigcache/v3"

type DataCache struct {
	cache *bigcache.BigCache
}

var Cache *DataCache

func NewDataCache(c *bigcache.BigCache) *DataCache {
	return &DataCache{
		cache: c,
	}
}

func NewCache(c *DataCache) {
	Cache = c
}

func (d *DataCache) Set(key string, value []byte) error {
	return d.cache.Set(key, value)
}

func (d *DataCache) Get(key string) ([]byte, error) {
	return d.cache.Get(key)
}
