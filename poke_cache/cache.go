package poke_cache

import (
	"sync"
	"time"
)

type Cache struct {
    caches map[string]CacheEntry
    ttl time.Duration
    mu *sync.RWMutex
}

type CacheEntry struct {
    at time.Time
    val []byte
}

func (cache *Cache) Add(url string, val []byte) {
    cache.mu.Lock()
    defer cache.mu.Unlock()

    cache.caches[url] = CacheEntry{
        time.Now(),
        val,
    }
}

func (cache *Cache) Get(url string) ([]byte, bool) {
    cache.mu.RLock()
    defer cache.mu.RUnlock()

    if cache.caches[url].val != nil {
        return cache.caches[url].val, true
    }

    return nil, false
}

func (cache *Cache) reapLoop() {
    interval := time.Tick(cache.ttl)
    for range interval {
        for name, entry := range cache.caches {
            if time.Now().Sub(entry.at) > cache.ttl {
                delete(cache.caches, name)
            }
        }
    }
}

func NewCache(interval time.Duration) Cache  {
    cache := Cache {
        make(map[string]CacheEntry),
        interval, 
        &sync.RWMutex{},
    }
    
    go cache.reapLoop()

    return cache
}



