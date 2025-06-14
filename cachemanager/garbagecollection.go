package cachemanager

import "time"

func (c *Cache) GC() {
	for {
		<-time.After(c.cleanupInterval)

		if c.items == nil {
			return
		}

		if keys := c.expiredKeys(); len(keys) > 0 {
			c.clearItems(keys)
		}
	}
}

func (c *Cache) StartGC() {
	go c.GC()
}

func (c *Cache) expiredKeys() (keys []string) {
	c.RLock()
	defer c.RUnlock()

	for k, v := range c.items {
		if time.Now().UnixNano() > v.Expiration && v.Expiration > 0 {
			keys = append(keys, k)
		}
	}

	return
}

func (c *Cache) clearItems(keys []string) {
	c.Lock()
	defer c.Unlock()

	for _, k := range keys {
		delete(c.items, k)
	}
}
