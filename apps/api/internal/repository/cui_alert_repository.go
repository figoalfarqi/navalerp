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
func (r *CuiAlertRepository) Get(ctx context.Context, id string) (*model.CuiAlert, error) {
	query := `SELECT alert_id, alert_code, cui_asset_id, alert_type, severity, detected_at, assigned_ship_id, status, ai_confidence, recommended_action, resolution_notes, resolved_at, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM cui_alerts WHERE alert_id = $1 AND deleted_at IS NULL`

	row := r.DB.QueryRow(ctx, query, id)
	var m model.CuiAlert
	err := row.Scan(&m.AlertId, &m.AlertCode, &m.CuiAssetId, &m.AlertType, &m.Severity, &m.DetectedAt, &m.AssignedShipId, &m.Status, &m.AiConfidence, &m.RecommendedAction, &m.ResolutionNotes, &m.ResolvedAt, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves cui_alert records with search, filter, and pagination
func (r *CuiAlertRepository) List(ctx context.Context, opts model.ListOptions) ([]model.CuiAlert, int, error) {
	baseQuery := `FROM cui_alerts WHERE 1=1`
	baseQuery += ` AND deleted_at IS NULL`
	var args []interface{}
	argIndex := 1

	if opts.Search != "" {
		baseQuery += fmt.Sprintf(" AND (alert_code ILIKE $%d OR recommended_action ILIKE $%d OR resolution_notes ILIKE $%d)", argIndex, argIndex, argIndex)
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
	sortCol := "alert_id"
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

	selectQuery := "SELECT alert_id, alert_code, cui_asset_id, alert_type, severity, detected_at, assigned_ship_id, status, ai_confidence, recommended_action, resolution_notes, resolved_at, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at " + baseQuery
	rows, err := r.DB.Query(ctx, selectQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.CuiAlert
	for rows.Next() {
		var m model.CuiAlert
		if err := rows.Scan(&m.AlertId, &m.AlertCode, &m.CuiAssetId, &m.AlertType, &m.Severity, &m.DetectedAt, &m.AssignedShipId, &m.Status, &m.AiConfidence, &m.RecommendedAction, &m.ResolutionNotes, &m.ResolvedAt, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
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
