package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EquipmentRepository struct {
	DB *pgxpool.Pool
}

func NewEquipmentRepository(db *pgxpool.Pool) *EquipmentRepository {
	return &EquipmentRepository{DB: db}
}

// Get retrieves a single equipment by equipment_id
func (r *EquipmentRepository) Get(ctx context.Context, id string) (*model.Equipment, error) {
	query := `SELECT equipment_id, system_id, serial_number, equipment_tag, equipment_name, manufacturer, model_number, country_of_origin, installation_date, total_operating_hours, design_life_hours, criticality_level, health_status, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM mro_equipments WHERE equipment_id = $1 AND deleted_at IS NULL`

	var m model.Equipment
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.EquipmentId, &m.SystemId, &m.SerialNumber, &m.EquipmentTag, &m.EquipmentName, &m.Manufacturer, &m.ModelNumber, &m.CountryOfOrigin, &m.InstallationDate, &m.TotalOperatingHours, &m.DesignLifeHours, &m.CriticalityLevel, &m.HealthStatus, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	// Load Parameters
	childRowsParameters, err := r.DB.Query(ctx, `SELECT param_id, equipment_id, recorded_at, rpm, temperature_celsius, pressure_bar, vibration_level, oil_pressure_bar, running_hours_snapshot, status_flag, recorded_by_user_id FROM mro_equipment_parameters WHERE equipment_id = $1`, id)
	if err == nil {
		defer childRowsParameters.Close()
		for childRowsParameters.Next() {
			var item model.EquipmentParameters
			if err := childRowsParameters.Scan(&item.ParamId, &item.EquipmentId, &item.RecordedAt, &item.Rpm, &item.TemperatureCelsius, &item.PressureBar, &item.VibrationLevel, &item.OilPressureBar, &item.RunningHoursSnapshot, &item.StatusFlag, &item.RecordedByUserId); err == nil {
				m.Parameters = append(m.Parameters, item)
			}
		}
	}

	return &m, nil
}

// List retrieves paginated equipment records
func (r *EquipmentRepository) List(ctx context.Context, opts model.ListOptions) ([]model.Equipment, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(serial_number ILIKE $%[1]d OR equipment_tag ILIKE $%[1]d OR equipment_name ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM mro_equipments WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT equipment_id, system_id, serial_number, equipment_tag, equipment_name, manufacturer, model_number, country_of_origin, installation_date, total_operating_hours, design_life_hours, criticality_level, health_status, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM mro_equipments WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.Equipment
	for rows.Next() {
		var m model.Equipment
		if err := rows.Scan(&m.EquipmentId, &m.SystemId, &m.SerialNumber, &m.EquipmentTag, &m.EquipmentName, &m.Manufacturer, &m.ModelNumber, &m.CountryOfOrigin, &m.InstallationDate, &m.TotalOperatingHours, &m.DesignLifeHours, &m.CriticalityLevel, &m.HealthStatus, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new equipment with optional child items
func (r *EquipmentRepository) Create(ctx context.Context, m *model.Equipment) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO mro_equipments (system_id, serial_number, equipment_tag, equipment_name, manufacturer, model_number, country_of_origin, installation_date, total_operating_hours, design_life_hours, criticality_level, health_status, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11::equipment_criticality_type, $12::equipment_health_type, $13, $14, $15) RETURNING equipment_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.SystemId, m.SerialNumber, m.EquipmentTag, m.EquipmentName, m.Manufacturer, m.ModelNumber, m.CountryOfOrigin, m.InstallationDate, m.TotalOperatingHours, m.DesignLifeHours, m.CriticalityLevel, m.HealthStatus, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.Parameters {
		_, err := tx.Exec(ctx, `INSERT INTO mro_equipment_parameters (equipment_id, recorded_at, rpm, temperature_celsius, pressure_bar, vibration_level, oil_pressure_bar, running_hours_snapshot, status_flag, recorded_by_user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9::telemetry_status_type, $10)`, newID, item.RecordedAt, item.Rpm, item.TemperatureCelsius, item.PressureBar, item.VibrationLevel, item.OilPressureBar, item.RunningHoursSnapshot, item.StatusFlag, item.RecordedByUserId)
		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing equipment
func (r *EquipmentRepository) Update(ctx context.Context, id string, m *model.Equipment) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE mro_equipments SET system_id = $1, serial_number = $2, equipment_tag = $3, equipment_name = $4, manufacturer = $5, model_number = $6, country_of_origin = $7, installation_date = $8, total_operating_hours = $9, design_life_hours = $10, criticality_level = $11::equipment_criticality_type, health_status = $12::equipment_health_type, updated_by = $13, updated_at = CURRENT_TIMESTAMP WHERE equipment_id = $14`
	_, err = tx.Exec(ctx, updateQuery, m.SystemId, m.SerialNumber, m.EquipmentTag, m.EquipmentName, m.Manufacturer, m.ModelNumber, m.CountryOfOrigin, m.InstallationDate, m.TotalOperatingHours, m.DesignLifeHours, m.CriticalityLevel, m.HealthStatus, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	if len(m.Parameters) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM mro_equipment_parameters WHERE equipment_id = $1`, id)
		for _, item := range m.Parameters {
			_, err := tx.Exec(ctx, `INSERT INTO mro_equipment_parameters (equipment_id, recorded_at, rpm, temperature_celsius, pressure_bar, vibration_level, oil_pressure_bar, running_hours_snapshot, status_flag, recorded_by_user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9::telemetry_status_type, $10)`, id, item.RecordedAt, item.Rpm, item.TemperatureCelsius, item.PressureBar, item.VibrationLevel, item.OilPressureBar, item.RunningHoursSnapshot, item.StatusFlag, item.RecordedByUserId)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes equipment
func (r *EquipmentRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE mro_equipments SET deleted_at = CURRENT_TIMESTAMP WHERE equipment_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
