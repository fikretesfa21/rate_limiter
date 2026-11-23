package limiter

import (
	"sync"
	"time"
)

type Config struct {
	Limit  int
	Window time.Duration
}

type Counter struct {
	mu          sync.Mutex
	count       int
	windowStart time.Time
}

func (c *Counter) resetIfExpired(window time.Duration) {
	if time.Since(c.windowStart) >= window {
		c.count = 0
		c.windowStart = time.Now()
	}
}

func (c *Counter) Allow(cfg Config) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.resetIfExpired(cfg.Window)

	if c.count < cfg.Limit {
		c.count++
		return true
	}

	return false
}
