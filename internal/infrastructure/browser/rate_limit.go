package browser

import "time"

type RateLimiter struct {
	interval time.Duration
	ticker   *time.Ticker
}

func NewRateLimiter(requestsPerSecond float64) *RateLimiter {
	interval := time.Duration(1000/requestsPerSecond) * time.Millisecond
	return &RateLimiter{
		interval: interval,
		ticker:   time.NewTicker(interval),
	}
}

func (r *RateLimiter) Wait() {
	<-r.ticker.C
}
