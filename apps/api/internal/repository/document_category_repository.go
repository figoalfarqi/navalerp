package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DocumentCategoryRepository struct {
	DB *pgxpool.Pool
}

func NewDocumentCategoryRepository(db *pgxpool.Pool) *DocumentCategoryRepository {
	return &DocumentCategoryRepository{DB: db}
}

// Get retrieves a single document_category by category_id
func (r *DocumentCategoryRepository) Get(ctx context.Context, id string) (*model.DocumentCategory, error) {
	query := `SELECT category_id, category_code, category_name, retention_years, confidentiality_level, description, created_at FROM doc_categories WHERE category_id = $1`

	var m model.DocumentCategory
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.CategoryId, &m.CategoryCode, &m.CategoryName, &m.RetentionYears, &m.ConfidentialityLevel, &m.Description, &m.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated document_category records
func (r *DocumentCategoryRepository) List(ctx context.Context, opts model.ListOptions) ([]model.DocumentCategory, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(category_code ILIKE $%[1]d OR category_name ILIKE $%[1]d OR description ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM doc_categories WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT category_id, category_code, category_name, retention_years, confidentiality_level, description, created_at FROM doc_categories WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.DocumentCategory
	for rows.Next() {
		var m model.DocumentCategory
		if err := rows.Scan(&m.CategoryId, &m.CategoryCode, &m.CategoryName, &m.RetentionYears, &m.ConfidentialityLevel, &m.Description, &m.CreatedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new document_category with optional child items
func (r *DocumentCategoryRepository) Create(ctx context.Context, m *model.DocumentCategory) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO doc_categories (category_code, category_name, retention_years, confidentiality_level, description) VALUES ($1, $2, $3, $4::document_confidentiality_level_type, $5) RETURNING category_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.CategoryCode, m.CategoryName, m.RetentionYears, m.ConfidentialityLevel, m.Description).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing document_category
func (r *DocumentCategoryRepository) Update(ctx context.Context, id string, m *model.DocumentCategory) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE doc_categories SET category_code = $1, category_name = $2, retention_years = $3, confidentiality_level = $4::document_confidentiality_level_type, description = $5 WHERE category_id = $6`
	_, err = tx.Exec(ctx, updateQuery, m.CategoryCode, m.CategoryName, m.RetentionYears, m.ConfidentialityLevel, m.Description, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes document_category
func (r *DocumentCategoryRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM doc_categories WHERE category_id = $1`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
