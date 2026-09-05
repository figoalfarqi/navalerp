package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GoodsReceiptRepository struct {
	DB *pgxpool.Pool
}

func NewGoodsReceiptRepository(db *pgxpool.Pool) *GoodsReceiptRepository {
	return &GoodsReceiptRepository{DB: db}
}

// Get retrieves a single goods_receipt by receipt_id
func (r *GoodsReceiptRepository) Get(ctx context.Context, id string) (*model.GoodsReceipt, error) {
	query := `SELECT t.receipt_id, t.receipt_number, t.po_id, COALESCE(j_po.po_number, ''), t.warehouse_id, COALESCE(j_wh.warehouse_name, ''), t.received_date, t.delivery_order_number, t.inspected_by_user_id, COALESCE(j_usr.full_name, ''), t.inspection_passed, t.remarks, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM proc_goods_receipts t
	LEFT JOIN proc_purchase_orders j_po ON j_po.po_id = t.po_id
	LEFT JOIN inv_warehouses j_wh ON j_wh.warehouse_id = t.warehouse_id
	LEFT JOIN sys_users j_usr ON j_usr.user_id = t.inspected_by_user_id
	WHERE t.receipt_id = $1 AND t.deleted_at IS NULL`

	var m model.GoodsReceipt
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.ReceiptId, &m.ReceiptNumber, &m.PoId, &m.PoNumber, &m.WarehouseId, &m.WarehouseName, &m.ReceivedDate, &m.DeliveryOrderNumber, &m.InspectedByUserId, &m.InspectorName, &m.InspectionPassed, &m.Remarks, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	// Load Items
	childRowsItems, err := r.DB.Query(ctx, `SELECT receipt_item_id, receipt_id, po_item_id, material_id, quantity_received, quantity_accepted, quantity_rejected, rejection_reason, created_at FROM proc_goods_receipt_items WHERE receipt_id = $1`, id)
	if err == nil {
		defer childRowsItems.Close()
		for childRowsItems.Next() {
			var item model.GoodsReceiptItems
			if err := childRowsItems.Scan(&item.ReceiptItemId, &item.ReceiptId, &item.PoItemId, &item.MaterialId, &item.QuantityReceived, &item.QuantityAccepted, &item.QuantityRejected, &item.RejectionReason, &item.CreatedAt); err == nil {
				m.Items = append(m.Items, item)
			}
		}
	}

	return &m, nil
}

// List retrieves paginated goods_receipt records
func (r *GoodsReceiptRepository) List(ctx context.Context, opts model.ListOptions) ([]model.GoodsReceipt, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "t.deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(receipt_number ILIKE $%[1]d OR delivery_order_number ILIKE $%[1]d OR remarks ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM proc_goods_receipts t WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf(`SELECT t.receipt_id, t.receipt_number, t.po_id, COALESCE(j_po.po_number, ''), t.warehouse_id, COALESCE(j_wh.warehouse_name, ''), t.received_date, t.delivery_order_number, t.inspected_by_user_id, COALESCE(j_usr.full_name, ''), t.inspection_passed, t.remarks, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM proc_goods_receipts t
	LEFT JOIN proc_purchase_orders j_po ON j_po.po_id = t.po_id
	LEFT JOIN inv_warehouses j_wh ON j_wh.warehouse_id = t.warehouse_id
	LEFT JOIN sys_users j_usr ON j_usr.user_id = t.inspected_by_user_id
	WHERE %s ORDER BY t.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.GoodsReceipt
	for rows.Next() {
		var m model.GoodsReceipt
		if err := rows.Scan(&m.ReceiptId, &m.ReceiptNumber, &m.PoId, &m.PoNumber, &m.WarehouseId, &m.WarehouseName, &m.ReceivedDate, &m.DeliveryOrderNumber, &m.InspectedByUserId, &m.InspectorName, &m.InspectionPassed, &m.Remarks, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new goods_receipt with optional child items
func (r *GoodsReceiptRepository) Create(ctx context.Context, m *model.GoodsReceipt) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO proc_goods_receipts (receipt_number, po_id, warehouse_id, received_date, delivery_order_number, inspected_by_user_id, inspection_passed, remarks, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING receipt_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.ReceiptNumber, m.PoId, m.WarehouseId, m.ReceivedDate, m.DeliveryOrderNumber, m.InspectedByUserId, m.InspectionPassed, m.Remarks, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.Items {
		_, err := tx.Exec(ctx, `INSERT INTO proc_goods_receipt_items (receipt_id, po_item_id, material_id, quantity_received, quantity_accepted, quantity_rejected, rejection_reason) VALUES ($1, $2, $3, $4, $5, $6, $7)`, newID, item.PoItemId, item.MaterialId, item.QuantityReceived, item.QuantityAccepted, item.QuantityRejected, item.RejectionReason)
		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing goods_receipt
func (r *GoodsReceiptRepository) Update(ctx context.Context, id string, m *model.GoodsReceipt) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE proc_goods_receipts SET receipt_number = $1, po_id = $2, warehouse_id = $3, received_date = $4, delivery_order_number = $5, inspected_by_user_id = $6, inspection_passed = $7, remarks = $8, updated_by = $9, updated_at = CURRENT_TIMESTAMP WHERE receipt_id = $10`
	_, err = tx.Exec(ctx, updateQuery, m.ReceiptNumber, m.PoId, m.WarehouseId, m.ReceivedDate, m.DeliveryOrderNumber, m.InspectedByUserId, m.InspectionPassed, m.Remarks, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	if len(m.Items) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM proc_goods_receipt_items WHERE receipt_id = $1`, id)
		for _, item := range m.Items {
			_, err := tx.Exec(ctx, `INSERT INTO proc_goods_receipt_items (receipt_id, po_item_id, material_id, quantity_received, quantity_accepted, quantity_rejected, rejection_reason) VALUES ($1, $2, $3, $4, $5, $6, $7)`, id, item.PoItemId, item.MaterialId, item.QuantityReceived, item.QuantityAccepted, item.QuantityRejected, item.RejectionReason)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes goods_receipt
func (r *GoodsReceiptRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE proc_goods_receipts SET deleted_at = CURRENT_TIMESTAMP WHERE receipt_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
