package app

import (
	"net/http"

	"github.com/figoalfarqi/navalerp/config"
	"github.com/figoalfarqi/navalerp/internal/app/middleware"
)

func SetupRouter(cfg *config.Config) http.Handler {
	mux := http.NewServeMux()
	handlers := buildRouteHandlers(cfg)

	mux.HandleFunc("GET /api/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"NavalERP API v1 is operational"}`))
	})
	registerApplicationRoutes(mux, cfg, handlers)

	return middleware.Logger(
		middleware.CORS(
			middleware.RateLimiter(mux),
		),
	)
}
