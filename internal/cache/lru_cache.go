package cache

import (
	"log"
	"sync"
	"time"
)

type LRUNode struct {
	key         string
	isValid     bool
	cardNetwork string
	expiration  time.Time
	prev        *LRUNode
	next        *LRUNode
}

type LRUMemCache struct {
	mu          sync.RWMutex
	cacheLogger *log.Logger
	items       map[string]LRUNode
	capacity    int
	// LRU queue management
	head *LRUNode
	tail *LRUNode
}

func NewLRUMemCache(size int, cacheLogger *log.Logger) *LRUMemCache {
	// LRU management
	head := &LRUNode{}
	tail := &LRUNode{}

	head.next = tail
	tail.prev = head

	cache := &LRUMemCache{
		cacheLogger: cacheLogger,
		items:       make(map[string]LRUNode),
		capacity:    size,
		head:        head,
		tail:        tail,
	}

	// To clean up expired cache lines
	go cache.cleanup()

	return cache
}

func (cache *LRUMemCache) Get(key string) (DataItem, bool) {
	if _, exists := cache.items[key]; !exists {
		cache.cacheLogger.Printf("Cache miss for key %s...\n", key)
		return DataItem{}, false
	}

	node := cache.items[key]
	cache.pop(&node)
	cache.insert(&node)

	cache.cacheLogger.Printf("Cache hit for key %s...\n", key)

	return DataItem{
		IsValid:     node.isValid,
		CardNetwork: node.cardNetwork,
	}, true
}

func (cache *LRUMemCache) Put(key string, isValid bool, cardNetwork string, ttl time.Duration) {
	// If full
	if len(cache.items) == cache.capacity {
		toRemove := cache.tail.prev

		cache.cacheLogger.Printf("Evicting key %s...\n", toRemove.key)

		cache.pop(toRemove)
		delete(cache.items, toRemove.key)
	}

	cache.cacheLogger.Printf("Inserting key %s...\n", key)

	newNode := LRUNode{
		key:         key,
		isValid:     isValid,
		cardNetwork: cardNetwork,
		expiration:  time.Now().Add(ttl),
	}

	cache.insert(&newNode)
	cache.items[key] = newNode
}

func (cache *LRUMemCache) cleanup() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		cache.mu.RLock()
		now := time.Now()

		for key, val := range cache.items {
			if val.expiration.Before(now) {
				delete(cache.items, key)
			}
		}

		cache.mu.RUnlock()
	}
}

// Insert at head of LRU DLL
func (cache *LRUMemCache) insert(node *LRUNode) {
	l, r := cache.head, cache.head.next
	l.next = node
	node.next = r
	r.prev = node
	node.prev = l
}

// Pop from tail of LRU DLL
func (cache *LRUMemCache) pop(node *LRUNode) {
	l, r := node.prev, node.next
	l.next = r
	r.prev = l
}
