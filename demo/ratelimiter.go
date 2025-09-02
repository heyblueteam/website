package demo

import (
	"sync"
	"time"
)

const (
	// MaxRequestsPerIPPerHour is the maximum number of requests per IP per hour
	MaxRequestsPerIPPerHour = 5
	// MaxTotalRequestsPerHour is the maximum total requests per hour
	MaxTotalRequestsPerHour = 100
	// CleanupInterval is how often to clean old entries
	CleanupInterval = 10 * time.Minute
)

// RateLimiter tracks and limits request rates
type RateLimiter struct {
	ipRequests    map[string][]time.Time
	totalRequests []time.Time
	mu            sync.Mutex
}

// NewRateLimiter creates a new rate limiter instance
func NewRateLimiter() *RateLimiter {
	rl := &RateLimiter{
		ipRequests:    make(map[string][]time.Time),
		totalRequests: []time.Time{},
	}
	
	// Start cleanup goroutine
	go rl.cleanup()
	
	return rl
}

// Allow checks if a request from the given IP should be allowed
func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	
	now := time.Now()
	hourAgo := now.Add(-time.Hour)
	
	// Clean old entries for this IP
	rl.ipRequests[ip] = filterRecentTimestamps(rl.ipRequests[ip], hourAgo)
	
	// Clean old total entries
	rl.totalRequests = filterRecentTimestamps(rl.totalRequests, hourAgo)
	
	// Check IP limit
	if len(rl.ipRequests[ip]) >= MaxRequestsPerIPPerHour {
		return false
	}
	
	// Check total limit
	if len(rl.totalRequests) >= MaxTotalRequestsPerHour {
		return false
	}
	
	// Add new request
	rl.ipRequests[ip] = append(rl.ipRequests[ip], now)
	rl.totalRequests = append(rl.totalRequests, now)
	
	return true
}

// cleanup periodically removes old entries to prevent memory growth
func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(CleanupInterval)
	defer ticker.Stop()
	
	for range ticker.C {
		rl.mu.Lock()
		hourAgo := time.Now().Add(-time.Hour)
		
		// Clean IP entries
		for ip, timestamps := range rl.ipRequests {
			filtered := filterRecentTimestamps(timestamps, hourAgo)
			if len(filtered) == 0 {
				delete(rl.ipRequests, ip)
			} else {
				rl.ipRequests[ip] = filtered
			}
		}
		
		// Clean total entries
		rl.totalRequests = filterRecentTimestamps(rl.totalRequests, hourAgo)
		
		rl.mu.Unlock()
	}
}

// filterRecentTimestamps returns only timestamps after the cutoff time
func filterRecentTimestamps(timestamps []time.Time, cutoff time.Time) []time.Time {
	var recent []time.Time
	for _, t := range timestamps {
		if t.After(cutoff) {
			recent = append(recent, t)
		}
	}
	return recent
}

// GetStats returns current rate limiter statistics
func (rl *RateLimiter) GetStats() (ipCount int, totalCount int) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	
	hourAgo := time.Now().Add(-time.Hour)
	
	// Count IPs with recent requests
	for _, timestamps := range rl.ipRequests {
		for _, t := range timestamps {
			if t.After(hourAgo) {
				ipCount++
				break
			}
		}
	}
	
	// Count total recent requests
	for _, t := range rl.totalRequests {
		if t.After(hourAgo) {
			totalCount++
		}
	}
	
	return ipCount, totalCount
}