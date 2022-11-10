package cache

import (
	"github.com/bradfitz/gomemcache/memcache"
)

type Cache struct {
	client *memcache.Client
}

func New(server ...string) *Cache {
	return &Cache{
		client: memcache.New(server...),
	}
}

func (c *Cache) Ping() error {
	return c.client.Ping()
}

func (c *Cache) Get(key string) (item *memcache.Item, err error) {
	item, err = c.client.Get(key)
	return
}

func (c *Cache) Add(item *memcache.Item) error {
	return c.client.Add(item)
}

func (c *Cache) Delete(key string) error {
	return c.client.Delete(key)
}
