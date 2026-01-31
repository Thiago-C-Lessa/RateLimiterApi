package limiter

import (
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	Limit           int64
	Tokens          float64
	LastTimeAllowed time.Time
	RefillRate      int64
	sync.Mutex
}

func NewRateLimiter(limit int) *RateLimiter {
	rl := &RateLimiter{
		Limit:           int64(limit),
		Tokens:          float64(limit),
		LastTimeAllowed: time.Now(),
		RefillRate:      int64(limit),
	}

	return rl
}

func (rl *RateLimiter) allow() bool {

	rl.Lock()
	defer rl.Unlock()

	elipsed := time.Since(rl.LastTimeAllowed)
	newTokens := float64(rl.RefillRate) * elipsed.Seconds()

	rl.Tokens = min(rl.Tokens+newTokens, float64(rl.Limit))
	rl.LastTimeAllowed = time.Now()

	if rl.Tokens < 1 {
		return false
	}

	rl.Tokens--
	return true
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !rl.allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
