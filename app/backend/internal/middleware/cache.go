package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// CacheEntry represents a cached response
type CacheEntry struct {
	Data      []byte
	Timestamp time.Time
	TTL       time.Duration
}

// CacheManager manages in-memory cache
type CacheManager struct {
	cache map[string]*CacheEntry
	mutex sync.RWMutex
}

// NewCacheManager creates a new cache manager
func NewCacheManager() *CacheManager {
	cm := &CacheManager{
		cache: make(map[string]*CacheEntry),
	}

	// Start cleanup goroutine
	go cm.cleanup()

	return cm
}

// Get retrieves a value from cache
func (cm *CacheManager) Get(key string) ([]byte, bool) {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	entry, exists := cm.cache[key]
	if !exists {
		return nil, false
	}

	// Check if entry has expired
	if time.Since(entry.Timestamp) > entry.TTL {
		return nil, false
	}

	return entry.Data, true
}

// Set stores a value in cache
func (cm *CacheManager) Set(key string, data []byte, ttl time.Duration) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	cm.cache[key] = &CacheEntry{
		Data:      data,
		Timestamp: time.Now(),
		TTL:       ttl,
	}
}

// Delete removes a value from cache
func (cm *CacheManager) Delete(key string) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	delete(cm.cache, key)
}

// cleanup removes expired entries periodically
func (cm *CacheManager) cleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		cm.mutex.Lock()
		now := time.Now()
		for key, entry := range cm.cache {
			if now.Sub(entry.Timestamp) > entry.TTL {
				delete(cm.cache, key)
			}
		}
		cm.mutex.Unlock()
	}
}

// Global cache manager instance
var cacheManager = NewCacheManager()

// Cache returns a Gin middleware for caching responses
func Cache(ttl time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Only cache GET requests
		if c.Request.Method != "GET" {
			c.Next()
			return
		}

		// Create cache key from request path and query parameters
		cacheKey := c.Request.URL.Path
		if c.Request.URL.RawQuery != "" {
			cacheKey += "?" + c.Request.URL.RawQuery
		}

		// Check if response is cached
		if cachedData, exists := cacheManager.Get(cacheKey); exists {
			c.Header("X-Cache", "HIT")
			c.Header("Content-Type", "application/json")
			c.Data(200, "application/json", cachedData)
			c.Abort()
			return
		}

		// Create a custom response writer to capture the response
		writer := &responseWriter{
			ResponseWriter: c.Writer,
			body:           make([]byte, 0),
		}
		c.Writer = writer

		// Process the request
		c.Next()

		// Cache the response if it was successful
		if c.Writer.Status() == 200 && len(writer.body) > 0 {
			cacheManager.Set(cacheKey, writer.body, ttl)
			c.Header("X-Cache", "MISS")
		}
	}
}

// responseWriter captures the response body
type responseWriter struct {
	gin.ResponseWriter
	body []byte
}

func (w *responseWriter) Write(data []byte) (int, error) {
	w.body = append(w.body, data...)
	return w.ResponseWriter.Write(data)
}

func (w *responseWriter) WriteString(s string) (int, error) {
	w.body = append(w.body, []byte(s)...)
	return w.ResponseWriter.WriteString(s)
}
