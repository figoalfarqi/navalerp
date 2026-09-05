package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DocumentRepository struct {
	DB *pgxpool.Pool
}

func NewDocumentRepository(db *pgxpool.Pool) *DocumentRepository {
	return &DocumentRepository{DB: db}
}

// Get retrieves a single document by document_id
func (r *DocumentRepository) Get(ctx context.Context, id string) (*model.Document, error) {
	query := `SELECT document_id, document_number, title, category_id, originating_unit_id, classification_level, effective_date, expiry_date, status, approved_by_user_id, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM doc_documents WHERE document_id = $1 AND deleted_at IS NULL`

	var m model.Document
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.DocumentId, &m.DocumentNumber, &m.Title, &m.CategoryId, &m.OriginatingUnitId, &m.ClassificationLevel, &m.EffectiveDate, &m.ExpiryDate, &m.Status, &m.ApprovedByUserId, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	// Load Versions
	childRowsVersions, err := r.DB.Query(ctx, `SELECT version_id, document_id, version_number, file_name, file_path, file_size_bytes, file_hash_sha256, mime_type, change_summary, uploaded_by_user_id, created_at FROM doc_document_versions WHERE document_id = $1`, id)
	if err == nil {
		defer childRowsVersions.Close()
		for childRowsVersions.Next() {
			var item model.DocumentVersions
			if err := childRowsVersions.Scan(&item.VersionId, &item.DocumentId, &item.VersionNumber, &item.FileName, &item.FilePath, &item.FileSizeBytes, &item.FileHashSha256, &item.MimeType, &item.ChangeSummary, &item.UploadedByUserId, &item.CreatedAt); err == nil {
				m.Versions = append(m.Versions, item)
			}
		}
	}

	// Load Links
	childRowsLinks, err := r.DB.Query(ctx, `SELECT link_id, document_id, entity_type, entity_id, link_purpose, created_at FROM doc_document_links WHERE document_id = $1`, id)
	if err == nil {
		defer childRowsLinks.Close()
		for childRowsLinks.Next() {
			var item model.DocumentLinks
			if err := childRowsLinks.Scan(&item.LinkId, &item.DocumentId, &item.EntityType, &item.EntityId, &item.LinkPurpose, &item.CreatedAt); err == nil {
				m.Links = append(m.Links, item)
			}
		}
	}

	return &m, nil
}

// List retrieves paginated document records
func (r *DocumentRepository) List(ctx context.Context, opts model.ListOptions) ([]model.Document, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(document_number ILIKE $%[1]d OR title ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM doc_documents WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT document_id, document_number, title, category_id, originating_unit_id, classification_level, effective_date, expiry_date, status, approved_by_user_id, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM doc_documents WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.Document
	for rows.Next() {
		var m model.Document
		if err := rows.Scan(&m.DocumentId, &m.DocumentNumber, &m.Title, &m.CategoryId, &m.OriginatingUnitId, &m.ClassificationLevel, &m.EffectiveDate, &m.ExpiryDate, &m.Status, &m.ApprovedByUserId, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new document with optional child items
func (r *DocumentRepository) Create(ctx context.Context, m *model.Document) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO doc_documents (document_number, title, category_id, originating_unit_id, classification_level, effective_date, expiry_date, status, approved_by_user_id, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5::document_confidentiality_level_type, $6, $7, $8::document_status_type, $9, $10, $11, $12) RETURNING document_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.DocumentNumber, m.Title, m.CategoryId, m.OriginatingUnitId, m.ClassificationLevel, m.EffectiveDate, m.ExpiryDate, m.Status, m.ApprovedByUserId, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.Versions {
		_, err := tx.Exec(ctx, `INSERT INTO doc_document_versions (document_id, version_number, file_name, file_path, file_size_bytes, file_hash_sha256, mime_type, change_summary, uploaded_by_user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`, newID, item.VersionNumber, item.FileName, item.FilePath, item.FileSizeBytes, item.FileHashSha256, item.MimeType, item.ChangeSummary, item.UploadedByUserId)
		if err != nil {
			return "", err
		}
	}

	for _, item := range m.Links {
		_, err := tx.Exec(ctx, `INSERT INTO doc_document_links (document_id, entity_type, entity_id, link_purpose) VALUES ($1, $2::document_linked_entity_type, $3, $4::document_link_purpose_type)`, newID, item.EntityType, item.EntityId, item.LinkPurpose)
		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing document
func (r *DocumentRepository) Update(ctx context.Context, id string, m *model.Document) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE doc_documents SET document_number = $1, title = $2, category_id = $3, originating_unit_id = $4, classification_level = $5::document_confidentiality_level_type, effective_date = $6, expiry_date = $7, status = $8::document_status_type, approved_by_user_id = $9, updated_by = $10, updated_at = CURRENT_TIMESTAMP WHERE document_id = $11`
	_, err = tx.Exec(ctx, updateQuery, m.DocumentNumber, m.Title, m.CategoryId, m.OriginatingUnitId, m.ClassificationLevel, m.EffectiveDate, m.ExpiryDate, m.Status, m.ApprovedByUserId, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	if len(m.Versions) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM doc_document_versions WHERE document_id = $1`, id)
		for _, item := range m.Versions {
			_, err := tx.Exec(ctx, `INSERT INTO doc_document_versions (document_id, version_number, file_name, file_path, file_size_bytes, file_hash_sha256, mime_type, change_summary, uploaded_by_user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`, id, item.VersionNumber, item.FileName, item.FilePath, item.FileSizeBytes, item.FileHashSha256, item.MimeType, item.ChangeSummary, item.UploadedByUserId)
			if err != nil {
				return err
			}
		}
	}

	if len(m.Links) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM doc_document_links WHERE document_id = $1`, id)
		for _, item := range m.Links {
			_, err := tx.Exec(ctx, `INSERT INTO doc_document_links (document_id, entity_type, entity_id, link_purpose) VALUES ($1, $2::document_linked_entity_type, $3, $4::document_link_purpose_type)`, id, item.EntityType, item.EntityId, item.LinkPurpose)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes document
func (r *DocumentRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE doc_documents SET deleted_at = CURRENT_TIMESTAMP WHERE document_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
