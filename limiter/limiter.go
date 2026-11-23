package limiter

import "time"

type Config struct {
	Limit  int
	Window time.Duration
}

type Counter struct {
	count       int
	windowStart time.Time
}
