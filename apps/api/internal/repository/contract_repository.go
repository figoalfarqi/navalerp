package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ContractRepository struct {
	DB *pgxpool.Pool
}

func NewContractRepository(db *pgxpool.Pool) *ContractRepository {
	return &ContractRepository{DB: db}
}

// Get retrieves a single contract by contract_id
func (r *ContractRepository) Get(ctx context.Context, id string) (*model.Contract, error) {
	query := `SELECT contract_id, tender_id, contract_number, vendor_id, contract_title, contract_value, currency, start_date, end_date, procurement_method, warranty_period_months, tot_clause_summary, status, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM proc_contracts WHERE contract_id = $1 AND deleted_at IS NULL`

	var m model.Contract
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.ContractId, &m.TenderId, &m.ContractNumber, &m.VendorId, &m.ContractTitle, &m.ContractValue, &m.Currency, &m.StartDate, &m.EndDate, &m.ProcurementMethod, &m.WarrantyPeriodMonths, &m.TotClauseSummary, &m.Status, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	// Load Amendments
	childRowsAmendments, err := r.DB.Query(ctx, `SELECT amendment_id, contract_id, amendment_number, amendment_date, description, additional_value, extended_end_date, created_by, created_at FROM proc_contract_amendments WHERE contract_id = $1`, id)
	if err == nil {
		defer childRowsAmendments.Close()
		for childRowsAmendments.Next() {
			var item model.ContractAmendments
			if err := childRowsAmendments.Scan(&item.AmendmentId, &item.ContractId, &item.AmendmentNumber, &item.AmendmentDate, &item.Description, &item.AdditionalValue, &item.ExtendedEndDate, &item.CreatedBy, &item.CreatedAt); err == nil {
				m.Amendments = append(m.Amendments, item)
			}
		}
	}

	return &m, nil
}

// List retrieves paginated contract records
func (r *ContractRepository) List(ctx context.Context, opts model.ListOptions) ([]model.Contract, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(contract_number ILIKE $%[1]d OR contract_title ILIKE $%[1]d OR currency ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM proc_contracts WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT contract_id, tender_id, contract_number, vendor_id, contract_title, contract_value, currency, start_date, end_date, procurement_method, warranty_period_months, tot_clause_summary, status, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM proc_contracts WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.Contract
	for rows.Next() {
		var m model.Contract
		if err := rows.Scan(&m.ContractId, &m.TenderId, &m.ContractNumber, &m.VendorId, &m.ContractTitle, &m.ContractValue, &m.Currency, &m.StartDate, &m.EndDate, &m.ProcurementMethod, &m.WarrantyPeriodMonths, &m.TotClauseSummary, &m.Status, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new contract with optional child items
func (r *ContractRepository) Create(ctx context.Context, m *model.Contract) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO proc_contracts (tender_id, contract_number, vendor_id, contract_title, contract_value, currency, start_date, end_date, procurement_method, warranty_period_months, tot_clause_summary, status, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9::procurement_method_type, $10, $11, $12::contract_status_type, $13, $14, $15) RETURNING contract_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.TenderId, m.ContractNumber, m.VendorId, m.ContractTitle, m.ContractValue, m.Currency, m.StartDate, m.EndDate, m.ProcurementMethod, m.WarrantyPeriodMonths, m.TotClauseSummary, m.Status, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.Amendments {
		_, err := tx.Exec(ctx, `INSERT INTO proc_contract_amendments (contract_id, amendment_number, amendment_date, description, additional_value, extended_end_date, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7)`, newID, item.AmendmentNumber, item.AmendmentDate, item.Description, item.AdditionalValue, item.ExtendedEndDate, item.CreatedBy)
		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing contract
func (r *ContractRepository) Update(ctx context.Context, id string, m *model.Contract) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE proc_contracts SET tender_id = $1, contract_number = $2, vendor_id = $3, contract_title = $4, contract_value = $5, currency = $6, start_date = $7, end_date = $8, procurement_method = $9::procurement_method_type, warranty_period_months = $10, tot_clause_summary = $11, status = $12::contract_status_type, updated_by = $13, updated_at = CURRENT_TIMESTAMP WHERE contract_id = $14`
	_, err = tx.Exec(ctx, updateQuery, m.TenderId, m.ContractNumber, m.VendorId, m.ContractTitle, m.ContractValue, m.Currency, m.StartDate, m.EndDate, m.ProcurementMethod, m.WarrantyPeriodMonths, m.TotClauseSummary, m.Status, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	if len(m.Amendments) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM proc_contract_amendments WHERE contract_id = $1`, id)
		for _, item := range m.Amendments {
			_, err := tx.Exec(ctx, `INSERT INTO proc_contract_amendments (contract_id, amendment_number, amendment_date, description, additional_value, extended_end_date, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7)`, id, item.AmendmentNumber, item.AmendmentDate, item.Description, item.AdditionalValue, item.ExtendedEndDate, item.CreatedBy)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes contract
func (r *ContractRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE proc_contracts SET deleted_at = CURRENT_TIMESTAMP WHERE contract_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
