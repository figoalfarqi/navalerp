package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StockBalanceRepository struct {
	DB *pgxpool.Pool
}

func NewStockBalanceRepository(db *pgxpool.Pool) *StockBalanceRepository {
	return &StockBalanceRepository{DB: db}
}

// Get retrieves a single stock_balance by balance_id
func (r *StockBalanceRepository) Get(ctx context.Context, id string) (*model.StockBalance, error) {
	query := `SELECT t.balance_id, t.warehouse_id, COALESCE(j_wh.warehouse_name, ''), t.location_id, t.material_id, COALESCE(j_mat.material_name, ''), t.quantity_on_hand, t.quantity_reserved, t.quantity_in_transit, t.quantity_available, t.last_count_date, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM inv_stock_balances t
	LEFT JOIN inv_warehouses j_wh ON j_wh.warehouse_id = t.warehouse_id
	LEFT JOIN inv_materials j_mat ON j_mat.material_id = t.material_id
	WHERE t.balance_id = $1 AND t.deleted_at IS NULL`

	var m model.StockBalance
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.BalanceId, &m.WarehouseId, &m.WarehouseName, &m.LocationId, &m.MaterialId, &m.MaterialName, &m.QuantityOnHand, &m.QuantityReserved, &m.QuantityInTransit, &m.QuantityAvailable, &m.LastCountDate, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated stock_balance records
func (r *StockBalanceRepository) List(ctx context.Context, opts model.ListOptions) ([]model.StockBalance, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "t.deleted_at IS NULL")
	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM inv_stock_balances t WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf(`SELECT t.balance_id, t.warehouse_id, COALESCE(j_wh.warehouse_name, ''), t.location_id, t.material_id, COALESCE(j_mat.material_name, ''), t.quantity_on_hand, t.quantity_reserved, t.quantity_in_transit, t.quantity_available, t.last_count_date, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM inv_stock_balances t
	LEFT JOIN inv_warehouses j_wh ON j_wh.warehouse_id = t.warehouse_id
	LEFT JOIN inv_materials j_mat ON j_mat.material_id = t.material_id
	WHERE %s ORDER BY t.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.StockBalance
	for rows.Next() {
		var m model.StockBalance
		if err := rows.Scan(&m.BalanceId, &m.WarehouseId, &m.WarehouseName, &m.LocationId, &m.MaterialId, &m.MaterialName, &m.QuantityOnHand, &m.QuantityReserved, &m.QuantityInTransit, &m.QuantityAvailable, &m.LastCountDate, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new stock_balance with optional child items
func (r *StockBalanceRepository) Create(ctx context.Context, m *model.StockBalance) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO inv_stock_balances (warehouse_id, location_id, material_id, quantity_on_hand, quantity_reserved, quantity_in_transit, last_count_date, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING balance_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.WarehouseId, m.LocationId, m.MaterialId, m.QuantityOnHand, m.QuantityReserved, m.QuantityInTransit, m.LastCountDate, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing stock_balance
func (r *StockBalanceRepository) Update(ctx context.Context, id string, m *model.StockBalance) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE inv_stock_balances SET warehouse_id = $1, location_id = $2, material_id = $3, quantity_on_hand = $4, quantity_reserved = $5, quantity_in_transit = $6, last_count_date = $7, updated_by = $8, updated_at = CURRENT_TIMESTAMP WHERE balance_id = $9`
	_, err = tx.Exec(ctx, updateQuery, m.WarehouseId, m.LocationId, m.MaterialId, m.QuantityOnHand, m.QuantityReserved, m.QuantityInTransit, m.LastCountDate, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes stock_balance
func (r *StockBalanceRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE inv_stock_balances SET deleted_at = CURRENT_TIMESTAMP WHERE balance_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
