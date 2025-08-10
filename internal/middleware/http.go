package middleware

import (
	"net/http"

	"github.com/rohankarn35/rate_limiter_golang/pkg/limiter"
)

func RateLimiterMiddleware(l limiter.Limiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !l.Allow() {
			http.Error(w, "Too Many Request", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
