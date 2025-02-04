package browser

import (
	"math/rand"
	"time"
)

type RateLimiter struct {
	interval time.Duration
	ticker   *time.Ticker
}

func NewBrowserRateLimiter(interval time.Duration) *RateLimiter {
	// Get a random duration between -5s and +5s
	// rand.Int63n(11) - 5  // 0-10 minus 5 gives us -5 to +5
	jitter := time.Duration(rand.Int63n(11)-5) * time.Second

	// Make sure we don't go negative
	finalInterval := interval
	if interval > 5*time.Second {
		finalInterval = interval + jitter
	} else {
		// For small intervals, only add positive jitter
		finalInterval = interval + time.Duration(rand.Int63n(5))*time.Second
	}

	return &RateLimiter{
		interval: finalInterval,
		ticker:   time.NewTicker(interval),
	}
}

func (r *RateLimiter) Wait() {
	<-r.ticker.C
}

func (r *RateLimiter) UpdateInterval(interval time.Duration) {
	r.interval = interval
}
