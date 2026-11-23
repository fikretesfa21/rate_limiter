package limiter

import (
	"testing"
	"time"
)

func TestAllowsRequestsWithinLimit(t *testing.T) {
	rl := New(3, time.Minute)

	for i := 1; i <= 3; i++ {
		if !rl.Allow() {
			t.Errorf("request %d should be allowed but was denied", i)
		}
	}
}

func TestBlocksRequestsOverLimit(t *testing.T) {
	rl := New(3, time.Minute)

	for i := 0; i < 3; i++ {
		rl.Allow()
	}

	if rl.Allow() {
		t.Error("4th request should be blocked but was allowed")
	}
}

func TestWindowResetsAfterExpiry(t *testing.T) {
	rl := New(2, 50*time.Millisecond)

	rl.Allow()
	rl.Allow()

	if rl.Allow() {
		t.Fatal("should be blocked before window resets")
	}

	time.Sleep(60 * time.Millisecond)

	if !rl.Allow() {
		t.Error("should be allowed after window resets")
	}
}
