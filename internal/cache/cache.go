package cache

import "time"

type DataItem struct {
	IsValid     bool
	CardNetwork string
}

type Cache interface {
	Get(key string) (DataItem, bool)
	Put(key string, isValid bool, cardNetwork string, ttl time.Duration)
}
