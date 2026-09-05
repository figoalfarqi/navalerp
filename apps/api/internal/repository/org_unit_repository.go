package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrgUnitRepository struct {
	DB *pgxpool.Pool
}

func NewOrgUnitRepository(db *pgxpool.Pool) *OrgUnitRepository {
	return &OrgUnitRepository{DB: db}
}

// Get retrieves a single org_unit by unit_id
func (r *OrgUnitRepository) Get(ctx context.Context, id string) (*model.OrgUnit, error) {
	query := `SELECT t.unit_id, t.parent_unit_id, COALESCE(j_punit.unit_name, ''), t.unit_code, t.unit_name, t.unit_type, t.description, t.command_level, t.latitude, t.longitude, t.address, t.phone, t.is_active, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM org_units t
	LEFT JOIN org_units j_punit ON j_punit.unit_id = t.parent_unit_id
	WHERE t.unit_id = $1 AND t.deleted_at IS NULL`

	var m model.OrgUnit
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.UnitId, &m.ParentUnitId, &m.ParentUnitName, &m.UnitCode, &m.UnitName, &m.UnitType, &m.Description, &m.CommandLevel, &m.Latitude, &m.Longitude, &m.Address, &m.Phone, &m.IsActive, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated org_unit records
func (r *OrgUnitRepository) List(ctx context.Context, opts model.ListOptions) ([]model.OrgUnit, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "t.deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(unit_code ILIKE $%[1]d OR unit_name ILIKE $%[1]d OR description ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM org_units t WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf(`SELECT t.unit_id, t.parent_unit_id, COALESCE(j_punit.unit_name, ''), t.unit_code, t.unit_name, t.unit_type, t.description, t.command_level, t.latitude, t.longitude, t.address, t.phone, t.is_active, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM org_units t
	LEFT JOIN org_units j_punit ON j_punit.unit_id = t.parent_unit_id
	WHERE %s ORDER BY t.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.OrgUnit
	for rows.Next() {
		var m model.OrgUnit
		if err := rows.Scan(&m.UnitId, &m.ParentUnitId, &m.ParentUnitName, &m.UnitCode, &m.UnitName, &m.UnitType, &m.Description, &m.CommandLevel, &m.Latitude, &m.Longitude, &m.Address, &m.Phone, &m.IsActive, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new org_unit with optional child items
func (r *OrgUnitRepository) Create(ctx context.Context, m *model.OrgUnit) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO org_units (parent_unit_id, unit_code, unit_name, unit_type, description, command_level, latitude, longitude, address, phone, is_active, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4::org_unit_type, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING unit_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.ParentUnitId, m.UnitCode, m.UnitName, m.UnitType, m.Description, m.CommandLevel, m.Latitude, m.Longitude, m.Address, m.Phone, m.IsActive, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing org_unit
func (r *OrgUnitRepository) Update(ctx context.Context, id string, m *model.OrgUnit) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE org_units SET parent_unit_id = $1, unit_code = $2, unit_name = $3, unit_type = $4::org_unit_type, description = $5, command_level = $6, latitude = $7, longitude = $8, address = $9, phone = $10, is_active = $11, updated_by = $12, updated_at = CURRENT_TIMESTAMP WHERE unit_id = $13`
	_, err = tx.Exec(ctx, updateQuery, m.ParentUnitId, m.UnitCode, m.UnitName, m.UnitType, m.Description, m.CommandLevel, m.Latitude, m.Longitude, m.Address, m.Phone, m.IsActive, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes org_unit
func (r *OrgUnitRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE org_units SET deleted_at = CURRENT_TIMESTAMP WHERE unit_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
