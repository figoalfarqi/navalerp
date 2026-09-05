package main

import (
	"log"
	"net/http"

	"github.com/figoalfarqi/apipml/config"
	"github.com/figoalfarqi/apipml/internal/app"
	"github.com/figoalfarqi/apipml/pkg/database"
	"github.com/joho/godotenv"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Config
	cfg := config.New()

	// Init DB
	dbPool := database.NewPostgresPool(cfg.DBURL())
	defer dbPool.Close()

	// Router
	router := app.SetupRouter(cfg)

	addr := ":" + cfg.AppPort
	log.Printf("🚀 Server running at %s\n", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
