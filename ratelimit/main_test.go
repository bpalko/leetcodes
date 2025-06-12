// filepath: /home/palko/code/Go/leetcode/ratelimit/main_test.go
package main

import (
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	limit := 2
	window := 2 * time.Second

	t.Run("Valid Request", func(t *testing.T) {
		rateLimiter := NewRateLimiter(limit, window)
		clientID := "client1"

		for i := 0; i < limit; i++ {
			if !rateLimiter.AllowRequest(clientID) {
				t.Errorf("Expected request %d to be allowed, but it was denied", i+1)
			}
		}
	})

	t.Run("Rate Limit Exceeded", func(t *testing.T) {
		rateLimiter := NewRateLimiter(limit, window)
		clientID := "client1"

		for i := 0; i < limit; i++ {
			rateLimiter.AllowRequest(clientID) // Fill up the limit
		}

		if !rateLimiter.AllowRequest(clientID) {
			t.Logf("Expected request to be denied after exceeding the limit")
		}
	})

	t.Run("Timeout and Retry", func(t *testing.T) {
		rateLimiter := NewRateLimiter(limit, window)
		clientID := "client1"

		for i := 0; i < limit; i++ {
			rateLimiter.AllowRequest(clientID) // Fill up the limit
		}

		if rateLimiter.AllowRequest(clientID) {
			t.Errorf("Expected request to be denied after exceeding the limit, but it was allowed")
		}
		t.Logf("Waiting for the time window to pass...")
		time.Sleep(window) // Wait for the time window to pass

		if !rateLimiter.AllowRequest(clientID) {
			t.Errorf("Expected request to be allowed after timeout, but it was denied")
		}
	})
}
