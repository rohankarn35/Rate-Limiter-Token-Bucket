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

func (tb *TokenBucket) startRefill() {
	ticker := time.NewTicker(tb.refillInterval)
	for {
		select {
		case <-ticker.C:
			tb.mutex.Lock()
			tb.tokens += tb.refillRate
			if tb.tokens > tb.capacity {
				tb.tokens = tb.capacity
			}
			tb.mutex.Unlock()
		case <-tb.quit:
			ticker.Stop()
			return
		}
	}
}

func (tb *TokenBucket) Allow() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

func (tb *TokenBucket) Stop() {
	close(tb.quit)
}
