package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
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
	EntityName  string
	TableName   string
	PKColumn    string
	ModuleNum   string
	ModuleName  string
	Title       string
	PluralTitle string
	Children    []ChildTableConfig
}

var cuiEntities = []EntityConfig{
	{
		EntityName: "cui_asset", TableName: "cui_assets", PKColumn: "cui_asset_id",
		ModuleNum: "11", ModuleName: "Infrastruktur Bawah Laut (CUI)", Title: "Aset Bawah Laut", PluralTitle: "Aset Bawah Laut Kritis (CUI)",
		Children: []ChildTableConfig{
			{TableName: "cui_monitoring_logs", FKColumn: "cui_asset_id", StructKey: "MonitoringLogs"},
			{TableName: "cui_alerts", FKColumn: "cui_asset_id", StructKey: "Alerts"},
			{TableName: "cui_inspections", FKColumn: "cui_asset_id", StructKey: "Inspections"},
		},
	},
	{
		EntityName: "cui_monitoring_log", TableName: "cui_monitoring_logs", PKColumn: "log_id",
		ModuleNum: "11", ModuleName: "Infrastruktur Bawah Laut (CUI)", Title: "Log Sensor CUI", PluralTitle: "Log Sensor Pemantauan Bawah Laut",
	},
	{
		EntityName: "cui_alert", TableName: "cui_alerts", PKColumn: "alert_id",
		ModuleNum: "11", ModuleName: "Infrastruktur Bawah Laut (CUI)", Title: "Peringatan CUI", PluralTitle: "Peringatan Anomali Bawah Laut",
	},
	{
		EntityName: "cui_inspection", TableName: "cui_inspections", PKColumn: "inspection_id",
		ModuleNum: "11", ModuleName: "Infrastruktur Bawah Laut (CUI)", Title: "Inspeksi CUI", PluralTitle: "Riwayat Inspeksi Bawah Laut",
	},
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
	for _, prefix := range []string{"cui_", "mro_", "inv_", "proc_", "fin_", "hcm_", "infra_", "log_", "doc_", "ops_"} {
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
	case strings.Contains(dt, "timestamp") || dt == "date":
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
		WHERE c.table_schema = 'public' AND c.table_name LIKE 'cui_%'
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

	fmt.Printf("Loaded metadata for %d CUI tables.\n", len(tableCols))

	repoRoot := "d:/stm/VirutalGate/navalerp"
	baseApiDir := filepath.Join(repoRoot, "apps", "api", "internal")
	baseWebDir := filepath.Join(repoRoot, "apps", "web", "src", "app", "admin", "data")

	for _, ent := range cuiEntities {
		generateModel(baseApiDir, ent, tableCols)
		generateRepository(baseApiDir, ent, tableCols)
		generateService(baseApiDir, ent, tableCols)
		generateHandler(baseApiDir, ent, tableCols)
		generateFrontend(baseWebDir, ent, tableCols)
		fmt.Printf("Generated code and CRUD for: %s\n", ent.EntityName)
	}

	fmt.Println("=== CUI GENERATION COMPLETE ===")
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
	selectFields := strings.Join(colNames, ", ")

	hasDeletedAt := false
	for _, col := range cols {
		if col.Name == "deleted_at" {
			hasDeletedAt = true
			break
		}
	}

	// 1. Get
	sb.WriteString(fmt.Sprintf("// Get retrieves a single %s by ID\n", ent.EntityName))
	sb.WriteString(fmt.Sprintf("func (r *%s) Get(ctx context.Context, id string) (*model.%s, error) {\n", repoName, structName))
	sb.WriteString(fmt.Sprintf("\tquery := `SELECT %s FROM %s WHERE %s = $1", selectFields, ent.TableName, ent.PKColumn))
	if hasDeletedAt {
		sb.WriteString(" AND deleted_at IS NULL")
	}
	sb.WriteString("`\n\n")

	sb.WriteString(fmt.Sprintf("\trow := r.DB.QueryRow(ctx, query, id)\n"))
	sb.WriteString(fmt.Sprintf("\tvar m model.%s\n", structName))
	var scanTargets []string
	for _, col := range cols {
		scanTargets = append(scanTargets, fmt.Sprintf("&m.%s", toPascalCase(col.Name)))
	}
	sb.WriteString(fmt.Sprintf("\terr := row.Scan(%s)\n", strings.Join(scanTargets, ", ")))
	sb.WriteString("\tif err != nil {\n\t\treturn nil, err\n\t}\n\n")

	// Query child tables if any
	for _, ch := range ent.Children {
		childCols := tableCols[ch.TableName]
		var cColNames []string
		for _, cc := range childCols {
			cColNames = append(cColNames, cc.Name)
		}
		cSelect := strings.Join(cColNames, ", ")
		sb.WriteString(fmt.Sprintf("\t// Fetch children: %s\n", ch.StructKey))
		sb.WriteString(fmt.Sprintf("\tcRows_%s, err := r.DB.Query(ctx, `SELECT %s FROM %s WHERE %s = $1", ch.StructKey, cSelect, ch.TableName, ch.FKColumn))
		cHasDeletedAt := false
		for _, cc := range childCols {
			if cc.Name == "deleted_at" {
				cHasDeletedAt = true
				break
			}
		}
		if cHasDeletedAt {
			sb.WriteString(" AND deleted_at IS NULL")
		}
		sb.WriteString("`, id)\n")
		sb.WriteString(fmt.Sprintf("\tif err == nil {\n\t\tdefer cRows_%s.Close()\n", ch.StructKey))
		childStructName := cleanChildStructName(ch.TableName)
		sb.WriteString(fmt.Sprintf("\t\tfor cRows_%s.Next() {\n", ch.StructKey))
		sb.WriteString(fmt.Sprintf("\t\t\tvar item model.%s\n", childStructName))
		var cScanTargets []string
		for _, cc := range childCols {
			cScanTargets = append(cScanTargets, fmt.Sprintf("&item.%s", toPascalCase(cc.Name)))
		}
		sb.WriteString(fmt.Sprintf("\t\t\tif scanErr := cRows_%s.Scan(%s); scanErr == nil {\n", ch.StructKey, strings.Join(cScanTargets, ", ")))
		sb.WriteString(fmt.Sprintf("\t\t\t\tm.%s = append(m.%s, item)\n", ch.StructKey, ch.StructKey))
		sb.WriteString("\t\t\t}\n\t\t}\n\t}\n\n")
	}

	sb.WriteString("\treturn &m, nil\n}\n\n")

	// 2. List
	sb.WriteString(fmt.Sprintf("// List retrieves %s records with search, filter, and pagination\n", ent.EntityName))
	sb.WriteString(fmt.Sprintf("func (r *%s) List(ctx context.Context, opts model.ListOptions) ([]model.%s, int, error) {\n", repoName, structName))
	sb.WriteString(fmt.Sprintf("\tbaseQuery := `FROM %s WHERE 1=1`\n", ent.TableName))
	if hasDeletedAt {
		sb.WriteString("\tbaseQuery += ` AND deleted_at IS NULL`\n")
	}
	sb.WriteString("\tvar args []interface{}\n\targIndex := 1\n\n")

	// Searchable columns
	var searchCols []string
	for _, col := range cols {
		dt := strings.ToLower(col.DataType)
		if strings.Contains(dt, "char") || strings.Contains(dt, "text") {
			if col.Name != "created_by" && col.Name != "updated_by" && col.Name != "deleted_by" && col.Name != "password_hash" {
				searchCols = append(searchCols, col.Name)
			}
		}
	}
	if len(searchCols) > 0 {
		var parts []string
		for _, sc := range searchCols {
			parts = append(parts, fmt.Sprintf("%s ILIKE $%s", sc, "%d"))
		}
		cond := strings.Join(parts, " OR ")
		sb.WriteString("\tif opts.Search != \"\" {\n")
		sb.WriteString(fmt.Sprintf("\t\tbaseQuery += fmt.Sprintf(\" AND (%s)\", ", cond))
		var argTokens []string
		for range searchCols {
			argTokens = append(argTokens, "argIndex")
		}
		sb.WriteString(strings.Join(argTokens, ", ") + ")\n")
		sb.WriteString("\t\targs = append(args, \"%\"+opts.Search+\"%\")\n")
		sb.WriteString("\t\targIndex++\n\t}\n\n")
	}

	sb.WriteString("\t// Generic filter support\n")
	sb.WriteString("\tfor k, v := range opts.Filters {\n")
	sb.WriteString("\t\tif v != \"\" {\n")
	sb.WriteString("\t\t\tbaseQuery += fmt.Sprintf(\" AND %s = $%d\", k, argIndex)\n")
	sb.WriteString("\t\t\targs = append(args, v)\n")
	sb.WriteString("\t\t\targIndex++\n\t\t}\n\t}\n\n")

	sb.WriteString("\t// Count total\n\tvar total int\n")
	sb.WriteString("\tcountQuery := \"SELECT COUNT(*) \" + baseQuery\n")
	sb.WriteString("\terr := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total)\n")
	sb.WriteString("\tif err != nil {\n\t\treturn nil, 0, err\n\t}\n\n")

	sb.WriteString("\t// Sorting and Pagination\n")
	sb.WriteString(fmt.Sprintf("\tsortCol := \"%s\"\n", ent.PKColumn))
	sb.WriteString("\tif opts.SortBy != \"\" {\n\t\tsortCol = opts.SortBy\n\t}\n")
	sb.WriteString("\torder := \"DESC\"\n\tif strings.ToUpper(opts.Order) == \"ASC\" {\n\t\torder = \"ASC\"\n\t}\n")
	sb.WriteString("\tbaseQuery += fmt.Sprintf(\" ORDER BY %s %s\", sortCol, order)\n\n")
	sb.WriteString("\tif opts.Limit > 0 {\n")
	sb.WriteString("\t\tbaseQuery += fmt.Sprintf(\" LIMIT $%d OFFSET $%d\", argIndex, argIndex+1)\n")
	sb.WriteString("\t\targs = append(args, opts.Limit, opts.Offset)\n")
	sb.WriteString("\t}\n\n")

	sb.WriteString(fmt.Sprintf("\tselectQuery := \"SELECT %s \" + baseQuery\n", selectFields))
	sb.WriteString("\trows, err := r.DB.Query(ctx, selectQuery, args...)\n")
	sb.WriteString("\tif err != nil {\n\t\treturn nil, 0, err\n\t}\n\tdefer rows.Close()\n\n")

	sb.WriteString(fmt.Sprintf("\tvar items []model.%s\n", structName))
	sb.WriteString("\tfor rows.Next() {\n")
	sb.WriteString(fmt.Sprintf("\t\tvar m model.%s\n", structName))
	var listScanTargets []string
	for _, col := range cols {
		listScanTargets = append(listScanTargets, fmt.Sprintf("&m.%s", toPascalCase(col.Name)))
	}
	sb.WriteString(fmt.Sprintf("\t\tif err := rows.Scan(%s); err != nil {\n\t\t\treturn nil, 0, err\n\t\t}\n", strings.Join(listScanTargets, ", ")))
	sb.WriteString("\t\titems = append(items, m)\n\t}\n\n")
	sb.WriteString("\treturn items, total, nil\n}\n\n")

	// 3. Create
	sb.WriteString(fmt.Sprintf("// Create inserts a new %s\n", ent.EntityName))
	sb.WriteString(fmt.Sprintf("func (r *%s) Create(ctx context.Context, m *model.%s) (string, error) {\n", repoName, structName))
	sb.WriteString("\ttx, err := r.DB.Begin(ctx)\n\tif err != nil {\n\t\treturn \"\", err\n\t}\n\tdefer tx.Rollback(ctx)\n\n")

	var insertCols []string
	var insertVals []string
	var insertArgs []string
	idx := 1
	for _, col := range cols {
		if col.IsPK && col.Default != nil {
			continue
		}
		if col.IsGenerated {
			continue
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

	sb.WriteString(fmt.Sprintf("\tquery := `INSERT INTO %s (%s) VALUES (%s) RETURNING %s`\n",
		ent.TableName, strings.Join(insertCols, ", "), strings.Join(insertVals, ", "), ent.PKColumn))
	sb.WriteString(fmt.Sprintf("\tvar newID string\n\terr = tx.QueryRow(ctx, query, %s).Scan(&newID)\n", strings.Join(insertArgs, ", ")))
	sb.WriteString("\tif err != nil {\n\t\treturn \"\", err\n\t}\n\n")

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

	sb.WriteString("\treturn newID, tx.Commit(ctx)\n}\n\n")

	// 4. Update
	sb.WriteString(fmt.Sprintf("// Update modifies an existing %s\n", ent.EntityName))
	sb.WriteString(fmt.Sprintf("func (r *%s) Update(ctx context.Context, id string, m *model.%s) error {\n", repoName, structName))
	sb.WriteString("\ttx, err := r.DB.Begin(ctx)\n\tif err != nil {\n\t\treturn err\n\t}\n\tdefer tx.Rollback(ctx)\n\n")

	var setClauses []string
	var updateArgs []string
	uIdx := 1
	for _, col := range cols {
		if col.IsPK || col.IsGenerated || col.Name == "created_at" || col.Name == "created_by" || col.Name == "deleted_at" || col.Name == "deleted_by" {
			continue
		}
		if col.Name == "updated_at" {
			setClauses = append(setClauses, "updated_at = CURRENT_TIMESTAMP")
			continue
		}
		if col.UDTName != "" && !strings.HasPrefix(col.UDTName, "int") && !strings.HasPrefix(col.UDTName, "varchar") &&
			!strings.HasPrefix(col.UDTName, "text") && !strings.HasPrefix(col.UDTName, "bool") &&
			!strings.HasPrefix(col.UDTName, "uuid") && !strings.HasPrefix(col.UDTName, "numeric") &&
			!strings.HasPrefix(col.UDTName, "timestamp") && !strings.HasPrefix(col.UDTName, "date") && !strings.HasPrefix(col.UDTName, "json") {
			setClauses = append(setClauses, fmt.Sprintf("%s = $%d::%s", col.Name, uIdx, col.UDTName))
		} else {
			setClauses = append(setClauses, fmt.Sprintf("%s = $%d", col.Name, uIdx))
		}
		updateArgs = append(updateArgs, fmt.Sprintf("m.%s", toPascalCase(col.Name)))
		uIdx++
	}
	updateArgs = append(updateArgs, "id")

	sb.WriteString(fmt.Sprintf("\tquery := `UPDATE %s SET %s WHERE %s = $%d`\n",
		ent.TableName, strings.Join(setClauses, ", "), ent.PKColumn, uIdx))
	sb.WriteString(fmt.Sprintf("\t_, err = tx.Exec(ctx, query, %s)\n\tif err != nil {\n\t\treturn err\n\t}\n\n", strings.Join(updateArgs, ", ")))

	for _, ch := range ent.Children {
		childCols := tableCols[ch.TableName]
		sb.WriteString(fmt.Sprintf("\t// Replace children: %s\n", ch.StructKey))
		sb.WriteString(fmt.Sprintf("\tif len(m.%s) > 0 {\n", ch.StructKey))
		sb.WriteString(fmt.Sprintf("\t\t_, err := tx.Exec(ctx, `DELETE FROM %s WHERE %s = $1`, id)\n", ch.TableName, ch.FKColumn))
		sb.WriteString("\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n")

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
	sb.WriteString("package service\n\nimport (\n\t\"context\"\n\t\"errors\"\n\n\t\"github.com/figoalfarqi/navalerp/internal/model\"\n\t\"github.com/figoalfarqi/navalerp/internal/repository\"\n)\n\n")

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
	for _, ch := range ent.Children {
		sb.WriteString(fmt.Sprintf("\tm.%s = req.%s\n", ch.StructKey, ch.StructKey))
	}
	sb.WriteString("\tid, err := s.Repo.Create(ctx, m)\n\tif err != nil {\n\t\treturn nil, err, nil\n\t}\n\titem, err := s.Repo.Get(ctx, id)\n\treturn item, err, nil\n}\n\n")

	// Update
	sb.WriteString(fmt.Sprintf("func (s *%s) Update(ctx context.Context, loginID string, id string, req *model.%sRequest) (*model.%s, error, map[string]string) {\n", svcName, structName, structName))
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
	for _, ch := range ent.Children {
		sb.WriteString(fmt.Sprintf("\tm.%s = req.%s\n", ch.StructKey, ch.StructKey))
	}
	sb.WriteString("\terr := s.Repo.Update(ctx, id, m)\n\tif err != nil {\n\t\treturn nil, err, nil\n\t}\n\titem, err := s.Repo.Get(ctx, id)\n\treturn item, err, nil\n}\n\n")

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
	sb.WriteString("package handler\n\nimport (\n\t\"encoding/json\"\n\t\"net/http\"\n\t\"strconv\"\n\n\t\"github.com/figoalfarqi/navalerp/config\"\n\t\"github.com/figoalfarqi/navalerp/internal/app/middleware\"\n\t\"github.com/figoalfarqi/navalerp/internal/model\"\n\t\"github.com/figoalfarqi/navalerp/internal/service\"\n\t\"github.com/figoalfarqi/navalerp/pkg/response\"\n)\n\n")

	sb.WriteString(fmt.Sprintf("type %s struct {\n\tSvc *service.%s\n\tCfg *config.Config\n}\n\n", handlerName, svcName))
	sb.WriteString(fmt.Sprintf("func New%s(svc *service.%s, cfg *config.Config) *%s {\n\treturn &%s{Svc: svc, Cfg: cfg}\n}\n\n", handlerName, svcName, handlerName, handlerName))

	// Get
	sb.WriteString(fmt.Sprintf("func (h *%s) Get(w http.ResponseWriter, r *http.Request) {\n", handlerName))
	sb.WriteString("\tid := r.PathValue(\"id\")\n\tif id != \"\" {\n")
	sb.WriteString(fmt.Sprintf("\t\titem, err := h.Svc.GetByID(r.Context(), id)\n\t\tif err != nil {\n\t\t\tresponse.JSON(w, http.StatusNotFound, \"%s not found\", nil, nil)\n\t\t\treturn\n\t\t}\n", ent.EntityName))
	sb.WriteString(fmt.Sprintf("\t\tresponse.JSON(w, http.StatusOK, \"%s retrieved\", item, nil)\n\t\treturn\n\t}\n\n", ent.EntityName))

	sb.WriteString("\tq := r.URL.Query()\n\tpage, _ := strconv.Atoi(q.Get(\"page\"))\n\tif page < 1 { page = 1 }\n")
	sb.WriteString("\tlimit, _ := strconv.Atoi(q.Get(\"limit\"))\n\tif limit < 1 { limit = 10 }\n")
	sb.WriteString("\toffset := (page - 1) * limit\n")
	sb.WriteString("\tsearch := q.Get(\"search\")\n\tsortBy := q.Get(\"sort_by\")\n\torder := q.Get(\"order\")\n")
	sb.WriteString("\tfilters := make(map[string]string)\n\tfor k, v := range q {\n\t\tif k != \"page\" && k != \"limit\" && k != \"search\" && k != \"sort_by\" && k != \"order\" && len(v) > 0 {\n\t\t\tfilters[k] = v[0]\n\t\t}\n\t}\n\n")
	sb.WriteString("\topts := model.ListOptions{Limit: limit, Offset: offset, Search: search, SortBy: sortBy, Order: order, Filters: filters}\n")
	sb.WriteString("\titems, total, err := h.Svc.List(r.Context(), opts)\n")
	sb.WriteString("\tif err != nil {\n\t\tresponse.JSON(w, http.StatusInternalServerError, err.Error(), nil, nil)\n\t\treturn\n\t}\n")
	sb.WriteString(fmt.Sprintf("\tresponse.JSON(w, http.StatusOK, \"%s list retrieved\", map[string]any{\n\t\t\"items\": items,\n\t\t\"total\": total,\n\t\t\"page\":  page,\n\t\t\"limit\": limit,\n\t}, nil)\n}\n\n", ent.EntityName))

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

func titleCase(s string) string {
	parts := strings.Fields(s)
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + strings.ToLower(p[1:])
		}
	}
	return strings.Join(parts, " ")
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
		if col.Name == "deleted_at" || col.Name == "deleted_by" || col.Name == "notes" || col.Name == "findings" || col.Name == "recommended_action" {
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
	var pb strings.Builder
	pb.WriteString("\"use client\";\n\n")
	pb.WriteString("import Table from \"@/components/table/Table\";\n")
	pb.WriteString("import { columns, entityEndpoint, entityName, entityTitle, filterFields, primaryKey } from \"./_config\";\n\n")
	pb.WriteString(fmt.Sprintf("export default function %sPage() {\n", toPascalCase(ent.EntityName)))
	pb.WriteString("  return (\n")
	pb.WriteString("    <Table\n")
	pb.WriteString("      title={`Data ${entityTitle}`}\n")
	pb.WriteString("      endpoint={entityEndpoint}\n")
	pb.WriteString("      columns={columns}\n")
	pb.WriteString("      filterFields={filterFields}\n")
	pb.WriteString("      key_table={primaryKey}\n")
	pb.WriteString("      table_web_url={entityName}\n")
	pb.WriteString("    />\n")
	pb.WriteString("  );\n")
	pb.WriteString("}\n")
	os.WriteFile(filepath.Join(dir, "page.tsx"), []byte(pb.String()), 0644)

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
