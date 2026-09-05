package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type InvoiceRepository struct {
	DB *pgxpool.Pool
}

func NewInvoiceRepository(db *pgxpool.Pool) *InvoiceRepository {
	return &InvoiceRepository{DB: db}
}

// Get retrieves a single invoice by invoice_id
func (r *InvoiceRepository) Get(ctx context.Context, id string) (*model.Invoice, error) {
	query := `SELECT t.invoice_id, t.invoice_number, t.vendor_id, COALESCE(j_vnd.vendor_name, ''), t.contract_id, COALESCE(j_ctr.contract_number, ''), t.po_id, COALESCE(j_po.po_number, ''), t.invoice_date, t.due_date, t.tax_invoice_number, t.subtotal, t.tax_amount, t.total_amount, t.verification_status, t.verified_by_user_id, COALESCE(j_usr.full_name, ''), t.payment_status, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM fin_invoices t
	LEFT JOIN proc_vendors j_vnd ON j_vnd.vendor_id = t.vendor_id
	LEFT JOIN proc_contracts j_ctr ON j_ctr.contract_id = t.contract_id
	LEFT JOIN proc_purchase_orders j_po ON j_po.po_id = t.po_id
	LEFT JOIN sys_users j_usr ON j_usr.user_id = t.verified_by_user_id
	WHERE t.invoice_id = $1 AND t.deleted_at IS NULL`

	var m model.Invoice
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.InvoiceId, &m.InvoiceNumber, &m.VendorId, &m.VendorName, &m.ContractId, &m.ContractNumber, &m.PoId, &m.PoNumber, &m.InvoiceDate, &m.DueDate, &m.TaxInvoiceNumber, &m.Subtotal, &m.TaxAmount, &m.TotalAmount, &m.VerificationStatus, &m.VerifiedByUserId, &m.VerifierName, &m.PaymentStatus, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated invoice records
func (r *InvoiceRepository) List(ctx context.Context, opts model.ListOptions) ([]model.Invoice, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "t.deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(invoice_number ILIKE $%[1]d OR tax_invoice_number ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM fin_invoices t WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf(`SELECT t.invoice_id, t.invoice_number, t.vendor_id, COALESCE(j_vnd.vendor_name, ''), t.contract_id, COALESCE(j_ctr.contract_number, ''), t.po_id, COALESCE(j_po.po_number, ''), t.invoice_date, t.due_date, t.tax_invoice_number, t.subtotal, t.tax_amount, t.total_amount, t.verification_status, t.verified_by_user_id, COALESCE(j_usr.full_name, ''), t.payment_status, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM fin_invoices t
	LEFT JOIN proc_vendors j_vnd ON j_vnd.vendor_id = t.vendor_id
	LEFT JOIN proc_contracts j_ctr ON j_ctr.contract_id = t.contract_id
	LEFT JOIN proc_purchase_orders j_po ON j_po.po_id = t.po_id
	LEFT JOIN sys_users j_usr ON j_usr.user_id = t.verified_by_user_id
	WHERE %s ORDER BY t.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.Invoice
	for rows.Next() {
		var m model.Invoice
		if err := rows.Scan(&m.InvoiceId, &m.InvoiceNumber, &m.VendorId, &m.VendorName, &m.ContractId, &m.ContractNumber, &m.PoId, &m.PoNumber, &m.InvoiceDate, &m.DueDate, &m.TaxInvoiceNumber, &m.Subtotal, &m.TaxAmount, &m.TotalAmount, &m.VerificationStatus, &m.VerifiedByUserId, &m.VerifierName, &m.PaymentStatus, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new invoice with optional child items
func (r *InvoiceRepository) Create(ctx context.Context, m *model.Invoice) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO fin_invoices (invoice_number, vendor_id, contract_id, po_id, invoice_date, due_date, tax_invoice_number, subtotal, tax_amount, verification_status, verified_by_user_id, payment_status, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10::invoice_verification_status_type, $11, $12::invoice_payment_status_type, $13, $14, $15) RETURNING invoice_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.InvoiceNumber, m.VendorId, m.ContractId, m.PoId, m.InvoiceDate, m.DueDate, m.TaxInvoiceNumber, m.Subtotal, m.TaxAmount, m.VerificationStatus, m.VerifiedByUserId, m.PaymentStatus, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing invoice
func (r *InvoiceRepository) Update(ctx context.Context, id string, m *model.Invoice) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE fin_invoices SET invoice_number = $1, vendor_id = $2, contract_id = $3, po_id = $4, invoice_date = $5, due_date = $6, tax_invoice_number = $7, subtotal = $8, tax_amount = $9, verification_status = $10::invoice_verification_status_type, verified_by_user_id = $11, payment_status = $12::invoice_payment_status_type, updated_by = $13, updated_at = CURRENT_TIMESTAMP WHERE invoice_id = $14`
	_, err = tx.Exec(ctx, updateQuery, m.InvoiceNumber, m.VendorId, m.ContractId, m.PoId, m.InvoiceDate, m.DueDate, m.TaxInvoiceNumber, m.Subtotal, m.TaxAmount, m.VerificationStatus, m.VerifiedByUserId, m.PaymentStatus, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes invoice
func (r *InvoiceRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE fin_invoices SET deleted_at = CURRENT_TIMESTAMP WHERE invoice_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
