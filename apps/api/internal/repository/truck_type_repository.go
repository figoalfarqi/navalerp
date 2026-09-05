package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/figoalfarqi/navalerp/internal/helper"
	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TruckTypeRepository struct {
	DB *pgxpool.Pool
}

func NewTruckTypeRepository(db *pgxpool.Pool) *TruckTypeRepository {
	return &TruckTypeRepository{DB: db}
}

// ==================================================
// Create
// ==================================================
func (r *TruckTypeRepository) Create(ctx context.Context, ft *model.TruckType) (int, error) {
	query := `
        INSERT INTO truck_type (
            truck_type_name,
            truck_type_description,
            truck_box_length,
            truck_box_width,
            truck_box_height,
            truck_capacity,
            is_active,
            created_by,
            updated_by
        ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
        RETURNING truck_type_id
    `

	var id int
	err := r.DB.QueryRow(
		ctx,
		query,
		ft.TruckTypeName,
		ft.TruckTypeDescription,
		ft.TruckBoxLength,
		ft.TruckBoxWidth,
		ft.TruckBoxHeight,
		ft.TruckCapacity,
		ft.IsActive,
		ft.CreatedBy,
		ft.UpdatedBy,
	).Scan(&id)

	return id, err
}

// ==================================================
// Update
// ==================================================
func (r *TruckTypeRepository) Update(ctx context.Context, id int, ft *model.TruckType) error {

	// conditional update is_active
	setIsActive := ""
	args := []interface{}{
		ft.TruckTypeName,        // $1
		ft.TruckTypeDescription, // $2
		ft.TruckBoxLength,       // $3
		ft.TruckBoxWidth,        // $4
		ft.TruckBoxHeight,       // $5
		ft.TruckCapacity,        // $6
	}

	argPos := len(args) + 1

	if ft.IsActive != -1 { // -1 artinya abaikan perubahan
		setIsActive = fmt.Sprintf(", is_active = $%d", argPos)
		args = append(args, ft.IsActive)
		argPos++
	}

	args = append(args, ft.UpdatedBy) // $argPos
	args = append(args, ft.UpdatedAt) // $argPos+1
	args = append(args, id)           // $argPos+2

	query := fmt.Sprintf(`
        UPDATE truck_type
        SET truck_type_name = $1,
            truck_type_description = $2,
            truck_box_length = $3,
            truck_box_width = $4,
            truck_box_height = $5,
            truck_capacity = $6
            %s,
            updated_by = $%d,
            updated_at = $%d
        WHERE truck_type_id = $%d AND deleted_at IS NULL
    `,
		setIsActive,
		argPos,   // updated_by
		argPos+1, // updated_at
		argPos+2, // id
	)

	_, err := r.DB.Exec(ctx, query, args...)
	return err
}

// ==================================================
// Soft Delete
// ==================================================
func (r *TruckTypeRepository) SoftDelete(ctx context.Context, id, deletedBy int) error {
	query := `
        UPDATE truck_type
        SET deleted_at = $1,
            deleted_by = $2,
            updated_at = $1,
            updated_by = $2,
            is_active = 0
        WHERE truck_type_id = $3 AND deleted_at IS NULL
    `
	_, err := r.DB.Exec(ctx, query, time.Now(), deletedBy, id)
	return err
}

// ==================================================
// Get By ID
// ==================================================
func (r *TruckTypeRepository) GetByID(ctx context.Context, id int) (*model.TruckType, error) {
	query := `
        SELECT 
            truck_type_id,
            truck_type_name,
            truck_type_description,
            truck_box_length,
            truck_box_width,
            truck_box_height,
            truck_capacity,
            is_active,
            created_by,
            updated_by,
            deleted_by,
            created_at,
            updated_at,
            deleted_at
        FROM truck_type
        WHERE truck_type_id = $1 AND deleted_at IS NULL
    `

	var ft model.TruckType
	err := r.DB.QueryRow(ctx, query, id).Scan(
		&ft.TruckTypeID,
		&ft.TruckTypeName,
		&ft.TruckTypeDescription,
		&ft.TruckBoxLength,
		&ft.TruckBoxWidth,
		&ft.TruckBoxHeight,
		&ft.TruckCapacity,
		&ft.IsActive,
		&ft.CreatedBy,
		&ft.UpdatedBy,
		&ft.DeletedBy,
		&ft.CreatedAt,
		&ft.UpdatedAt,
		&ft.DeletedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return &ft, nil
}

// ==================================================
// List + Dynamic Filters
// ==================================================
func (r *TruckTypeRepository) List(
	ctx context.Context,
	cursorValue interface{}, cursorKey *int, limit int, filters map[string]string,
	orderBy, sort string,
) ([]model.TruckType, error) {
	tableKey := "truck_type_id"
	parsedOrderBy := orderBy
	baseQuery := `
        SELECT 
            truck_type_id,
            truck_type_name,
            truck_type_description,
            truck_box_length,
            truck_box_width,
            truck_box_height,
            truck_capacity,
            is_active,
            created_by,
            updated_by,
            deleted_by,
            created_at,
            updated_at,
            deleted_at
        FROM truck_type
        WHERE deleted_at IS NULL
    `

	args := []interface{}{}
	argPos := 1

	// Cursor Pagination
	cursorClause, cursorArgs, argPos := helper.BuildCursorCondition(parsedOrderBy, tableKey, sort, cursorValue, cursorKey, argPos)
	baseQuery += cursorClause
	args = append(args, cursorArgs...)

	// LIKE filters
	likeFilters := map[string]string{
		"truck_type_name":        "truck_type_name ILIKE $%d",
		"truck_type_description": "truck_type_description ILIKE $%d",
		"truck_box_length":       "truck_box_length ILIKE $%d",
		"truck_box_width":        "truck_box_width ILIKE $%d",
		"truck_box_height":       "truck_box_height ILIKE $%d",
		"truck_capacity":         "truck_capacity ILIKE $%d",
	}
	for key, clause := range likeFilters {
		if v := filters[key]; v != "" {
			baseQuery += " AND " + fmt.Sprintf(clause, argPos)
			args = append(args, "%"+v+"%")
			argPos++
		}
	}

	// exact filters
	exactFilters := []string{"is_active", "created_by", "updated_by"}
	for _, key := range exactFilters {
		if v := filters[key]; v != "" {
			baseQuery += fmt.Sprintf(" AND %s = $%d", key, argPos)
			args = append(args, v)
			argPos++
		}
	}

	// time filters
	timeFilters := map[string]string{
		"created_at_after":  "created_at >= $%d",
		"created_at_before": "created_at <= $%d",
		"updated_at_after":  "updated_at >= $%d",
		"updated_at_before": "updated_at <= $%d",
	}
	for key, clause := range timeFilters {
		if v := filters[key]; v != "" {
			baseQuery += " AND " + fmt.Sprintf(clause, argPos)
			args = append(args, v)
			argPos++
		}
	}

	baseQuery += fmt.Sprintf(" ORDER BY %s %s, %s %s LIMIT $%d", parsedOrderBy, sort, tableKey, sort, argPos)

	args = append(args, limit)

	rows, err := r.DB.Query(ctx, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.TruckType
	for rows.Next() {
		var ft model.TruckType
		err := rows.Scan(
			&ft.TruckTypeID,
			&ft.TruckTypeName,
			&ft.TruckTypeDescription,
			&ft.TruckBoxLength,
			&ft.TruckBoxWidth,
			&ft.TruckBoxHeight,
			&ft.TruckCapacity,
			&ft.IsActive,
			&ft.CreatedBy,
			&ft.UpdatedBy,
			&ft.DeletedBy,
			&ft.CreatedAt,
			&ft.UpdatedAt,
			&ft.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, ft)
	}

	return list, rows.Err()
}
