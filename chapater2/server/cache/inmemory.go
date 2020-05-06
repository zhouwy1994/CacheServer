package cache

import (
	"sync"
)

type inMemoryCache struct {
	c map[string][]byte
	mutex sync.RWMutex
	Stat
}

func newInMemoryCache() *inMemoryCache {
	return &inMemoryCache{c:make(map[string][]byte)}
}

func (c *inMemoryCache) Set(k string, v []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	tmp, exist := c.c[k]
	if exist {
		c.del(k, tmp)
	}
	c.c[k] = v
	c.add(k, v)
	return nil
}

func (c *inMemoryCache) Get(k string) ([]byte, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.c[k], nil
}

func (c *inMemoryCache) Del(k string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	temp, exist := c.c[k]
	if exist {
		delete(c.c, k)
		c.del(k, temp)
	}
	return nil
}

func (c *inMemoryCache) GetStat() Stat {
	return c.Stat
}
