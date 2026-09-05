package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ShipClassRepository struct {
	DB *pgxpool.Pool
}

func NewShipClassRepository(db *pgxpool.Pool) *ShipClassRepository {
	return &ShipClassRepository{DB: db}
}

// Get retrieves a single ship_class by class_id
func (r *ShipClassRepository) Get(ctx context.Context, id string) (*model.ShipClass, error) {
	query := `SELECT class_id, class_code, class_name, category, specifications, builder, total_built, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM mro_ship_classes WHERE class_id = $1 AND deleted_at IS NULL`

	var m model.ShipClass
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.ClassId, &m.ClassCode, &m.ClassName, &m.Category, &m.Specifications, &m.Builder, &m.TotalBuilt, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated ship_class records
func (r *ShipClassRepository) List(ctx context.Context, opts model.ListOptions) ([]model.ShipClass, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(class_code ILIKE $%[1]d OR class_name ILIKE $%[1]d OR builder ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM mro_ship_classes WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT class_id, class_code, class_name, category, specifications, builder, total_built, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM mro_ship_classes WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.ShipClass
	for rows.Next() {
		var m model.ShipClass
		if err := rows.Scan(&m.ClassId, &m.ClassCode, &m.ClassName, &m.Category, &m.Specifications, &m.Builder, &m.TotalBuilt, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new ship_class with optional child items
func (r *ShipClassRepository) Create(ctx context.Context, m *model.ShipClass) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO mro_ship_classes (class_code, class_name, category, specifications, builder, total_built, created_by, updated_by, deleted_by) VALUES ($1, $2, $3::ship_category_type, $4, $5, $6, $7, $8, $9) RETURNING class_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.ClassCode, m.ClassName, m.Category, m.Specifications, m.Builder, m.TotalBuilt, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing ship_class
func (r *ShipClassRepository) Update(ctx context.Context, id string, m *model.ShipClass) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE mro_ship_classes SET class_code = $1, class_name = $2, category = $3::ship_category_type, specifications = $4, builder = $5, total_built = $6, updated_by = $7, updated_at = CURRENT_TIMESTAMP WHERE class_id = $8`
	_, err = tx.Exec(ctx, updateQuery, m.ClassCode, m.ClassName, m.Category, m.Specifications, m.Builder, m.TotalBuilt, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes ship_class
func (r *ShipClassRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE mro_ship_classes SET deleted_at = CURRENT_TIMESTAMP WHERE class_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
