package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BudgetCommitmentRepository struct {
	DB *pgxpool.Pool
}

func NewBudgetCommitmentRepository(db *pgxpool.Pool) *BudgetCommitmentRepository {
	return &BudgetCommitmentRepository{DB: db}
}

// Get retrieves a single budget_commitment by commitment_id
func (r *BudgetCommitmentRepository) Get(ctx context.Context, id string) (*model.BudgetCommitment, error) {
	query := `SELECT commitment_id, commitment_number, allocation_id, contract_id, po_id, work_order_id, committed_amount, commitment_date, status, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM fin_budget_commitments WHERE commitment_id = $1 AND deleted_at IS NULL`

	var m model.BudgetCommitment
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.CommitmentId, &m.CommitmentNumber, &m.AllocationId, &m.ContractId, &m.PoId, &m.WorkOrderId, &m.CommittedAmount, &m.CommitmentDate, &m.Status, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated budget_commitment records
func (r *BudgetCommitmentRepository) List(ctx context.Context, opts model.ListOptions) ([]model.BudgetCommitment, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(commitment_number ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM fin_budget_commitments WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT commitment_id, commitment_number, allocation_id, contract_id, po_id, work_order_id, committed_amount, commitment_date, status, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM fin_budget_commitments WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.BudgetCommitment
	for rows.Next() {
		var m model.BudgetCommitment
		if err := rows.Scan(&m.CommitmentId, &m.CommitmentNumber, &m.AllocationId, &m.ContractId, &m.PoId, &m.WorkOrderId, &m.CommittedAmount, &m.CommitmentDate, &m.Status, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new budget_commitment with optional child items
func (r *BudgetCommitmentRepository) Create(ctx context.Context, m *model.BudgetCommitment) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO fin_budget_commitments (commitment_number, allocation_id, contract_id, po_id, work_order_id, committed_amount, commitment_date, status, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8::budget_commitment_status_type, $9, $10, $11) RETURNING commitment_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.CommitmentNumber, m.AllocationId, m.ContractId, m.PoId, m.WorkOrderId, m.CommittedAmount, m.CommitmentDate, m.Status, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing budget_commitment
func (r *BudgetCommitmentRepository) Update(ctx context.Context, id string, m *model.BudgetCommitment) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE fin_budget_commitments SET commitment_number = $1, allocation_id = $2, contract_id = $3, po_id = $4, work_order_id = $5, committed_amount = $6, commitment_date = $7, status = $8::budget_commitment_status_type, updated_by = $9, updated_at = CURRENT_TIMESTAMP WHERE commitment_id = $10`
	_, err = tx.Exec(ctx, updateQuery, m.CommitmentNumber, m.AllocationId, m.ContractId, m.PoId, m.WorkOrderId, m.CommittedAmount, m.CommitmentDate, m.Status, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes budget_commitment
func (r *BudgetCommitmentRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE fin_budget_commitments SET deleted_at = CURRENT_TIMESTAMP WHERE commitment_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
