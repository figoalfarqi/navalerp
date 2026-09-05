package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ChartOfAccountRepository struct {
	DB *pgxpool.Pool
}

func NewChartOfAccountRepository(db *pgxpool.Pool) *ChartOfAccountRepository {
	return &ChartOfAccountRepository{DB: db}
}

// Get retrieves a single chart_of_account by account_id
func (r *ChartOfAccountRepository) Get(ctx context.Context, id string) (*model.ChartOfAccount, error) {
	query := `SELECT t.account_id, t.account_code, t.account_name, t.account_type, t.parent_account_id, COALESCE(j_pacc.account_name, ''), t.is_active, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM fin_chart_of_accounts t
	LEFT JOIN fin_chart_of_accounts j_pacc ON j_pacc.account_id = t.parent_account_id
	WHERE t.account_id = $1 AND t.deleted_at IS NULL`

	var m model.ChartOfAccount
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.AccountId, &m.AccountCode, &m.AccountName, &m.AccountType, &m.ParentAccountId, &m.ParentAccountName, &m.IsActive, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated chart_of_account records
func (r *ChartOfAccountRepository) List(ctx context.Context, opts model.ListOptions) ([]model.ChartOfAccount, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "t.deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(account_code ILIKE $%[1]d OR account_name ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM fin_chart_of_accounts t WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf(`SELECT t.account_id, t.account_code, t.account_name, t.account_type, t.parent_account_id, COALESCE(j_pacc.account_name, ''), t.is_active, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM fin_chart_of_accounts t
	LEFT JOIN fin_chart_of_accounts j_pacc ON j_pacc.account_id = t.parent_account_id
	WHERE %s ORDER BY t.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.ChartOfAccount
	for rows.Next() {
		var m model.ChartOfAccount
		if err := rows.Scan(&m.AccountId, &m.AccountCode, &m.AccountName, &m.AccountType, &m.ParentAccountId, &m.ParentAccountName, &m.IsActive, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new chart_of_account with optional child items
func (r *ChartOfAccountRepository) Create(ctx context.Context, m *model.ChartOfAccount) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO fin_chart_of_accounts (account_code, account_name, account_type, parent_account_id, is_active, created_by, updated_by, deleted_by) VALUES ($1, $2, $3::coa_account_type, $4, $5, $6, $7, $8) RETURNING account_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.AccountCode, m.AccountName, m.AccountType, m.ParentAccountId, m.IsActive, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing chart_of_account
func (r *ChartOfAccountRepository) Update(ctx context.Context, id string, m *model.ChartOfAccount) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE fin_chart_of_accounts SET account_code = $1, account_name = $2, account_type = $3::coa_account_type, parent_account_id = $4, is_active = $5, updated_by = $6, updated_at = CURRENT_TIMESTAMP WHERE account_id = $7`
	_, err = tx.Exec(ctx, updateQuery, m.AccountCode, m.AccountName, m.AccountType, m.ParentAccountId, m.IsActive, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes chart_of_account
func (r *ChartOfAccountRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE fin_chart_of_accounts SET deleted_at = CURRENT_TIMESTAMP WHERE account_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
