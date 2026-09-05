package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FailureReportRepository struct {
	DB *pgxpool.Pool
}

func NewFailureReportRepository(db *pgxpool.Pool) *FailureReportRepository {
	return &FailureReportRepository{DB: db}
}

// Get retrieves a single failure_report by report_id
func (r *FailureReportRepository) Get(ctx context.Context, id string) (*model.FailureReport, error) {
	query := `SELECT t.report_id, t.equipment_id, COALESCE(j_eqp.equipment_name, ''), t.reported_by_user_id, COALESCE(j_usr.full_name, ''), t.report_number, t.incident_date, t.severity, t.failure_mode, t.description, t.operational_impact, t.immediate_action_taken, t.status, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM mro_failure_reports t
	LEFT JOIN mro_equipments j_eqp ON j_eqp.equipment_id = t.equipment_id
	LEFT JOIN sys_users j_usr ON j_usr.user_id = t.reported_by_user_id
	WHERE t.report_id = $1 AND t.deleted_at IS NULL`

	var m model.FailureReport
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.ReportId, &m.EquipmentId, &m.EquipmentName, &m.ReportedByUserId, &m.ReporterName, &m.ReportNumber, &m.IncidentDate, &m.Severity, &m.FailureMode, &m.Description, &m.OperationalImpact, &m.ImmediateActionTaken, &m.Status, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated failure_report records
func (r *FailureReportRepository) List(ctx context.Context, opts model.ListOptions) ([]model.FailureReport, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "t.deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(report_number ILIKE $%[1]d OR failure_mode ILIKE $%[1]d OR description ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM mro_failure_reports t WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf(`SELECT t.report_id, t.equipment_id, COALESCE(j_eqp.equipment_name, ''), t.reported_by_user_id, COALESCE(j_usr.full_name, ''), t.report_number, t.incident_date, t.severity, t.failure_mode, t.description, t.operational_impact, t.immediate_action_taken, t.status, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM mro_failure_reports t
	LEFT JOIN mro_equipments j_eqp ON j_eqp.equipment_id = t.equipment_id
	LEFT JOIN sys_users j_usr ON j_usr.user_id = t.reported_by_user_id
	WHERE %s ORDER BY t.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.FailureReport
	for rows.Next() {
		var m model.FailureReport
		if err := rows.Scan(&m.ReportId, &m.EquipmentId, &m.EquipmentName, &m.ReportedByUserId, &m.ReporterName, &m.ReportNumber, &m.IncidentDate, &m.Severity, &m.FailureMode, &m.Description, &m.OperationalImpact, &m.ImmediateActionTaken, &m.Status, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new failure_report with optional child items
func (r *FailureReportRepository) Create(ctx context.Context, m *model.FailureReport) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO mro_failure_reports (equipment_id, reported_by_user_id, report_number, incident_date, severity, failure_mode, description, operational_impact, immediate_action_taken, status, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5::failure_severity_type, $6, $7, $8, $9, $10::failure_status_type, $11, $12, $13) RETURNING report_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.EquipmentId, m.ReportedByUserId, m.ReportNumber, m.IncidentDate, m.Severity, m.FailureMode, m.Description, m.OperationalImpact, m.ImmediateActionTaken, m.Status, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing failure_report
func (r *FailureReportRepository) Update(ctx context.Context, id string, m *model.FailureReport) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE mro_failure_reports SET equipment_id = $1, reported_by_user_id = $2, report_number = $3, incident_date = $4, severity = $5::failure_severity_type, failure_mode = $6, description = $7, operational_impact = $8, immediate_action_taken = $9, status = $10::failure_status_type, updated_by = $11, updated_at = CURRENT_TIMESTAMP WHERE report_id = $12`
	_, err = tx.Exec(ctx, updateQuery, m.EquipmentId, m.ReportedByUserId, m.ReportNumber, m.IncidentDate, m.Severity, m.FailureMode, m.Description, m.OperationalImpact, m.ImmediateActionTaken, m.Status, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes failure_report
func (r *FailureReportRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE mro_failure_reports SET deleted_at = CURRENT_TIMESTAMP WHERE report_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
