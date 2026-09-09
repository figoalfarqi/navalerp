package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransportUnitRepository struct {
	DB *pgxpool.Pool
}

func NewTransportUnitRepository(db *pgxpool.Pool) *TransportUnitRepository {
	return &TransportUnitRepository{DB: db}
}

// Get retrieves a single transport_unit by transport_unit_id
func (r *TransportUnitRepository) Get(ctx context.Context, id string) (*model.TransportUnit, error) {
	query := `SELECT t.transport_unit_id, t.unit_code, t.unit_name, t.transport_type, t.cargo_capacity_tons, t.fuel_capacity_liters, t.operating_unit_id, COALESCE(j_unit.unit_name, ''), t.status, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM log_transport_units t
	LEFT JOIN org_units j_unit ON j_unit.unit_id = t.operating_unit_id
	WHERE t.transport_unit_id = $1 AND t.deleted_at IS NULL`

	var m model.TransportUnit
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.TransportUnitId, &m.UnitCode, &m.UnitName, &m.TransportType, &m.CargoCapacityTons, &m.FuelCapacityLiters, &m.OperatingUnitId, &m.OperatingUnitName, &m.Status, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated transport_unit records
func (r *TransportUnitRepository) List(ctx context.Context, opts model.ListOptions) ([]model.TransportUnit, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "t.deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(unit_code ILIKE $%[1]d OR unit_name ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM log_transport_units t WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 {
		limit = 10
	}
	offset := opts.Offset
	if offset < 0 {
		offset = 0
	}

	listQuery := fmt.Sprintf(`SELECT t.transport_unit_id, t.unit_code, t.unit_name, t.transport_type, t.cargo_capacity_tons, t.fuel_capacity_liters, t.operating_unit_id, COALESCE(j_unit.unit_name, ''), t.status, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM log_transport_units t
	LEFT JOIN org_units j_unit ON j_unit.unit_id = t.operating_unit_id
	WHERE %s ORDER BY t.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.TransportUnit
	for rows.Next() {
		var m model.TransportUnit
		if err := rows.Scan(&m.TransportUnitId, &m.UnitCode, &m.UnitName, &m.TransportType, &m.CargoCapacityTons, &m.FuelCapacityLiters, &m.OperatingUnitId, &m.OperatingUnitName, &m.Status, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new transport_unit with optional child items
func (r *TransportUnitRepository) Create(ctx context.Context, m *model.TransportUnit) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO log_transport_units (unit_code, unit_name, transport_type, cargo_capacity_tons, fuel_capacity_liters, operating_unit_id, status, created_by, updated_by, deleted_by) VALUES ($1, $2, $3::transport_unit_type, $4, $5, $6, $7::transport_status_type, $8, $9, $10) RETURNING transport_unit_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.UnitCode, m.UnitName, m.TransportType, m.CargoCapacityTons, m.FuelCapacityLiters, m.OperatingUnitId, m.Status, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing transport_unit
func (r *TransportUnitRepository) Update(ctx context.Context, id string, m *model.TransportUnit) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE log_transport_units SET unit_code = $1, unit_name = $2, transport_type = $3::transport_unit_type, cargo_capacity_tons = $4, fuel_capacity_liters = $5, operating_unit_id = $6, status = $7::transport_status_type, updated_by = $8, updated_at = CURRENT_TIMESTAMP WHERE transport_unit_id = $9`
	_, err = tx.Exec(ctx, updateQuery, m.UnitCode, m.UnitName, m.TransportType, m.CargoCapacityTons, m.FuelCapacityLiters, m.OperatingUnitId, m.Status, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes transport_unit
func (r *TransportUnitRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE log_transport_units SET deleted_at = CURRENT_TIMESTAMP WHERE transport_unit_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
