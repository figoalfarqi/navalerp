package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/figoalfarqi/apipml/internal/helper"
	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TruckRepository struct {
	DB *pgxpool.Pool
}

func NewTruckRepository(db *pgxpool.Pool) *TruckRepository {
	return &TruckRepository{DB: db}
}

// ==================================================
// Create (TX)
// ==================================================
func (r *TruckRepository) Create(
	ctx context.Context,
	m *model.Truck,
) (int, error) {

	query := `
        INSERT INTO truck (
            truck_type_id,
            truck_merk_id,
            driver_id,
            vendor_id,
            license_plate,
            ownership_status_id,
            production_year,
            number_of_tires,
            is_active,
            created_by,
            updated_by
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
        RETURNING truck_id
    `

	var id int
	err := r.DB.QueryRow(ctx, query,
		m.TruckTypeID,
		m.TruckMerkID,
		m.DriverID,
		m.VendorID,
		m.LicensePlate,
		m.OwnershipStatusID,
		m.ProductionYear,
		m.NumberOfTires,
		m.IsActive,
		m.CreatedBy,
		m.UpdatedBy,
	).Scan(&id)
	return id, err
}

// ==================================================
// Update (TX)
// ==================================================
func (r *TruckRepository) Update(
	ctx context.Context,
	id int,
	m *model.Truck,
) error {

	// Field yang selalu di-update
	args := []interface{}{
		m.TruckTypeID,       // $1
		m.TruckMerkID,       // $2
		m.DriverID,          // $3
		m.VendorID,          // $4
		m.LicensePlate,      // $5
		m.OwnershipStatusID, // $6
		m.ProductionYear,    // $7
		m.NumberOfTires,     // $8
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
        UPDATE truck
        SET
            truck_type_id = $1,
            truck_merk_id = $2,
            driver_id = $3,
            vendor_id = $4,
            license_plate = $5,
            ownership_status_id = $6,
            production_year = $7,
            number_of_tires = $8
            %s,
            updated_by = $%d,
            updated_at = $%d
        WHERE truck_id = $%d
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
func (r *TruckRepository) SoftDelete(ctx context.Context, deletedBy, id int) error {
	query := `
        UPDATE truck
        SET deleted_at = $1,
            deleted_by = $2,
            updated_by = $2,
            updated_at = $1
        WHERE truck_id = $3
          AND deleted_at IS NULL
    `

	_, err := r.DB.Exec(ctx, query, time.Now(), deletedBy, id)
	return err
}

// ==================================================
// Get By ID
// ==================================================
func (r *TruckRepository) GetByID(ctx context.Context, id int) (*model.Truck, error) {

	query := `
        SELECT
            t.truck_id,
            t.truck_type_id,
            t.truck_merk_id,
            t.driver_id,
            t.vendor_id,
            t.license_plate,
            t.ownership_status_id,
            t.production_year,
            t.number_of_tires,
            t.is_active,
            t.created_by,
            t.updated_by,
            t.created_at,
            t.updated_at,
            t.deleted_at,
            t.deleted_by,

            tt.truck_type_id,
            tt.truck_type_name,
            tt.truck_type_description,
            tt.truck_box_length,
            tt.truck_box_width,
            tt.truck_box_height,
            tt.truck_capacity,
            tt.is_active as tt_is_active,

            tm.truck_merk_id,
            tm.truck_merk_name,
            tm.is_active as tm_is_active,

            d.app_user_id,
            d.app_user_name,
            d.app_user_phone,
            d.username,

            v.vendor_id,
            v.vendor_name,
            v.vendor_email,
            v.vendor_phone,
            v.is_active as v_is_active

        FROM truck t
        LEFT JOIN truck_type tt ON tt.truck_type_id = t.truck_type_id AND tt.deleted_at IS NULL
        LEFT JOIN truck_merk tm ON tm.truck_merk_id = t.truck_merk_id AND tm.deleted_at IS NULL
        LEFT JOIN app_user d ON d.app_user_id = t.driver_id AND d.deleted_at IS NULL
        LEFT JOIN vendor v ON v.vendor_id = t.vendor_id AND v.deleted_at IS NULL

        WHERE t.truck_id = $1
          AND t.deleted_at IS NULL
    `

	var m model.Truck
	var tt model.TruckType
	var tm model.TruckMerk
	var d model.AppUser
	var v model.VendorNullable

	err := r.DB.QueryRow(ctx, query, id).Scan(
		&m.TruckID,
		&m.TruckTypeID,
		&m.TruckMerkID,
		&m.DriverID,
		&m.VendorID,
		&m.LicensePlate,
		&m.OwnershipStatusID,
		&m.ProductionYear,
		&m.NumberOfTires,
		&m.IsActive,
		&m.CreatedBy,
		&m.UpdatedBy,
		&m.CreatedAt,
		&m.UpdatedAt,
		&m.DeletedAt,
		&m.DeletedBy,

		&tt.TruckTypeID,
		&tt.TruckTypeName,
		&tt.TruckTypeDescription,
		&tt.TruckBoxLength,
		&tt.TruckBoxWidth,
		&tt.TruckBoxHeight,
		&tt.TruckCapacity,
		&tt.IsActive,

		&tm.TruckMerkID,
		&tm.TruckMerkName,
		&tm.IsActive,

		&d.AppUserID,
		&d.AppUserName,
		&d.AppUserPhone,
		&d.Username,

		&v.VendorID,
		&v.VendorName,
		&v.VendorEmail,
		&v.VendorPhone,
		&v.IsActive,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	// Set TruckType if exists
	if tt.TruckTypeID != 0 {
		m.TruckType = &tt
	}

	// Set TruckMerk if exists
	if tm.TruckMerkID != 0 {
		m.TruckMerk = &tm
	}

	// Set Driver if exists
	if d.AppUserID != 0 {
		m.Driver = &d
	}

	// Set Vendor if exists
	if v.VendorID != nil {
		m.Vendor = v.ToNotNullable()
	}

	return &m, nil
}

// ==================================================
// List
// ==================================================
func (r *TruckRepository) List(
	ctx context.Context,
	cursorValue interface{},
	cursorKey *int,
	limit int,
	filters map[string]string,
	orderBy, sort string,
) ([]model.Truck, error) {

	tableKey := "truck_id"
	parsedOrderBy := "t." + orderBy

	baseQuery := `
        SELECT
            t.truck_id,
            t.truck_type_id,
            t.truck_merk_id,
            t.driver_id,
            t.vendor_id,
            t.license_plate,
            t.ownership_status_id,
            t.production_year,
            t.number_of_tires,
            t.is_active,
            t.created_by,
            t.updated_by,
            t.created_at,
            t.updated_at,
            t.deleted_at,
            t.deleted_by,

            tt.truck_type_id,
            tt.truck_type_name,
            tt.truck_box_length,
            tt.truck_box_width,
            tt.truck_box_height,
            tt.truck_capacity,
            tt.is_active as tt_is_active,

            tm.truck_merk_id,
            tm.truck_merk_name,
            tm.is_active as tm_is_active,

            d.app_user_id,
            d.app_user_name,
            d.app_user_phone,

            v.vendor_id,
            v.vendor_name,
            v.is_active as v_is_active

        FROM truck t
        LEFT JOIN truck_type tt ON tt.truck_type_id = t.truck_type_id AND tt.deleted_at IS NULL
        LEFT JOIN truck_merk tm ON tm.truck_merk_id = t.truck_merk_id AND tm.deleted_at IS NULL
        LEFT JOIN app_user d ON d.app_user_id = t.driver_id AND d.deleted_at IS NULL
        LEFT JOIN vendor v ON v.vendor_id = t.vendor_id AND v.deleted_at IS NULL

        WHERE t.deleted_at IS NULL
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
		"truck_type_id",
		"truck_merk_id",
		"driver_id",
		"vendor_id",
		"ownership_status_id",
		"is_active",
	}

	for _, key := range exactFilters {
		if v, ok := filters[key]; ok && v != "" {
			baseQuery += fmt.Sprintf(" AND t.%s = $%d", key, argPos)
			args = append(args, v)
			argPos++
		}
	}

	// Time filters
	timeFilters := map[string]string{
		"created_at_after":  "t.created_at >= $%d",
		"created_at_before": "t.created_at <= $%d",
		"updated_at_after":  "t.updated_at >= $%d",
		"updated_at_before": "t.updated_at <= $%d",
	}

	for key, clause := range timeFilters {
		if v, ok := filters[key]; ok && v != "" {
			baseQuery += " AND " + fmt.Sprintf(clause, argPos)
			args = append(args, v)
			argPos++
		}
	}

	// Search by license plate
	if v, ok := filters["license_plate"]; ok && v != "" {
		baseQuery += fmt.Sprintf(" AND t.license_plate ILIKE $%d", argPos)
		args = append(args, "%"+v+"%")
		argPos++
	}

	// Filter by production year range
	if v, ok := filters["production_year_min"]; ok && v != "" {
		baseQuery += fmt.Sprintf(" AND t.production_year >= $%d", argPos)
		args = append(args, v)
		argPos++
	}
	if v, ok := filters["production_year_max"]; ok && v != "" {
		baseQuery += fmt.Sprintf(" AND t.production_year <= $%d", argPos)
		args = append(args, v)
		argPos++
	}

	// Filter by number of tires
	if v, ok := filters["number_of_tires"]; ok && v != "" {
		baseQuery += fmt.Sprintf(" AND t.number_of_tires = $%d", argPos)
		args = append(args, v)
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

	result := []model.Truck{}

	for rows.Next() {
		var m model.Truck
		var tt model.TruckType
		var tm model.TruckMerk
		var d model.AppUser
		var v model.Vendor

		err := rows.Scan(
			&m.TruckID,
			&m.TruckTypeID,
			&m.TruckMerkID,
			&m.DriverID,
			&m.VendorID,
			&m.LicensePlate,
			&m.OwnershipStatusID,
			&m.ProductionYear,
			&m.NumberOfTires,
			&m.IsActive,
			&m.CreatedBy,
			&m.UpdatedBy,
			&m.CreatedAt,
			&m.UpdatedAt,
			&m.DeletedAt,
			&m.DeletedBy,

			&tt.TruckTypeID,
			&tt.TruckTypeName,
			&tt.TruckBoxLength,
			&tt.TruckBoxWidth,
			&tt.TruckBoxHeight,
			&tt.TruckCapacity,
			&tt.IsActive,

			&tm.TruckMerkID,
			&tm.TruckMerkName,
			&tm.IsActive,

			&d.AppUserID,
			&d.AppUserName,
			&d.AppUserPhone,

			&v.VendorID,
			&v.VendorName,
			&v.IsActive,
		)
		if err != nil {
			return nil, err
		}

		// Set TruckType if exists
		if tt.TruckTypeID != 0 {
			m.TruckType = &tt
		}

		// Set TruckMerk if exists
		if tm.TruckMerkID != 0 {
			m.TruckMerk = &tm
		}

		// Set Driver if exists
		if d.AppUserID != 0 {
			m.Driver = &d
		}

		// Set Vendor if exists
		if v.VendorID != 0 {
			m.Vendor = &v
		}

		result = append(result, m)
	}

	return result, nil
}
