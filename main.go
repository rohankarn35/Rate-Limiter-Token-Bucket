package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/rohankarn35/rate_limiter_golang/limiter"
)

func main() {

	capacity := 10
	refillRate := 5
	refillInterval := 1 * time.Second
	tb := limiter.NewTokenBucket(capacity, refillRate, refillInterval)
	var wg sync.WaitGroup
	totalRequest := 50
	wg.Add(totalRequest)

	for i := 1; i <= totalRequest; i++ {
		go func(id int) {
			defer wg.Done()
			if tb.Allow() {
				fmt.Println("Request allowed")
			} else {
				fmt.Println("request rejected")
			}

		}(i)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Wait()
	tb.Stop()
}
