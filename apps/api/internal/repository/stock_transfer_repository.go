package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StockTransferRepository struct {
	DB *pgxpool.Pool
}

func NewStockTransferRepository(db *pgxpool.Pool) *StockTransferRepository {
	return &StockTransferRepository{DB: db}
}

// Get retrieves a single stock_transfer by transfer_id
func (r *StockTransferRepository) Get(ctx context.Context, id string) (*model.StockTransfer, error) {
	query := `SELECT transfer_id, transfer_number, from_warehouse_id, to_warehouse_id, movement_type, scheduled_departure, actual_departure, scheduled_arrival, actual_arrival, transporter_unit, status, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM inv_stock_transfers WHERE transfer_id = $1 AND deleted_at IS NULL`

	var m model.StockTransfer
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.TransferId, &m.TransferNumber, &m.FromWarehouseId, &m.ToWarehouseId, &m.MovementType, &m.ScheduledDeparture, &m.ActualDeparture, &m.ScheduledArrival, &m.ActualArrival, &m.TransporterUnit, &m.Status, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	// Load Items
	childRowsItems, err := r.DB.Query(ctx, `SELECT transfer_item_id, transfer_id, material_id, quantity_shipped, quantity_received, condition_on_receipt, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM inv_stock_transfer_items WHERE transfer_id = $1`, id)
	if err == nil {
		defer childRowsItems.Close()
		for childRowsItems.Next() {
			var item model.StockTransferItems
			if err := childRowsItems.Scan(&item.TransferItemId, &item.TransferId, &item.MaterialId, &item.QuantityShipped, &item.QuantityReceived, &item.ConditionOnReceipt, &item.CreatedBy, &item.UpdatedBy, &item.DeletedBy, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt); err == nil {
				m.Items = append(m.Items, item)
			}
		}
	}

	return &m, nil
}

// List retrieves paginated stock_transfer records
func (r *StockTransferRepository) List(ctx context.Context, opts model.ListOptions) ([]model.StockTransfer, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(transfer_number ILIKE $%[1]d OR transporter_unit ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM inv_stock_transfers WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT transfer_id, transfer_number, from_warehouse_id, to_warehouse_id, movement_type, scheduled_departure, actual_departure, scheduled_arrival, actual_arrival, transporter_unit, status, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM inv_stock_transfers WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.StockTransfer
	for rows.Next() {
		var m model.StockTransfer
		if err := rows.Scan(&m.TransferId, &m.TransferNumber, &m.FromWarehouseId, &m.ToWarehouseId, &m.MovementType, &m.ScheduledDeparture, &m.ActualDeparture, &m.ScheduledArrival, &m.ActualArrival, &m.TransporterUnit, &m.Status, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new stock_transfer with optional child items
func (r *StockTransferRepository) Create(ctx context.Context, m *model.StockTransfer) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO inv_stock_transfers (transfer_number, from_warehouse_id, to_warehouse_id, movement_type, scheduled_departure, actual_departure, scheduled_arrival, actual_arrival, transporter_unit, status, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4::stock_movement_type, $5, $6, $7, $8, $9, $10::transfer_status_type, $11, $12, $13) RETURNING transfer_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.TransferNumber, m.FromWarehouseId, m.ToWarehouseId, m.MovementType, m.ScheduledDeparture, m.ActualDeparture, m.ScheduledArrival, m.ActualArrival, m.TransporterUnit, m.Status, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.Items {
		_, err := tx.Exec(ctx, `INSERT INTO inv_stock_transfer_items (transfer_id, material_id, quantity_shipped, quantity_received, condition_on_receipt, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5::item_condition_type, $6, $7, $8)`, newID, item.MaterialId, item.QuantityShipped, item.QuantityReceived, item.ConditionOnReceipt, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing stock_transfer
func (r *StockTransferRepository) Update(ctx context.Context, id string, m *model.StockTransfer) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE inv_stock_transfers SET transfer_number = $1, from_warehouse_id = $2, to_warehouse_id = $3, movement_type = $4::stock_movement_type, scheduled_departure = $5, actual_departure = $6, scheduled_arrival = $7, actual_arrival = $8, transporter_unit = $9, status = $10::transfer_status_type, updated_by = $11, updated_at = CURRENT_TIMESTAMP WHERE transfer_id = $12`
	_, err = tx.Exec(ctx, updateQuery, m.TransferNumber, m.FromWarehouseId, m.ToWarehouseId, m.MovementType, m.ScheduledDeparture, m.ActualDeparture, m.ScheduledArrival, m.ActualArrival, m.TransporterUnit, m.Status, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	if len(m.Items) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM inv_stock_transfer_items WHERE transfer_id = $1`, id)
		for _, item := range m.Items {
			_, err := tx.Exec(ctx, `INSERT INTO inv_stock_transfer_items (transfer_id, material_id, quantity_shipped, quantity_received, condition_on_receipt, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5::item_condition_type, $6, $7, $8)`, id, item.MaterialId, item.QuantityShipped, item.QuantityReceived, item.ConditionOnReceipt, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes stock_transfer
func (r *StockTransferRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE inv_stock_transfers SET deleted_at = CURRENT_TIMESTAMP WHERE transfer_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
