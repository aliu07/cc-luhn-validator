package cache

import (
	"fmt"
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
	mu       sync.RWMutex
	items    map[string]LRUNode
	capacity int

	// LRU queue management
	head *LRUNode
	tail *LRUNode
}

func NewLRUMemCache(size int) *LRUMemCache {
	// LRU management
	head := &LRUNode{}
	tail := &LRUNode{}

	head.next = tail
	tail.prev = head

	cache := &LRUMemCache{
		items:    make(map[string]LRUNode),
		capacity: size,
		head:     head,
		tail:     tail,
	}

	// To clean up expired cache lines
	go cache.cleanup()

	return cache
}

func (c *LRUMemCache) Get(key string) (DataItem, bool) {
	if _, exists := c.items[key]; !exists {
		fmt.Printf("[CACHE] Cache miss for key %s...\n", key)
		return DataItem{}, false
	}

	node := c.items[key]
	c.pop(&node)
	c.insert(&node)

	fmt.Printf("[CACHE] Cache hit for key %s...\n", key)

	return DataItem{
		IsValid:     node.isValid,
		CardNetwork: node.cardNetwork,
	}, true
}

func (c *LRUMemCache) Put(key string, isValid bool, cardNetwork string, ttl time.Duration) {
	// If full
	if len(c.items) == c.capacity {
		toRemove := c.tail.prev

		fmt.Printf("[CACHE] Evicting key %s...\n", toRemove.key)

		c.pop(toRemove)
		delete(c.items, toRemove.key)
	}

	fmt.Printf("[CACHE] Inserting key %s...\n", key)

	newNode := LRUNode{
		key:         key,
		isValid:     isValid,
		cardNetwork: cardNetwork,
		expiration:  time.Now().Add(ttl),
	}
	c.insert(&newNode)
	c.items[key] = newNode
}

func (c *LRUMemCache) cleanup() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.RLock()
		now := time.Now()

		for k, v := range c.items {
			if v.expiration.Before(now) {
				delete(c.items, k)
			}
		}

		c.mu.RUnlock()
	}
}

// Insert at head of LRU DLL
func (c *LRUMemCache) insert(node *LRUNode) {
	l, r := c.head, c.head.next
	l.next = node
	node.next = r
	r.prev = node
	node.prev = l
}

// Pop from tail of LRU DLL
func (c *LRUMemCache) pop(node *LRUNode) {
	l, r := node.prev, node.next
	l.next = r
	r.prev = l
}
