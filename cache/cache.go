package cache

import "time"

// Cacher 缓存接口
type Cacher interface {
	Init()
	Get(key string, target interface{}) bool
	GetString(key string) string
	GetBool(key string) bool
	GetFloat64(key string) float64
	GetInt(key string) int
	Exists(key string) bool
	Delete(key string)
	// Set 放置值
	Set(key string, value interface{})
	SetExpired(key string, value interface{}, t time.Duration)
	Close()
}
