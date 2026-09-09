package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CuiMonitoringLogRepository struct {
	DB *pgxpool.Pool
}

func NewCuiMonitoringLogRepository(db *pgxpool.Pool) *CuiMonitoringLogRepository {
	return &CuiMonitoringLogRepository{DB: db}
}

// Get retrieves a single cui_monitoring_log by ID
func (r *CuiMonitoringLogRepository) Get(ctx context.Context, id string) (*model.CuiMonitoringLog, error) {
	query := `SELECT log_id, cui_asset_id, sensor_code, sensor_type, log_time, metric_value, metric_unit, status, vessel_proximity_mmsi, anomaly_score, description, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM cui_monitoring_logs WHERE log_id = $1 AND deleted_at IS NULL`

	row := r.DB.QueryRow(ctx, query, id)
	var m model.CuiMonitoringLog
	err := row.Scan(&m.LogId, &m.CuiAssetId, &m.SensorCode, &m.SensorType, &m.LogTime, &m.MetricValue, &m.MetricUnit, &m.Status, &m.VesselProximityMmsi, &m.AnomalyScore, &m.Description, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves cui_monitoring_log records with search, filter, and pagination
func (r *CuiMonitoringLogRepository) List(ctx context.Context, opts model.ListOptions) ([]model.CuiMonitoringLog, int, error) {
	baseQuery := `FROM cui_monitoring_logs WHERE 1=1`
	baseQuery += ` AND deleted_at IS NULL`
	var args []interface{}
	argIndex := 1

	if opts.Search != "" {
		baseQuery += fmt.Sprintf(" AND (sensor_code ILIKE $%d OR sensor_type ILIKE $%d OR metric_unit ILIKE $%d OR status ILIKE $%d OR vessel_proximity_mmsi ILIKE $%d OR description ILIKE $%d)", argIndex, argIndex, argIndex, argIndex, argIndex, argIndex)
		args = append(args, "%"+opts.Search+"%")
		argIndex++
	}

	// Generic filter support
	for k, v := range opts.Filters {
		if v != "" {
			baseQuery += fmt.Sprintf(" AND %s = $%d", k, argIndex)
			args = append(args, v)
			argIndex++
		}
	}

	// Count total
	var total int
	countQuery := "SELECT COUNT(*) " + baseQuery
	err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Sorting and Pagination
	sortCol := "log_id"
	if opts.SortBy != "" {
		sortCol = opts.SortBy
	}
	order := "DESC"
	if strings.ToUpper(opts.Order) == "ASC" {
		order = "ASC"
	}
	baseQuery += fmt.Sprintf(" ORDER BY %s %s", sortCol, order)

	if opts.Limit > 0 {
		baseQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
		args = append(args, opts.Limit, opts.Offset)
	}

	selectQuery := "SELECT log_id, cui_asset_id, sensor_code, sensor_type, log_time, metric_value, metric_unit, status, vessel_proximity_mmsi, anomaly_score, description, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at " + baseQuery
	rows, err := r.DB.Query(ctx, selectQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.CuiMonitoringLog
	for rows.Next() {
		var m model.CuiMonitoringLog
		if err := rows.Scan(&m.LogId, &m.CuiAssetId, &m.SensorCode, &m.SensorType, &m.LogTime, &m.MetricValue, &m.MetricUnit, &m.Status, &m.VesselProximityMmsi, &m.AnomalyScore, &m.Description, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new cui_monitoring_log
func (r *CuiMonitoringLogRepository) Create(ctx context.Context, m *model.CuiMonitoringLog) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO cui_monitoring_logs (cui_asset_id, sensor_code, sensor_type, log_time, metric_value, metric_unit, status, vessel_proximity_mmsi, anomaly_score, description, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING log_id`
	var newID string
	err = tx.QueryRow(ctx, query, m.CuiAssetId, m.SensorCode, m.SensorType, m.LogTime, m.MetricValue, m.MetricUnit, m.Status, m.VesselProximityMmsi, m.AnomalyScore, m.Description, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	return newID, tx.Commit(ctx)
}

// Update modifies an existing cui_monitoring_log
func (r *CuiMonitoringLogRepository) Update(ctx context.Context, id string, m *model.CuiMonitoringLog) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `UPDATE cui_monitoring_logs SET cui_asset_id = $1, sensor_code = $2, sensor_type = $3, log_time = $4, metric_value = $5, metric_unit = $6, status = $7, vessel_proximity_mmsi = $8, anomaly_score = $9, description = $10, updated_by = $11, updated_at = CURRENT_TIMESTAMP WHERE log_id = $12`
	_, err = tx.Exec(ctx, query, m.CuiAssetId, m.SensorCode, m.SensorType, m.LogTime, m.MetricValue, m.MetricUnit, m.Status, m.VesselProximityMmsi, m.AnomalyScore, m.Description, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes cui_monitoring_log
func (r *CuiMonitoringLogRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE cui_monitoring_logs SET deleted_at = CURRENT_TIMESTAMP WHERE log_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
