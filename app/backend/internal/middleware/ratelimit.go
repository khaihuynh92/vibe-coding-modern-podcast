package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimiter manages rate limiting per IP
type RateLimiter struct {
	requests map[string][]time.Time
	mutex    sync.RWMutex
	limit    int
	window   time.Duration
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}

	// Start cleanup goroutine
	go rl.cleanup()

	return rl
}

// IsAllowed checks if a request is allowed for the given IP
func (rl *RateLimiter) IsAllowed(ip string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now()
	cutoff := now.Add(-rl.window)

	// Get existing requests for this IP
	requests, exists := rl.requests[ip]
	if !exists {
		requests = make([]time.Time, 0)
	}

	// Remove old requests outside the window
	var validRequests []time.Time
	for _, reqTime := range requests {
		if reqTime.After(cutoff) {
			validRequests = append(validRequests, reqTime)
		}
	}

	// Check if we're under the limit
	if len(validRequests) >= rl.limit {
		rl.requests[ip] = validRequests
		return false
	}

	// Add current request
	validRequests = append(validRequests, now)
	rl.requests[ip] = validRequests

	return true
}

// cleanup removes old entries periodically
func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		rl.mutex.Lock()
		now := time.Now()
		cutoff := now.Add(-rl.window * 2) // Keep some buffer

		for ip, requests := range rl.requests {
			var validRequests []time.Time
			for _, reqTime := range requests {
				if reqTime.After(cutoff) {
					validRequests = append(validRequests, reqTime)
				}
			}

			if len(validRequests) == 0 {
				delete(rl.requests, ip)
			} else {
				rl.requests[ip] = validRequests
			}
		}
		rl.mutex.Unlock()
	}
}

// Global rate limiter instance
var rateLimiter = NewRateLimiter(100, time.Minute) // 100 requests per minute

// RateLimit returns a Gin middleware for rate limiting
func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		if !rateLimiter.IsAllowed(clientIP) {
			c.Header("X-RateLimit-Limit", "100")
			c.Header("X-RateLimit-Remaining", "0")
			c.Header("X-RateLimit-Reset", time.Now().Add(time.Minute).Format(time.RFC3339))
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":   "rate_limit_exceeded",
				"message": "Too many requests. Please try again later.",
				"code":    http.StatusTooManyRequests,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// RateLimitWithConfig returns a Gin middleware with custom rate limiting configuration
func RateLimitWithConfig(limit int, window time.Duration) gin.HandlerFunc {
	rl := NewRateLimiter(limit, window)

	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		if !rl.IsAllowed(clientIP) {
			c.Header("X-RateLimit-Limit", string(rune(limit)))
			c.Header("X-RateLimit-Remaining", "0")
			c.Header("X-RateLimit-Reset", time.Now().Add(window).Format(time.RFC3339))
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":   "rate_limit_exceeded",
				"message": "Too many requests. Please try again later.",
				"code":    http.StatusTooManyRequests,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
