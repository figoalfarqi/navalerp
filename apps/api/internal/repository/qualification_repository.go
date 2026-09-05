package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type QualificationRepository struct {
	DB *pgxpool.Pool
}

func NewQualificationRepository(db *pgxpool.Pool) *QualificationRepository {
	return &QualificationRepository{DB: db}
}

// Get retrieves a single qualification by qualification_id
func (r *QualificationRepository) Get(ctx context.Context, id string) (*model.Qualification, error) {
	query := `SELECT qualification_id, qualification_code, qualification_name, qualification_category, issuing_institution, validity_years, description, created_at FROM hcm_qualifications WHERE qualification_id = $1`

	var m model.Qualification
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.QualificationId, &m.QualificationCode, &m.QualificationName, &m.QualificationCategory, &m.IssuingInstitution, &m.ValidityYears, &m.Description, &m.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated qualification records
func (r *QualificationRepository) List(ctx context.Context, opts model.ListOptions) ([]model.Qualification, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(qualification_code ILIKE $%[1]d OR qualification_name ILIKE $%[1]d OR issuing_institution ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM hcm_qualifications WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT qualification_id, qualification_code, qualification_name, qualification_category, issuing_institution, validity_years, description, created_at FROM hcm_qualifications WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.Qualification
	for rows.Next() {
		var m model.Qualification
		if err := rows.Scan(&m.QualificationId, &m.QualificationCode, &m.QualificationName, &m.QualificationCategory, &m.IssuingInstitution, &m.ValidityYears, &m.Description, &m.CreatedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new qualification with optional child items
func (r *QualificationRepository) Create(ctx context.Context, m *model.Qualification) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO hcm_qualifications (qualification_code, qualification_name, qualification_category, issuing_institution, validity_years, description) VALUES ($1, $2, $3::qualification_category_type, $4, $5, $6) RETURNING qualification_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.QualificationCode, m.QualificationName, m.QualificationCategory, m.IssuingInstitution, m.ValidityYears, m.Description).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing qualification
func (r *QualificationRepository) Update(ctx context.Context, id string, m *model.Qualification) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE hcm_qualifications SET qualification_code = $1, qualification_name = $2, qualification_category = $3::qualification_category_type, issuing_institution = $4, validity_years = $5, description = $6 WHERE qualification_id = $7`
	_, err = tx.Exec(ctx, updateQuery, m.QualificationCode, m.QualificationName, m.QualificationCategory, m.IssuingInstitution, m.ValidityYears, m.Description, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes qualification
func (r *QualificationRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM hcm_qualifications WHERE qualification_id = $1`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
