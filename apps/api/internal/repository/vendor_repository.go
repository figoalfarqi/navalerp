package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type VendorRepository struct {
	DB *pgxpool.Pool
}

func NewVendorRepository(db *pgxpool.Pool) *VendorRepository {
	return &VendorRepository{DB: db}
}

// Get retrieves a single vendor by vendor_id
func (r *VendorRepository) Get(ctx context.Context, id string) (*model.Vendor, error) {
	query := `SELECT vendor_id, vendor_code, vendor_name, tax_number, security_clearance_level, defence_industry_license_no, country, contact_person, email, phone, bank_account_name, bank_account_no, bank_name, performance_rating, is_approved, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM proc_vendors WHERE vendor_id = $1 AND deleted_at IS NULL`

	var m model.Vendor
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.VendorId, &m.VendorCode, &m.VendorName, &m.TaxNumber, &m.SecurityClearanceLevel, &m.DefenceIndustryLicenseNo, &m.Country, &m.ContactPerson, &m.Email, &m.Phone, &m.BankAccountName, &m.BankAccountNo, &m.BankName, &m.PerformanceRating, &m.IsApproved, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	// Load Ratings
	childRowsRatings, err := r.DB.Query(ctx, `SELECT rating_id, vendor_id, evaluation_date, evaluator_user_id, quality_score, delivery_time_score, service_score, price_score, overall_score, remarks, created_at FROM proc_vendor_ratings WHERE vendor_id = $1`, id)
	if err == nil {
		defer childRowsRatings.Close()
		for childRowsRatings.Next() {
			var item model.VendorRatings
			if err := childRowsRatings.Scan(&item.RatingId, &item.VendorId, &item.EvaluationDate, &item.EvaluatorUserId, &item.QualityScore, &item.DeliveryTimeScore, &item.ServiceScore, &item.PriceScore, &item.OverallScore, &item.Remarks, &item.CreatedAt); err == nil {
				m.Ratings = append(m.Ratings, item)
			}
		}
	}

	return &m, nil
}

// List retrieves paginated vendor records
func (r *VendorRepository) List(ctx context.Context, opts model.ListOptions) ([]model.Vendor, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(vendor_code ILIKE $%[1]d OR vendor_name ILIKE $%[1]d OR tax_number ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM proc_vendors WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT vendor_id, vendor_code, vendor_name, tax_number, security_clearance_level, defence_industry_license_no, country, contact_person, email, phone, bank_account_name, bank_account_no, bank_name, performance_rating, is_approved, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM proc_vendors WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.Vendor
	for rows.Next() {
		var m model.Vendor
		if err := rows.Scan(&m.VendorId, &m.VendorCode, &m.VendorName, &m.TaxNumber, &m.SecurityClearanceLevel, &m.DefenceIndustryLicenseNo, &m.Country, &m.ContactPerson, &m.Email, &m.Phone, &m.BankAccountName, &m.BankAccountNo, &m.BankName, &m.PerformanceRating, &m.IsApproved, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new vendor with optional child items
func (r *VendorRepository) Create(ctx context.Context, m *model.Vendor) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO proc_vendors (vendor_code, vendor_name, tax_number, security_clearance_level, defence_industry_license_no, country, contact_person, email, phone, bank_account_name, bank_account_no, bank_name, performance_rating, is_approved, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4::security_clearance_type, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) RETURNING vendor_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.VendorCode, m.VendorName, m.TaxNumber, m.SecurityClearanceLevel, m.DefenceIndustryLicenseNo, m.Country, m.ContactPerson, m.Email, m.Phone, m.BankAccountName, m.BankAccountNo, m.BankName, m.PerformanceRating, m.IsApproved, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.Ratings {
		_, err := tx.Exec(ctx, `INSERT INTO proc_vendor_ratings (vendor_id, evaluation_date, evaluator_user_id, quality_score, delivery_time_score, service_score, price_score, overall_score, remarks) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`, newID, item.EvaluationDate, item.EvaluatorUserId, item.QualityScore, item.DeliveryTimeScore, item.ServiceScore, item.PriceScore, item.OverallScore, item.Remarks)
		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing vendor
func (r *VendorRepository) Update(ctx context.Context, id string, m *model.Vendor) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE proc_vendors SET vendor_code = $1, vendor_name = $2, tax_number = $3, security_clearance_level = $4::security_clearance_type, defence_industry_license_no = $5, country = $6, contact_person = $7, email = $8, phone = $9, bank_account_name = $10, bank_account_no = $11, bank_name = $12, performance_rating = $13, is_approved = $14, updated_by = $15, updated_at = CURRENT_TIMESTAMP WHERE vendor_id = $16`
	_, err = tx.Exec(ctx, updateQuery, m.VendorCode, m.VendorName, m.TaxNumber, m.SecurityClearanceLevel, m.DefenceIndustryLicenseNo, m.Country, m.ContactPerson, m.Email, m.Phone, m.BankAccountName, m.BankAccountNo, m.BankName, m.PerformanceRating, m.IsApproved, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	if len(m.Ratings) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM proc_vendor_ratings WHERE vendor_id = $1`, id)
		for _, item := range m.Ratings {
			_, err := tx.Exec(ctx, `INSERT INTO proc_vendor_ratings (vendor_id, evaluation_date, evaluator_user_id, quality_score, delivery_time_score, service_score, price_score, overall_score, remarks) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`, id, item.EvaluationDate, item.EvaluatorUserId, item.QualityScore, item.DeliveryTimeScore, item.ServiceScore, item.PriceScore, item.OverallScore, item.Remarks)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes vendor
func (r *VendorRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE proc_vendors SET deleted_at = CURRENT_TIMESTAMP WHERE vendor_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
