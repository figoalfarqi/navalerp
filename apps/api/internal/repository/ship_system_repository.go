package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ShipSystemRepository struct {
	DB *pgxpool.Pool
}

func NewShipSystemRepository(db *pgxpool.Pool) *ShipSystemRepository {
	return &ShipSystemRepository{DB: db}
}

// Get retrieves a single ship_system by system_id
func (r *ShipSystemRepository) Get(ctx context.Context, id string) (*model.ShipSystem, error) {
	query := `SELECT system_id, ship_id, parent_system_id, system_code, system_name, system_category, system_level, description, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM mro_systems WHERE system_id = $1 AND deleted_at IS NULL`

	var m model.ShipSystem
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.SystemId, &m.ShipId, &m.ParentSystemId, &m.SystemCode, &m.SystemName, &m.SystemCategory, &m.SystemLevel, &m.Description, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated ship_system records
func (r *ShipSystemRepository) List(ctx context.Context, opts model.ListOptions) ([]model.ShipSystem, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(system_code ILIKE $%[1]d OR system_name ILIKE $%[1]d OR description ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM mro_systems WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT system_id, ship_id, parent_system_id, system_code, system_name, system_category, system_level, description, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM mro_systems WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.ShipSystem
	for rows.Next() {
		var m model.ShipSystem
		if err := rows.Scan(&m.SystemId, &m.ShipId, &m.ParentSystemId, &m.SystemCode, &m.SystemName, &m.SystemCategory, &m.SystemLevel, &m.Description, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new ship_system with optional child items
func (r *ShipSystemRepository) Create(ctx context.Context, m *model.ShipSystem) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO mro_systems (ship_id, parent_system_id, system_code, system_name, system_category, system_level, description, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5::system_category_type, $6::system_level_type, $7, $8, $9, $10) RETURNING system_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.ShipId, m.ParentSystemId, m.SystemCode, m.SystemName, m.SystemCategory, m.SystemLevel, m.Description, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing ship_system
func (r *ShipSystemRepository) Update(ctx context.Context, id string, m *model.ShipSystem) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE mro_systems SET ship_id = $1, parent_system_id = $2, system_code = $3, system_name = $4, system_category = $5::system_category_type, system_level = $6::system_level_type, description = $7, updated_by = $8, updated_at = CURRENT_TIMESTAMP WHERE system_id = $9`
	_, err = tx.Exec(ctx, updateQuery, m.ShipId, m.ParentSystemId, m.SystemCode, m.SystemName, m.SystemCategory, m.SystemLevel, m.Description, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes ship_system
func (r *ShipSystemRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE mro_systems SET deleted_at = CURRENT_TIMESTAMP WHERE system_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
