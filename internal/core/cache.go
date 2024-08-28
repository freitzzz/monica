package core

import "time"

type CacheValue[T any] struct {
	Value         T
	cacheHitTime  time.Time
	cacheDuration time.Duration
}

func Cached[T any](value T) *CacheValue[T] {
	cd := time.Duration(5) * time.Second

	return &CacheValue[T]{
		Value:         value,
		cacheDuration: cd,
		cacheHitTime:  time.Now().Add(cd),
	}
}

func (c *CacheValue[T]) LookupOrRecache(cb func() T) T {
	if c.cacheHitTime.Before(time.Now()) {
		c.Value = cb()
		c.cacheHitTime = time.Now().Add(c.cacheDuration)
	}

	return c.Value
}
