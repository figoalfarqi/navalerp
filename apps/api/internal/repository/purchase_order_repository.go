package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PurchaseOrderRepository struct {
	DB *pgxpool.Pool
}

func NewPurchaseOrderRepository(db *pgxpool.Pool) *PurchaseOrderRepository {
	return &PurchaseOrderRepository{DB: db}
}

// Get retrieves a single purchase_order by po_id
func (r *PurchaseOrderRepository) Get(ctx context.Context, id string) (*model.PurchaseOrder, error) {
	query := `SELECT po_id, po_number, contract_id, vendor_id, issuing_unit_id, order_date, delivery_deadline, destination_warehouse_id, total_amount, tax_amount, grand_total, status, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM proc_purchase_orders WHERE po_id = $1 AND deleted_at IS NULL`

	var m model.PurchaseOrder
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.PoId, &m.PoNumber, &m.ContractId, &m.VendorId, &m.IssuingUnitId, &m.OrderDate, &m.DeliveryDeadline, &m.DestinationWarehouseId, &m.TotalAmount, &m.TaxAmount, &m.GrandTotal, &m.Status, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	// Load Items
	childRowsItems, err := r.DB.Query(ctx, `SELECT po_item_id, po_id, material_id, quantity, unit_price, total_price, notes, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM proc_purchase_order_items WHERE po_id = $1`, id)
	if err == nil {
		defer childRowsItems.Close()
		for childRowsItems.Next() {
			var item model.PurchaseOrderItems
			if err := childRowsItems.Scan(&item.PoItemId, &item.PoId, &item.MaterialId, &item.Quantity, &item.UnitPrice, &item.TotalPrice, &item.Notes, &item.CreatedBy, &item.UpdatedBy, &item.DeletedBy, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt); err == nil {
				m.Items = append(m.Items, item)
			}
		}
	}

	return &m, nil
}

// List retrieves paginated purchase_order records
func (r *PurchaseOrderRepository) List(ctx context.Context, opts model.ListOptions) ([]model.PurchaseOrder, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(po_number ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM proc_purchase_orders WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT po_id, po_number, contract_id, vendor_id, issuing_unit_id, order_date, delivery_deadline, destination_warehouse_id, total_amount, tax_amount, grand_total, status, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM proc_purchase_orders WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.PurchaseOrder
	for rows.Next() {
		var m model.PurchaseOrder
		if err := rows.Scan(&m.PoId, &m.PoNumber, &m.ContractId, &m.VendorId, &m.IssuingUnitId, &m.OrderDate, &m.DeliveryDeadline, &m.DestinationWarehouseId, &m.TotalAmount, &m.TaxAmount, &m.GrandTotal, &m.Status, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new purchase_order with optional child items
func (r *PurchaseOrderRepository) Create(ctx context.Context, m *model.PurchaseOrder) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO proc_purchase_orders (po_number, contract_id, vendor_id, issuing_unit_id, order_date, delivery_deadline, destination_warehouse_id, total_amount, tax_amount, status, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10::po_status_type, $11, $12, $13) RETURNING po_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.PoNumber, m.ContractId, m.VendorId, m.IssuingUnitId, m.OrderDate, m.DeliveryDeadline, m.DestinationWarehouseId, m.TotalAmount, m.TaxAmount, m.Status, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.Items {
		_, err := tx.Exec(ctx, `INSERT INTO proc_purchase_order_items (po_id, material_id, quantity, unit_price, notes, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, newID, item.MaterialId, item.Quantity, item.UnitPrice, item.Notes, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing purchase_order
func (r *PurchaseOrderRepository) Update(ctx context.Context, id string, m *model.PurchaseOrder) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE proc_purchase_orders SET po_number = $1, contract_id = $2, vendor_id = $3, issuing_unit_id = $4, order_date = $5, delivery_deadline = $6, destination_warehouse_id = $7, total_amount = $8, tax_amount = $9, status = $10::po_status_type, updated_by = $11, updated_at = CURRENT_TIMESTAMP WHERE po_id = $12`
	_, err = tx.Exec(ctx, updateQuery, m.PoNumber, m.ContractId, m.VendorId, m.IssuingUnitId, m.OrderDate, m.DeliveryDeadline, m.DestinationWarehouseId, m.TotalAmount, m.TaxAmount, m.Status, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	if len(m.Items) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM proc_purchase_order_items WHERE po_id = $1`, id)
		for _, item := range m.Items {
			_, err := tx.Exec(ctx, `INSERT INTO proc_purchase_order_items (po_id, material_id, quantity, unit_price, notes, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, id, item.MaterialId, item.Quantity, item.UnitPrice, item.Notes, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes purchase_order
func (r *PurchaseOrderRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE proc_purchase_orders SET deleted_at = CURRENT_TIMESTAMP WHERE po_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
