package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StockAdjustmentRepository struct {
	DB *pgxpool.Pool
}

func NewStockAdjustmentRepository(db *pgxpool.Pool) *StockAdjustmentRepository {
	return &StockAdjustmentRepository{DB: db}
}

// Get retrieves a single stock_adjustment by adjustment_id
func (r *StockAdjustmentRepository) Get(ctx context.Context, id string) (*model.StockAdjustment, error) {
	query := `SELECT adjustment_id, warehouse_id, adjustment_number, adjustment_date, conducted_by_user_id, reason, status, remarks, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM inv_stock_adjustments WHERE adjustment_id = $1 AND deleted_at IS NULL`

	var m model.StockAdjustment
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.AdjustmentId, &m.WarehouseId, &m.AdjustmentNumber, &m.AdjustmentDate, &m.ConductedByUserId, &m.Reason, &m.Status, &m.Remarks, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	// Load Items
	childRowsItems, err := r.DB.Query(ctx, `SELECT adj_item_id, adjustment_id, material_id, book_quantity, physical_quantity, difference_quantity, unit_cost, total_adjustment_value, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM inv_stock_adjustment_items WHERE adjustment_id = $1`, id)
	if err == nil {
		defer childRowsItems.Close()
		for childRowsItems.Next() {
			var item model.StockAdjustmentItems
			if err := childRowsItems.Scan(&item.AdjItemId, &item.AdjustmentId, &item.MaterialId, &item.BookQuantity, &item.PhysicalQuantity, &item.DifferenceQuantity, &item.UnitCost, &item.TotalAdjustmentValue, &item.CreatedBy, &item.UpdatedBy, &item.DeletedBy, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt); err == nil {
				m.Items = append(m.Items, item)
			}
		}
	}

	return &m, nil
}

// List retrieves paginated stock_adjustment records
func (r *StockAdjustmentRepository) List(ctx context.Context, opts model.ListOptions) ([]model.StockAdjustment, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(adjustment_number ILIKE $%[1]d OR remarks ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM inv_stock_adjustments WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT adjustment_id, warehouse_id, adjustment_number, adjustment_date, conducted_by_user_id, reason, status, remarks, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM inv_stock_adjustments WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.StockAdjustment
	for rows.Next() {
		var m model.StockAdjustment
		if err := rows.Scan(&m.AdjustmentId, &m.WarehouseId, &m.AdjustmentNumber, &m.AdjustmentDate, &m.ConductedByUserId, &m.Reason, &m.Status, &m.Remarks, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new stock_adjustment with optional child items
func (r *StockAdjustmentRepository) Create(ctx context.Context, m *model.StockAdjustment) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO inv_stock_adjustments (warehouse_id, adjustment_number, adjustment_date, conducted_by_user_id, reason, status, remarks, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5::adjustment_reason_type, $6::adjustment_status_type, $7, $8, $9, $10) RETURNING adjustment_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.WarehouseId, m.AdjustmentNumber, m.AdjustmentDate, m.ConductedByUserId, m.Reason, m.Status, m.Remarks, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.Items {
		_, err := tx.Exec(ctx, `INSERT INTO inv_stock_adjustment_items (adjustment_id, material_id, book_quantity, physical_quantity, unit_cost, total_adjustment_value, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`, newID, item.MaterialId, item.BookQuantity, item.PhysicalQuantity, item.UnitCost, item.TotalAdjustmentValue, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing stock_adjustment
func (r *StockAdjustmentRepository) Update(ctx context.Context, id string, m *model.StockAdjustment) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE inv_stock_adjustments SET warehouse_id = $1, adjustment_number = $2, adjustment_date = $3, conducted_by_user_id = $4, reason = $5::adjustment_reason_type, status = $6::adjustment_status_type, remarks = $7, updated_by = $8, updated_at = CURRENT_TIMESTAMP WHERE adjustment_id = $9`
	_, err = tx.Exec(ctx, updateQuery, m.WarehouseId, m.AdjustmentNumber, m.AdjustmentDate, m.ConductedByUserId, m.Reason, m.Status, m.Remarks, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	if len(m.Items) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM inv_stock_adjustment_items WHERE adjustment_id = $1`, id)
		for _, item := range m.Items {
			_, err := tx.Exec(ctx, `INSERT INTO inv_stock_adjustment_items (adjustment_id, material_id, book_quantity, physical_quantity, unit_cost, total_adjustment_value, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`, id, item.MaterialId, item.BookQuantity, item.PhysicalQuantity, item.UnitCost, item.TotalAdjustmentValue, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes stock_adjustment
func (r *StockAdjustmentRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE inv_stock_adjustments SET deleted_at = CURRENT_TIMESTAMP WHERE adjustment_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
