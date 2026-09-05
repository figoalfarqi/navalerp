package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FuelBunkerRepository struct {
	DB *pgxpool.Pool
}

func NewFuelBunkerRepository(db *pgxpool.Pool) *FuelBunkerRepository {
	return &FuelBunkerRepository{DB: db}
}

// Get retrieves a single fuel_bunker by bunker_id
func (r *FuelBunkerRepository) Get(ctx context.Context, id string) (*model.FuelBunker, error) {
	query := `SELECT bunker_id, ship_id, facility_id, fuel_type, quantity_liters, density_15c, flow_rate_lph, bunkering_start_time, bunkering_end_time, receipt_voucher_no, authorised_by_user_id, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM infra_fuel_bunker_records WHERE bunker_id = $1 AND deleted_at IS NULL`

	var m model.FuelBunker
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.BunkerId, &m.ShipId, &m.FacilityId, &m.FuelType, &m.QuantityLiters, &m.Density15c, &m.FlowRateLph, &m.BunkeringStartTime, &m.BunkeringEndTime, &m.ReceiptVoucherNo, &m.AuthorisedByUserId, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated fuel_bunker records
func (r *FuelBunkerRepository) List(ctx context.Context, opts model.ListOptions) ([]model.FuelBunker, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(receipt_voucher_no ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM infra_fuel_bunker_records WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT bunker_id, ship_id, facility_id, fuel_type, quantity_liters, density_15c, flow_rate_lph, bunkering_start_time, bunkering_end_time, receipt_voucher_no, authorised_by_user_id, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM infra_fuel_bunker_records WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.FuelBunker
	for rows.Next() {
		var m model.FuelBunker
		if err := rows.Scan(&m.BunkerId, &m.ShipId, &m.FacilityId, &m.FuelType, &m.QuantityLiters, &m.Density15c, &m.FlowRateLph, &m.BunkeringStartTime, &m.BunkeringEndTime, &m.ReceiptVoucherNo, &m.AuthorisedByUserId, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new fuel_bunker with optional child items
func (r *FuelBunkerRepository) Create(ctx context.Context, m *model.FuelBunker) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO infra_fuel_bunker_records (ship_id, facility_id, fuel_type, quantity_liters, density_15c, flow_rate_lph, bunkering_start_time, bunkering_end_time, receipt_voucher_no, authorised_by_user_id, created_by, updated_by, deleted_by) VALUES ($1, $2, $3::fuel_bunker_type, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING bunker_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.ShipId, m.FacilityId, m.FuelType, m.QuantityLiters, m.Density15c, m.FlowRateLph, m.BunkeringStartTime, m.BunkeringEndTime, m.ReceiptVoucherNo, m.AuthorisedByUserId, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing fuel_bunker
func (r *FuelBunkerRepository) Update(ctx context.Context, id string, m *model.FuelBunker) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE infra_fuel_bunker_records SET ship_id = $1, facility_id = $2, fuel_type = $3::fuel_bunker_type, quantity_liters = $4, density_15c = $5, flow_rate_lph = $6, bunkering_start_time = $7, bunkering_end_time = $8, receipt_voucher_no = $9, authorised_by_user_id = $10, updated_by = $11, updated_at = CURRENT_TIMESTAMP WHERE bunker_id = $12`
	_, err = tx.Exec(ctx, updateQuery, m.ShipId, m.FacilityId, m.FuelType, m.QuantityLiters, m.Density15c, m.FlowRateLph, m.BunkeringStartTime, m.BunkeringEndTime, m.ReceiptVoucherNo, m.AuthorisedByUserId, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes fuel_bunker
func (r *FuelBunkerRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE infra_fuel_bunker_records SET deleted_at = CURRENT_TIMESTAMP WHERE bunker_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
