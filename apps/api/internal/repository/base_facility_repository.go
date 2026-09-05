package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BaseFacilityRepository struct {
	DB *pgxpool.Pool
}

func NewBaseFacilityRepository(db *pgxpool.Pool) *BaseFacilityRepository {
	return &BaseFacilityRepository{DB: db}
}

// Get retrieves a single base_facility by facility_id
func (r *BaseFacilityRepository) Get(ctx context.Context, id string) (*model.BaseFacility, error) {
	query := `SELECT t.facility_id, t.base_unit_id, COALESCE(j_unit.unit_name, ''), t.facility_code, t.facility_name, t.facility_type, t.length_meters, t.draft_depth_meters, t.max_displacement_tonnage, t.has_shore_power, t.has_fresh_water, t.has_fuel_bunker_line, t.status, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM infra_facilities t
	LEFT JOIN org_units j_unit ON j_unit.unit_id = t.base_unit_id
	WHERE t.facility_id = $1 AND t.deleted_at IS NULL`

	var m model.BaseFacility
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.FacilityId, &m.BaseUnitId, &m.UnitName, &m.FacilityCode, &m.FacilityName, &m.FacilityType, &m.LengthMeters, &m.DraftDepthMeters, &m.MaxDisplacementTonnage, &m.HasShorePower, &m.HasFreshWater, &m.HasFuelBunkerLine, &m.Status, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	// Load Maintenances
	childRowsMaintenances, err := r.DB.Query(ctx, `SELECT maint_id, facility_id, maintenance_type, start_date, end_date, cost, performed_by, status, remarks, created_by, created_at FROM infra_facility_maintenances WHERE facility_id = $1`, id)
	if err == nil {
		defer childRowsMaintenances.Close()
		for childRowsMaintenances.Next() {
			var item model.FacilityMaintenances
			if err := childRowsMaintenances.Scan(&item.MaintId, &item.FacilityId, &item.MaintenanceType, &item.StartDate, &item.EndDate, &item.Cost, &item.PerformedBy, &item.Status, &item.Remarks, &item.CreatedBy, &item.CreatedAt); err == nil {
				m.Maintenances = append(m.Maintenances, item)
			}
		}
	}

	return &m, nil
}

// List retrieves paginated base_facility records
func (r *BaseFacilityRepository) List(ctx context.Context, opts model.ListOptions) ([]model.BaseFacility, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "t.deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(facility_code ILIKE $%[1]d OR facility_name ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM infra_facilities t WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf(`SELECT t.facility_id, t.base_unit_id, COALESCE(j_unit.unit_name, ''), t.facility_code, t.facility_name, t.facility_type, t.length_meters, t.draft_depth_meters, t.max_displacement_tonnage, t.has_shore_power, t.has_fresh_water, t.has_fuel_bunker_line, t.status, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM infra_facilities t
	LEFT JOIN org_units j_unit ON j_unit.unit_id = t.base_unit_id
	WHERE %s ORDER BY t.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.BaseFacility
	for rows.Next() {
		var m model.BaseFacility
		if err := rows.Scan(&m.FacilityId, &m.BaseUnitId, &m.UnitName, &m.FacilityCode, &m.FacilityName, &m.FacilityType, &m.LengthMeters, &m.DraftDepthMeters, &m.MaxDisplacementTonnage, &m.HasShorePower, &m.HasFreshWater, &m.HasFuelBunkerLine, &m.Status, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new base_facility with optional child items
func (r *BaseFacilityRepository) Create(ctx context.Context, m *model.BaseFacility) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO infra_facilities (base_unit_id, facility_code, facility_name, facility_type, length_meters, draft_depth_meters, max_displacement_tonnage, has_shore_power, has_fresh_water, has_fuel_bunker_line, status, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4::facility_type_enum, $5, $6, $7, $8, $9, $10, $11::facility_status_type, $12, $13, $14) RETURNING facility_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.BaseUnitId, m.FacilityCode, m.FacilityName, m.FacilityType, m.LengthMeters, m.DraftDepthMeters, m.MaxDisplacementTonnage, m.HasShorePower, m.HasFreshWater, m.HasFuelBunkerLine, m.Status, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.Maintenances {
		_, err := tx.Exec(ctx, `INSERT INTO infra_facility_maintenances (facility_id, maintenance_type, start_date, end_date, cost, performed_by, status, remarks, created_by) VALUES ($1, $2::facility_maint_type, $3, $4, $5, $6, $7::facility_maint_status_type, $8, $9)`, newID, item.MaintenanceType, item.StartDate, item.EndDate, item.Cost, item.PerformedBy, item.Status, item.Remarks, item.CreatedBy)
		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing base_facility
func (r *BaseFacilityRepository) Update(ctx context.Context, id string, m *model.BaseFacility) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE infra_facilities SET base_unit_id = $1, facility_code = $2, facility_name = $3, facility_type = $4::facility_type_enum, length_meters = $5, draft_depth_meters = $6, max_displacement_tonnage = $7, has_shore_power = $8, has_fresh_water = $9, has_fuel_bunker_line = $10, status = $11::facility_status_type, updated_by = $12, updated_at = CURRENT_TIMESTAMP WHERE facility_id = $13`
	_, err = tx.Exec(ctx, updateQuery, m.BaseUnitId, m.FacilityCode, m.FacilityName, m.FacilityType, m.LengthMeters, m.DraftDepthMeters, m.MaxDisplacementTonnage, m.HasShorePower, m.HasFreshWater, m.HasFuelBunkerLine, m.Status, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	if len(m.Maintenances) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM infra_facility_maintenances WHERE facility_id = $1`, id)
		for _, item := range m.Maintenances {
			_, err := tx.Exec(ctx, `INSERT INTO infra_facility_maintenances (facility_id, maintenance_type, start_date, end_date, cost, performed_by, status, remarks, created_by) VALUES ($1, $2::facility_maint_type, $3, $4, $5, $6, $7::facility_maint_status_type, $8, $9)`, id, item.MaintenanceType, item.StartDate, item.EndDate, item.Cost, item.PerformedBy, item.Status, item.Remarks, item.CreatedBy)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes base_facility
func (r *BaseFacilityRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE infra_facilities SET deleted_at = CURRENT_TIMESTAMP WHERE facility_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
