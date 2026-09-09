package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CuiAlertRepository struct {
	DB *pgxpool.Pool
}

func NewCuiAlertRepository(db *pgxpool.Pool) *CuiAlertRepository {
	return &CuiAlertRepository{DB: db}
}

// Get retrieves a single cui_alert by ID
// Get retrieves a single cui_alert by ID
func (r *CuiAlertRepository) Get(ctx context.Context, id string) (*model.CuiAlert, error) {
	query := `SELECT 
		l.alert_id, l.alert_code, l.cui_asset_id, l.alert_type, l.severity, l.detected_at, 
		l.assigned_ship_id, l.status, l.ai_confidence, l.recommended_action, l.resolution_notes, 
		l.resolved_at, l.created_by, l.updated_by, l.deleted_by, l.created_at, l.updated_at, l.deleted_at,
		a.asset_name, a.asset_code, s.ship_name, s.hull_number
	FROM cui_alerts l
	LEFT JOIN cui_assets a ON a.cui_asset_id = l.cui_asset_id
	LEFT JOIN mro_ships s ON s.ship_id = l.assigned_ship_id
	WHERE l.alert_id = $1 AND l.deleted_at IS NULL`

	row := r.DB.QueryRow(ctx, query, id)
	var m model.CuiAlert
	err := row.Scan(&m.AlertId, &m.AlertCode, &m.CuiAssetId, &m.AlertType, &m.Severity, &m.DetectedAt, &m.AssignedShipId, &m.Status, &m.AiConfidence, &m.RecommendedAction, &m.ResolutionNotes, &m.ResolvedAt, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt, &m.AssetName, &m.AssetCode, &m.ShipName, &m.HullNumber)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves cui_alert records with search, filter, and pagination
func (r *CuiAlertRepository) List(ctx context.Context, opts model.ListOptions) ([]model.CuiAlert, int, error) {
	whereQuery := `FROM cui_alerts l
	LEFT JOIN cui_assets a ON a.cui_asset_id = l.cui_asset_id
	LEFT JOIN mro_ships s ON s.ship_id = l.assigned_ship_id
	WHERE 1=1 AND l.deleted_at IS NULL`
	var args []interface{}
	argIndex := 1

	if opts.Search != "" {
		whereQuery += fmt.Sprintf(" AND (l.alert_code ILIKE $%d OR l.recommended_action ILIKE $%d OR l.resolution_notes ILIKE $%d OR a.asset_name ILIKE $%d OR a.asset_code ILIKE $%d OR s.ship_name ILIKE $%d)", argIndex, argIndex, argIndex, argIndex, argIndex, argIndex)
		args = append(args, "%"+opts.Search+"%")
		argIndex++
	}

	allowedCols := map[string]bool{
		"alert_id":           true,
		"alert_code":         true,
		"cui_asset_id":       true,
		"alert_type":         true,
		"severity":           true,
		"detected_at":        true,
		"assigned_ship_id":   true,
		"status":             true,
		"ai_confidence":      true,
		"recommended_action": true,
		"resolution_notes":   true,
		"resolved_at":        true,
		"created_at":         true,
		"updated_at":         true,
	}

	// Generic filter support with column whitelisting
	for k, v := range opts.Filters {
		if v != "" && allowedCols[k] {
			whereQuery += fmt.Sprintf(" AND l.%s = $%d", k, argIndex)
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
	whereQuery += fmt.Sprintf(" ORDER BY l.%s %s", sortCol, order)

	if opts.Limit > 0 {
		whereQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
		args = append(args, opts.Limit, opts.Offset)
	}

	selectQuery := `SELECT 
		l.alert_id, l.alert_code, l.cui_asset_id, l.alert_type, l.severity, l.detected_at, 
		l.assigned_ship_id, l.status, l.ai_confidence, l.recommended_action, l.resolution_notes, 
		l.resolved_at, l.created_by, l.updated_by, l.deleted_by, l.created_at, l.updated_at, l.deleted_at,
		a.asset_name, a.asset_code, s.ship_name, s.hull_number ` + whereQuery
	rows, err := r.DB.Query(ctx, selectQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.CuiAlert
	for rows.Next() {
		var m model.CuiAlert
		if err := rows.Scan(&m.AlertId, &m.AlertCode, &m.CuiAssetId, &m.AlertType, &m.Severity, &m.DetectedAt, &m.AssignedShipId, &m.Status, &m.AiConfidence, &m.RecommendedAction, &m.ResolutionNotes, &m.ResolvedAt, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt, &m.AssetName, &m.AssetCode, &m.ShipName, &m.HullNumber); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new cui_alert
func (r *CuiAlertRepository) Create(ctx context.Context, m *model.CuiAlert) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO cui_alerts (alert_code, cui_asset_id, alert_type, severity, detected_at, assigned_ship_id, status, ai_confidence, recommended_action, resolution_notes, resolved_at, created_by, updated_by, deleted_by) VALUES ($1, $2, $3::cui_alert_type_enum, $4::cui_alert_severity_enum, $5, $6, $7::cui_alert_status_enum, $8, $9, $10, $11, $12, $13, $14) RETURNING alert_id`
	var newID string
	err = tx.QueryRow(ctx, query, m.AlertCode, m.CuiAssetId, m.AlertType, m.Severity, m.DetectedAt, m.AssignedShipId, m.Status, m.AiConfidence, m.RecommendedAction, m.ResolutionNotes, m.ResolvedAt, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	return newID, tx.Commit(ctx)
}

// Update modifies an existing cui_alert
func (r *CuiAlertRepository) Update(ctx context.Context, id string, m *model.CuiAlert) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `UPDATE cui_alerts SET alert_code = $1, cui_asset_id = $2, alert_type = $3::cui_alert_type_enum, severity = $4::cui_alert_severity_enum, detected_at = $5, assigned_ship_id = $6, status = $7::cui_alert_status_enum, ai_confidence = $8, recommended_action = $9, resolution_notes = $10, resolved_at = $11, updated_by = $12, updated_at = CURRENT_TIMESTAMP WHERE alert_id = $13`
	_, err = tx.Exec(ctx, query, m.AlertCode, m.CuiAssetId, m.AlertType, m.Severity, m.DetectedAt, m.AssignedShipId, m.Status, m.AiConfidence, m.RecommendedAction, m.ResolutionNotes, m.ResolvedAt, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes cui_alert
func (r *CuiAlertRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE cui_alerts SET deleted_at = CURRENT_TIMESTAMP WHERE alert_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
