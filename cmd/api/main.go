package main

import (
	"log"
	"net/http"
	"runtime"

	"RateLimiterApi/internal/api"
	"RateLimiterApi/internal/config"
	"RateLimiterApi/internal/limiter"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	cfg := config.LoadConfig()

	rl := limiter.NewRateLimiter(cfg.RateLimitPerSecond)

	r := api.NewRouter(rl, cfg.CORSAllowedOrigins)

	log.Printf("Server running on %s\n", cfg.Port)
	if err := http.ListenAndServe(cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
