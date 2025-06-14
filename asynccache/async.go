package asynccache

import (
	"awesomeProject1/cachemanager"
	"context"
	"errors"
	"time"
)

var ErrTimeout = errors.New("time expired")

type Cache struct {
	cache *cachemanager.Cache
}

func InitAsyncCache(defaultExpiration, cleanupInterval time.Duration) *Cache {
	return &Cache{
		cache: cachemanager.InitCache(defaultExpiration, cleanupInterval),
	}
}

func (c *Cache) Get(ctx context.Context, key string) (interface{}, error) {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		v, ok := c.cache.Get(key)
		if ok {
			ch <- v
		}
	}()

	select {
	case <-ctx.Done():
		return nil, ErrTimeout
	case x, ok := <-ch:
		if ok {
			return x, nil
		}
		return nil, errors.New("not found")
	}
}

func (c *Cache) Add(ctx context.Context, key string, value interface{}) error {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		c.cache.Set(key, value, 0)
	}()

	select {
	case <-ctx.Done():
		return ErrTimeout
	case <-ch:
		return nil
	}
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		ok := c.cache.Delete(key)
		if ok {
			ch <- ok
		}
	}()

	select {
	case <-ctx.Done():
		return ErrTimeout
	case ok := <-ch:
		if ok.(bool) {
			return nil
		}
		return errors.New("not found")
	}
}
