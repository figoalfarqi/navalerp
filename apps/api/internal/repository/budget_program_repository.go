package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BudgetProgramRepository struct {
	DB *pgxpool.Pool
}

func NewBudgetProgramRepository(db *pgxpool.Pool) *BudgetProgramRepository {
	return &BudgetProgramRepository{DB: db}
}

// Get retrieves a single budget_program by program_id
func (r *BudgetProgramRepository) Get(ctx context.Context, id string) (*model.BudgetProgram, error) {
	query := `SELECT t.program_id, t.fiscal_year, t.dipa_number, t.program_code, t.program_name, t.total_budget, t.responsible_unit_id, COALESCE(j_unit.unit_name, ''), t.status, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM fin_budget_programs t
	LEFT JOIN org_units j_unit ON j_unit.unit_id = t.responsible_unit_id
	WHERE t.program_id = $1 AND t.deleted_at IS NULL`

	var m model.BudgetProgram
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.ProgramId, &m.FiscalYear, &m.DipaNumber, &m.ProgramCode, &m.ProgramName, &m.TotalBudget, &m.ResponsibleUnitId, &m.ResponsibleUnitName, &m.Status, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	// Load Allocations
	childRowsAllocations, err := r.DB.Query(ctx, `SELECT allocation_id, program_id, activity_code, activity_name, target_unit_id, ship_id, account_id, allocated_amount, absorbed_amount, remaining_amount, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM fin_budget_allocations WHERE program_id = $1`, id)
	if err == nil {
		defer childRowsAllocations.Close()
		for childRowsAllocations.Next() {
			var item model.BudgetAllocations
			if err := childRowsAllocations.Scan(&item.AllocationId, &item.ProgramId, &item.ActivityCode, &item.ActivityName, &item.TargetUnitId, &item.ShipId, &item.AccountId, &item.AllocatedAmount, &item.AbsorbedAmount, &item.RemainingAmount, &item.CreatedBy, &item.UpdatedBy, &item.DeletedBy, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt); err == nil {
				m.Allocations = append(m.Allocations, item)
			}
		}
	}

	return &m, nil
}

// List retrieves paginated budget_program records
func (r *BudgetProgramRepository) List(ctx context.Context, opts model.ListOptions) ([]model.BudgetProgram, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "t.deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(dipa_number ILIKE $%[1]d OR program_code ILIKE $%[1]d OR program_name ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM fin_budget_programs t WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf(`SELECT t.program_id, t.fiscal_year, t.dipa_number, t.program_code, t.program_name, t.total_budget, t.responsible_unit_id, COALESCE(j_unit.unit_name, ''), t.status, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM fin_budget_programs t
	LEFT JOIN org_units j_unit ON j_unit.unit_id = t.responsible_unit_id
	WHERE %s ORDER BY t.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.BudgetProgram
	for rows.Next() {
		var m model.BudgetProgram
		if err := rows.Scan(&m.ProgramId, &m.FiscalYear, &m.DipaNumber, &m.ProgramCode, &m.ProgramName, &m.TotalBudget, &m.ResponsibleUnitId, &m.ResponsibleUnitName, &m.Status, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new budget_program with optional child items
func (r *BudgetProgramRepository) Create(ctx context.Context, m *model.BudgetProgram) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO fin_budget_programs (fiscal_year, dipa_number, program_code, program_name, total_budget, responsible_unit_id, status, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7::budget_program_status_type, $8, $9, $10) RETURNING program_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.FiscalYear, m.DipaNumber, m.ProgramCode, m.ProgramName, m.TotalBudget, m.ResponsibleUnitId, m.Status, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.Allocations {
		_, err := tx.Exec(ctx, `INSERT INTO fin_budget_allocations (program_id, activity_code, activity_name, target_unit_id, ship_id, account_id, allocated_amount, absorbed_amount, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, newID, item.ActivityCode, item.ActivityName, item.TargetUnitId, item.ShipId, item.AccountId, item.AllocatedAmount, item.AbsorbedAmount, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing budget_program
func (r *BudgetProgramRepository) Update(ctx context.Context, id string, m *model.BudgetProgram) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE fin_budget_programs SET fiscal_year = $1, dipa_number = $2, program_code = $3, program_name = $4, total_budget = $5, responsible_unit_id = $6, status = $7::budget_program_status_type, updated_by = $8, updated_at = CURRENT_TIMESTAMP WHERE program_id = $9`
	_, err = tx.Exec(ctx, updateQuery, m.FiscalYear, m.DipaNumber, m.ProgramCode, m.ProgramName, m.TotalBudget, m.ResponsibleUnitId, m.Status, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	if len(m.Allocations) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM fin_budget_allocations WHERE program_id = $1`, id)
		for _, item := range m.Allocations {
			_, err := tx.Exec(ctx, `INSERT INTO fin_budget_allocations (program_id, activity_code, activity_name, target_unit_id, ship_id, account_id, allocated_amount, absorbed_amount, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, id, item.ActivityCode, item.ActivityName, item.TargetUnitId, item.ShipId, item.AccountId, item.AllocatedAmount, item.AbsorbedAmount, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes budget_program
func (r *BudgetProgramRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE fin_budget_programs SET deleted_at = CURRENT_TIMESTAMP WHERE program_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
