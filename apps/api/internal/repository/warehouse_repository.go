package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WarehouseRepository struct {
	DB *pgxpool.Pool
}

func NewWarehouseRepository(db *pgxpool.Pool) *WarehouseRepository {
	return &WarehouseRepository{DB: db}
}

// Get retrieves a single warehouse by warehouse_id
func (r *WarehouseRepository) Get(ctx context.Context, id string) (*model.Warehouse, error) {
	query := `SELECT w.warehouse_id, w.unit_id, COALESCE(u_unit.unit_name, ''), w.warehouse_code, w.warehouse_name, w.warehouse_type, w.capacity_m3, w.manager_user_id, COALESCE(u_mgr.full_name, ''), w.location_address, w.is_active, w.created_by, w.updated_by, w.deleted_by, w.created_at, w.updated_at, w.deleted_at
	FROM inv_warehouses w
	LEFT JOIN org_units u_unit ON u_unit.unit_id = w.unit_id
	LEFT JOIN sys_users u_mgr ON u_mgr.user_id = w.manager_user_id
	WHERE w.warehouse_id = $1 AND w.deleted_at IS NULL`

	var m model.Warehouse
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.WarehouseId, &m.UnitId, &m.UnitName, &m.WarehouseCode, &m.WarehouseName, &m.WarehouseType, &m.CapacityM3, &m.ManagerUserId, &m.ManagerName, &m.LocationAddress, &m.IsActive, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	// Load Locations
	childRowsLocations, err := r.DB.Query(ctx, `SELECT location_id, warehouse_id, zone_name, aisle, rack, shelf, bin_code, capacity_kg, is_hazardous_zone, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM inv_storage_locations WHERE warehouse_id = $1`, id)
	if err == nil {
		defer childRowsLocations.Close()
		for childRowsLocations.Next() {
			var item model.StorageLocations
			if err := childRowsLocations.Scan(&item.LocationId, &item.WarehouseId, &item.ZoneName, &item.Aisle, &item.Rack, &item.Shelf, &item.BinCode, &item.CapacityKg, &item.IsHazardousZone, &item.CreatedBy, &item.UpdatedBy, &item.DeletedBy, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt); err == nil {
				m.Locations = append(m.Locations, item)
			}
		}
	}

	return &m, nil
}

// List retrieves paginated warehouse records
func (r *WarehouseRepository) List(ctx context.Context, opts model.ListOptions) ([]model.Warehouse, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "w.deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(w.warehouse_code ILIKE $%[1]d OR w.warehouse_name ILIKE $%[1]d OR w.location_address ILIKE $%[1]d OR u_unit.unit_name ILIKE $%[1]d OR u_mgr.full_name ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf(`SELECT COUNT(*) FROM inv_warehouses w
	LEFT JOIN org_units u_unit ON u_unit.unit_id = w.unit_id
	LEFT JOIN sys_users u_mgr ON u_mgr.user_id = w.manager_user_id
	WHERE %s`, whereSql)
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

	listQuery := fmt.Sprintf(`SELECT w.warehouse_id, w.unit_id, COALESCE(u_unit.unit_name, ''), w.warehouse_code, w.warehouse_name, w.warehouse_type, w.capacity_m3, w.manager_user_id, COALESCE(u_mgr.full_name, ''), w.location_address, w.is_active, w.created_by, w.updated_by, w.deleted_by, w.created_at, w.updated_at, w.deleted_at
	FROM inv_warehouses w
	LEFT JOIN org_units u_unit ON u_unit.unit_id = w.unit_id
	LEFT JOIN sys_users u_mgr ON u_mgr.user_id = w.manager_user_id
	WHERE %s ORDER BY w.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.Warehouse
	for rows.Next() {
		var m model.Warehouse
		if err := rows.Scan(&m.WarehouseId, &m.UnitId, &m.UnitName, &m.WarehouseCode, &m.WarehouseName, &m.WarehouseType, &m.CapacityM3, &m.ManagerUserId, &m.ManagerName, &m.LocationAddress, &m.IsActive, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new warehouse with optional child items
func (r *WarehouseRepository) Create(ctx context.Context, m *model.Warehouse) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO inv_warehouses (unit_id, warehouse_code, warehouse_name, warehouse_type, capacity_m3, manager_user_id, location_address, is_active, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4::warehouse_type_enum, $5, $6, $7, $8, $9, $10, $11) RETURNING warehouse_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.UnitId, m.WarehouseCode, m.WarehouseName, m.WarehouseType, m.CapacityM3, m.ManagerUserId, m.LocationAddress, m.IsActive, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.Locations {
		_, err := tx.Exec(ctx, `INSERT INTO inv_storage_locations (warehouse_id, zone_name, aisle, rack, shelf, bin_code, capacity_kg, is_hazardous_zone, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, newID, item.ZoneName, item.Aisle, item.Rack, item.Shelf, item.BinCode, item.CapacityKg, item.IsHazardousZone, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing warehouse
func (r *WarehouseRepository) Update(ctx context.Context, id string, m *model.Warehouse) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE inv_warehouses SET unit_id = $1, warehouse_code = $2, warehouse_name = $3, warehouse_type = $4::warehouse_type_enum, capacity_m3 = $5, manager_user_id = $6, location_address = $7, is_active = $8, updated_by = $9, updated_at = CURRENT_TIMESTAMP WHERE warehouse_id = $10`
	_, err = tx.Exec(ctx, updateQuery, m.UnitId, m.WarehouseCode, m.WarehouseName, m.WarehouseType, m.CapacityM3, m.ManagerUserId, m.LocationAddress, m.IsActive, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	if len(m.Locations) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM inv_storage_locations WHERE warehouse_id = $1`, id)
		for _, item := range m.Locations {
			_, err := tx.Exec(ctx, `INSERT INTO inv_storage_locations (warehouse_id, zone_name, aisle, rack, shelf, bin_code, capacity_kg, is_hazardous_zone, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, id, item.ZoneName, item.Aisle, item.Rack, item.Shelf, item.BinCode, item.CapacityKg, item.IsHazardousZone, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes warehouse
func (r *WarehouseRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE inv_warehouses SET deleted_at = CURRENT_TIMESTAMP WHERE warehouse_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
