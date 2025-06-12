package main

import "time"

func main() {
	// Example usage
	limiter := NewRateLimiter(5, 10*time.Second)

	clientID := "client1"
	for i := 0; i < 10; i++ {
		if limiter.AllowRequest(clientID) {
			println("Request allowed")
		} else {
			println("Rate limit exceeded")
		}
		time.Sleep(1 * time.Second) // Simulate time between requests
	}
	// Simulate waiting for the time window to pass
	time.Sleep(11 * time.Second)
	if limiter.AllowRequest(clientID) {
		println("Request allowed after waiting")
	} else {
		println("Rate limit still exceeded after waiting")
	}
}

type ClientInfo struct {
	ID           string
	RequestTimes []time.Time
}

type RateLimiter struct {
	Clients map[string]*ClientInfo
	Limit   int
	Window  time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		Clients: make(map[string]*ClientInfo),
		Limit:   limit,
		Window:  window,
	}
}

// Returns true if request is allowed, false if rate limit exceeded
func (r *RateLimiter) AllowRequest(clientID string) bool {
    now := time.Now()
    client, exists := r.Clients[clientID]
    if !exists {
        client = &ClientInfo{ID: clientID, RequestTimes: []time.Time{}}
        r.Clients[clientID] = client
    }

    // Remove expired timestamps directly from the front of the slice
    for len(client.RequestTimes) > 0 && now.Sub(client.RequestTimes[0]) > r.Window {
        client.RequestTimes = client.RequestTimes[1:] // Remove the oldest timestamp
    }

    if len(client.RequestTimes) >= r.Limit {
        return false
    }

    // Allow this request, record timestamp
    client.RequestTimes = append(client.RequestTimes, now)
    return true
}