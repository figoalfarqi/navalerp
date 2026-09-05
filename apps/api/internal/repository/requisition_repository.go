package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RequisitionRepository struct {
	DB *pgxpool.Pool
}

func NewRequisitionRepository(db *pgxpool.Pool) *RequisitionRepository {
	return &RequisitionRepository{DB: db}
}

// Get retrieves a single requisition by requisition_id
func (r *RequisitionRepository) Get(ctx context.Context, id string) (*model.Requisition, error) {
	query := `SELECT t.requisition_id, t.requisition_number, t.origin_unit_id, COALESCE(j_unit.unit_name, ''), t.work_order_id, t.priority, t.requested_date, t.required_by_date, t.approval_status, t.approved_by_user_id, COALESCE(j_usr.full_name, ''), t.approved_at, t.total_estimated_cost, t.justification, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM proc_requisitions t
	LEFT JOIN org_units j_unit ON j_unit.unit_id = t.origin_unit_id
	LEFT JOIN sys_users j_usr ON j_usr.user_id = t.approved_by_user_id
	WHERE t.requisition_id = $1 AND t.deleted_at IS NULL`

	var m model.Requisition
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.RequisitionId, &m.RequisitionNumber, &m.OriginUnitId, &m.UnitName, &m.WorkOrderId, &m.Priority, &m.RequestedDate, &m.RequiredByDate, &m.ApprovalStatus, &m.ApprovedByUserId, &m.ApproverName, &m.ApprovedAt, &m.TotalEstimatedCost, &m.Justification, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	// Load Items
	childRowsItems, err := r.DB.Query(ctx, `SELECT req_item_id, requisition_id, material_id, quantity, estimated_unit_price, estimated_total_price, notes, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM proc_requisition_items WHERE requisition_id = $1`, id)
	if err == nil {
		defer childRowsItems.Close()
		for childRowsItems.Next() {
			var item model.RequisitionItems
			if err := childRowsItems.Scan(&item.ReqItemId, &item.RequisitionId, &item.MaterialId, &item.Quantity, &item.EstimatedUnitPrice, &item.EstimatedTotalPrice, &item.Notes, &item.CreatedBy, &item.UpdatedBy, &item.DeletedBy, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt); err == nil {
				m.Items = append(m.Items, item)
			}
		}
	}

	return &m, nil
}

// List retrieves paginated requisition records
func (r *RequisitionRepository) List(ctx context.Context, opts model.ListOptions) ([]model.Requisition, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "t.deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(requisition_number ILIKE $%[1]d OR justification ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM proc_requisitions t WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf(`SELECT t.requisition_id, t.requisition_number, t.origin_unit_id, COALESCE(j_unit.unit_name, ''), t.work_order_id, t.priority, t.requested_date, t.required_by_date, t.approval_status, t.approved_by_user_id, COALESCE(j_usr.full_name, ''), t.approved_at, t.total_estimated_cost, t.justification, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM proc_requisitions t
	LEFT JOIN org_units j_unit ON j_unit.unit_id = t.origin_unit_id
	LEFT JOIN sys_users j_usr ON j_usr.user_id = t.approved_by_user_id
	WHERE %s ORDER BY t.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.Requisition
	for rows.Next() {
		var m model.Requisition
		if err := rows.Scan(&m.RequisitionId, &m.RequisitionNumber, &m.OriginUnitId, &m.UnitName, &m.WorkOrderId, &m.Priority, &m.RequestedDate, &m.RequiredByDate, &m.ApprovalStatus, &m.ApprovedByUserId, &m.ApproverName, &m.ApprovedAt, &m.TotalEstimatedCost, &m.Justification, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new requisition with optional child items
func (r *RequisitionRepository) Create(ctx context.Context, m *model.Requisition) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO proc_requisitions (requisition_number, origin_unit_id, work_order_id, priority, requested_date, required_by_date, approval_status, approved_by_user_id, approved_at, total_estimated_cost, justification, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4::requisition_priority_type, $5, $6, $7::requisition_approval_status_type, $8, $9, $10, $11, $12, $13, $14) RETURNING requisition_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.RequisitionNumber, m.OriginUnitId, m.WorkOrderId, m.Priority, m.RequestedDate, m.RequiredByDate, m.ApprovalStatus, m.ApprovedByUserId, m.ApprovedAt, m.TotalEstimatedCost, m.Justification, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.Items {
		_, err := tx.Exec(ctx, `INSERT INTO proc_requisition_items (requisition_id, material_id, quantity, estimated_unit_price, notes, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, newID, item.MaterialId, item.Quantity, item.EstimatedUnitPrice, item.Notes, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing requisition
func (r *RequisitionRepository) Update(ctx context.Context, id string, m *model.Requisition) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE proc_requisitions SET requisition_number = $1, origin_unit_id = $2, work_order_id = $3, priority = $4::requisition_priority_type, requested_date = $5, required_by_date = $6, approval_status = $7::requisition_approval_status_type, approved_by_user_id = $8, approved_at = $9, total_estimated_cost = $10, justification = $11, updated_by = $12, updated_at = CURRENT_TIMESTAMP WHERE requisition_id = $13`
	_, err = tx.Exec(ctx, updateQuery, m.RequisitionNumber, m.OriginUnitId, m.WorkOrderId, m.Priority, m.RequestedDate, m.RequiredByDate, m.ApprovalStatus, m.ApprovedByUserId, m.ApprovedAt, m.TotalEstimatedCost, m.Justification, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	if len(m.Items) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM proc_requisition_items WHERE requisition_id = $1`, id)
		for _, item := range m.Items {
			_, err := tx.Exec(ctx, `INSERT INTO proc_requisition_items (requisition_id, material_id, quantity, estimated_unit_price, notes, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, id, item.MaterialId, item.Quantity, item.EstimatedUnitPrice, item.Notes, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes requisition
func (r *RequisitionRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE proc_requisitions SET deleted_at = CURRENT_TIMESTAMP WHERE requisition_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
