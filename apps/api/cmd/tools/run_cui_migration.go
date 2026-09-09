package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	connStr := "postgres://dutakasih:dvt4k4s1h@72.62.122.31:5432/navalerp?sslmode=disable"
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	// Read and run tables.sql
	tablesSQL, err := os.ReadFile("d:/stm/VirutalGate/navalerp/db_reference/11_cui_underwater_infrastructure/tables.sql")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read tables.sql: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Applying CUI tables.sql...")
	if _, err := pool.Exec(ctx, string(tablesSQL)); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute tables.sql: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("CUI tables ready!")

	// Read and run insert.sql
	insertSQL, err := os.ReadFile("d:/stm/VirutalGate/navalerp/db_reference/11_cui_underwater_infrastructure/insert.sql")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read insert.sql: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Applying CUI insert.sql...")
	if _, err := pool.Exec(ctx, string(insertSQL)); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute insert.sql: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("CUI seed data applied successfully!")

	// Verify counts
	var aCount, lCount, altCount, iCount int
	_ = pool.QueryRow(ctx, "SELECT COUNT(*) FROM cui_assets;").Scan(&aCount)
	_ = pool.QueryRow(ctx, "SELECT COUNT(*) FROM cui_monitoring_logs;").Scan(&lCount)
	_ = pool.QueryRow(ctx, "SELECT COUNT(*) FROM cui_alerts;").Scan(&altCount)
	_ = pool.QueryRow(ctx, "SELECT COUNT(*) FROM cui_inspections;").Scan(&iCount)

	fmt.Printf("Verification: cui_assets=%d, logs=%d, alerts=%d, inspections=%d\n", aCount, lCount, altCount, iCount)
}
