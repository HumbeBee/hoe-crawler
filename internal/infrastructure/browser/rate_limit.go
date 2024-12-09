package browser

import "time"

type RateLimiter struct {
	interval time.Duration
	ticker   *time.Ticker
}

func NewBrowserRateLimiter(interval time.Duration) *RateLimiter {
	return &RateLimiter{
		interval: interval,
		ticker:   time.NewTicker(interval),
	}
}

func (r *RateLimiter) Wait() {
	<-r.ticker.C
}

func (r *RateLimiter) UpdateInterval(interval time.Duration) {
	r.interval = interval
}
