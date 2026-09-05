package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PaymentRepository struct {
	DB *pgxpool.Pool
}

func NewPaymentRepository(db *pgxpool.Pool) *PaymentRepository {
	return &PaymentRepository{DB: db}
}

// Get retrieves a single payment by payment_id
func (r *PaymentRepository) Get(ctx context.Context, id string) (*model.Payment, error) {
	query := `SELECT payment_id, payment_reference_no, spp_number, spm_number, invoice_id, payment_date, amount_paid, payment_method, bank_source_account, authorised_by_user_id, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM fin_payments WHERE payment_id = $1 AND deleted_at IS NULL`

	var m model.Payment
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.PaymentId, &m.PaymentReferenceNo, &m.SppNumber, &m.SpmNumber, &m.InvoiceId, &m.PaymentDate, &m.AmountPaid, &m.PaymentMethod, &m.BankSourceAccount, &m.AuthorisedByUserId, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated payment records
func (r *PaymentRepository) List(ctx context.Context, opts model.ListOptions) ([]model.Payment, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(payment_reference_no ILIKE $%[1]d OR spp_number ILIKE $%[1]d OR spm_number ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM fin_payments WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT payment_id, payment_reference_no, spp_number, spm_number, invoice_id, payment_date, amount_paid, payment_method, bank_source_account, authorised_by_user_id, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM fin_payments WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.Payment
	for rows.Next() {
		var m model.Payment
		if err := rows.Scan(&m.PaymentId, &m.PaymentReferenceNo, &m.SppNumber, &m.SpmNumber, &m.InvoiceId, &m.PaymentDate, &m.AmountPaid, &m.PaymentMethod, &m.BankSourceAccount, &m.AuthorisedByUserId, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new payment with optional child items
func (r *PaymentRepository) Create(ctx context.Context, m *model.Payment) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO fin_payments (payment_reference_no, spp_number, spm_number, invoice_id, payment_date, amount_paid, payment_method, bank_source_account, authorised_by_user_id, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7::payment_method_type, $8, $9, $10, $11, $12) RETURNING payment_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.PaymentReferenceNo, m.SppNumber, m.SpmNumber, m.InvoiceId, m.PaymentDate, m.AmountPaid, m.PaymentMethod, m.BankSourceAccount, m.AuthorisedByUserId, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing payment
func (r *PaymentRepository) Update(ctx context.Context, id string, m *model.Payment) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE fin_payments SET payment_reference_no = $1, spp_number = $2, spm_number = $3, invoice_id = $4, payment_date = $5, amount_paid = $6, payment_method = $7::payment_method_type, bank_source_account = $8, authorised_by_user_id = $9, updated_by = $10, updated_at = CURRENT_TIMESTAMP WHERE payment_id = $11`
	_, err = tx.Exec(ctx, updateQuery, m.PaymentReferenceNo, m.SppNumber, m.SpmNumber, m.InvoiceId, m.PaymentDate, m.AmountPaid, m.PaymentMethod, m.BankSourceAccount, m.AuthorisedByUserId, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes payment
func (r *PaymentRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE fin_payments SET deleted_at = CURRENT_TIMESTAMP WHERE payment_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
