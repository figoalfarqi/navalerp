package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CuiAssetRepository struct {
	DB *pgxpool.Pool
}

func NewCuiAssetRepository(db *pgxpool.Pool) *CuiAssetRepository {
	return &CuiAssetRepository{DB: db}
}

// Get retrieves a single cui_asset by ID
func (r *CuiAssetRepository) Get(ctx context.Context, id string) (*model.CuiAsset, error) {
	query := `SELECT 
		c.cui_asset_id, c.asset_code, c.asset_name, c.asset_type, c.operator_name, 
		c.theater_id, c.depth_meters, c.length_km, c.latitude, c.longitude, 
		c.start_coordinates, c.end_coordinates, c.status, c.health_score, c.protection_priority, 
		c.last_inspected_at, c.next_inspection_due, c.notes, c.created_by, c.updated_by, 
		c.deleted_by, c.created_at, c.updated_at, c.deleted_at,
		t.theater_name, t.theater_code
	FROM cui_assets c
	LEFT JOIN ops_theaters t ON t.theater_id = c.theater_id
	WHERE c.cui_asset_id = $1 AND c.deleted_at IS NULL`

	row := r.DB.QueryRow(ctx, query, id)
	var m model.CuiAsset
	err := row.Scan(&m.CuiAssetId, &m.AssetCode, &m.AssetName, &m.AssetType, &m.OperatorName, &m.TheaterId, &m.DepthMeters, &m.LengthKm, &m.Latitude, &m.Longitude, &m.StartCoordinates, &m.EndCoordinates, &m.Status, &m.HealthScore, &m.ProtectionPriority, &m.LastInspectedAt, &m.NextInspectionDue, &m.Notes, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt, &m.TheaterName, &m.TheaterCode)
	if err != nil {
		return nil, err
	}

	// Fetch children: MonitoringLogs
	cRows_MonitoringLogs, err := r.DB.Query(ctx, `SELECT log_id, cui_asset_id, sensor_code, sensor_type, log_time, metric_value, metric_unit, status, vessel_proximity_mmsi, anomaly_score, description, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM cui_monitoring_logs WHERE cui_asset_id = $1 AND deleted_at IS NULL`, id)
	if err == nil {
		defer cRows_MonitoringLogs.Close()
		for cRows_MonitoringLogs.Next() {
			var item model.MonitoringLogs
			if scanErr := cRows_MonitoringLogs.Scan(&item.LogId, &item.CuiAssetId, &item.SensorCode, &item.SensorType, &item.LogTime, &item.MetricValue, &item.MetricUnit, &item.Status, &item.VesselProximityMmsi, &item.AnomalyScore, &item.Description, &item.CreatedBy, &item.UpdatedBy, &item.DeletedBy, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt); scanErr == nil {
				m.MonitoringLogs = append(m.MonitoringLogs, item)
			}
		}
	}

	// Fetch children: Alerts
	cRows_Alerts, err := r.DB.Query(ctx, `SELECT alert_id, alert_code, cui_asset_id, alert_type, severity, detected_at, assigned_ship_id, status, ai_confidence, recommended_action, resolution_notes, resolved_at, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM cui_alerts WHERE cui_asset_id = $1 AND deleted_at IS NULL`, id)
	if err == nil {
		defer cRows_Alerts.Close()
		for cRows_Alerts.Next() {
			var item model.Alerts
			if scanErr := cRows_Alerts.Scan(&item.AlertId, &item.AlertCode, &item.CuiAssetId, &item.AlertType, &item.Severity, &item.DetectedAt, &item.AssignedShipId, &item.Status, &item.AiConfidence, &item.RecommendedAction, &item.ResolutionNotes, &item.ResolvedAt, &item.CreatedBy, &item.UpdatedBy, &item.DeletedBy, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt); scanErr == nil {
				m.Alerts = append(m.Alerts, item)
			}
		}
	}

	// Fetch children: Inspections
	cRows_Inspections, err := r.DB.Query(ctx, `SELECT inspection_id, inspection_number, cui_asset_id, ship_id, inspection_date, inspector_officer_id, method, condition_rating, findings, remedial_action_required, next_inspection_date, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM cui_inspections WHERE cui_asset_id = $1 AND deleted_at IS NULL`, id)
	if err == nil {
		defer cRows_Inspections.Close()
		for cRows_Inspections.Next() {
			var item model.Inspections
			if scanErr := cRows_Inspections.Scan(&item.InspectionId, &item.InspectionNumber, &item.CuiAssetId, &item.ShipId, &item.InspectionDate, &item.InspectorOfficerId, &item.Method, &item.ConditionRating, &item.Findings, &item.RemedialActionRequired, &item.NextInspectionDate, &item.CreatedBy, &item.UpdatedBy, &item.DeletedBy, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt); scanErr == nil {
				m.Inspections = append(m.Inspections, item)
			}
		}
	}

	return &m, nil
}

// List retrieves cui_asset records with search, filter, and pagination
func (r *CuiAssetRepository) List(ctx context.Context, opts model.ListOptions) ([]model.CuiAsset, int, error) {
	whereQuery := `FROM cui_assets c
	LEFT JOIN ops_theaters t ON t.theater_id = c.theater_id
	WHERE 1=1 AND c.deleted_at IS NULL`
	var args []interface{}
	argIndex := 1

	if opts.Search != "" {
		whereQuery += fmt.Sprintf(" AND (c.asset_code ILIKE $%d OR c.asset_name ILIKE $%d OR c.operator_name ILIKE $%d OR c.start_coordinates ILIKE $%d OR c.end_coordinates ILIKE $%d OR c.notes ILIKE $%d OR t.theater_name ILIKE $%d)", argIndex, argIndex, argIndex, argIndex, argIndex, argIndex, argIndex)
		args = append(args, "%"+opts.Search+"%")
		argIndex++
	}

	allowedCols := map[string]bool{
		"cui_asset_id":        true,
		"asset_code":          true,
		"asset_name":          true,
		"asset_type":          true,
		"operator_name":       true,
		"theater_id":          true,
		"depth_meters":        true,
		"length_km":           true,
		"latitude":            true,
		"longitude":           true,
		"start_coordinates":   true,
		"end_coordinates":     true,
		"status":              true,
		"health_score":        true,
		"protection_priority": true,
		"last_inspected_at":   true,
		"next_inspection_due": true,
		"created_at":          true,
		"updated_at":          true,
	}

	// Generic filter support with column whitelisting
	for k, v := range opts.Filters {
		if v != "" && allowedCols[k] {
			whereQuery += fmt.Sprintf(" AND c.%s = $%d", k, argIndex)
			args = append(args, v)
			argIndex++
		}
	}

	// Count total
	var total int
	countQuery := "SELECT COUNT(*) " + whereQuery
	err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Sorting and Pagination
	sortCol := "created_at"
	if opts.SortBy != "" && allowedCols[opts.SortBy] {
		sortCol = opts.SortBy
	}
	order := "DESC"
	if strings.ToUpper(opts.Order) == "ASC" || strings.ToUpper(opts.Sort) == "ASC" {
		order = "ASC"
	}
	whereQuery += fmt.Sprintf(" ORDER BY c.%s %s", sortCol, order)

	if opts.Limit > 0 {
		whereQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
		args = append(args, opts.Limit, opts.Offset)
	}

	selectQuery := `SELECT 
		c.cui_asset_id, c.asset_code, c.asset_name, c.asset_type, c.operator_name, 
		c.theater_id, c.depth_meters, c.length_km, c.latitude, c.longitude, 
		c.start_coordinates, c.end_coordinates, c.status, c.health_score, c.protection_priority, 
		c.last_inspected_at, c.next_inspection_due, c.notes, c.created_by, c.updated_by, 
		c.deleted_by, c.created_at, c.updated_at, c.deleted_at,
		t.theater_name, t.theater_code ` + whereQuery
	rows, err := r.DB.Query(ctx, selectQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.CuiAsset
	for rows.Next() {
		var m model.CuiAsset
		if err := rows.Scan(&m.CuiAssetId, &m.AssetCode, &m.AssetName, &m.AssetType, &m.OperatorName, &m.TheaterId, &m.DepthMeters, &m.LengthKm, &m.Latitude, &m.Longitude, &m.StartCoordinates, &m.EndCoordinates, &m.Status, &m.HealthScore, &m.ProtectionPriority, &m.LastInspectedAt, &m.NextInspectionDue, &m.Notes, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt, &m.TheaterName, &m.TheaterCode); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new cui_asset
func (r *CuiAssetRepository) Create(ctx context.Context, m *model.CuiAsset) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO cui_assets (asset_code, asset_name, asset_type, operator_name, theater_id, depth_meters, length_km, latitude, longitude, start_coordinates, end_coordinates, status, health_score, protection_priority, last_inspected_at, next_inspection_due, notes, created_by, updated_by, deleted_by) VALUES ($1, $2, $3::cui_asset_type_enum, $4, $5, $6, $7, $8, $9, $10, $11, $12::cui_status_enum, $13, $14::cui_priority_enum, $15, $16, $17, $18, $19, $20) RETURNING cui_asset_id`
	var newID string
	err = tx.QueryRow(ctx, query, m.AssetCode, m.AssetName, m.AssetType, m.OperatorName, m.TheaterId, m.DepthMeters, m.LengthKm, m.Latitude, m.Longitude, m.StartCoordinates, m.EndCoordinates, m.Status, m.HealthScore, m.ProtectionPriority, m.LastInspectedAt, m.NextInspectionDue, m.Notes, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.MonitoringLogs {
		_, err := tx.Exec(ctx, `INSERT INTO cui_monitoring_logs (cui_asset_id, sensor_code, sensor_type, log_time, metric_value, metric_unit, status, vessel_proximity_mmsi, anomaly_score, description, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`, newID, item.SensorCode, item.SensorType, item.LogTime, item.MetricValue, item.MetricUnit, item.Status, item.VesselProximityMmsi, item.AnomalyScore, item.Description, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
		if err != nil {
			return "", err
		}
	}

	for _, item := range m.Alerts {
		_, err := tx.Exec(ctx, `INSERT INTO cui_alerts (alert_code, cui_asset_id, alert_type, severity, detected_at, assigned_ship_id, status, ai_confidence, recommended_action, resolution_notes, resolved_at, created_by, updated_by, deleted_by) VALUES ($1, $2, $3::cui_alert_type_enum, $4::cui_alert_severity_enum, $5, $6, $7::cui_alert_status_enum, $8, $9, $10, $11, $12, $13, $14)`, item.AlertCode, newID, item.AlertType, item.Severity, item.DetectedAt, item.AssignedShipId, item.Status, item.AiConfidence, item.RecommendedAction, item.ResolutionNotes, item.ResolvedAt, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
		if err != nil {
			return "", err
		}
	}

	for _, item := range m.Inspections {
		_, err := tx.Exec(ctx, `INSERT INTO cui_inspections (inspection_number, cui_asset_id, ship_id, inspection_date, inspector_officer_id, method, condition_rating, findings, remedial_action_required, next_inspection_date, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6::cui_inspection_method_enum, $7::cui_condition_enum, $8, $9, $10, $11, $12, $13)`, item.InspectionNumber, newID, item.ShipId, item.InspectionDate, item.InspectorOfficerId, item.Method, item.ConditionRating, item.Findings, item.RemedialActionRequired, item.NextInspectionDate, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
		if err != nil {
			return "", err
		}
	}

	return newID, tx.Commit(ctx)
}

// Update modifies an existing cui_asset
func (r *CuiAssetRepository) Update(ctx context.Context, id string, m *model.CuiAsset) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `UPDATE cui_assets SET asset_code = $1, asset_name = $2, asset_type = $3::cui_asset_type_enum, operator_name = $4, theater_id = $5, depth_meters = $6, length_km = $7, latitude = $8, longitude = $9, start_coordinates = $10, end_coordinates = $11, status = $12::cui_status_enum, health_score = $13, protection_priority = $14::cui_priority_enum, last_inspected_at = $15, next_inspection_due = $16, notes = $17, updated_by = $18, updated_at = CURRENT_TIMESTAMP WHERE cui_asset_id = $19`
	_, err = tx.Exec(ctx, query, m.AssetCode, m.AssetName, m.AssetType, m.OperatorName, m.TheaterId, m.DepthMeters, m.LengthKm, m.Latitude, m.Longitude, m.StartCoordinates, m.EndCoordinates, m.Status, m.HealthScore, m.ProtectionPriority, m.LastInspectedAt, m.NextInspectionDue, m.Notes, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	// Replace children: MonitoringLogs
	if len(m.MonitoringLogs) > 0 {
		_, err := tx.Exec(ctx, `DELETE FROM cui_monitoring_logs WHERE cui_asset_id = $1`, id)
		if err != nil {
			return err
		}
		for _, item := range m.MonitoringLogs {
			_, err := tx.Exec(ctx, `INSERT INTO cui_monitoring_logs (cui_asset_id, sensor_code, sensor_type, log_time, metric_value, metric_unit, status, vessel_proximity_mmsi, anomaly_score, description, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`, id, item.SensorCode, item.SensorType, item.LogTime, item.MetricValue, item.MetricUnit, item.Status, item.VesselProximityMmsi, item.AnomalyScore, item.Description, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
			if err != nil {
				return err
			}
		}
	}

	// Replace children: Alerts
	if len(m.Alerts) > 0 {
		_, err := tx.Exec(ctx, `DELETE FROM cui_alerts WHERE cui_asset_id = $1`, id)
		if err != nil {
			return err
		}
		for _, item := range m.Alerts {
			_, err := tx.Exec(ctx, `INSERT INTO cui_alerts (alert_code, cui_asset_id, alert_type, severity, detected_at, assigned_ship_id, status, ai_confidence, recommended_action, resolution_notes, resolved_at, created_by, updated_by, deleted_by) VALUES ($1, $2, $3::cui_alert_type_enum, $4::cui_alert_severity_enum, $5, $6, $7::cui_alert_status_enum, $8, $9, $10, $11, $12, $13, $14)`, item.AlertCode, id, item.AlertType, item.Severity, item.DetectedAt, item.AssignedShipId, item.Status, item.AiConfidence, item.RecommendedAction, item.ResolutionNotes, item.ResolvedAt, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
			if err != nil {
				return err
			}
		}
	}

	// Replace children: Inspections
	if len(m.Inspections) > 0 {
		_, err := tx.Exec(ctx, `DELETE FROM cui_inspections WHERE cui_asset_id = $1`, id)
		if err != nil {
			return err
		}
		for _, item := range m.Inspections {
			_, err := tx.Exec(ctx, `INSERT INTO cui_inspections (inspection_number, cui_asset_id, ship_id, inspection_date, inspector_officer_id, method, condition_rating, findings, remedial_action_required, next_inspection_date, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6::cui_inspection_method_enum, $7::cui_condition_enum, $8, $9, $10, $11, $12, $13)`, item.InspectionNumber, id, item.ShipId, item.InspectionDate, item.InspectorOfficerId, item.Method, item.ConditionRating, item.Findings, item.RemedialActionRequired, item.NextInspectionDate, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes cui_asset
func (r *CuiAssetRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE cui_assets SET deleted_at = CURRENT_TIMESTAMP WHERE cui_asset_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
