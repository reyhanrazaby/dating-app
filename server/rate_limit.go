package server

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type RateLimiter struct {
	visitors map[string]*rate.Limiter
	mu       sync.Mutex
}

func (r *RateLimiter) addVisitor(ip string) *rate.Limiter {
	limiter := rate.NewLimiter(10, 5) // limit to 10 RPS with a burst of 5
	r.mu.Lock()
	r.visitors[ip] = limiter
	r.mu.Unlock()
	return limiter
}

func (r *RateLimiter) getVisitor(ip string) *rate.Limiter {
	r.mu.Lock()
	limiter, exists := r.visitors[ip]
	if !exists {
		r.mu.Unlock()
		return r.addVisitor(ip)
	}
	r.mu.Unlock()
	return limiter
}

func rateLimit() gin.HandlerFunc {
	rateLimiter := &RateLimiter{
		visitors: make(map[string]*rate.Limiter),
	}
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := rateLimiter.getVisitor(ip)

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"message": "Too many requests"})
			return
		}

		c.Next()
	}
}
