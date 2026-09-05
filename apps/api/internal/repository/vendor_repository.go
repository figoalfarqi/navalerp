package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/figoalfarqi/navalerp/internal/helper"
	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type VendorRepository struct {
	DB *pgxpool.Pool
}

func NewVendorRepository(db *pgxpool.Pool) *VendorRepository {
	return &VendorRepository{DB: db}
}

// ==================================================
// Create (TX)
// ==================================================
func (r *VendorRepository) Create(
	ctx context.Context,
	m *model.Vendor,
) (int, error) {

	query := `
        INSERT INTO vendor (
            vendor_type_id,
            bank_merk_id,
            vendor_name,
            vendor_email,
            vendor_phone,
            vendor_tin,
            city_id,
            vendor_address,
            bank_account_number,
            bank_account_name,
            is_active,
            created_by,
            updated_by
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
        RETURNING vendor_id
    `

	var id int
	err := r.DB.QueryRow(ctx, query,
		m.VendorTypeID,
		m.BankMerkID,
		m.VendorName,
		m.VendorEmail,
		m.VendorPhone,
		m.VendorTin,
		m.CityID,
		m.VendorAddress,
		m.BankAccountNumber,
		m.BankAccountName,
		m.IsActive,
		m.CreatedBy,
		m.UpdatedBy,
	).Scan(&id)

	return id, err
}

// ==================================================
// Update (TX)
// ==================================================
func (r *VendorRepository) Update(
	ctx context.Context,
	id int,
	m *model.Vendor,
) error {

	// Field yang selalu di-update
	args := []interface{}{
		m.VendorTypeID,      // $1
		m.BankMerkID,        // $2
		m.VendorName,        // $3
		m.VendorEmail,       // $4
		m.VendorPhone,       // $5
		m.VendorTin,         // $6
		m.CityID,            // $7
		m.VendorAddress,     // $8
		m.BankAccountNumber, // $9
		m.BankAccountName,   // $10
	}

	// Mulai parameter posisi
	argPos := len(args) + 1

	// Optional field: is_active
	setIsActive := ""
	if m.IsActive != -1 { // jika bukan -1, update
		setIsActive = fmt.Sprintf(", is_active = $%d", argPos)
		args = append(args, m.IsActive)
		argPos++
	}

	// updated_by dan updated_at
	args = append(args,
		m.UpdatedBy, // $argPos
		m.UpdatedAt, // $argPos+1
		id,          // $argPos+2
	)

	query := fmt.Sprintf(`
        UPDATE vendor
        SET
            vendor_type_id = $1,
            bank_merk_id = $2,
            vendor_name = $3,
            vendor_email = $4,
            vendor_phone = $5,
            vendor_tin = $6,
            city_id = $7,
            vendor_address = $8,
            bank_account_number = $9,
            bank_account_name = $10
            %s,
            updated_by = $%d,
            updated_at = $%d
        WHERE vendor_id = $%d
          AND deleted_at IS NULL
        `,
		setIsActive,
		argPos,   // updated_by
		argPos+1, // updated_at
		argPos+2, // id
	)

	cmd, err := r.DB.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

// ==================================================
// Soft Delete
// ==================================================
func (r *VendorRepository) SoftDelete(ctx context.Context, deletedBy, id int) error {
	query := `
        UPDATE vendor
        SET deleted_at = $1,
            deleted_by = $2,
            updated_by = $2,
            updated_at = $1
        WHERE vendor_id = $3
          AND deleted_at IS NULL
    `

	_, err := r.DB.Exec(ctx, query, time.Now(), deletedBy, id)
	return err
}

// ==================================================
// Get By ID
// ==================================================
func (r *VendorRepository) GetByID(ctx context.Context, id int) (*model.Vendor, error) {

	query := `
        SELECT
            v.vendor_id,
            v.vendor_type_id,
            v.bank_merk_id,
            v.vendor_name,
            v.vendor_email,
            v.vendor_phone,
            v.vendor_tin,
            v.city_id,
            v.vendor_address,
            v.bank_account_number,
            v.bank_account_name,
            v.is_active,
            v.created_by,
            v.updated_by,
            v.created_at,
            v.updated_at,
            v.deleted_at,
            v.deleted_by,

            vt.vendor_type_id,
            vt.vendor_type_name,
            vt.vendor_type_description,
            vt.is_active as vt_is_active,

            bm.bank_merk_id,
            bm.bank_merk_name,
            bm.bank_merk_description,
            bm.is_active as bm_is_active,

            ci.city_id,
            ci.province_id,
            ci.city_name,
            ci.is_active as ci_is_active,
            ci.created_at as ci_created_at,
            ci.updated_at as ci_updated_at,

            COALESCE(trucks.data, '[]')

        FROM vendor v
        LEFT JOIN vendor_type vt ON vt.vendor_type_id = v.vendor_type_id AND vt.deleted_at IS NULL
        LEFT JOIN bank_merk bm ON bm.bank_merk_id = v.bank_merk_id AND bm.deleted_at IS NULL
        LEFT JOIN city ci ON ci.city_id = v.city_id AND ci.deleted_at IS NULL

        LEFT JOIN LATERAL (
            SELECT json_agg(json_build_object(
                'truck_id', t.truck_id,
                'license_plate', t.license_plate,
                'ownership_status_id', t.ownership_status_id,
                'production_year', t.production_year,
                'number_of_tires', t.number_of_tires,
                'is_active', t.is_active,
                'created_at', t.created_at,
                'updated_at', t.updated_at
            )) AS data
            FROM truck t
            WHERE t.vendor_id = v.vendor_id
              AND t.deleted_at IS NULL
        ) trucks ON TRUE

        WHERE v.vendor_id = $1
          AND v.deleted_at IS NULL
    `

	var m model.Vendor
	var vt model.VendorType
	var bm model.BankMerk
	var city model.CityNullable
	var trucksJSON []byte

	err := r.DB.QueryRow(ctx, query, id).Scan(
		&m.VendorID,
		&m.VendorTypeID,
		&m.BankMerkID,
		&m.VendorName,
		&m.VendorEmail,
		&m.VendorPhone,
		&m.VendorTin,
		&m.CityID,
		&m.VendorAddress,
		&m.BankAccountNumber,
		&m.BankAccountName,
		&m.IsActive,
		&m.CreatedBy,
		&m.UpdatedBy,
		&m.CreatedAt,
		&m.UpdatedAt,
		&m.DeletedAt,
		&m.DeletedBy,

		&vt.VendorTypeID,
		&vt.VendorTypeName,
		&vt.VendorTypeDescription,
		&vt.IsActive,

		&bm.BankMerkID,
		&bm.BankMerkName,
		&bm.BankMerkDescription,
		&bm.IsActive,

		&city.CityID,
		&city.ProvinceID,
		&city.CityName,
		&city.IsActive,
		&city.CreatedAt,
		&city.UpdatedAt,

		&trucksJSON,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	// Set VendorType if exists
	if vt.VendorTypeID != 0 {
		m.VendorType = &vt
	}

	// Set BankMerk if exists
	if bm.BankMerkID != 0 {
		m.BankMerk = &bm
	}

	// Set City if exists
	if city.CityID != nil {
		m.City = city.ToNotNullable()
	}

	// Parse Trucks
	if len(trucksJSON) > 0 && string(trucksJSON) != "null" {
		var trucks []model.Truck
		_ = json.Unmarshal(trucksJSON, &trucks)
		m.Trucks = trucks
	}

	return &m, nil
}

// ==================================================
// List
// ==================================================
func (r *VendorRepository) List(
	ctx context.Context,
	cursorValue interface{},
	cursorKey *int,
	limit int,
	filters map[string]string,
	orderBy, sort string,
) ([]model.Vendor, error) {

	tableKey := "vendor_id"
	parsedOrderBy := "v." + orderBy

	baseQuery := `
        SELECT
            v.vendor_id,
            v.vendor_type_id,
            v.bank_merk_id,
            v.vendor_name,
            v.vendor_email,
            v.vendor_phone,
            v.vendor_tin,
            v.city_id,
            v.vendor_address,
            v.bank_account_number,
            v.bank_account_name,
            v.is_active,
            v.created_by,
            v.updated_by,
            v.created_at,
            v.updated_at,
            v.deleted_at,
            v.deleted_by,

            vt.vendor_type_id,
            vt.vendor_type_name,
            vt.is_active as vt_is_active,

            bm.bank_merk_id,
            bm.bank_merk_name,
            bm.is_active as bm_is_active,

            ci.city_id,
            ci.province_id,
            ci.city_name,
            ci.is_active as ci_is_active,

            COALESCE(trucks.data, '[]')

        FROM vendor v
        LEFT JOIN vendor_type vt ON vt.vendor_type_id = v.vendor_type_id AND vt.deleted_at IS NULL
        LEFT JOIN bank_merk bm ON bm.bank_merk_id = v.bank_merk_id AND bm.deleted_at IS NULL
        LEFT JOIN city ci ON ci.city_id = v.city_id AND ci.deleted_at IS NULL

        LEFT JOIN LATERAL (
            SELECT json_agg(json_build_object(
                'truck_id', t.truck_id,
                'license_plate', t.license_plate,
                'is_active', t.is_active
            )) AS data
            FROM truck t
            WHERE t.vendor_id = v.vendor_id
              AND t.deleted_at IS NULL
            LIMIT 10
        ) trucks ON TRUE

        WHERE v.deleted_at IS NULL
    `

	args := []any{}
	argPos := 1

	// Cursor pagination
	cursorClause, cursorArgs, argPos := helper.BuildCursorCondition(
		parsedOrderBy,
		tableKey,
		sort,
		cursorValue,
		cursorKey,
		argPos,
	)
	baseQuery += cursorClause
	args = append(args, cursorArgs...)

	// Exact match filters
	exactFilters := []string{
		"vendor_type_id",
		"bank_merk_id",
		"city_id",
		"is_active",
	}

	for _, key := range exactFilters {
		if v, ok := filters[key]; ok && v != "" {
			baseQuery += fmt.Sprintf(" AND v.%s = $%d", key, argPos)
			args = append(args, v)
			argPos++
		}
	}

	// Time filters
	timeFilters := map[string]string{
		"created_at_after":  "v.created_at >= $%d",
		"created_at_before": "v.created_at <= $%d",
		"updated_at_after":  "v.updated_at >= $%d",
		"updated_at_before": "v.updated_at <= $%d",
	}

	for key, clause := range timeFilters {
		if v, ok := filters[key]; ok && v != "" {
			baseQuery += " AND " + fmt.Sprintf(clause, argPos)
			args = append(args, v)
			argPos++
		}
	}

	// Search by vendor name
	if v, ok := filters["vendor_name"]; ok && v != "" {
		baseQuery += fmt.Sprintf(" AND v.vendor_name ILIKE $%d", argPos)
		args = append(args, "%"+v+"%")
		argPos++
	}

	// Search by email
	if v, ok := filters["vendor_email"]; ok && v != "" {
		baseQuery += fmt.Sprintf(" AND v.vendor_email ILIKE $%d", argPos)
		args = append(args, "%"+v+"%")
		argPos++
	}

	// Search by phone
	if v, ok := filters["vendor_phone"]; ok && v != "" {
		baseQuery += fmt.Sprintf(" AND v.vendor_phone ILIKE $%d", argPos)
		args = append(args, "%"+v+"%")
		argPos++
	}

	// Search by TIN
	if v, ok := filters["vendor_tin"]; ok && v != "" {
		baseQuery += fmt.Sprintf(" AND v.vendor_tin ILIKE $%d", argPos)
		args = append(args, "%"+v+"%")
		argPos++
	}

	// Filter by vendor type name
	if v, ok := filters["vendor_type_name"]; ok && v != "" {
		baseQuery += fmt.Sprintf(" AND vt.vendor_type_name ILIKE $%d", argPos)
		args = append(args, "%"+v+"%")
		argPos++
	}

	// Filter by city name
	if v, ok := filters["city_name"]; ok && v != "" {
		baseQuery += fmt.Sprintf(" AND ci.city_name ILIKE $%d", argPos)
		args = append(args, "%"+v+"%")
		argPos++
	}

	// Order + limit
	baseQuery += fmt.Sprintf(" ORDER BY %s %s, %s %s LIMIT $%d",
		parsedOrderBy, sort, tableKey, sort, argPos)

	args = append(args, limit)

	rows, err := r.DB.Query(ctx, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []model.Vendor{}

	for rows.Next() {
		var m model.Vendor
		var vt model.VendorType
		var bm model.BankMerk
		var city model.CityNullable
		var trucksJSON []byte

		err := rows.Scan(
			&m.VendorID,
			&m.VendorTypeID,
			&m.BankMerkID,
			&m.VendorName,
			&m.VendorEmail,
			&m.VendorPhone,
			&m.VendorTin,
			&m.CityID,
			&m.VendorAddress,
			&m.BankAccountNumber,
			&m.BankAccountName,
			&m.IsActive,
			&m.CreatedBy,
			&m.UpdatedBy,
			&m.CreatedAt,
			&m.UpdatedAt,
			&m.DeletedAt,
			&m.DeletedBy,

			&vt.VendorTypeID,
			&vt.VendorTypeName,
			&vt.IsActive,

			&bm.BankMerkID,
			&bm.BankMerkName,
			&bm.IsActive,

			&city.CityID,
			&city.ProvinceID,
			&city.CityName,
			&city.IsActive,

			&trucksJSON,
		)
		if err != nil {
			return nil, err
		}

		// Set VendorType if exists
		if vt.VendorTypeID != 0 {
			m.VendorType = &vt
		}

		// Set BankMerk if exists
		if bm.BankMerkID != 0 {
			m.BankMerk = &bm
		}

		// Set City if exists
		if city.CityID != nil {
			m.City = city.ToNotNullable()
		}

		// Parse Trucks
		if len(trucksJSON) > 0 && string(trucksJSON) != "null" {
			var trucks []model.Truck
			_ = json.Unmarshal(trucksJSON, &trucks)
			m.Trucks = trucks
		}

		result = append(result, m)
	}

	return result, nil
}
