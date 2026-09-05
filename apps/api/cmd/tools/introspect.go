package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
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

	rows, err := pool.Query(ctx, `
		SELECT user_id, username, password_hash, full_name, role, auth_version, is_active
		FROM sys_users;
	`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	for rows.Next() {
		var uid, uname, pwd, name, role string
		var authVer int
		var active bool
		if err := rows.Scan(&uid, &uname, &pwd, &name, &role, &authVer, &active); err != nil {
			fmt.Fprintf(os.Stderr, "Scan failed: %v\n", err)
			continue
		}
		fmt.Printf("User: %s | %s | %s\n", uname, name, role)
		for _, testPwd := range []string{"admin", "admin123", "dutakasih", "password", "navalerp", "123456"} {
			if bcrypt.CompareHashAndPassword([]byte(pwd), []byte(testPwd)) == nil {
				fmt.Printf("Found password for %s: '%s'\n", uname, testPwd)
				break
			}
		}
	}
}
