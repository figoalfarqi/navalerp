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

type VendorTypeRepository struct {
	DB *pgxpool.Pool
}

func NewVendorTypeRepository(db *pgxpool.Pool) *VendorTypeRepository {
	return &VendorTypeRepository{DB: db}
}

// ==================================================
// Create VendorType
// ==================================================
func (r *VendorTypeRepository) Create(ctx context.Context, m *model.VendorType) (int, error) {
	query := `
		INSERT INTO vendor_type (
			vendor_type_name,
			vendor_type_description,
			is_active,
			created_by,
			updated_by
		) VALUES ($1,$2,$3,$4,$5)
		RETURNING vendor_type_id
	`

	var id int
	err := r.DB.QueryRow(ctx, query,
		m.VendorTypeName,
		m.VendorTypeDescription,
		m.IsActive,
		m.CreatedBy,
		m.UpdatedBy,
	).Scan(&id)

	return id, err
}

// ==================================================
// Update VendorType
// ==================================================
func (r *VendorTypeRepository) Update(ctx context.Context, id int, m *model.VendorType) error {

	setIsActive := ""
	args := []interface{}{
		m.VendorTypeName,        // $1
		m.VendorTypeDescription, // $2
	}
	argPos := len(args) + 1

	if m.IsActive != -1 {
		setIsActive = fmt.Sprintf(", is_active = $%d", argPos)
		args = append(args, m.IsActive)
		argPos++
	}

	args = append(args,
		m.UpdatedBy, // $argPos
		m.UpdatedAt, // $argPos+1
		id,          // $argPos+2
	)

	query := fmt.Sprintf(`
		UPDATE vendor_type
		SET vendor_type_name = $1,
			vendor_type_description = $2
			%s,
			updated_by = $%d,
			updated_at = $%d
		WHERE vendor_type_id = $%d AND deleted_at IS NULL
	`,
		setIsActive,
		argPos,   // updated_by
		argPos+1, // updated_at
		argPos+2, // where id
	)

	_, err := r.DB.Exec(ctx, query, args...)
	return err
}

// ==================================================
// Soft Delete VendorType
// ==================================================
func (r *VendorTypeRepository) SoftDelete(ctx context.Context, id, deletedBy int) error {
	query := `
		UPDATE vendor_type
		SET deleted_at = $1,
			deleted_by = $2,
			is_active = 0,
			updated_at = $1,
			updated_by = $2
		WHERE vendor_type_id = $3 AND deleted_at IS NULL
	`
	_, err := r.DB.Exec(ctx, query, time.Now(), deletedBy, id)
	return err
}

// ==================================================
// Get By ID
// ==================================================
func (r *VendorTypeRepository) GetByID(ctx context.Context, id int) (*model.VendorType, error) {
	query := `
		SELECT 
			vendor_type_id,
			vendor_type_name,
			vendor_type_description,
			is_active,
			created_by,
			updated_by,
			deleted_by,
			created_at,
			updated_at,
			deleted_at
		FROM vendor_type
		WHERE vendor_type_id = $1 AND deleted_at IS NULL
	`

	var m model.VendorType
	err := r.DB.QueryRow(ctx, query, id).Scan(
		&m.VendorTypeID,
		&m.VendorTypeName,
		&m.VendorTypeDescription,
		&m.IsActive,
		&m.CreatedBy,
		&m.UpdatedBy,
		&m.DeletedBy,
		&m.CreatedAt,
		&m.UpdatedAt,
		&m.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return &m, nil
}

// ==================================================
// List with Cursor + Filters
// ==================================================
func (r *VendorTypeRepository) List(
	ctx context.Context,
	cursorValue interface{}, cursorKey *int, limit int, filters map[string]string,
	orderBy, sort string,
) ([]model.VendorType, error) {
	tableKey := "vendor_type_id"
	parsedOrderBy := orderBy
	baseQuery := `
		SELECT 
			vendor_type_id,
			vendor_type_name,
			vendor_type_description,
			is_active,
			created_by,
			updated_by,
			deleted_by,
			created_at,
			updated_at,
			deleted_at
		FROM vendor_type
		WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	argPos := 1

	// ====================================
	// Cursor pagination
	// ====================================
	cursorClause, cursorArgs, argPos := helper.BuildCursorCondition(parsedOrderBy, tableKey, sort, cursorValue, cursorKey, argPos)
	baseQuery += cursorClause
	args = append(args, cursorArgs...)

	// ====================================
	// LIKE filters
	// ====================================
	likeFilters := map[string]string{
		"vendor_type_name":        "vendor_type_name ILIKE $%d",
		"vendor_type_description": "vendor_type_description ILIKE $%d",
	}
	for key, clause := range likeFilters {
		if v, ok := filters[key]; ok && v != "" {
			baseQuery += " AND " + fmt.Sprintf(clause, argPos)
			args = append(args, "%"+v+"%")
			argPos++
		}
	}

	// ==================================================
	// Exact match filters
	// ==================================================
	exactFilters := []string{
		"created_by",
		"updated_by",
		"is_active",
	}
	for _, key := range exactFilters {
		if v, ok := filters[key]; ok && v != "" {
			baseQuery += fmt.Sprintf(" AND %s = $%d", key, argPos)
			args = append(args, v)
			argPos++
		}
	}

	// ====================================
	// Time range filters
	// ====================================
	timeFilters := map[string]string{
		"created_at_after":  "created_at >= $%d",
		"created_at_before": "created_at <= $%d",
		"updated_at_after":  "updated_at >= $%d",
		"updated_at_before": "updated_at <= $%d",
	}

	for key, clause := range timeFilters {
		if v, ok := filters[key]; ok && v != "" {
			baseQuery += " AND " + fmt.Sprintf(clause, argPos)
			args = append(args, v)
			argPos++
		}
	}

	// ====================================
	// Order + limit
	// ====================================
	baseQuery += fmt.Sprintf(" ORDER BY %s %s, %s %s LIMIT $%d", parsedOrderBy, sort, tableKey, sort, argPos)

	args = append(args, limit)

	rows, err := r.DB.Query(ctx, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.VendorType

	for rows.Next() {
		var m model.VendorType
		if err := rows.Scan(
			&m.VendorTypeID,
			&m.VendorTypeName,
			&m.VendorTypeDescription,
			&m.IsActive,
			&m.CreatedBy,
			&m.UpdatedBy,
			&m.DeletedBy,
			&m.CreatedAt,
			&m.UpdatedAt,
			&m.DeletedAt,
		); err != nil {
			return nil, err
		}

		list = append(list, m)
	}

	return list, rows.Err()
}
