package app

import (
	"net/http"

	"github.com/figoalfarqi/navalerp/config"
	"github.com/figoalfarqi/navalerp/internal/app/middleware"
	"github.com/figoalfarqi/navalerp/internal/handler"
)

func SetupRouter(cfg *config.Config) http.Handler {
	mux := http.NewServeMux()
	handlers := buildRouteHandlers(cfg)

	mux.HandleFunc("GET /api/v1/hello", handler.HelloHandler)
	registerApplicationRoutes(mux, cfg, handlers)

	return middleware.Logger(
		middleware.CORS(
			middleware.RateLimiter(mux),
		),
	)
}
