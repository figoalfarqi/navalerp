package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BerthBookingRepository struct {
	DB *pgxpool.Pool
}

func NewBerthBookingRepository(db *pgxpool.Pool) *BerthBookingRepository {
	return &BerthBookingRepository{DB: db}
}

// Get retrieves a single berth_booking by booking_id
func (r *BerthBookingRepository) Get(ctx context.Context, id string) (*model.BerthBooking, error) {
	query := `SELECT t.booking_id, t.facility_id, COALESCE(j_fac.facility_name, ''), t.ship_id, COALESCE(j_ship.ship_name, ''), t.booking_purpose, t.eta, t.etd, t.actual_berth_time, t.actual_unberth_time, t.shore_power_kwh_used, t.fresh_water_ton_used, t.status, t.approved_by_user_id, t.remarks, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM infra_berth_bookings t
	LEFT JOIN infra_facilities j_fac ON j_fac.facility_id = t.facility_id
	LEFT JOIN mro_ships j_ship ON j_ship.ship_id = t.ship_id
	WHERE t.booking_id = $1 AND t.deleted_at IS NULL`

	var m model.BerthBooking
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.BookingId, &m.FacilityId, &m.FacilityName, &m.ShipId, &m.ShipName, &m.BookingPurpose, &m.Eta, &m.Etd, &m.ActualBerthTime, &m.ActualUnberthTime, &m.ShorePowerKwhUsed, &m.FreshWaterTonUsed, &m.Status, &m.ApprovedByUserId, &m.Remarks, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated berth_booking records
func (r *BerthBookingRepository) List(ctx context.Context, opts model.ListOptions) ([]model.BerthBooking, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "t.deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(remarks ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM infra_berth_bookings t WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf(`SELECT t.booking_id, t.facility_id, COALESCE(j_fac.facility_name, ''), t.ship_id, COALESCE(j_ship.ship_name, ''), t.booking_purpose, t.eta, t.etd, t.actual_berth_time, t.actual_unberth_time, t.shore_power_kwh_used, t.fresh_water_ton_used, t.status, t.approved_by_user_id, t.remarks, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM infra_berth_bookings t
	LEFT JOIN infra_facilities j_fac ON j_fac.facility_id = t.facility_id
	LEFT JOIN mro_ships j_ship ON j_ship.ship_id = t.ship_id
	WHERE %s ORDER BY t.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.BerthBooking
	for rows.Next() {
		var m model.BerthBooking
		if err := rows.Scan(&m.BookingId, &m.FacilityId, &m.FacilityName, &m.ShipId, &m.ShipName, &m.BookingPurpose, &m.Eta, &m.Etd, &m.ActualBerthTime, &m.ActualUnberthTime, &m.ShorePowerKwhUsed, &m.FreshWaterTonUsed, &m.Status, &m.ApprovedByUserId, &m.Remarks, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new berth_booking with optional child items
func (r *BerthBookingRepository) Create(ctx context.Context, m *model.BerthBooking) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO infra_berth_bookings (facility_id, ship_id, booking_purpose, eta, etd, actual_berth_time, actual_unberth_time, shore_power_kwh_used, fresh_water_ton_used, status, approved_by_user_id, remarks, created_by, updated_by, deleted_by) VALUES ($1, $2, $3::berth_booking_purpose_type, $4, $5, $6, $7, $8, $9, $10::berth_booking_status_type, $11, $12, $13, $14, $15) RETURNING booking_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.FacilityId, m.ShipId, m.BookingPurpose, m.Eta, m.Etd, m.ActualBerthTime, m.ActualUnberthTime, m.ShorePowerKwhUsed, m.FreshWaterTonUsed, m.Status, m.ApprovedByUserId, m.Remarks, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing berth_booking
func (r *BerthBookingRepository) Update(ctx context.Context, id string, m *model.BerthBooking) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE infra_berth_bookings SET facility_id = $1, ship_id = $2, booking_purpose = $3::berth_booking_purpose_type, eta = $4, etd = $5, actual_berth_time = $6, actual_unberth_time = $7, shore_power_kwh_used = $8, fresh_water_ton_used = $9, status = $10::berth_booking_status_type, approved_by_user_id = $11, remarks = $12, updated_by = $13, updated_at = CURRENT_TIMESTAMP WHERE booking_id = $14`
	_, err = tx.Exec(ctx, updateQuery, m.FacilityId, m.ShipId, m.BookingPurpose, m.Eta, m.Etd, m.ActualBerthTime, m.ActualUnberthTime, m.ShorePowerKwhUsed, m.FreshWaterTonUsed, m.Status, m.ApprovedByUserId, m.Remarks, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes berth_booking
func (r *BerthBookingRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE infra_berth_bookings SET deleted_at = CURRENT_TIMESTAMP WHERE booking_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
