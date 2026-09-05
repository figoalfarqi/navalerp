package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/figoalfarqi/navalerp/internal/helper"
	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TruckMerkRepository struct {
	DB *pgxpool.Pool
}

func NewTruckMerkRepository(db *pgxpool.Pool) *TruckMerkRepository {
	return &TruckMerkRepository{DB: db}
}

// ==================================================
// Create TruckMerk
// ==================================================
func (r *TruckMerkRepository) Create(ctx context.Context, m *model.TruckMerk) (int, error) {
	query := `
		INSERT INTO truck_merk (
			truck_merk_name,
			is_active,
			created_by,
			updated_by
		) VALUES ($1,$2,$3,$4)
		RETURNING truck_merk_id
	`

	var id int
	err := r.DB.QueryRow(ctx, query,
		m.TruckMerkName,
		m.IsActive,
		m.CreatedBy,
		m.UpdatedBy,
	).Scan(&id)

	return id, err
}

// ==================================================
// Update TruckMerk
// ==================================================
func (r *TruckMerkRepository) Update(ctx context.Context, id int, m *model.TruckMerk) error {

	setIsActive := ""
	args := []interface{}{
		m.TruckMerkName, // $1
	}
	argPos := 2

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
		UPDATE truck_merk
		SET truck_merk_name = $1
			%s,
			updated_by = $%d,
			updated_at = $%d
		WHERE truck_merk_id = $%d AND deleted_at IS NULL
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
// Soft Delete TruckMerk
// ==================================================
func (r *TruckMerkRepository) SoftDelete(ctx context.Context, id, deletedBy int) error {
	query := `
		UPDATE truck_merk
		SET deleted_at = $1,
			deleted_by = $2,
			is_active = 0,
			updated_at = $1,
			updated_by = $2
		WHERE truck_merk_id = $3 AND deleted_at IS NULL
	`
	_, err := r.DB.Exec(ctx, query, time.Now(), deletedBy, id)
	return err
}

// ==================================================
// Get By ID
// ==================================================
func (r *TruckMerkRepository) GetByID(ctx context.Context, id int) (*model.TruckMerk, error) {
	query := `
		SELECT 
			truck_merk_id,
			truck_merk_name,
			is_active,
			created_by,
			updated_by,
			deleted_by,
			created_at,
			updated_at,
			deleted_at
		FROM truck_merk
		WHERE truck_merk_id = $1 AND deleted_at IS NULL
	`

	var m model.TruckMerk
	err := r.DB.QueryRow(ctx, query, id).Scan(
		&m.TruckMerkID,
		&m.TruckMerkName,
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
func (r *TruckMerkRepository) List(
	ctx context.Context,
	cursorValue interface{}, cursorKey *int, limit int, filters map[string]string,
	orderBy, sort string,
) ([]model.TruckMerk, error) {
	tableKey := "truck_merk_id"
	parsedOrderBy := orderBy
	baseQuery := `
		SELECT 
			truck_merk_id,
			truck_merk_name,
			is_active,
			created_by,
			updated_by,
			deleted_by,
			created_at,
			updated_at,
			deleted_at
		FROM truck_merk
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
		"truck_merk_name": "truck_merk_name ILIKE $%d",
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

	var list []model.TruckMerk

	for rows.Next() {
		var m model.TruckMerk
		if err := rows.Scan(
			&m.TruckMerkID,
			&m.TruckMerkName,
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
