package demo

import (
	"fmt"
	"testing"
	"time"
)

func TestRateLimiter_IPLimit(t *testing.T) {
	rl := NewRateLimiter()
	testIP := "192.168.1.1"
	
	// First 5 requests should pass
	for i := 0; i < 5; i++ {
		if !rl.Allow(testIP) {
			t.Errorf("Request %d should have been allowed", i+1)
		}
	}
	
	// 6th request should fail
	if rl.Allow(testIP) {
		t.Error("6th request should have been blocked")
	}
}

func TestRateLimiter_TotalLimit(t *testing.T) {
	rl := NewRateLimiter()
	
	// Simulate requests from many different IPs
	// Each IP gets only 1 request, but total should hit 100
	for i := 0; i < 100; i++ {
		ip := fmt.Sprintf("192.168.1.%d", i+1)
		if !rl.Allow(ip) {
			t.Errorf("Request %d should have been allowed", i+1)
		}
	}
	
	// 101st request should fail even from new IP
	if rl.Allow("10.0.0.1") {
		t.Error("101st request should have been blocked due to total limit")
	}
}

func TestRateLimiter_TimeWindow(t *testing.T) {
	// This test would take too long to run in practice
	// but demonstrates the concept
	t.Skip("Skipping time-based test for speed")
	
	rl := NewRateLimiter()
	testIP := "192.168.1.1"
	
	// Use up the limit
	for i := 0; i < 5; i++ {
		rl.Allow(testIP)
	}
	
	// Should be blocked
	if rl.Allow(testIP) {
		t.Error("Should be blocked after 5 requests")
	}
	
	// Wait for window to expire
	time.Sleep(1 * time.Hour)
	
	// Should be allowed again
	if !rl.Allow(testIP) {
		t.Error("Should be allowed after time window expires")
	}
}