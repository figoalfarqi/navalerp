package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TenderRepository struct {
	DB *pgxpool.Pool
}

func NewTenderRepository(db *pgxpool.Pool) *TenderRepository {
	return &TenderRepository{DB: db}
}

// Get retrieves a single tender by tender_id
func (r *TenderRepository) Get(ctx context.Context, id string) (*model.Tender, error) {
	query := `SELECT t.tender_id, t.tender_number, t.title, t.procurement_category, t.estimated_budget, t.procurement_method, t.start_date, t.closing_date, t.status, t.winner_vendor_id, COALESCE(j_vnd.vendor_name, ''), t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM proc_tenders t
	LEFT JOIN proc_vendors j_vnd ON j_vnd.vendor_id = t.winner_vendor_id
	WHERE t.tender_id = $1 AND t.deleted_at IS NULL`

	var m model.Tender
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.TenderId, &m.TenderNumber, &m.Title, &m.ProcurementCategory, &m.EstimatedBudget, &m.ProcurementMethod, &m.StartDate, &m.ClosingDate, &m.Status, &m.WinnerVendorId, &m.WinnerVendorName, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	// Load Bids
	childRowsBids, err := r.DB.Query(ctx, `SELECT bid_id, tender_id, vendor_id, bid_amount, submission_date, technical_score, commercial_score, is_winner, remarks, created_at FROM proc_tender_bids WHERE tender_id = $1`, id)
	if err == nil {
		defer childRowsBids.Close()
		for childRowsBids.Next() {
			var item model.TenderBids
			if err := childRowsBids.Scan(&item.BidId, &item.TenderId, &item.VendorId, &item.BidAmount, &item.SubmissionDate, &item.TechnicalScore, &item.CommercialScore, &item.IsWinner, &item.Remarks, &item.CreatedAt); err == nil {
				m.Bids = append(m.Bids, item)
			}
		}
	}

	return &m, nil
}

// List retrieves paginated tender records
func (r *TenderRepository) List(ctx context.Context, opts model.ListOptions) ([]model.Tender, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "t.deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(tender_number ILIKE $%[1]d OR title ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM proc_tenders t WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf(`SELECT t.tender_id, t.tender_number, t.title, t.procurement_category, t.estimated_budget, t.procurement_method, t.start_date, t.closing_date, t.status, t.winner_vendor_id, COALESCE(j_vnd.vendor_name, ''), t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM proc_tenders t
	LEFT JOIN proc_vendors j_vnd ON j_vnd.vendor_id = t.winner_vendor_id
	WHERE %s ORDER BY t.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.Tender
	for rows.Next() {
		var m model.Tender
		if err := rows.Scan(&m.TenderId, &m.TenderNumber, &m.Title, &m.ProcurementCategory, &m.EstimatedBudget, &m.ProcurementMethod, &m.StartDate, &m.ClosingDate, &m.Status, &m.WinnerVendorId, &m.WinnerVendorName, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new tender with optional child items
func (r *TenderRepository) Create(ctx context.Context, m *model.Tender) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO proc_tenders (tender_number, title, procurement_category, estimated_budget, procurement_method, start_date, closing_date, status, winner_vendor_id, created_by, updated_by, deleted_by) VALUES ($1, $2, $3::procurement_category_type, $4, $5::procurement_method_type, $6, $7, $8::tender_status_type, $9, $10, $11, $12) RETURNING tender_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.TenderNumber, m.Title, m.ProcurementCategory, m.EstimatedBudget, m.ProcurementMethod, m.StartDate, m.ClosingDate, m.Status, m.WinnerVendorId, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.Bids {
		_, err := tx.Exec(ctx, `INSERT INTO proc_tender_bids (tender_id, vendor_id, bid_amount, submission_date, technical_score, commercial_score, is_winner, remarks) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, newID, item.VendorId, item.BidAmount, item.SubmissionDate, item.TechnicalScore, item.CommercialScore, item.IsWinner, item.Remarks)
		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing tender
func (r *TenderRepository) Update(ctx context.Context, id string, m *model.Tender) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE proc_tenders SET tender_number = $1, title = $2, procurement_category = $3::procurement_category_type, estimated_budget = $4, procurement_method = $5::procurement_method_type, start_date = $6, closing_date = $7, status = $8::tender_status_type, winner_vendor_id = $9, updated_by = $10, updated_at = CURRENT_TIMESTAMP WHERE tender_id = $11`
	_, err = tx.Exec(ctx, updateQuery, m.TenderNumber, m.Title, m.ProcurementCategory, m.EstimatedBudget, m.ProcurementMethod, m.StartDate, m.ClosingDate, m.Status, m.WinnerVendorId, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	if len(m.Bids) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM proc_tender_bids WHERE tender_id = $1`, id)
		for _, item := range m.Bids {
			_, err := tx.Exec(ctx, `INSERT INTO proc_tender_bids (tender_id, vendor_id, bid_amount, submission_date, technical_score, commercial_score, is_winner, remarks) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, id, item.VendorId, item.BidAmount, item.SubmissionDate, item.TechnicalScore, item.CommercialScore, item.IsWinner, item.Remarks)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes tender
func (r *TenderRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE proc_tenders SET deleted_at = CURRENT_TIMESTAMP WHERE tender_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
