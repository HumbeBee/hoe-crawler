package browser

import (
	"math/rand"
	"time"
)

type RateLimiter struct {
	interval time.Duration
	//ticker   *time.Ticker
	timer *time.Timer
}

func NewBrowserRateLimiter(interval time.Duration) *RateLimiter {
	return &RateLimiter{
		interval: interval,
		//ticker:   time.NewTicker(interval),
		timer: time.NewTimer(interval),
	}
}

func (r *RateLimiter) Wait() {
	<-r.timer.C

	// Calculate new random interval
	jitter := time.Duration(rand.Int63n(11)-5) * time.Second
	nextInterval := r.interval
	if r.interval > 5*time.Second {
		nextInterval = r.interval + jitter
	} else {
		nextInterval = r.interval + time.Duration(rand.Int63n(6))*time.Second
	}

	// Reset timer with new random interval
	r.timer.Reset(nextInterval)
}

func (r *RateLimiter) UpdateInterval(interval time.Duration) {
	r.interval = interval
}
