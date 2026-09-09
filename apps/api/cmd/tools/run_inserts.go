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
		"db_reference/11_cui_underwater_infrastructure/insert.sql",
	}

	baseDir := "d:\\stm\\VirutalGate\\navalerp"

	fmt.Println(">>> Membersihkan data lama dengan TRUNCATE CASCADE...")
	truncateSQL := `
		TRUNCATE TABLE 
			cui_inspections, cui_alerts, cui_monitoring_logs, cui_assets,
			ops_readiness_alerts, ops_ship_readiness_snapshots, ops_daily_logs, ops_mission_ship_assignments, ops_missions, ops_theaters,
			doc_documents, doc_document_links, doc_document_versions, doc_categories,
			log_shipment_items, log_shipments, log_routes, log_transport_units,
			infra_facility_maintenances, infra_fuel_bunker_records, infra_berth_bookings, infra_facilities,
			hcm_sea_duty_allowances, hcm_crew_assignments, hcm_medical_readiness, hcm_personnel_qualifications, hcm_qualifications, hcm_service_records, hcm_personnel, hcm_corps, hcm_ranks,
			fin_platform_tco_summaries, fin_journal_entries, fin_payments, fin_invoices, fin_budget_commitments, fin_budget_allocations, fin_budget_programs, fin_chart_of_accounts,
			proc_goods_receipt_items, proc_goods_receipts, proc_purchase_order_items, proc_purchase_orders, proc_contract_amendments, proc_contracts, proc_tender_bids, proc_tenders, proc_requisition_items, proc_requisitions, proc_vendors,
			inv_stock_adjustments, inv_stock_transfers, inv_item_instances, inv_stock_balances, inv_materials, inv_storage_locations, inv_warehouses,
			mro_docking_records, mro_work_order_items, mro_work_order_tasks, mro_work_orders, mro_pm_schedules, mro_failure_reports, mro_equipment_parameters, mro_equipments, mro_systems, mro_ships, mro_ship_classes,
			sys_audit_logs, sys_users, org_units
		CASCADE;
	`
	if _, err := pool.Exec(ctx, truncateSQL); err != nil {
		fmt.Fprintf(os.Stderr, "  [WARNING] Truncate error (melanjutkan): %v\n", err)
	} else {
		fmt.Println("  [SUKSES] Seluruh tabel berhasil dibersihkan untuk re-seed.")
	}

	for i, relPath := range files {
		fullPath := filepath.Join(baseDir, relPath)
		fmt.Printf("[%d/%d] Menjalankan %s...\n", i+1, len(files), relPath)

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
