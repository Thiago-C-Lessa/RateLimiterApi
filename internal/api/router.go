package api

import (
	"net/http"

	"RateLimiterApi/internal/limiter"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func NewRouter(rl *limiter.RateLimiter, allowedOrigins []string) http.Handler {
	r := chi.NewRouter()

	// Configurar CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(middleware.Logger)

	// Middleware de rate limiter
	r.Use(rl.Middleware)

	// Rotas
	r.Get("/ping", PingHandler)

	return r
}
