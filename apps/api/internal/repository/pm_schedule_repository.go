package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PmScheduleRepository struct {
	DB *pgxpool.Pool
}

func NewPmScheduleRepository(db *pgxpool.Pool) *PmScheduleRepository {
	return &PmScheduleRepository{DB: db}
}

// Get retrieves a single pm_schedule by pm_id
func (r *PmScheduleRepository) Get(ctx context.Context, id string) (*model.PmSchedule, error) {
	query := `SELECT t.pm_id, t.equipment_id, COALESCE(j_eqp.equipment_name, ''), t.pm_code, t.pm_title, t.interval_hours, t.interval_days, t.last_performed_at, t.next_due_at, t.task_instructions, t.estimated_duration_hours, t.is_active, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM mro_pm_schedules t
	LEFT JOIN mro_equipments j_eqp ON j_eqp.equipment_id = t.equipment_id
	WHERE t.pm_id = $1 AND t.deleted_at IS NULL`

	var m model.PmSchedule
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.PmId, &m.EquipmentId, &m.EquipmentName, &m.PmCode, &m.PmTitle, &m.IntervalHours, &m.IntervalDays, &m.LastPerformedAt, &m.NextDueAt, &m.TaskInstructions, &m.EstimatedDurationHours, &m.IsActive, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated pm_schedule records
func (r *PmScheduleRepository) List(ctx context.Context, opts model.ListOptions) ([]model.PmSchedule, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "t.deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(pm_code ILIKE $%[1]d OR pm_title ILIKE $%[1]d OR task_instructions ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM mro_pm_schedules t WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf(`SELECT t.pm_id, t.equipment_id, COALESCE(j_eqp.equipment_name, ''), t.pm_code, t.pm_title, t.interval_hours, t.interval_days, t.last_performed_at, t.next_due_at, t.task_instructions, t.estimated_duration_hours, t.is_active, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM mro_pm_schedules t
	LEFT JOIN mro_equipments j_eqp ON j_eqp.equipment_id = t.equipment_id
	WHERE %s ORDER BY t.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.PmSchedule
	for rows.Next() {
		var m model.PmSchedule
		if err := rows.Scan(&m.PmId, &m.EquipmentId, &m.EquipmentName, &m.PmCode, &m.PmTitle, &m.IntervalHours, &m.IntervalDays, &m.LastPerformedAt, &m.NextDueAt, &m.TaskInstructions, &m.EstimatedDurationHours, &m.IsActive, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new pm_schedule with optional child items
func (r *PmScheduleRepository) Create(ctx context.Context, m *model.PmSchedule) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO mro_pm_schedules (equipment_id, pm_code, pm_title, interval_hours, interval_days, last_performed_at, next_due_at, task_instructions, estimated_duration_hours, is_active, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING pm_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.EquipmentId, m.PmCode, m.PmTitle, m.IntervalHours, m.IntervalDays, m.LastPerformedAt, m.NextDueAt, m.TaskInstructions, m.EstimatedDurationHours, m.IsActive, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing pm_schedule
func (r *PmScheduleRepository) Update(ctx context.Context, id string, m *model.PmSchedule) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE mro_pm_schedules SET equipment_id = $1, pm_code = $2, pm_title = $3, interval_hours = $4, interval_days = $5, last_performed_at = $6, next_due_at = $7, task_instructions = $8, estimated_duration_hours = $9, is_active = $10, updated_by = $11, updated_at = CURRENT_TIMESTAMP WHERE pm_id = $12`
	_, err = tx.Exec(ctx, updateQuery, m.EquipmentId, m.PmCode, m.PmTitle, m.IntervalHours, m.IntervalDays, m.LastPerformedAt, m.NextDueAt, m.TaskInstructions, m.EstimatedDurationHours, m.IsActive, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes pm_schedule
func (r *PmScheduleRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE mro_pm_schedules SET deleted_at = CURRENT_TIMESTAMP WHERE pm_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
