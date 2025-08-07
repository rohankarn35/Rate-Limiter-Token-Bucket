package limiter

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity       int
	tokens         int
	refillRate     int
	refillInterval time.Duration
	mutex          sync.Mutex
	quit           chan struct{}
}

func (t *TokenBucket) startRefill() {
	ticker := time.NewTicker(t.refillInterval)
	for {
		select {
		case <-ticker.C:
			t.mutex.Lock()
			t.tokens += t.refillRate
			if t.tokens > t.capacity {
				t.tokens = t.capacity
			}
			t.mutex.Unlock()
		case <-t.quit:
			ticker.Stop()
			return
		}
	}
}
func (t *TokenBucket) Allow() bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if t.tokens > 0 {
		t.tokens--
		return true
	}
	return false
}

func (t *TokenBucket) Stop() {
	close(t.quit)
}

func NewTokenBucket(capacity, refillRate int, refillInterval time.Duration) *TokenBucket {
	tb := &TokenBucket{
		capacity:       capacity,
		tokens:         capacity,
		refillRate:     refillRate,
		refillInterval: refillInterval,
		quit:           make(chan struct{}),
	}
	go tb.startRefill()
	return tb
}
