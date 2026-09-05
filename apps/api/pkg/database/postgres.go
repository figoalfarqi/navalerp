package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func NewPostgresPool(dbURL string) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	p, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}
	if err := p.Ping(ctx); err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	if _, err := p.Exec(ctx, `
		ALTER TABLE app_user
		ADD COLUMN IF NOT EXISTS auth_version INTEGER NOT NULL DEFAULT 0
	`); err != nil {
		p.Close()
		log.Fatalf("Unable to ensure authentication schema: %v", err)
	}
	pool = p
	log.Println("✅ Connected to PostgreSQL")
	return pool
}

func GetPool() *pgxpool.Pool {
	return pool
}
