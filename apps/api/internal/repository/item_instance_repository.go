package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ItemInstanceRepository struct {
	DB *pgxpool.Pool
}

func NewItemInstanceRepository(db *pgxpool.Pool) *ItemInstanceRepository {
	return &ItemInstanceRepository{DB: db}
}

// Get retrieves a single item_instance by instance_id
func (r *ItemInstanceRepository) Get(ctx context.Context, id string) (*model.ItemInstance, error) {
	query := `SELECT t.instance_id, t.warehouse_id, COALESCE(j_wh.warehouse_name, ''), t.location_id, t.material_id, COALESCE(j_mat.material_name, ''), t.batch_number, t.serial_number, t.lot_number, t.expiry_date, t.manufactured_date, t.condition, t.inspection_due_date, t.quantity, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM inv_item_instances t
	LEFT JOIN inv_warehouses j_wh ON j_wh.warehouse_id = t.warehouse_id
	LEFT JOIN inv_materials j_mat ON j_mat.material_id = t.material_id
	WHERE t.instance_id = $1 AND t.deleted_at IS NULL`

	var m model.ItemInstance
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.InstanceId, &m.WarehouseId, &m.WarehouseName, &m.LocationId, &m.MaterialId, &m.MaterialName, &m.BatchNumber, &m.SerialNumber, &m.LotNumber, &m.ExpiryDate, &m.ManufacturedDate, &m.Condition, &m.InspectionDueDate, &m.Quantity, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated item_instance records
func (r *ItemInstanceRepository) List(ctx context.Context, opts model.ListOptions) ([]model.ItemInstance, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "t.deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(batch_number ILIKE $%[1]d OR serial_number ILIKE $%[1]d OR lot_number ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM inv_item_instances t WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf(`SELECT t.instance_id, t.warehouse_id, COALESCE(j_wh.warehouse_name, ''), t.location_id, t.material_id, COALESCE(j_mat.material_name, ''), t.batch_number, t.serial_number, t.lot_number, t.expiry_date, t.manufactured_date, t.condition, t.inspection_due_date, t.quantity, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM inv_item_instances t
	LEFT JOIN inv_warehouses j_wh ON j_wh.warehouse_id = t.warehouse_id
	LEFT JOIN inv_materials j_mat ON j_mat.material_id = t.material_id
	WHERE %s ORDER BY t.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.ItemInstance
	for rows.Next() {
		var m model.ItemInstance
		if err := rows.Scan(&m.InstanceId, &m.WarehouseId, &m.WarehouseName, &m.LocationId, &m.MaterialId, &m.MaterialName, &m.BatchNumber, &m.SerialNumber, &m.LotNumber, &m.ExpiryDate, &m.ManufacturedDate, &m.Condition, &m.InspectionDueDate, &m.Quantity, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new item_instance with optional child items
func (r *ItemInstanceRepository) Create(ctx context.Context, m *model.ItemInstance) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO inv_item_instances (warehouse_id, location_id, material_id, batch_number, serial_number, lot_number, expiry_date, manufactured_date, condition, inspection_due_date, quantity, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9::item_condition_type, $10, $11, $12, $13, $14) RETURNING instance_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.WarehouseId, m.LocationId, m.MaterialId, m.BatchNumber, m.SerialNumber, m.LotNumber, m.ExpiryDate, m.ManufacturedDate, m.Condition, m.InspectionDueDate, m.Quantity, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing item_instance
func (r *ItemInstanceRepository) Update(ctx context.Context, id string, m *model.ItemInstance) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE inv_item_instances SET warehouse_id = $1, location_id = $2, material_id = $3, batch_number = $4, serial_number = $5, lot_number = $6, expiry_date = $7, manufactured_date = $8, condition = $9::item_condition_type, inspection_due_date = $10, quantity = $11, updated_by = $12, updated_at = CURRENT_TIMESTAMP WHERE instance_id = $13`
	_, err = tx.Exec(ctx, updateQuery, m.WarehouseId, m.LocationId, m.MaterialId, m.BatchNumber, m.SerialNumber, m.LotNumber, m.ExpiryDate, m.ManufacturedDate, m.Condition, m.InspectionDueDate, m.Quantity, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes item_instance
func (r *ItemInstanceRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE inv_item_instances SET deleted_at = CURRENT_TIMESTAMP WHERE instance_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
