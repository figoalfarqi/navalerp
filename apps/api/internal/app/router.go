package app

import (
	"net/http"

	"github.com/figoalfarqi/apipml/config"
	"github.com/figoalfarqi/apipml/internal/app/middleware"
	"github.com/figoalfarqi/apipml/internal/handler"
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
