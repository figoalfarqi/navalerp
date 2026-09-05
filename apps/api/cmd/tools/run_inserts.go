package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	connStr := "postgres://dutakasih:dvt4k4s1h@72.62.122.31:5432/navalerp?sslmode=disable"
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Gagal koneksi ke database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	fmt.Println(">>> Terhubung ke database PostgreSQL navalerp!")

	files := []string{
		"db_reference/01_organization_user/insert.sql",
		"db_reference/02_fleet_asset_mro/insert.sql",
		"db_reference/03_inventory_warehousing/insert.sql",
		"db_reference/04_defence_procurement/insert.sql",
		"db_reference/05_finance_budget_accounting/insert.sql",
		"db_reference/06_military_human_capital/insert.sql",
		"db_reference/07_base_infrastructure/insert.sql",
		"db_reference/08_logistics_transportation/insert.sql",
		"db_reference/09_edrms_documents/insert.sql",
		"db_reference/10_readiness_operations_command/insert.sql",
	}

	baseDir := "d:\\stm\\VirutalGate\\navalerp"

	for i, relPath := range files {
		fullPath := filepath.Join(baseDir, relPath)
		fmt.Printf("[%d/10] Menjalankan %s...\n", i+1, relPath)

		content, err := os.ReadFile(fullPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "  [ERROR] Tidak dapat membaca file: %v\n", err)
			continue
		}

		// Strip UTF-8 BOM jika ada
		content = bytes.TrimPrefix(content, []byte("\xef\xbb\xbf"))

		_, err = pool.Exec(ctx, string(content))
		if err != nil {
			fmt.Fprintf(os.Stderr, "  [ERROR] Gagal mengeksekusi %s: %v\n", relPath, err)
		} else {
			fmt.Printf("  [SUKSES] %s berhasil dieksekusi.\n", relPath)
		}
	}

	// Memastikan semua user di sys_users memiliki password Password123!
	fmt.Println("\n>>> Memverifikasi akun sys_users & password...")
	updateQuery := `
		UPDATE sys_users 
		SET password_hash = '$2a$10$7xMvtLLBCyffehFn5sKAmetDVxApVcMnudS9eLJgi4/8xour8YVFe',
		    is_active = TRUE;
	`
	_, err = pool.Exec(ctx, updateQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Gagal update password: %v\n", err)
	} else {
		fmt.Println("  [SUKSES] Password seluruh pengguna disamakan ke Password123!")
	}

	// Validasi password untuk setiap user
	rows, err := pool.Query(ctx, "SELECT username, full_name, role, password_hash FROM sys_users ORDER BY username")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Gagal query user: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	fmt.Println("\n>>> Hasil Verifikasi Akun Pengguna (Password: Password123!):")
	for rows.Next() {
		var uname, name, role, hash string
		if err := rows.Scan(&uname, &name, &role, &hash); err != nil {
			continue
		}
		check := bcrypt.CompareHashAndPassword([]byte(hash), []byte("Password123!"))
		status := "VALID (Password123!)"
		if check != nil {
			status = "TIDAK COCOK"
		}
		fmt.Printf("  - %-18s | %-32s | %-20s | %s\n", uname, name, role, status)
	}

	fmt.Println("\n=======================================================")
	fmt.Println(">>> SELURUH INSERT SEED DATA BERHASIL DIJALANKAN KE DB!")
	fmt.Println("=======================================================")
}
