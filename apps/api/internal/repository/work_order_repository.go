package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WorkOrderRepository struct {
	DB *pgxpool.Pool
}

func NewWorkOrderRepository(db *pgxpool.Pool) *WorkOrderRepository {
	return &WorkOrderRepository{DB: db}
}

// Get retrieves a single work_order by work_order_id
func (r *WorkOrderRepository) Get(ctx context.Context, id string) (*model.WorkOrder, error) {
	query := `SELECT t.work_order_id, t.failure_report_id, t.pm_schedule_id, t.equipment_id, COALESCE(j_eqp.equipment_name, ''), t.work_order_number, t.work_order_type, t.priority, t.scheduled_start_date, t.scheduled_end_date, t.actual_start_date, t.actual_end_date, t.lead_engineer_user_id, COALESCE(j_usr.full_name, ''), t.assigned_facility, t.status, t.total_labor_hours, t.estimated_cost, t.actual_cost, t.completion_notes, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM mro_work_orders t
	LEFT JOIN mro_equipments j_eqp ON j_eqp.equipment_id = t.equipment_id
	LEFT JOIN sys_users j_usr ON j_usr.user_id = t.lead_engineer_user_id
	WHERE t.work_order_id = $1 AND t.deleted_at IS NULL`

	var m model.WorkOrder
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.WorkOrderId, &m.FailureReportId, &m.PmScheduleId, &m.EquipmentId, &m.EquipmentName, &m.WorkOrderNumber, &m.WorkOrderType, &m.Priority, &m.ScheduledStartDate, &m.ScheduledEndDate, &m.ActualStartDate, &m.ActualEndDate, &m.LeadEngineerUserId, &m.EngineerName, &m.AssignedFacility, &m.Status, &m.TotalLaborHours, &m.EstimatedCost, &m.ActualCost, &m.CompletionNotes, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	// Load Tasks
	childRowsTasks, err := r.DB.Query(ctx, `SELECT task_id, work_order_id, step_number, task_description, estimated_minutes, actual_minutes, is_completed, completed_by_user_id, notes, created_at FROM mro_work_order_tasks WHERE work_order_id = $1`, id)
	if err == nil {
		defer childRowsTasks.Close()
		for childRowsTasks.Next() {
			var item model.WorkOrderTasks
			if err := childRowsTasks.Scan(&item.TaskId, &item.WorkOrderId, &item.StepNumber, &item.TaskDescription, &item.EstimatedMinutes, &item.ActualMinutes, &item.IsCompleted, &item.CompletedByUserId, &item.Notes, &item.CreatedAt); err == nil {
				m.Tasks = append(m.Tasks, item)
			}
		}
	}

	// Load Items
	childRowsItems, err := r.DB.Query(ctx, `SELECT wo_item_id, work_order_id, material_id, quantity_required, quantity_issued, unit_cost, total_cost, is_critical_spare, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM mro_work_order_items WHERE work_order_id = $1`, id)
	if err == nil {
		defer childRowsItems.Close()
		for childRowsItems.Next() {
			var item model.WorkOrderItems
			if err := childRowsItems.Scan(&item.WoItemId, &item.WorkOrderId, &item.MaterialId, &item.QuantityRequired, &item.QuantityIssued, &item.UnitCost, &item.TotalCost, &item.IsCriticalSpare, &item.CreatedBy, &item.UpdatedBy, &item.DeletedBy, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt); err == nil {
				m.Items = append(m.Items, item)
			}
		}
	}

	return &m, nil
}

// List retrieves paginated work_order records
func (r *WorkOrderRepository) List(ctx context.Context, opts model.ListOptions) ([]model.WorkOrder, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "t.deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(work_order_number ILIKE $%[1]d OR assigned_facility ILIKE $%[1]d OR completion_notes ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM mro_work_orders t WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf(`SELECT t.work_order_id, t.failure_report_id, t.pm_schedule_id, t.equipment_id, COALESCE(j_eqp.equipment_name, ''), t.work_order_number, t.work_order_type, t.priority, t.scheduled_start_date, t.scheduled_end_date, t.actual_start_date, t.actual_end_date, t.lead_engineer_user_id, COALESCE(j_usr.full_name, ''), t.assigned_facility, t.status, t.total_labor_hours, t.estimated_cost, t.actual_cost, t.completion_notes, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM mro_work_orders t
	LEFT JOIN mro_equipments j_eqp ON j_eqp.equipment_id = t.equipment_id
	LEFT JOIN sys_users j_usr ON j_usr.user_id = t.lead_engineer_user_id
	WHERE %s ORDER BY t.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.WorkOrder
	for rows.Next() {
		var m model.WorkOrder
		if err := rows.Scan(&m.WorkOrderId, &m.FailureReportId, &m.PmScheduleId, &m.EquipmentId, &m.EquipmentName, &m.WorkOrderNumber, &m.WorkOrderType, &m.Priority, &m.ScheduledStartDate, &m.ScheduledEndDate, &m.ActualStartDate, &m.ActualEndDate, &m.LeadEngineerUserId, &m.EngineerName, &m.AssignedFacility, &m.Status, &m.TotalLaborHours, &m.EstimatedCost, &m.ActualCost, &m.CompletionNotes, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new work_order with optional child items
func (r *WorkOrderRepository) Create(ctx context.Context, m *model.WorkOrder) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO mro_work_orders (failure_report_id, pm_schedule_id, equipment_id, work_order_number, work_order_type, priority, scheduled_start_date, scheduled_end_date, actual_start_date, actual_end_date, lead_engineer_user_id, assigned_facility, status, total_labor_hours, estimated_cost, actual_cost, completion_notes, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5::work_order_type_enum, $6::work_order_priority_type, $7, $8, $9, $10, $11, $12, $13::work_order_status_type, $14, $15, $16, $17, $18, $19, $20) RETURNING work_order_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.FailureReportId, m.PmScheduleId, m.EquipmentId, m.WorkOrderNumber, m.WorkOrderType, m.Priority, m.ScheduledStartDate, m.ScheduledEndDate, m.ActualStartDate, m.ActualEndDate, m.LeadEngineerUserId, m.AssignedFacility, m.Status, m.TotalLaborHours, m.EstimatedCost, m.ActualCost, m.CompletionNotes, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.Tasks {
		_, err := tx.Exec(ctx, `INSERT INTO mro_work_order_tasks (work_order_id, step_number, task_description, estimated_minutes, actual_minutes, is_completed, completed_by_user_id, notes) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, newID, item.StepNumber, item.TaskDescription, item.EstimatedMinutes, item.ActualMinutes, item.IsCompleted, item.CompletedByUserId, item.Notes)
		if err != nil {
			return "", err
		}
	}

	for _, item := range m.Items {
		_, err := tx.Exec(ctx, `INSERT INTO mro_work_order_items (work_order_id, material_id, quantity_required, quantity_issued, unit_cost, total_cost, is_critical_spare, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`, newID, item.MaterialId, item.QuantityRequired, item.QuantityIssued, item.UnitCost, item.TotalCost, item.IsCriticalSpare, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing work_order
func (r *WorkOrderRepository) Update(ctx context.Context, id string, m *model.WorkOrder) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE mro_work_orders SET failure_report_id = $1, pm_schedule_id = $2, equipment_id = $3, work_order_number = $4, work_order_type = $5::work_order_type_enum, priority = $6::work_order_priority_type, scheduled_start_date = $7, scheduled_end_date = $8, actual_start_date = $9, actual_end_date = $10, lead_engineer_user_id = $11, assigned_facility = $12, status = $13::work_order_status_type, total_labor_hours = $14, estimated_cost = $15, actual_cost = $16, completion_notes = $17, updated_by = $18, updated_at = CURRENT_TIMESTAMP WHERE work_order_id = $19`
	_, err = tx.Exec(ctx, updateQuery, m.FailureReportId, m.PmScheduleId, m.EquipmentId, m.WorkOrderNumber, m.WorkOrderType, m.Priority, m.ScheduledStartDate, m.ScheduledEndDate, m.ActualStartDate, m.ActualEndDate, m.LeadEngineerUserId, m.AssignedFacility, m.Status, m.TotalLaborHours, m.EstimatedCost, m.ActualCost, m.CompletionNotes, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	if len(m.Tasks) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM mro_work_order_tasks WHERE work_order_id = $1`, id)
		for _, item := range m.Tasks {
			_, err := tx.Exec(ctx, `INSERT INTO mro_work_order_tasks (work_order_id, step_number, task_description, estimated_minutes, actual_minutes, is_completed, completed_by_user_id, notes) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, id, item.StepNumber, item.TaskDescription, item.EstimatedMinutes, item.ActualMinutes, item.IsCompleted, item.CompletedByUserId, item.Notes)
			if err != nil {
				return err
			}
		}
	}

	if len(m.Items) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM mro_work_order_items WHERE work_order_id = $1`, id)
		for _, item := range m.Items {
			_, err := tx.Exec(ctx, `INSERT INTO mro_work_order_items (work_order_id, material_id, quantity_required, quantity_issued, unit_cost, total_cost, is_critical_spare, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`, id, item.MaterialId, item.QuantityRequired, item.QuantityIssued, item.UnitCost, item.TotalCost, item.IsCriticalSpare, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes work_order
func (r *WorkOrderRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE mro_work_orders SET deleted_at = CURRENT_TIMESTAMP WHERE work_order_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
