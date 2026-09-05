package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ColumnInfo struct {
	Name        string
	DataType    string
	UDTName     string
	IsNullable  bool
	Default     *string
	IsPK        bool
	IsGenerated bool
}

type ChildTableConfig struct {
	TableName string
	FKColumn  string
	StructKey string
}

type EntityConfig struct {
	EntityName  string // e.g. "work_order"
	TableName   string // e.g. "mro_work_orders"
	PKColumn    string // e.g. "work_order_id"
	ModuleNum   string // e.g. "02"
	ModuleName  string // e.g. "Armada & MRO"
	Title       string // e.g. "Perintah Kerja MRO"
	PluralTitle string // e.g. "Daftar Perintah Kerja MRO"
	Children    []ChildTableConfig
}

var entities = []EntityConfig{
	// Module 1: Organisasi & Pengguna
	{
		EntityName: "org_unit", TableName: "org_units", PKColumn: "unit_id",
		ModuleNum: "01", ModuleName: "Organisasi & Pengguna", Title: "Satuan Kerja", PluralTitle: "Daftar Satuan Kerja",
	},
	{
		EntityName: "sys_user", TableName: "sys_users", PKColumn: "user_id",
		ModuleNum: "01", ModuleName: "Organisasi & Pengguna", Title: "Pengguna Sistem", PluralTitle: "Daftar Pengguna Sistem",
	},
	{
		EntityName: "audit_log", TableName: "sys_audit_logs", PKColumn: "log_id",
		ModuleNum: "01", ModuleName: "Organisasi & Pengguna", Title: "Audit Log", PluralTitle: "Log Aktivitas Sistem",
	},

	// Module 2: Armada Kapal & MRO
	{
		EntityName: "ship_class", TableName: "mro_ship_classes", PKColumn: "class_id",
		ModuleNum: "02", ModuleName: "Armada Kapal & MRO", Title: "Kelas Kapal", PluralTitle: "Daftar Kelas Kapal",
	},
	{
		EntityName: "ship", TableName: "mro_ships", PKColumn: "ship_id",
		ModuleNum: "02", ModuleName: "Armada Kapal & MRO", Title: "Kapal Perang KRI", PluralTitle: "Daftar Kapal Perang KRI",
	},
	{
		EntityName: "ship_system", TableName: "mro_systems", PKColumn: "system_id",
		ModuleNum: "02", ModuleName: "Armada Kapal & MRO", Title: "Sistem Kapal", PluralTitle: "Daftar Sistem Kapal",
	},
	{
		EntityName: "equipment", TableName: "mro_equipments", PKColumn: "equipment_id",
		ModuleNum: "02", ModuleName: "Armada Kapal & MRO", Title: "Peralatan Mesin", PluralTitle: "Daftar Peralatan & Sensor",
		Children: []ChildTableConfig{
			{TableName: "mro_equipment_parameters", FKColumn: "equipment_id", StructKey: "Parameters"},
		},
	},
	{
		EntityName: "pm_schedule", TableName: "mro_pm_schedules", PKColumn: "pm_id",
		ModuleNum: "02", ModuleName: "Armada Kapal & MRO", Title: "Jadwal PMS", PluralTitle: "Jadwal Pemeliharaan Preventif",
	},
	{
		EntityName: "failure_report", TableName: "mro_failure_reports", PKColumn: "report_id",
		ModuleNum: "02", ModuleName: "Armada Kapal & MRO", Title: "Laporan Kerusakan", PluralTitle: "Laporan Kerusakan Alutsista",
	},
	{
		EntityName: "work_order", TableName: "mro_work_orders", PKColumn: "work_order_id",
		ModuleNum: "02", ModuleName: "Armada Kapal & MRO", Title: "Perintah Kerja MRO", PluralTitle: "Perintah Kerja MRO",
		Children: []ChildTableConfig{
			{TableName: "mro_work_order_tasks", FKColumn: "work_order_id", StructKey: "Tasks"},
			{TableName: "mro_work_order_items", FKColumn: "work_order_id", StructKey: "Items"},
		},
	},
	{
		EntityName: "docking_record", TableName: "mro_docking_records", PKColumn: "docking_id",
		ModuleNum: "02", ModuleName: "Armada Kapal & MRO", Title: "Riwayat Docking", PluralTitle: "Riwayat Docking Galangan",
	},

	// Module 3: Logistik & Pergudangan
	{
		EntityName: "warehouse", TableName: "inv_warehouses", PKColumn: "warehouse_id",
		ModuleNum: "03", ModuleName: "Logistik & Pergudangan", Title: "Gudang Militer", PluralTitle: "Daftar Gudang Militer",
		Children: []ChildTableConfig{
			{TableName: "inv_storage_locations", FKColumn: "warehouse_id", StructKey: "Locations"},
		},
	},
	{
		EntityName: "material", TableName: "inv_materials", PKColumn: "material_id",
		ModuleNum: "03", ModuleName: "Logistik & Pergudangan", Title: "Material & Suku Cadang", PluralTitle: "Katalog Suku Cadang & NSN",
		Children: []ChildTableConfig{
			{TableName: "inv_material_equipment_links", FKColumn: "material_id", StructKey: "EquipmentLinks"},
		},
	},
	{
		EntityName: "stock_balance", TableName: "inv_stock_balances", PKColumn: "balance_id",
		ModuleNum: "03", ModuleName: "Logistik & Pergudangan", Title: "Saldo Stok", PluralTitle: "Saldo Stok Bebekal",
	},
	{
		EntityName: "item_instance", TableName: "inv_item_instances", PKColumn: "instance_id",
		ModuleNum: "03", ModuleName: "Logistik & Pergudangan", Title: "Nomor Seri / Batch", PluralTitle: "Nomor Seri & Batch Amunisi",
	},
	{
		EntityName: "stock_transfer", TableName: "inv_stock_transfers", PKColumn: "transfer_id",
		ModuleNum: "03", ModuleName: "Logistik & Pergudangan", Title: "Transfer Bebekal", PluralTitle: "Transfer Antar Gudang/KRI",
		Children: []ChildTableConfig{
			{TableName: "inv_stock_transfer_items", FKColumn: "transfer_id", StructKey: "Items"},
		},
	},
	{
		EntityName: "stock_adjustment", TableName: "inv_stock_adjustments", PKColumn: "adjustment_id",
		ModuleNum: "03", ModuleName: "Logistik & Pergudangan", Title: "Penyesuaian Stok", PluralTitle: "Stock Opname & Adjustment",
		Children: []ChildTableConfig{
			{TableName: "inv_stock_adjustment_items", FKColumn: "adjustment_id", StructKey: "Items"},
		},
	},

	// Module 4: Pengadaan Pertahanan
	{
		EntityName: "vendor", TableName: "proc_vendors", PKColumn: "vendor_id",
		ModuleNum: "04", ModuleName: "Pengadaan Pertahanan", Title: "Rekanan Industri", PluralTitle: "Rekanan Industri Pertahanan",
		Children: []ChildTableConfig{
			{TableName: "proc_vendor_ratings", FKColumn: "vendor_id", StructKey: "Ratings"},
		},
	},
	{
		EntityName: "requisition", TableName: "proc_requisitions", PKColumn: "requisition_id",
		ModuleNum: "04", ModuleName: "Pengadaan Pertahanan", Title: "Permintaan Pengadaan", PluralTitle: "Purchase Requisition (PR)",
		Children: []ChildTableConfig{
			{TableName: "proc_requisition_items", FKColumn: "requisition_id", StructKey: "Items"},
		},
	},
	{
		EntityName: "tender", TableName: "proc_tenders", PKColumn: "tender_id",
		ModuleNum: "04", ModuleName: "Pengadaan Pertahanan", Title: "Tender & Lelang", PluralTitle: "Tender & Pengadaan Alutsista",
		Children: []ChildTableConfig{
			{TableName: "proc_tender_bids", FKColumn: "tender_id", StructKey: "Bids"},
		},
	},
	{
		EntityName: "contract", TableName: "proc_contracts", PKColumn: "contract_id",
		ModuleNum: "04", ModuleName: "Pengadaan Pertahanan", Title: "Kontrak Militer", PluralTitle: "Kontrak Pengadaan Militer",
		Children: []ChildTableConfig{
			{TableName: "proc_contract_amendments", FKColumn: "contract_id", StructKey: "Amendments"},
		},
	},
	{
		EntityName: "purchase_order", TableName: "proc_purchase_orders", PKColumn: "po_id",
		ModuleNum: "04", ModuleName: "Pengadaan Pertahanan", Title: "Purchase Order", PluralTitle: "Purchase Order (PO)",
		Children: []ChildTableConfig{
			{TableName: "proc_purchase_order_items", FKColumn: "po_id", StructKey: "Items"},
		},
	},
	{
		EntityName: "goods_receipt", TableName: "proc_goods_receipts", PKColumn: "receipt_id",
		ModuleNum: "04", ModuleName: "Pengadaan Pertahanan", Title: "Penerimaan BAPHP", PluralTitle: "Berita Acara Penerimaan (BAPHP)",
		Children: []ChildTableConfig{
			{TableName: "proc_goods_receipt_items", FKColumn: "receipt_id", StructKey: "Items"},
		},
	},

	// Module 5: Keuangan & Anggaran
	{
		EntityName: "chart_of_account", TableName: "fin_chart_of_accounts", PKColumn: "account_id",
		ModuleNum: "05", ModuleName: "Keuangan & Anggaran", Title: "Bagan Akun (COA)", PluralTitle: "Bagan Akun Standar Militer",
	},
	{
		EntityName: "budget_program", TableName: "fin_budget_programs", PKColumn: "program_id",
		ModuleNum: "05", ModuleName: "Keuangan & Anggaran", Title: "Program DIPA", PluralTitle: "Program Anggaran DIPA",
		Children: []ChildTableConfig{
			{TableName: "fin_budget_allocations", FKColumn: "program_id", StructKey: "Allocations"},
		},
	},
	{
		EntityName: "budget_commitment", TableName: "fin_budget_commitments", PKColumn: "commitment_id",
		ModuleNum: "05", ModuleName: "Keuangan & Anggaran", Title: "Komitmen Anggaran", PluralTitle: "Komitmen Pagu Anggaran",
	},
	{
		EntityName: "invoice", TableName: "fin_invoices", PKColumn: "invoice_id",
		ModuleNum: "05", ModuleName: "Keuangan & Anggaran", Title: "Tagihan Rekanan", PluralTitle: "Tagihan & Faktur Rekanan",
	},
	{
		EntityName: "payment", TableName: "fin_payments", PKColumn: "payment_id",
		ModuleNum: "05", ModuleName: "Keuangan & Anggaran", Title: "Pembayaran SP2D", PluralTitle: "Pembayaran & SP2D",
	},
	{
		EntityName: "journal_entry", TableName: "fin_journal_entries", PKColumn: "journal_id",
		ModuleNum: "05", ModuleName: "Keuangan & Anggaran", Title: "Jurnal Akuntansi", PluralTitle: "Jurnal Akuntansi Militer",
		Children: []ChildTableConfig{
			{TableName: "fin_journal_lines", FKColumn: "journal_id", StructKey: "Lines"},
		},
	},
	{
		EntityName: "platform_tco", TableName: "fin_platform_tco_summaries", PKColumn: "tco_id",
		ModuleNum: "05", ModuleName: "Keuangan & Anggaran", Title: "Total Cost of Ownership", PluralTitle: "Analisis TCO Alutsista",
	},

	// Module 6: SDM Militer (HCM)
	{
		EntityName: "military_rank", TableName: "hcm_ranks", PKColumn: "rank_id",
		ModuleNum: "06", ModuleName: "SDM Militer (HCM)", Title: "Pangkat Militer", PluralTitle: "Jenjang Kepangkatan TNI AL",
	},
	{
		EntityName: "military_corps", TableName: "hcm_corps", PKColumn: "corps_id",
		ModuleNum: "06", ModuleName: "SDM Militer (HCM)", Title: "Korps Militer", PluralTitle: "Daftar Korps TNI AL",
	},
	{
		EntityName: "qualification", TableName: "hcm_qualifications", PKColumn: "qualification_id",
		ModuleNum: "06", ModuleName: "SDM Militer (HCM)", Title: "Kualifikasi & Brevet", PluralTitle: "Standar Brevet & Sertifikasi",
	},
	{
		EntityName: "personnel", TableName: "hcm_personnel", PKColumn: "personnel_id",
		ModuleNum: "06", ModuleName: "SDM Militer (HCM)", Title: "Prajurit TNI AL", PluralTitle: "Data Induk Prajurit TNI AL",
		Children: []ChildTableConfig{
			{TableName: "hcm_service_records", FKColumn: "personnel_id", StructKey: "ServiceRecords"},
			{TableName: "hcm_personnel_qualifications", FKColumn: "personnel_id", StructKey: "Qualifications"},
			{TableName: "hcm_medical_readiness", FKColumn: "personnel_id", StructKey: "MedicalReadiness"},
		},
	},
	{
		EntityName: "crew_assignment", TableName: "hcm_crew_assignments", PKColumn: "assignment_id",
		ModuleNum: "06", ModuleName: "SDM Militer (HCM)", Title: "Awak KRI", PluralTitle: "Manifest Awak Kapal KRI",
		Children: []ChildTableConfig{
			{TableName: "hcm_sea_duty_allowances", FKColumn: "assignment_id", StructKey: "Allowances"},
		},
	},

	// Module 7: Pangkalan & Infrastruktur
	{
		EntityName: "base_facility", TableName: "infra_facilities", PKColumn: "facility_id",
		ModuleNum: "07", ModuleName: "Pangkalan & Fasilitas", Title: "Fasilitas Pangkalan", PluralTitle: "Dermaga & Fasilitas Pangkalan",
		Children: []ChildTableConfig{
			{TableName: "infra_facility_maintenances", FKColumn: "facility_id", StructKey: "Maintenances"},
		},
	},
	{
		EntityName: "berth_booking", TableName: "infra_berth_bookings", PKColumn: "booking_id",
		ModuleNum: "07", ModuleName: "Pangkalan & Fasilitas", Title: "Penjadwalan Sandar", PluralTitle: "Jadwal Sandar & Labuh KRI",
	},
	{
		EntityName: "fuel_bunker", TableName: "infra_fuel_bunker_records", PKColumn: "bunker_id",
		ModuleNum: "07", ModuleName: "Pangkalan & Fasilitas", Title: "Bunker BBM", PluralTitle: "Catatan Bunkering BBM Laut",
	},

	// Module 8: Transportasi Militer
	{
		EntityName: "transport_unit", TableName: "log_transport_units", PKColumn: "transport_unit_id",
		ModuleNum: "08", ModuleName: "Transportasi Militer", Title: "Armada Transportasi", PluralTitle: "Armada Angkut Militer",
	},
	{
		EntityName: "route", TableName: "log_routes", PKColumn: "route_id",
		ModuleNum: "08", ModuleName: "Transportasi Militer", Title: "Rute Pelayaran", PluralTitle: "Rute Pelayaran Logistik",
	},
	{
		EntityName: "shipment", TableName: "log_shipments", PKColumn: "shipment_id",
		ModuleNum: "08", ModuleName: "Transportasi Militer", Title: "Pengiriman Konvoi", PluralTitle: "Manifest Pengiriman Bebekal",
		Children: []ChildTableConfig{
			{TableName: "log_shipment_items", FKColumn: "shipment_id", StructKey: "Items"},
		},
	},

	// Module 9: EDRMS Dokumen Digital
	{
		EntityName: "document_category", TableName: "doc_categories", PKColumn: "category_id",
		ModuleNum: "09", ModuleName: "EDRMS Dokumen Digital", Title: "Kategori Dokumen", PluralTitle: "Kategori Dokumen Digital",
	},
	{
		EntityName: "document", TableName: "doc_documents", PKColumn: "document_id",
		ModuleNum: "09", ModuleName: "EDRMS Dokumen Digital", Title: "Dokumen Militer", PluralTitle: "Arsip & Dokumen Elektronik",
		Children: []ChildTableConfig{
			{TableName: "doc_document_versions", FKColumn: "document_id", StructKey: "Versions"},
			{TableName: "doc_document_links", FKColumn: "document_id", StructKey: "Links"},
		},
	},

	// Module 10: Kesiapan Operasi & Komando Tempur
	{
		EntityName: "theater", TableName: "ops_theaters", PKColumn: "theater_id",
		ModuleNum: "10", ModuleName: "Komando & Kesiapan", Title: "Teater Operasi", PluralTitle: "Wilayah Operasi & Pertahanan",
	},
	{
		EntityName: "mission", TableName: "ops_missions", PKColumn: "mission_id",
		ModuleNum: "10", ModuleName: "Komando & Kesiapan", Title: "Misi Tempur", PluralTitle: "Operasi & Misi Tempur KRI",
		Children: []ChildTableConfig{
			{TableName: "ops_mission_ship_assignments", FKColumn: "mission_id", StructKey: "AssignedShips"},
		},
	},
	{
		EntityName: "daily_log", TableName: "ops_daily_logs", PKColumn: "log_id",
		ModuleNum: "10", ModuleName: "Komando & Kesiapan", Title: "Log Harian KRI", PluralTitle: "Logbook Navigasi & Operasi KRI",
	},
	{
		EntityName: "readiness_report", TableName: "ops_ship_readiness_snapshots", PKColumn: "snapshot_id",
		ModuleNum: "10", ModuleName: "Komando & Kesiapan", Title: "Kesiapan KRI", PluralTitle: "Indeks Kesiapan KRI C-1..C-4",
	},
	{
		EntityName: "readiness_alert", TableName: "ops_readiness_alerts", PKColumn: "alert_id",
		ModuleNum: "10", ModuleName: "Komando & Kesiapan", Title: "Peringatan Dini", PluralTitle: "Peringatan Dini Kesiapan CASREP",
	},
}

func findRepoRoot() string {
	dir, _ := os.Getwd()
	for {
		if _, err := os.Stat(filepath.Join(dir, "db_reference")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "."
}

func toPascalCase(s string) string {
	parts := strings.Split(s, "_")
	var res string
	for _, p := range parts {
		if len(p) > 0 {
			res += strings.ToUpper(p[:1]) + strings.ToLower(p[1:])
		}
	}
	return res
}

func toCamelCase(s string) string {
	p := toPascalCase(s)
	if len(p) == 0 {
		return ""
	}
	return strings.ToLower(p[:1]) + p[1:]
}

func cleanChildStructName(tableName string) string {
	s := tableName
	for _, prefix := range []string{"mro_", "inv_", "proc_", "fin_", "hcm_", "infra_", "log_", "doc_", "ops_"} {
		s = strings.TrimPrefix(s, prefix)
	}
	return toPascalCase(s)
}

func postgresToGoType(col ColumnInfo) (goType string, jsonTag string) {
	jsonTag = fmt.Sprintf("`json:\"%s\"`", col.Name)
	if col.IsNullable && col.Name != "deleted_at" && col.Name != "deleted_by" {
		jsonTag = fmt.Sprintf("`json:\"%s,omitempty\"`", col.Name)
	}

	dt := strings.ToLower(col.DataType)
	udt := strings.ToLower(col.UDTName)

	switch {
	case dt == "integer" || dt == "int" || udt == "int4" || udt == "smallint":
		if col.IsNullable {
			goType = "*int"
		} else {
			goType = "int"
		}
	case dt == "bigint" || udt == "int8":
		if col.IsNullable {
			goType = "*int64"
		} else {
			goType = "int64"
		}
	case dt == "numeric" || dt == "decimal" || dt == "real" || dt == "double precision" || udt == "float8" || udt == "float4":
		if col.IsNullable {
			goType = "*float64"
		} else {
			goType = "float64"
		}
	case dt == "boolean" || udt == "bool":
		if col.IsNullable {
			goType = "*bool"
		} else {
			goType = "bool"
		}
	case dt == "date" || strings.Contains(dt, "timestamp"):
		if col.IsNullable {
			goType = "*time.Time"
		} else {
			goType = "time.Time"
		}
	case dt == "json" || dt == "jsonb":
		goType = "any"
	default: // varchar, text, uuid, user-defined enums
		if col.IsNullable {
			goType = "*string"
		} else {
			goType = "string"
		}
	}
	return
}

func main() {
	connStr := "postgres://dutakasih:dvt4k4s1h@72.62.122.31:5432/navalerp?sslmode=disable"
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	query := `
		SELECT 
			c.table_name,
			c.column_name,
			c.data_type,
			c.udt_name,
			(c.is_nullable = 'YES') AS is_nullable,
			c.column_default,
			COALESCE(pk.is_pk, false) AS is_pk,
			(COALESCE(c.is_generated, 'NEVER') = 'ALWAYS') AS is_generated
		FROM information_schema.columns c
		LEFT JOIN (
			SELECT ku.table_name, ku.column_name, true AS is_pk
			FROM information_schema.table_constraints tc
			JOIN information_schema.key_column_usage ku 
				ON tc.constraint_name = ku.constraint_name AND tc.table_schema = ku.table_schema
			WHERE tc.constraint_type = 'PRIMARY KEY' AND tc.table_schema = 'public'
		) pk ON c.table_name = pk.table_name AND c.column_name = pk.column_name
		WHERE c.table_schema = 'public'
		ORDER BY c.table_name, c.ordinal_position;
	`
	rows, err := pool.Query(ctx, query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Columns query failed: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	tableCols := make(map[string][]ColumnInfo)
	for rows.Next() {
		var (
			tableName string
			col       ColumnInfo
		)
		if err := rows.Scan(&tableName, &col.Name, &col.DataType, &col.UDTName, &col.IsNullable, &col.Default, &col.IsPK, &col.IsGenerated); err != nil {
			fmt.Fprintf(os.Stderr, "Columns scan failed: %v\n", err)
			os.Exit(1)
		}
		tableCols[tableName] = append(tableCols[tableName], col)
	}

	fmt.Printf("Loaded metadata for %d tables.\n", len(tableCols))

	repoRoot := findRepoRoot()
	baseApiDir := filepath.Join(repoRoot, "apps", "api", "internal")
	baseWebDir := filepath.Join(repoRoot, "apps", "web", "src", "app", "admin", "data")
	navDir := filepath.Join(repoRoot, "apps", "web", "src", "components", "admin")
	_ = navDir

	fmt.Println("Target base API dir:", baseApiDir)
	fmt.Println("Target base Web dir:", baseWebDir)

	for _, ent := range entities {
		generateModel(baseApiDir, ent, tableCols)
		generateRepository(baseApiDir, ent, tableCols)
		generateService(baseApiDir, ent, tableCols)
		generateHandler(baseApiDir, ent, tableCols)
		generateFrontend(baseWebDir, ent, tableCols)
	}

	generateRouter(baseApiDir, entities)
	// generateNavigation(navDir, entities)

	fmt.Println("=== FULL MODULAR GENERATION COMPLETE ===")
}

func generateModel(baseDir string, ent EntityConfig, tableCols map[string][]ColumnInfo) {
	cols := tableCols[ent.TableName]
	structName := toPascalCase(ent.EntityName)

	var sb strings.Builder
	sb.WriteString("package model\n\nimport (\n\t\"time\"\n)\n\n")

	// Main Struct
	sb.WriteString(fmt.Sprintf("type %s struct {\n", structName))
	for _, col := range cols {
		fieldName := toPascalCase(col.Name)
		goType, jsonTag := postgresToGoType(col)
		sb.WriteString(fmt.Sprintf("\t%s %s %s\n", fieldName, goType, jsonTag))
	}

	// Attached children
	for _, ch := range ent.Children {
		childStructName := cleanChildStructName(ch.TableName)
		sb.WriteString(fmt.Sprintf("\t%s []%s `json:\"%s,omitempty\"`\n", ch.StructKey, childStructName, strings.ToLower(ch.StructKey)))
	}
	sb.WriteString("}\n\n")

	// Child structs if any
	for _, ch := range ent.Children {
		childCols := tableCols[ch.TableName]
		childStructName := cleanChildStructName(ch.TableName)
		sb.WriteString(fmt.Sprintf("type %s struct {\n", childStructName))
		for _, col := range childCols {
			fieldName := toPascalCase(col.Name)
			goType, jsonTag := postgresToGoType(col)
			sb.WriteString(fmt.Sprintf("\t%s %s %s\n", fieldName, goType, jsonTag))
		}
		sb.WriteString("}\n\n")
	}

	// Request Struct
	sb.WriteString(fmt.Sprintf("type %sRequest struct {\n", structName))
	for _, col := range cols {
		if col.IsPK || col.IsGenerated || col.Name == "created_at" || col.Name == "updated_at" || col.Name == "deleted_at" ||
			col.Name == "created_by" || col.Name == "updated_by" || col.Name == "deleted_by" {
			continue
		}
		fieldName := toPascalCase(col.Name)
		goType, _ := postgresToGoType(col)
		if !col.IsNullable && !strings.HasPrefix(goType, "*") && goType != "any" {
			goType = "*" + goType
		}
		sb.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\"`\n", fieldName, goType, col.Name))
	}
	for _, ch := range ent.Children {
		childStructName := cleanChildStructName(ch.TableName)
		sb.WriteString(fmt.Sprintf("\t%s []%s `json:\"%s\"`\n", ch.StructKey, childStructName, strings.ToLower(ch.StructKey)))
	}
	sb.WriteString("}\n")

	filePath := filepath.Join(baseDir, "model", ent.EntityName+".go")
	os.WriteFile(filePath, []byte(sb.String()), 0644)
}

func generateRepository(baseDir string, ent EntityConfig, tableCols map[string][]ColumnInfo) {
	cols := tableCols[ent.TableName]
	structName := toPascalCase(ent.EntityName)
	repoName := structName + "Repository"

	var sb strings.Builder
	sb.WriteString("package repository\n\nimport (\n\t\"context\"\n\t\"fmt\"\n\t\"strings\"\n\n\t\"github.com/figoalfarqi/navalerp/internal/model\"\n\t\"github.com/jackc/pgx/v5/pgxpool\"\n)\n\n")

	sb.WriteString(fmt.Sprintf("type %s struct {\n\tDB *pgxpool.Pool\n}\n\n", repoName))
	sb.WriteString(fmt.Sprintf("func New%s(db *pgxpool.Pool) *%s {\n\treturn &%s{DB: db}\n}\n\n", repoName, repoName, repoName))

	var colNames []string
	for _, col := range cols {
		colNames = append(colNames, col.Name)
	}
	colsList := strings.Join(colNames, ", ")

	hasDeletedAt := false
	for _, col := range cols {
		if col.Name == "deleted_at" {
			hasDeletedAt = true
			break
		}
	}

	// 1. Get
	sb.WriteString(fmt.Sprintf("// Get retrieves a single %s by %s\n", ent.EntityName, ent.PKColumn))
	sb.WriteString(fmt.Sprintf("func (r *%s) Get(ctx context.Context, id string) (*model.%s, error) {\n", repoName, structName))
	sb.WriteString(fmt.Sprintf("\tquery := `SELECT %s FROM %s WHERE %s = $1", colsList, ent.TableName, ent.PKColumn))
	if hasDeletedAt {
		sb.WriteString(" AND deleted_at IS NULL")
	}
	sb.WriteString("`\n\n")

	sb.WriteString(fmt.Sprintf("\tvar m model.%s\n", structName))
	var scanArgs []string
	for _, col := range cols {
		scanArgs = append(scanArgs, fmt.Sprintf("&m.%s", toPascalCase(col.Name)))
	}
	sb.WriteString(fmt.Sprintf("\terr := r.DB.QueryRow(ctx, query, id).Scan(%s)\n", strings.Join(scanArgs, ", ")))
	sb.WriteString("\tif err != nil {\n\t\treturn nil, err\n\t}\n\n")

	// Load children
	for _, ch := range ent.Children {
		childCols := tableCols[ch.TableName]
		var childColNames []string
		for _, cc := range childCols {
			childColNames = append(childColNames, cc.Name)
		}
		childColsList := strings.Join(childColNames, ", ")
		childStructName := cleanChildStructName(ch.TableName)

		sb.WriteString(fmt.Sprintf("\t// Load %s\n", ch.StructKey))
		sb.WriteString(fmt.Sprintf("\tchildRows%s, err := r.DB.Query(ctx, `SELECT %s FROM %s WHERE %s = $1`, id)\n", ch.StructKey, childColsList, ch.TableName, ch.FKColumn))
		sb.WriteString("\tif err == nil {\n")
		sb.WriteString(fmt.Sprintf("\t\tdefer childRows%s.Close()\n", ch.StructKey))
		sb.WriteString(fmt.Sprintf("\t\tfor childRows%s.Next() {\n", ch.StructKey))
		sb.WriteString(fmt.Sprintf("\t\t\tvar item model.%s\n", childStructName))
		var childScanArgs []string
		for _, cc := range childCols {
			childScanArgs = append(childScanArgs, fmt.Sprintf("&item.%s", toPascalCase(cc.Name)))
		}
		sb.WriteString(fmt.Sprintf("\t\t\tif err := childRows%s.Scan(%s); err == nil {\n", ch.StructKey, strings.Join(childScanArgs, ", ")))
		sb.WriteString(fmt.Sprintf("\t\t\t\tm.%s = append(m.%s, item)\n", ch.StructKey, ch.StructKey))
		sb.WriteString("\t\t\t}\n\t\t}\n\t}\n\n")
	}

	sb.WriteString("\treturn &m, nil\n}\n\n")

	// 2. List
	sb.WriteString(fmt.Sprintf("// List retrieves paginated %s records\n", ent.EntityName))
	sb.WriteString(fmt.Sprintf("func (r *%s) List(ctx context.Context, opts model.ListOptions) ([]model.%s, int, error) {\n", repoName, structName))
	sb.WriteString("\twhereClauses := []string{\"1=1\"}\n")
	sb.WriteString("\tvar args []any\n")
	sb.WriteString("\targPos := 1\n\n")

	if hasDeletedAt {
		sb.WriteString("\twhereClauses = append(whereClauses, \"deleted_at IS NULL\")\n")
	}

	var searchCols []string
	for _, col := range cols {
		if strings.Contains(strings.ToLower(col.DataType), "char") || strings.ToLower(col.DataType) == "text" {
			searchCols = append(searchCols, col.Name)
		}
	}
	if len(searchCols) > 0 {
		if len(searchCols) > 3 {
			searchCols = searchCols[:3]
		}
		sb.WriteString("\tif opts.Search != \"\" {\n")
		sb.WriteString("\t\tsearchPattern := \"%\" + opts.Search + \"%\"\n")
		cond := strings.Join(searchCols, " ILIKE $%[1]d OR ") + " ILIKE $%[1]d"
		sb.WriteString(fmt.Sprintf("\t\twhereClauses = append(whereClauses, fmt.Sprintf(\"(%s)\", argPos))\n", cond))
		sb.WriteString("\t\targs = append(args, searchPattern)\n")
		sb.WriteString("\t\targPos++\n")
		sb.WriteString("\t}\n\n")
	}

	sb.WriteString("\twhereSql := strings.Join(whereClauses, \" AND \")\n")
	sb.WriteString(fmt.Sprintf("\tcountQuery := fmt.Sprintf(\"SELECT COUNT(*) FROM %s WHERE %%s\", whereSql)\n", ent.TableName))
	sb.WriteString("\tvar total int\n")
	sb.WriteString("\tif err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {\n\t\treturn nil, 0, err\n\t}\n\n")

	sb.WriteString("\tlimit := opts.Limit\n\tif limit <= 0 { limit = 10 }\n")
	sb.WriteString("\toffset := opts.Offset\n\tif offset < 0 { offset = 0 }\n\n")

	orderBy := ent.PKColumn
	for _, col := range cols {
		if col.Name == "created_at" {
			orderBy = "created_at DESC"
			break
		}
	}
	sb.WriteString(fmt.Sprintf("\tlistQuery := fmt.Sprintf(\"SELECT %s FROM %s WHERE %%s ORDER BY %s LIMIT $%%d OFFSET $%%d\", whereSql, argPos, argPos+1)\n", colsList, ent.TableName, orderBy))
	sb.WriteString("\targs = append(args, limit, offset)\n\n")

	sb.WriteString("\trows, err := r.DB.Query(ctx, listQuery, args...)\n")
	sb.WriteString("\tif err != nil {\n\t\treturn nil, 0, err\n\t}\n")
	sb.WriteString("\tdefer rows.Close()\n\n")

	sb.WriteString(fmt.Sprintf("\tvar items []model.%s\n", structName))
	sb.WriteString("\tfor rows.Next() {\n")
	sb.WriteString(fmt.Sprintf("\t\tvar m model.%s\n", structName))
	sb.WriteString(fmt.Sprintf("\t\tif err := rows.Scan(%s); err != nil {\n\t\t\treturn nil, 0, err\n\t\t}\n", strings.Join(scanArgs, ", ")))
	sb.WriteString("\t\titems = append(items, m)\n\t}\n\n")
	sb.WriteString("\treturn items, total, nil\n}\n\n")

	// 3. Create (inside TX)
	sb.WriteString(fmt.Sprintf("// Create inserts a new %s with optional child items\n", ent.EntityName))
	sb.WriteString(fmt.Sprintf("func (r *%s) Create(ctx context.Context, m *model.%s) (string, error) {\n", repoName, structName))
	sb.WriteString("\ttx, err := r.DB.Begin(ctx)\n\tif err != nil {\n\t\treturn \"\", err\n\t}\n\tdefer tx.Rollback(ctx)\n\n")

	var insertCols []string
	var insertVals []string
	var insertArgs []string
	idx := 1
	for _, col := range cols {
		if col.IsPK && col.Default != nil {
			continue // generated UUID
		}
		if col.IsGenerated {
			continue // generated always column
		}
		if col.Name == "created_at" || col.Name == "updated_at" || col.Name == "deleted_at" {
			continue
		}
		insertCols = append(insertCols, col.Name)

		if col.UDTName != "" && !strings.HasPrefix(col.UDTName, "int") && !strings.HasPrefix(col.UDTName, "varchar") &&
			!strings.HasPrefix(col.UDTName, "text") && !strings.HasPrefix(col.UDTName, "bool") &&
			!strings.HasPrefix(col.UDTName, "uuid") && !strings.HasPrefix(col.UDTName, "numeric") &&
			!strings.HasPrefix(col.UDTName, "timestamp") && !strings.HasPrefix(col.UDTName, "date") && !strings.HasPrefix(col.UDTName, "json") {
			insertVals = append(insertVals, fmt.Sprintf("$%d::%s", idx, col.UDTName))
		} else {
			insertVals = append(insertVals, fmt.Sprintf("$%d", idx))
		}
		insertArgs = append(insertArgs, fmt.Sprintf("m.%s", toPascalCase(col.Name)))
		idx++
	}

	sb.WriteString(fmt.Sprintf("\tinsertQuery := `INSERT INTO %s (%s) VALUES (%s) RETURNING %s`\n",
		ent.TableName, strings.Join(insertCols, ", "), strings.Join(insertVals, ", "), ent.PKColumn))
	sb.WriteString("\tvar newID string\n")
	sb.WriteString(fmt.Sprintf("\terr = tx.QueryRow(ctx, insertQuery, %s).Scan(&newID)\n", strings.Join(insertArgs, ", ")))
	sb.WriteString("\tif err != nil {\n\t\treturn \"\", err\n\t}\n\n")

	// Insert Children
	for _, ch := range ent.Children {
		childCols := tableCols[ch.TableName]
		var cCols []string
		var cVals []string
		var cArgs []string
		cIdx := 1
		for _, cc := range childCols {
			if cc.IsPK && cc.Default != nil {
				continue
			}
			if cc.IsGenerated {
				continue
			}
			if cc.Name == "created_at" || cc.Name == "updated_at" || cc.Name == "deleted_at" {
				continue
			}
			cCols = append(cCols, cc.Name)
			if cc.Name == ch.FKColumn {
				cVals = append(cVals, fmt.Sprintf("$%d", cIdx))
				cArgs = append(cArgs, "newID")
			} else {
				if cc.UDTName != "" && !strings.HasPrefix(cc.UDTName, "int") && !strings.HasPrefix(cc.UDTName, "varchar") &&
					!strings.HasPrefix(cc.UDTName, "text") && !strings.HasPrefix(cc.UDTName, "bool") &&
					!strings.HasPrefix(cc.UDTName, "uuid") && !strings.HasPrefix(cc.UDTName, "numeric") &&
					!strings.HasPrefix(cc.UDTName, "timestamp") && !strings.HasPrefix(cc.UDTName, "date") && !strings.HasPrefix(cc.UDTName, "json") {
					cVals = append(cVals, fmt.Sprintf("$%d::%s", cIdx, cc.UDTName))
				} else {
					cVals = append(cVals, fmt.Sprintf("$%d", cIdx))
				}
				cArgs = append(cArgs, fmt.Sprintf("item.%s", toPascalCase(cc.Name)))
			}
			cIdx++
		}

		sb.WriteString(fmt.Sprintf("\tfor _, item := range m.%s {\n", ch.StructKey))
		sb.WriteString(fmt.Sprintf("\t\t_, err := tx.Exec(ctx, `INSERT INTO %s (%s) VALUES (%s)`, %s)\n",
			ch.TableName, strings.Join(cCols, ", "), strings.Join(cVals, ", "), strings.Join(cArgs, ", ")))
		sb.WriteString("\t\tif err != nil {\n\t\t\treturn \"\", err\n\t\t}\n\t}\n\n")
	}

	sb.WriteString("\tif err := tx.Commit(ctx); err != nil {\n\t\treturn \"\", err\n\t}\n")
	sb.WriteString("\treturn newID, nil\n}\n\n")

	// 4. Update
	sb.WriteString(fmt.Sprintf("// Update updates an existing %s\n", ent.EntityName))
	sb.WriteString(fmt.Sprintf("func (r *%s) Update(ctx context.Context, id string, m *model.%s) error {\n", repoName, structName))
	sb.WriteString("\ttx, err := r.DB.Begin(ctx)\n\tif err != nil {\n\t\treturn err\n\t}\n\tdefer tx.Rollback(ctx)\n\n")

	var updateSets []string
	var updateArgs []string
	uIdx := 1
	for _, col := range cols {
		if col.IsPK || col.IsGenerated || col.Name == "created_at" || col.Name == "created_by" || col.Name == "deleted_at" || col.Name == "deleted_by" {
			continue
		}
		if col.Name == "updated_at" {
			updateSets = append(updateSets, "updated_at = CURRENT_TIMESTAMP")
			continue
		}
		if col.UDTName != "" && !strings.HasPrefix(col.UDTName, "int") && !strings.HasPrefix(col.UDTName, "varchar") &&
			!strings.HasPrefix(col.UDTName, "text") && !strings.HasPrefix(col.UDTName, "bool") &&
			!strings.HasPrefix(col.UDTName, "uuid") && !strings.HasPrefix(col.UDTName, "numeric") &&
			!strings.HasPrefix(col.UDTName, "timestamp") && !strings.HasPrefix(col.UDTName, "date") && !strings.HasPrefix(col.UDTName, "json") {
			updateSets = append(updateSets, fmt.Sprintf("%s = $%d::%s", col.Name, uIdx, col.UDTName))
		} else {
			updateSets = append(updateSets, fmt.Sprintf("%s = $%d", col.Name, uIdx))
		}
		updateArgs = append(updateArgs, fmt.Sprintf("m.%s", toPascalCase(col.Name)))
		uIdx++
	}
	updateArgs = append(updateArgs, "id")

	sb.WriteString(fmt.Sprintf("\tupdateQuery := `UPDATE %s SET %s WHERE %s = $%d`\n",
		ent.TableName, strings.Join(updateSets, ", "), ent.PKColumn, uIdx))
	sb.WriteString(fmt.Sprintf("\t_, err = tx.Exec(ctx, updateQuery, %s)\n", strings.Join(updateArgs, ", ")))
	sb.WriteString("\tif err != nil {\n\t\treturn err\n\t}\n\n")

	// Update Children
	for _, ch := range ent.Children {
		sb.WriteString(fmt.Sprintf("\tif len(m.%s) > 0 {\n", ch.StructKey))
		sb.WriteString(fmt.Sprintf("\t\t_, _ = tx.Exec(ctx, `DELETE FROM %s WHERE %s = $1`, id)\n", ch.TableName, ch.FKColumn))
		childCols := tableCols[ch.TableName]
		var cCols []string
		var cVals []string
		var cArgs []string
		cIdx := 1
		for _, cc := range childCols {
			if cc.IsPK && cc.Default != nil {
				continue
			}
			if cc.IsGenerated {
				continue
			}
			if cc.Name == "created_at" || cc.Name == "updated_at" || cc.Name == "deleted_at" {
				continue
			}
			cCols = append(cCols, cc.Name)
			if cc.Name == ch.FKColumn {
				cVals = append(cVals, fmt.Sprintf("$%d", cIdx))
				cArgs = append(cArgs, "id")
			} else {
				if cc.UDTName != "" && !strings.HasPrefix(cc.UDTName, "int") && !strings.HasPrefix(cc.UDTName, "varchar") &&
					!strings.HasPrefix(cc.UDTName, "text") && !strings.HasPrefix(cc.UDTName, "bool") &&
					!strings.HasPrefix(cc.UDTName, "uuid") && !strings.HasPrefix(cc.UDTName, "numeric") &&
					!strings.HasPrefix(cc.UDTName, "timestamp") && !strings.HasPrefix(cc.UDTName, "date") && !strings.HasPrefix(cc.UDTName, "json") {
					cVals = append(cVals, fmt.Sprintf("$%d::%s", cIdx, cc.UDTName))
				} else {
					cVals = append(cVals, fmt.Sprintf("$%d", cIdx))
				}
				cArgs = append(cArgs, fmt.Sprintf("item.%s", toPascalCase(cc.Name)))
			}
			cIdx++
		}
		sb.WriteString(fmt.Sprintf("\t\tfor _, item := range m.%s {\n", ch.StructKey))
		sb.WriteString(fmt.Sprintf("\t\t\t_, err := tx.Exec(ctx, `INSERT INTO %s (%s) VALUES (%s)`, %s)\n",
			ch.TableName, strings.Join(cCols, ", "), strings.Join(cVals, ", "), strings.Join(cArgs, ", ")))
		sb.WriteString("\t\t\tif err != nil {\n\t\t\t\treturn err\n\t\t\t}\n\t\t}\n\t}\n\n")
	}

	sb.WriteString("\treturn tx.Commit(ctx)\n}\n\n")

	// 5. Delete
	sb.WriteString(fmt.Sprintf("// Delete removes or soft-deletes %s\n", ent.EntityName))
	sb.WriteString(fmt.Sprintf("func (r *%s) Delete(ctx context.Context, id string) error {\n", repoName))
	if hasDeletedAt {
		sb.WriteString(fmt.Sprintf("\tquery := `UPDATE %s SET deleted_at = CURRENT_TIMESTAMP WHERE %s = $1 AND deleted_at IS NULL`\n", ent.TableName, ent.PKColumn))
		sb.WriteString("\t_, err := r.DB.Exec(ctx, query, id)\n\treturn err\n")
	} else {
		sb.WriteString(fmt.Sprintf("\tquery := `DELETE FROM %s WHERE %s = $1`\n", ent.TableName, ent.PKColumn))
		sb.WriteString("\t_, err := r.DB.Exec(ctx, query, id)\n\treturn err\n")
	}
	sb.WriteString("}\n")

	filePath := filepath.Join(baseDir, "repository", ent.EntityName+"_repository.go")
	os.WriteFile(filePath, []byte(sb.String()), 0644)
}

func generateService(baseDir string, ent EntityConfig, tableCols map[string][]ColumnInfo) {
	cols := tableCols[ent.TableName]
	structName := toPascalCase(ent.EntityName)
	svcName := structName + "Service"
	repoName := structName + "Repository"

	var sb strings.Builder
	sb.WriteString("package service\n\nimport (\n\t\"context\"\n\t\"errors\"\n\n")
	if ent.EntityName == "sys_user" {
		sb.WriteString("\t\"github.com/figoalfarqi/navalerp/internal/auth\"\n")
	}
	sb.WriteString("\t\"github.com/figoalfarqi/navalerp/internal/model\"\n\t\"github.com/figoalfarqi/navalerp/internal/repository\"\n)\n\n")

	sb.WriteString(fmt.Sprintf("type %s struct {\n\tRepo *repository.%s\n}\n\n", svcName, repoName))
	sb.WriteString(fmt.Sprintf("func New%s(repo *repository.%s) *%s {\n\treturn &%s{Repo: repo}\n}\n\n", svcName, repoName, svcName, svcName))

	sb.WriteString(fmt.Sprintf("func (s *%s) GetByID(ctx context.Context, id string) (*model.%s, error) {\n\tif id == \"\" {\n\t\treturn nil, errors.New(\"id is required\")\n\t}\n\treturn s.Repo.Get(ctx, id)\n}\n\n", svcName, structName))
	sb.WriteString(fmt.Sprintf("func (s *%s) List(ctx context.Context, opts model.ListOptions) ([]model.%s, int, error) {\n\treturn s.Repo.List(ctx, opts)\n}\n\n", svcName, structName))

	// Create
	sb.WriteString(fmt.Sprintf("func (s *%s) Create(ctx context.Context, loginID string, req *model.%sRequest) (*model.%s, error, map[string]string) {\n", svcName, structName, structName))
	sb.WriteString(fmt.Sprintf("\tm := &model.%s{}\n", structName))
	for _, col := range cols {
		if col.IsPK || col.IsGenerated || col.Name == "created_at" || col.Name == "updated_at" || col.Name == "deleted_at" ||
			col.Name == "created_by" || col.Name == "updated_by" || col.Name == "deleted_by" {
			continue
		}
		fName := toPascalCase(col.Name)
		goType, _ := postgresToGoType(col)
		if goType == "any" {
			sb.WriteString(fmt.Sprintf("\tif req.%s != nil { m.%s = req.%s }\n", fName, fName, fName))
		} else if strings.HasPrefix(goType, "*") {
			sb.WriteString(fmt.Sprintf("\tm.%s = req.%s\n", fName, fName))
		} else {
			sb.WriteString(fmt.Sprintf("\tif req.%s != nil { m.%s = *req.%s }\n", fName, fName, fName))
		}
	}
	if ent.EntityName == "sys_user" {
		sb.WriteString("\tif req.PasswordHash != nil && *req.PasswordHash != \"\" {\n")
		sb.WriteString("\t\thashed, err := auth.HashPassword(*req.PasswordHash)\n")
		sb.WriteString("\t\tif err == nil { m.PasswordHash = hashed }\n")
		sb.WriteString("\t}\n")
	}
	for _, ch := range ent.Children {
		sb.WriteString(fmt.Sprintf("\tm.%s = req.%s\n", ch.StructKey, ch.StructKey))
	}
	sb.WriteString("\tid, err := s.Repo.Create(ctx, m)\n\tif err != nil {\n\t\treturn nil, err, nil\n\t}\n\titem, err := s.Repo.Get(ctx, id)\n\treturn item, err, nil\n}\n\n")

	// Update
	sb.WriteString(fmt.Sprintf("func (s *%s) Update(ctx context.Context, loginID string, id string, req *model.%sRequest) (*model.%s, error, map[string]string) {\n", svcName, structName, structName))
	sb.WriteString("\tm, err := s.Repo.Get(ctx, id)\n\tif err != nil {\n\t\treturn nil, err, nil\n\t}\n")
	for _, col := range cols {
		if col.IsPK || col.IsGenerated || col.Name == "created_at" || col.Name == "updated_at" || col.Name == "deleted_at" ||
			col.Name == "created_by" || col.Name == "updated_by" || col.Name == "deleted_by" {
			continue
		}
		fName := toPascalCase(col.Name)
		goType, _ := postgresToGoType(col)
		if goType == "any" {
			sb.WriteString(fmt.Sprintf("\tif req.%s != nil { m.%s = req.%s }\n", fName, fName, fName))
		} else if strings.HasPrefix(goType, "*") {
			sb.WriteString(fmt.Sprintf("\tif req.%s != nil { m.%s = req.%s }\n", fName, fName, fName))
		} else {
			sb.WriteString(fmt.Sprintf("\tif req.%s != nil { m.%s = *req.%s }\n", fName, fName, fName))
		}
	}
	if ent.EntityName == "sys_user" {
		sb.WriteString("\tif req.PasswordHash != nil && *req.PasswordHash != \"\" {\n")
		sb.WriteString("\t\thashed, err := auth.HashPassword(*req.PasswordHash)\n")
		sb.WriteString("\t\tif err == nil { m.PasswordHash = hashed }\n")
		sb.WriteString("\t}\n")
	}
	for _, ch := range ent.Children {
		sb.WriteString(fmt.Sprintf("\tif req.%[1]s != nil { m.%[1]s = req.%[1]s }\n", ch.StructKey))
	}
	sb.WriteString("\tif err := s.Repo.Update(ctx, id, m); err != nil {\n\t\treturn nil, err, nil\n\t}\n\titem, err := s.Repo.Get(ctx, id)\n\treturn item, err, nil\n}\n\n")

	// Delete
	sb.WriteString(fmt.Sprintf("func (s *%s) Delete(ctx context.Context, loginID string, id string) error {\n", svcName))
	sb.WriteString("\treturn s.Repo.Delete(ctx, id)\n}\n")

	filePath := filepath.Join(baseDir, "service", ent.EntityName+"_service.go")
	os.WriteFile(filePath, []byte(sb.String()), 0644)
}

func generateHandler(baseDir string, ent EntityConfig, tableCols map[string][]ColumnInfo) {
	structName := toPascalCase(ent.EntityName)
	handlerName := structName + "Handler"
	svcName := structName + "Service"

	var sb strings.Builder
	sb.WriteString("package handler\n\nimport (\n\t\"encoding/json\"\n\t\"net/http\"\n\t\"strconv\"\n\t\"strings\"\n\n\t\"github.com/figoalfarqi/navalerp/config\"\n\t\"github.com/figoalfarqi/navalerp/internal/app/middleware\"\n\t\"github.com/figoalfarqi/navalerp/internal/model\"\n\t\"github.com/figoalfarqi/navalerp/internal/service\"\n\t\"github.com/figoalfarqi/navalerp/pkg/response\"\n)\n\n")

	sb.WriteString(fmt.Sprintf("type %s struct {\n\tSvc *service.%s\n\tCfg *config.Config\n}\n\n", handlerName, svcName))
	sb.WriteString(fmt.Sprintf("func New%s(s *service.%s, c *config.Config) *%s {\n\treturn &%s{Svc: s, Cfg: c}\n}\n\n", handlerName, svcName, handlerName, handlerName))

	// Create
	sb.WriteString(fmt.Sprintf("func (h *%s) Create(w http.ResponseWriter, r *http.Request) {\n", handlerName))
	sb.WriteString("\tloginID, _ := r.Context().Value(middleware.CtxUserID).(string)\n")
	sb.WriteString(fmt.Sprintf("\tvar req model.%sRequest\n", structName))
	sb.WriteString("\tif err := json.NewDecoder(r.Body).Decode(&req); err != nil {\n")
	sb.WriteString("\t\tresponse.JSON(w, http.StatusBadRequest, \"invalid json format\", nil, map[string]string{\"error\": err.Error()})\n\t\treturn\n\t}\n")
	sb.WriteString("\titem, err, valErr := h.Svc.Create(r.Context(), loginID, &req)\n")
	sb.WriteString("\tif valErr != nil {\n\t\tresponse.JSON(w, http.StatusBadRequest, \"validation error\", nil, valErr)\n\t\treturn\n\t}\n")
	sb.WriteString("\tif err != nil {\n\t\tresponse.JSON(w, http.StatusInternalServerError, err.Error(), nil, nil)\n\t\treturn\n\t}\n")
	sb.WriteString(fmt.Sprintf("\tresponse.JSON(w, http.StatusCreated, \"%s created\", item, nil)\n}\n\n", ent.EntityName))

	// Get
	sb.WriteString(fmt.Sprintf("func (h *%s) Get(w http.ResponseWriter, r *http.Request) {\n", handlerName))
	sb.WriteString("\tid := r.PathValue(\"id\")\n")
	sb.WriteString("\tif id != \"\" {\n")
	sb.WriteString("\t\titem, err := h.Svc.GetByID(r.Context(), id)\n")
	sb.WriteString("\t\tif err != nil {\n\t\t\tresponse.JSON(w, http.StatusNotFound, \"not found\", nil, nil)\n\t\t\treturn\n\t\t}\n")
	sb.WriteString("\t\tresponse.JSON(w, http.StatusOK, \"success\", item, nil)\n\t\treturn\n\t}\n\n")

	sb.WriteString("\tquery := r.URL.Query()\n")
	sb.WriteString("\tpage, _ := strconv.Atoi(query.Get(\"page\"))\n\tif page <= 0 { page = 1 }\n")
	sb.WriteString("\tlimit, _ := strconv.Atoi(query.Get(\"limit\"))\n\tif limit <= 0 { limit = 10 }\n")
	sb.WriteString("\toffset := (page - 1) * limit\n")
	sb.WriteString("\tsearch := strings.TrimSpace(query.Get(\"search\"))\n\n")

	sb.WriteString("\topts := model.ListOptions{Limit: limit, Offset: offset, Search: search}\n")
	sb.WriteString("\titems, total, err := h.Svc.List(r.Context(), opts)\n")
	sb.WriteString("\tif err != nil {\n\t\tresponse.JSON(w, http.StatusInternalServerError, err.Error(), nil, nil)\n\t\treturn\n\t}\n")
	sb.WriteString("\tresponse.JSON(w, http.StatusOK, \"success\", map[string]any{\n")
	sb.WriteString("\t\t\"items\": items,\n\t\t\"total\": total,\n\t\t\"page\": page,\n\t\t\"limit\": limit,\n\t}, nil)\n}\n\n")

	// Update
	sb.WriteString(fmt.Sprintf("func (h *%s) Update(w http.ResponseWriter, r *http.Request) {\n", handlerName))
	sb.WriteString("\tid := r.PathValue(\"id\")\n")
	sb.WriteString("\tif id == \"\" {\n\t\tresponse.JSON(w, http.StatusBadRequest, \"id is required\", nil, nil)\n\t\treturn\n\t}\n")
	sb.WriteString("\tloginID, _ := r.Context().Value(middleware.CtxUserID).(string)\n")
	sb.WriteString(fmt.Sprintf("\tvar req model.%sRequest\n", structName))
	sb.WriteString("\tif err := json.NewDecoder(r.Body).Decode(&req); err != nil {\n")
	sb.WriteString("\t\tresponse.JSON(w, http.StatusBadRequest, \"invalid json format\", nil, map[string]string{\"error\": err.Error()})\n\t\treturn\n\t}\n")
	sb.WriteString("\titem, err, valErr := h.Svc.Update(r.Context(), loginID, id, &req)\n")
	sb.WriteString("\tif valErr != nil {\n\t\tresponse.JSON(w, http.StatusBadRequest, \"validation error\", nil, valErr)\n\t\treturn\n\t}\n")
	sb.WriteString("\tif err != nil {\n\t\tresponse.JSON(w, http.StatusInternalServerError, err.Error(), nil, nil)\n\t\treturn\n\t}\n")
	sb.WriteString(fmt.Sprintf("\tresponse.JSON(w, http.StatusOK, \"%s updated\", item, nil)\n}\n\n", ent.EntityName))

	// Delete
	sb.WriteString(fmt.Sprintf("func (h *%s) Delete(w http.ResponseWriter, r *http.Request) {\n", handlerName))
	sb.WriteString("\tid := r.PathValue(\"id\")\n")
	sb.WriteString("\tif id == \"\" {\n\t\tresponse.JSON(w, http.StatusBadRequest, \"id is required\", nil, nil)\n\t\treturn\n\t}\n")
	sb.WriteString("\tloginID, _ := r.Context().Value(middleware.CtxUserID).(string)\n")
	sb.WriteString("\tif err := h.Svc.Delete(r.Context(), loginID, id); err != nil {\n")
	sb.WriteString("\t\tresponse.JSON(w, http.StatusInternalServerError, err.Error(), nil, nil)\n\t\treturn\n\t}\n")
	sb.WriteString(fmt.Sprintf("\tresponse.JSON(w, http.StatusOK, \"%s deleted\", nil, nil)\n}\n", ent.EntityName))

	filePath := filepath.Join(baseDir, "handler", ent.EntityName+"_handler.go")
	os.WriteFile(filePath, []byte(sb.String()), 0644)
}

func generateRouter(baseDir string, ents []EntityConfig) {
	// router_resources.go
	var rsb strings.Builder
	rsb.WriteString("package app\n\nimport (\n\t\"github.com/figoalfarqi/navalerp/config\"\n\t\"github.com/figoalfarqi/navalerp/internal/handler\"\n\t\"github.com/figoalfarqi/navalerp/internal/repository\"\n\t\"github.com/figoalfarqi/navalerp/internal/service\"\n\t\"github.com/figoalfarqi/navalerp/pkg/database\"\n)\n\n")

	rsb.WriteString("type routeHandlers struct {\n")
	rsb.WriteString("\tfileUpload *handler.FileUploadHandler\n")
	rsb.WriteString("\tdashboard  *handler.DashboardHandler\n")
	rsb.WriteString("\treport     *handler.ReportHandler\n")
	rsb.WriteString("\tnavalAuth  *handler.NavalAuthHandler\n")
	for _, ent := range ents {
		rsb.WriteString(fmt.Sprintf("\t%s *handler.%sHandler\n", toCamelCase(ent.EntityName), toPascalCase(ent.EntityName)))
	}
	rsb.WriteString("}\n\n")

	rsb.WriteString("func buildRouteHandlers(cfg *config.Config) *routeHandlers {\n")
	rsb.WriteString("\tdb := database.GetPool()\n")
	rsb.WriteString("\tfileService := service.NewFileUploadService(cfg)\n")
	rsb.WriteString("\tdashboardRepo := repository.NewDashboardRepository(db)\n")
	rsb.WriteString("\tdashboardService := service.NewDashboardService(dashboardRepo)\n")
	rsb.WriteString("\treportRepo := repository.NewReportRepository(db)\n")
	rsb.WriteString("\treportService := service.NewReportService(reportRepo)\n\n")

	for _, ent := range ents {
		pName := toPascalCase(ent.EntityName)
		cName := toCamelCase(ent.EntityName)
		rsb.WriteString(fmt.Sprintf("\t%sRepo := repository.New%sRepository(db)\n", cName, pName))
		rsb.WriteString(fmt.Sprintf("\t%sService := service.New%sService(%sRepo)\n", cName, pName, cName))
	}
	rsb.WriteString("\n\treturn &routeHandlers{\n")
	rsb.WriteString("\t\tfileUpload: handler.NewFileUploadHandler(fileService, cfg),\n")
	rsb.WriteString("\t\tdashboard:  handler.NewDashboardHandler(dashboardService, cfg),\n")
	rsb.WriteString("\t\treport:     handler.NewReportHandler(reportService, cfg),\n")
	rsb.WriteString("\t\tnavalAuth:  handler.NewNavalAuthHandler(db, cfg),\n")
	for _, ent := range ents {
		pName := toPascalCase(ent.EntityName)
		cName := toCamelCase(ent.EntityName)
		rsb.WriteString(fmt.Sprintf("\t\t%s: handler.New%sHandler(%sService, cfg),\n", cName, pName, cName))
	}
	rsb.WriteString("\t}\n}\n")

	os.WriteFile(filepath.Join(baseDir, "app", "router_resources.go"), []byte(rsb.String()), 0644)

	// router_registration.go
	var regsb strings.Builder
	regsb.WriteString("package app\n\nimport (\n\t\"net/http\"\n\t\"slices\"\n\t\"strings\"\n\n\t\"github.com/figoalfarqi/navalerp/config\"\n\t\"github.com/figoalfarqi/navalerp/internal/app/middleware\"\n\t\"github.com/figoalfarqi/navalerp/internal/helper\"\n)\n\n")

	regsb.WriteString("type crudHandler interface {\n\tCreate(http.ResponseWriter, *http.Request)\n\tGet(http.ResponseWriter, *http.Request)\n\tUpdate(http.ResponseWriter, *http.Request)\n\tDelete(http.ResponseWriter, *http.Request)\n}\n\n")
	regsb.WriteString("func authenticated(cfg *config.Config, roles []int, fn http.HandlerFunc) http.Handler {\n\treturn middleware.AuthJWT(cfg, roles, fn)\n}\n\n")

	regsb.WriteString("func registerRead(mux *http.ServeMux, cfg *config.Config, path string, roles []int, fn http.HandlerFunc) {\n")
	regsb.WriteString("\tif strings.HasPrefix(path, \"/api/v1/admin/\") {\n\t\tfor _, roleID := range []int{helper.AppRoleOwner, helper.AppRoleITDev} {\n\t\t\tif !slices.Contains(roles, roleID) {\n\t\t\t\troles = append(roles, roleID)\n\t\t\t}\n\t\t}\n\t}\n")
	regsb.WriteString("\tmux.Handle(\"GET \"+path, authenticated(cfg, roles, fn))\n")
	regsb.WriteString("\tmux.Handle(\"GET \"+path+\"/{id}\", authenticated(cfg, roles, fn))\n}\n\n")

	regsb.WriteString("func registerCRUD(mux *http.ServeMux, cfg *config.Config, path string, roles []int, h crudHandler) {\n")
	regsb.WriteString("\tregisterRead(mux, cfg, path, roles, h.Get)\n")
	regsb.WriteString("\tmux.Handle(\"POST \"+path, authenticated(cfg, roles, h.Create))\n")
	regsb.WriteString("\tmux.Handle(\"PUT \"+path+\"/{id}\", authenticated(cfg, roles, h.Update))\n")
	regsb.WriteString("\tmux.Handle(\"PATCH \"+path+\"/{id}\", authenticated(cfg, roles, h.Update))\n")
	regsb.WriteString("\tmux.Handle(\"DELETE \"+path+\"/{id}\", authenticated(cfg, roles, h.Delete))\n}\n\n")

	regsb.WriteString("func registerApplicationRoutes(mux *http.ServeMux, cfg *config.Config, h *routeHandlers) {\n")
	regsb.WriteString("\tdashboardRoles := helper.GetAppRoleIDsByRoleTypeName(\"admin_dashboard\")\n")
	regsb.WriteString("\treportRoles := helper.GetAppRoleIDsByRoleTypeName(\"admin_report\")\n")
	regsb.WriteString("\tallRoles := helper.GetAppRoleIDsByRoleTypeName(\"allrole\")\n\n")

	regsb.WriteString("\tmux.HandleFunc(\"POST /api/v1/admin/login\", h.navalAuth.Login)\n")
	regsb.WriteString("\tmux.HandleFunc(\"POST /api/v1/auth/login\", h.navalAuth.Login)\n")
	regsb.WriteString("\tmux.HandleFunc(\"POST /api/v1/auth/logout\", h.navalAuth.Logout)\n")
	regsb.WriteString("\tmux.Handle(\"GET /api/v1/auth/me\", authenticated(cfg, allRoles, h.navalAuth.Me))\n\n")

	regsb.WriteString("\tmux.Handle(\"GET /api/v1/admin/dashboard\", authenticated(cfg, dashboardRoles, h.dashboard.Get))\n")
	regsb.WriteString("\tmux.Handle(\"GET /api/v1/admin/report\", authenticated(cfg, reportRoles, h.report.Get))\n\n")

	regsb.WriteString("\t// 48 Modular NavalERP CRUD Routes\n")
	for _, ent := range ents {
		cName := toCamelCase(ent.EntityName)
		regsb.WriteString(fmt.Sprintf("\tregisterCRUD(mux, cfg, \"/api/v1/admin/%s\", allRoles, h.%s)\n", ent.EntityName, cName))
	}
	regsb.WriteString("}\n")

	os.WriteFile(filepath.Join(baseDir, "app", "router_registration.go"), []byte(regsb.String()), 0644)
}

func titleCase(s string) string {
	parts := strings.Split(s, " ")
	var res []string
	for _, p := range parts {
		if len(p) > 0 {
			res = append(res, strings.ToUpper(p[:1])+strings.ToLower(p[1:]))
		}
	}
	return strings.Join(res, " ")
}

func generateFrontend(baseDir string, ent EntityConfig, tableCols map[string][]ColumnInfo) {
	cols := tableCols[ent.TableName]
	dir := filepath.Join(baseDir, ent.EntityName)
	os.MkdirAll(dir, 0755)
	os.MkdirAll(filepath.Join(dir, "create"), 0755)
	os.MkdirAll(filepath.Join(dir, "[id]", "[mode]"), 0755)

	// 1. _config.tsx
	var sb strings.Builder
	sb.WriteString("/* eslint-disable @typescript-eslint/no-explicit-any */\n")
	sb.WriteString("import { ColumnField } from \"@/components/table/Table\";\n")
	sb.WriteString("import { FormField } from \"@/components/formCrud/FormCrud\";\n")
	sb.WriteString("import { FilterField } from \"@/components/table/FilterFormTable\";\n\n")

	sb.WriteString(fmt.Sprintf("export const entityName = \"%s\";\n", ent.EntityName))
	sb.WriteString(fmt.Sprintf("export const entityTitle = \"%s\";\n", ent.Title))
	sb.WriteString(fmt.Sprintf("export const entityEndpoint = \"/admin/%s\";\n", ent.EntityName))
	sb.WriteString(fmt.Sprintf("export const primaryKey = \"%s\";\n\n", ent.PKColumn))

	// Columns
	sb.WriteString("export const columns: ColumnField[] = [\n")
	colCount := 0
	for _, col := range cols {
		if col.Name == "deleted_at" || col.Name == "deleted_by" || col.Name == "password_hash" {
			continue
		}
		if colCount >= 7 {
			break
		}
		label := titleCase(strings.ReplaceAll(col.Name, "_", " "))
		sb.WriteString(fmt.Sprintf("  { key: \"%s\", label: \"%s\" },\n", col.Name, label))
		colCount++
	}
	sb.WriteString("];\n\n")

	// Filter Fields
	sb.WriteString("export const filterFields: FilterField[] = [\n")
	sb.WriteString(fmt.Sprintf("  { name: \"search\", label: \"Cari %s\", fieldType: \"text\", col: \"left\", placeHolder: \"Ketik kata kunci pencarian...\" },\n", ent.Title))
	sb.WriteString("];\n\n")

	// Form Fields
	sb.WriteString("export const formFields = (mode: string): FormField[] => [\n")
	for _, col := range cols {
		if col.IsPK || col.IsGenerated || col.Name == "created_at" || col.Name == "updated_at" || col.Name == "deleted_at" ||
			col.Name == "created_by" || col.Name == "updated_by" || col.Name == "deleted_by" {
			continue
		}
		label := titleCase(strings.ReplaceAll(col.Name, "_", " "))
		fieldType := "text"
		if col.Name == "password_hash" {
			label = "Password"
			fieldType = "password"
		} else {
			dt := strings.ToLower(col.DataType)
			switch {
			case strings.Contains(dt, "int") || strings.Contains(dt, "numeric") || strings.Contains(dt, "decimal") || strings.Contains(dt, "double"):
				fieldType = "number"
			case strings.Contains(dt, "bool"):
				fieldType = "boolean"
			case strings.Contains(dt, "date") || strings.Contains(dt, "time"):
				fieldType = "date"
			case dt == "text":
				fieldType = "textarea"
			}
		}
		req := !col.IsNullable && col.Default == nil
		sb.WriteString(fmt.Sprintf("  {\n    name: \"%s\",\n    label: \"%s\",\n    fieldType: \"%s\",\n    required: %t,\n    disabled: mode === \"view\",\n  },\n",
			col.Name, label, fieldType, req))
	}
	sb.WriteString("];\n\n")

	// Payload builder
	sb.WriteString("export const buildPayload = (data: any) => {\n")
	sb.WriteString("  const payload: any = { ...data };\n")
	sb.WriteString(fmt.Sprintf("  delete payload.%s;\n", ent.PKColumn))
	sb.WriteString("  delete payload.created_at;\n")
	sb.WriteString("  delete payload.updated_at;\n")
	sb.WriteString("  delete payload.deleted_at;\n")
	sb.WriteString("  delete payload.created_by;\n")
	sb.WriteString("  delete payload.updated_by;\n")
	sb.WriteString("  delete payload.deleted_by;\n")
	sb.WriteString("  return payload;\n")
	sb.WriteString("};\n")

	os.WriteFile(filepath.Join(dir, "_config.tsx"), []byte(sb.String()), 0644)

	// 2. page.tsx
	var psb strings.Builder
	psb.WriteString("\"use client\";\n\n")
	psb.WriteString("import Table from \"@/components/table/Table\";\n")
	psb.WriteString("import { columns, entityEndpoint, entityName, entityTitle, filterFields, primaryKey } from \"./_config\";\n\n")
	psb.WriteString(fmt.Sprintf("export default function %sListPage() {\n", toPascalCase(ent.EntityName)))
	psb.WriteString("  return (\n")
	psb.WriteString("    <Table\n")
	psb.WriteString("      title={entityTitle}\n")
	psb.WriteString("      url={entityEndpoint}\n")
	psb.WriteString("      table_name={entityName}\n")
	psb.WriteString("      table_url={`/admin/data/${entityName}`}\n")
	psb.WriteString("      columns={columns}\n")
	psb.WriteString("      primaryKey={primaryKey}\n")
	psb.WriteString("      filters={filterFields}\n")
	psb.WriteString("    />\n")
	psb.WriteString("  );\n")
	psb.WriteString("}\n")
	os.WriteFile(filepath.Join(dir, "page.tsx"), []byte(psb.String()), 0644)

	// 3. create/page.tsx
	var csb strings.Builder
	csb.WriteString("\"use client\";\n\n")
	csb.WriteString("import { useRouter } from \"next/navigation\";\n")
	csb.WriteString("import FormCrud from \"@/components/formCrud/FormCrud\";\n")
	csb.WriteString("import { buildPayload, entityEndpoint, entityTitle, formFields } from \"../_config\";\n\n")
	csb.WriteString(fmt.Sprintf("export default function %sCreatePage() {\n", toPascalCase(ent.EntityName)))
	csb.WriteString("  const router = useRouter();\n")
	csb.WriteString("  return (\n")
	csb.WriteString("    <FormCrud\n")
	csb.WriteString("      title={`Tambah ${entityTitle}`}\n")
	csb.WriteString("      url={entityEndpoint}\n")
	csb.WriteString("      mode=\"add\"\n")
	csb.WriteString("      fields={formFields(\"add\")}\n")
	csb.WriteString("      buildPayload={buildPayload}\n")
	csb.WriteString("      onSuccess={() => router.push(`/admin/data/${entityEndpoint.replace('/admin/', '')}`)}\n")
	csb.WriteString("    />\n")
	csb.WriteString("  );\n")
	csb.WriteString("}\n")
	os.WriteFile(filepath.Join(dir, "create", "page.tsx"), []byte(csb.String()), 0644)

	// 4. [id]/[mode]/page.tsx
	var dsb strings.Builder
	dsb.WriteString("\"use client\";\n\n")
	dsb.WriteString("import { useParams, useRouter } from \"next/navigation\";\n")
	dsb.WriteString("import AdminRecordState from \"@/components/admin/crud/AdminRecordState\";\n")
	dsb.WriteString("import { crudTitle, type AdminCrudMode } from \"@/components/admin/crud/adminCrud\";\n")
	dsb.WriteString("import { useAdminRecord } from \"@/components/admin/crud/useAdminRecord\";\n")
	dsb.WriteString("import FormCrud from \"@/components/formCrud/FormCrud\";\n")
	dsb.WriteString("import { buildPayload, entityEndpoint, entityTitle, formFields, primaryKey } from \"../../_config\";\n\n")
	dsb.WriteString(fmt.Sprintf("export default function %sDetailPage() {\n", toPascalCase(ent.EntityName)))
	dsb.WriteString("  const { id, mode } = useParams<{ id: string; mode: Exclude<AdminCrudMode, \"add\"> }>();\n")
	dsb.WriteString("  const router = useRouter();\n")
	dsb.WriteString("  const { record, error } = useAdminRecord({\n")
	dsb.WriteString("    endpoint: entityEndpoint,\n")
	dsb.WriteString("    id,\n")
	dsb.WriteString("    mode,\n")
	dsb.WriteString("    primaryKey,\n")
	dsb.WriteString("  });\n\n")
	dsb.WriteString("  if (!record) return <AdminRecordState error={error} />;\n\n")
	dsb.WriteString("  return (\n")
	dsb.WriteString("    <FormCrud\n")
	dsb.WriteString("      title={crudTitle(mode, entityTitle)}\n")
	dsb.WriteString("      url={mode === \"edit\" ? `${entityEndpoint}/${id}` : entityEndpoint}\n")
	dsb.WriteString("      mode={mode}\n")
	dsb.WriteString("      initialData={record}\n")
	dsb.WriteString("      fields={formFields(mode)}\n")
	dsb.WriteString("      hideSubmit={mode === \"view\"}\n")
	dsb.WriteString("      buildPayload={buildPayload}\n")
	dsb.WriteString("      onSuccess={() => router.push(`/admin/data/${entityEndpoint.replace('/admin/', '')}`)}\n")
	dsb.WriteString("    />\n")
	dsb.WriteString("  );\n")
	dsb.WriteString("}\n")
	os.WriteFile(filepath.Join(dir, "[id]", "[mode]", "page.tsx"), []byte(dsb.String()), 0644)
}

func generateNavigation(adminDir string, ents []EntityConfig) {
	type NavGroup struct {
		Number string
		Name   string
		Items  []EntityConfig
	}

	groupsMap := make(map[string]*NavGroup)
	var groupOrder []string

	for _, ent := range ents {
		if _, ok := groupsMap[ent.ModuleNum]; !ok {
			groupsMap[ent.ModuleNum] = &NavGroup{
				Number: ent.ModuleNum,
				Name:   ent.ModuleName,
			}
			groupOrder = append(groupOrder, ent.ModuleNum)
		}
		groupsMap[ent.ModuleNum].Items = append(groupsMap[ent.ModuleNum].Items, ent)
	}
	sort.Strings(groupOrder)

	var sb strings.Builder
	sb.WriteString("export interface AdminNavItem {\n  name: string;\n  url: string;\n}\n\n")
	sb.WriteString("export interface AdminNavGroup {\n  title: string;\n  items: AdminNavItem[];\n}\n\n")
	sb.WriteString("export const adminNavigationGroups: AdminNavGroup[] = [\n")
	sb.WriteString("  {\n    title: \"Overview\",\n    items: [\n      { name: \"Dashboard\", url: \"/admin\" },\n      { name: \"Laporan\", url: \"/admin/report\" },\n    ],\n  },\n")

	for _, gNum := range groupOrder {
		g := groupsMap[gNum]
		sb.WriteString(fmt.Sprintf("  {\n    title: \"%s. %s\",\n    items: [\n", g.Number, g.Name))
		for _, item := range g.Items {
			sb.WriteString(fmt.Sprintf("      { name: \"%s\", url: \"/admin/data/%s\" },\n", item.Title, item.EntityName))
		}
		sb.WriteString("    ],\n  },\n")
	}
	sb.WriteString("];\n")

	os.WriteFile(filepath.Join(adminDir, "adminNavigation.tsx"), []byte(sb.String()), 0644)
}
