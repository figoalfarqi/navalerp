package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ReadinessAlertRepository struct {
	DB *pgxpool.Pool
}

func NewReadinessAlertRepository(db *pgxpool.Pool) *ReadinessAlertRepository {
	return &ReadinessAlertRepository{DB: db}
}

// Get retrieves a single readiness_alert by alert_id
func (r *ReadinessAlertRepository) Get(ctx context.Context, id string) (*model.ReadinessAlert, error) {
	query := `SELECT alert_id, ship_id, equipment_id, severity, alert_type, alert_message, is_acknowledged, acknowledged_by_user_id, acknowledged_at, created_at FROM ops_readiness_alerts WHERE alert_id = $1`

	var m model.ReadinessAlert
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.AlertId, &m.ShipId, &m.EquipmentId, &m.Severity, &m.AlertType, &m.AlertMessage, &m.IsAcknowledged, &m.AcknowledgedByUserId, &m.AcknowledgedAt, &m.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated readiness_alert records
func (r *ReadinessAlertRepository) List(ctx context.Context, opts model.ListOptions) ([]model.ReadinessAlert, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(alert_message ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM ops_readiness_alerts WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT alert_id, ship_id, equipment_id, severity, alert_type, alert_message, is_acknowledged, acknowledged_by_user_id, acknowledged_at, created_at FROM ops_readiness_alerts WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.ReadinessAlert
	for rows.Next() {
		var m model.ReadinessAlert
		if err := rows.Scan(&m.AlertId, &m.ShipId, &m.EquipmentId, &m.Severity, &m.AlertType, &m.AlertMessage, &m.IsAcknowledged, &m.AcknowledgedByUserId, &m.AcknowledgedAt, &m.CreatedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new readiness_alert with optional child items
func (r *ReadinessAlertRepository) Create(ctx context.Context, m *model.ReadinessAlert) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO ops_readiness_alerts (ship_id, equipment_id, severity, alert_type, alert_message, is_acknowledged, acknowledged_by_user_id, acknowledged_at) VALUES ($1, $2, $3::alert_severity_type, $4::alert_type_enum, $5, $6, $7, $8) RETURNING alert_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.ShipId, m.EquipmentId, m.Severity, m.AlertType, m.AlertMessage, m.IsAcknowledged, m.AcknowledgedByUserId, m.AcknowledgedAt).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing readiness_alert
func (r *ReadinessAlertRepository) Update(ctx context.Context, id string, m *model.ReadinessAlert) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE ops_readiness_alerts SET ship_id = $1, equipment_id = $2, severity = $3::alert_severity_type, alert_type = $4::alert_type_enum, alert_message = $5, is_acknowledged = $6, acknowledged_by_user_id = $7, acknowledged_at = $8 WHERE alert_id = $9`
	_, err = tx.Exec(ctx, updateQuery, m.ShipId, m.EquipmentId, m.Severity, m.AlertType, m.AlertMessage, m.IsAcknowledged, m.AcknowledgedByUserId, m.AcknowledgedAt, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes readiness_alert
func (r *ReadinessAlertRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM ops_readiness_alerts WHERE alert_id = $1`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
