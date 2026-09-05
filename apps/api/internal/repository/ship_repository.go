package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ShipRepository struct {
	DB *pgxpool.Pool
}

func NewShipRepository(db *pgxpool.Pool) *ShipRepository {
	return &ShipRepository{DB: db}
}

// Get retrieves a single ship by ship_id
func (r *ShipRepository) Get(ctx context.Context, id string) (*model.Ship, error) {
	query := `SELECT s.ship_id, s.class_id, COALESCE(c.class_name, ''), s.assigned_unit_id, COALESCE(u.unit_name, ''), s.hull_number, s.ship_name, s.call_sign, s.commission_date, s.home_port, s.length_m, s.beam_m, s.draft_m, s.displacement_tons, s.max_speed_knots, s.cruise_range_nm, s.crew_capacity, s.fuel_capacity_liters, s.fresh_water_capacity_liters, s.status, s.current_readiness_status, s.created_by, s.updated_by, s.deleted_by, s.created_at, s.updated_at, s.deleted_at
	FROM mro_ships s
	LEFT JOIN mro_ship_classes c ON c.class_id = s.class_id
	LEFT JOIN org_units u ON u.unit_id = s.assigned_unit_id
	WHERE s.ship_id = $1 AND s.deleted_at IS NULL`

	var m model.Ship
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.ShipId, &m.ClassId, &m.ClassName, &m.AssignedUnitId, &m.UnitName, &m.HullNumber, &m.ShipName, &m.CallSign, &m.CommissionDate, &m.HomePort, &m.LengthM, &m.BeamM, &m.DraftM, &m.DisplacementTons, &m.MaxSpeedKnots, &m.CruiseRangeNm, &m.CrewCapacity, &m.FuelCapacityLiters, &m.FreshWaterCapacityLiters, &m.Status, &m.CurrentReadinessStatus, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated ship records
func (r *ShipRepository) List(ctx context.Context, opts model.ListOptions) ([]model.Ship, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "s.deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(s.hull_number ILIKE $%[1]d OR s.ship_name ILIKE $%[1]d OR s.call_sign ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM mro_ships s WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf(`SELECT s.ship_id, s.class_id, COALESCE(c.class_name, ''), s.assigned_unit_id, COALESCE(u.unit_name, ''), s.hull_number, s.ship_name, s.call_sign, s.commission_date, s.home_port, s.length_m, s.beam_m, s.draft_m, s.displacement_tons, s.max_speed_knots, s.cruise_range_nm, s.crew_capacity, s.fuel_capacity_liters, s.fresh_water_capacity_liters, s.status, s.current_readiness_status, s.created_by, s.updated_by, s.deleted_by, s.created_at, s.updated_at, s.deleted_at
	FROM mro_ships s
	LEFT JOIN mro_ship_classes c ON c.class_id = s.class_id
	LEFT JOIN org_units u ON u.unit_id = s.assigned_unit_id
	WHERE %s ORDER BY s.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.Ship
	for rows.Next() {
		var m model.Ship
		if err := rows.Scan(&m.ShipId, &m.ClassId, &m.ClassName, &m.AssignedUnitId, &m.UnitName, &m.HullNumber, &m.ShipName, &m.CallSign, &m.CommissionDate, &m.HomePort, &m.LengthM, &m.BeamM, &m.DraftM, &m.DisplacementTons, &m.MaxSpeedKnots, &m.CruiseRangeNm, &m.CrewCapacity, &m.FuelCapacityLiters, &m.FreshWaterCapacityLiters, &m.Status, &m.CurrentReadinessStatus, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new ship with optional child items
func (r *ShipRepository) Create(ctx context.Context, m *model.Ship) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO mro_ships (class_id, assigned_unit_id, hull_number, ship_name, call_sign, commission_date, home_port, length_m, beam_m, draft_m, displacement_tons, max_speed_knots, cruise_range_nm, crew_capacity, fuel_capacity_liters, fresh_water_capacity_liters, status, current_readiness_status, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17::ship_status_type, $18::ship_readiness_status_type, $19, $20, $21) RETURNING ship_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.ClassId, m.AssignedUnitId, m.HullNumber, m.ShipName, m.CallSign, m.CommissionDate, m.HomePort, m.LengthM, m.BeamM, m.DraftM, m.DisplacementTons, m.MaxSpeedKnots, m.CruiseRangeNm, m.CrewCapacity, m.FuelCapacityLiters, m.FreshWaterCapacityLiters, m.Status, m.CurrentReadinessStatus, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing ship
func (r *ShipRepository) Update(ctx context.Context, id string, m *model.Ship) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE mro_ships SET class_id = $1, assigned_unit_id = $2, hull_number = $3, ship_name = $4, call_sign = $5, commission_date = $6, home_port = $7, length_m = $8, beam_m = $9, draft_m = $10, displacement_tons = $11, max_speed_knots = $12, cruise_range_nm = $13, crew_capacity = $14, fuel_capacity_liters = $15, fresh_water_capacity_liters = $16, status = $17::ship_status_type, current_readiness_status = $18::ship_readiness_status_type, updated_by = $19, updated_at = CURRENT_TIMESTAMP WHERE ship_id = $20`
	_, err = tx.Exec(ctx, updateQuery, m.ClassId, m.AssignedUnitId, m.HullNumber, m.ShipName, m.CallSign, m.CommissionDate, m.HomePort, m.LengthM, m.BeamM, m.DraftM, m.DisplacementTons, m.MaxSpeedKnots, m.CruiseRangeNm, m.CrewCapacity, m.FuelCapacityLiters, m.FreshWaterCapacityLiters, m.Status, m.CurrentReadinessStatus, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes ship
func (r *ShipRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE mro_ships SET deleted_at = CURRENT_TIMESTAMP WHERE ship_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
