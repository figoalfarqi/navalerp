package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/helper"
	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type JournalEntryRepository struct {
	DB *pgxpool.Pool
}

func NewJournalEntryRepository(db *pgxpool.Pool) *JournalEntryRepository {
	return &JournalEntryRepository{DB: db}
}

// Get retrieves a single journal_entry by journal_id
func (r *JournalEntryRepository) Get(ctx context.Context, id string) (*model.JournalEntry, error) {
	query := `SELECT journal_id, entry_number, entry_date, description, source_module, source_reference_id, is_posted, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM fin_journal_entries WHERE journal_id = $1 AND deleted_at IS NULL`

	var m model.JournalEntry
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.JournalId, &m.EntryNumber, &m.EntryDate, &m.Description, &m.SourceModule, &m.SourceReferenceId, &m.IsPosted, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	// Load Lines
	childRowsLines, err := r.DB.Query(ctx, `SELECT l.line_id, l.journal_id, l.account_id, COALESCE(coa.account_code, ''), COALESCE(coa.account_name, ''), l.debit, l.credit, l.memo, l.created_at FROM fin_journal_lines l LEFT JOIN fin_chart_of_accounts coa ON coa.account_id = l.account_id WHERE l.journal_id = $1`, id)
	if err == nil {
		defer childRowsLines.Close()
		for childRowsLines.Next() {
			var item model.JournalLines
			if err := childRowsLines.Scan(&item.LineId, &item.JournalId, &item.AccountId, &item.AccountCode, &item.AccountName, &item.Debit, &item.Credit, &item.Memo, &item.CreatedAt); err == nil {
				m.Lines = append(m.Lines, item)
			}
		}
	}

	return &m, nil
}

// List retrieves paginated journal_entry records
func (r *JournalEntryRepository) List(ctx context.Context, opts model.ListOptions) ([]model.JournalEntry, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(entry_number ILIKE $%[1]d OR description ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM fin_journal_entries WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 {
		limit = 10
	}
	offset := opts.Offset
	if offset < 0 {
		offset = 0
	}

	listQuery := fmt.Sprintf("SELECT journal_id, entry_number, entry_date, description, source_module, source_reference_id, is_posted, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM fin_journal_entries WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.JournalEntry
	for rows.Next() {
		var m model.JournalEntry
		if err := rows.Scan(&m.JournalId, &m.EntryNumber, &m.EntryDate, &m.Description, &m.SourceModule, &m.SourceReferenceId, &m.IsPosted, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new journal_entry with optional child items
func (r *JournalEntryRepository) Create(ctx context.Context, m *model.JournalEntry) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	if m.EntryNumber == "" {
		m.EntryNumber, _ = helper.GenerateNextNumber(ctx, r.DB, "fin_journal_entries", "entry_number", "JRN")
	}

	insertQuery := `INSERT INTO fin_journal_entries (entry_number, entry_date, description, source_module, source_reference_id, is_posted, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4::journal_source_type, $5, $6, $7, $8, $9) RETURNING journal_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.EntryNumber, m.EntryDate, m.Description, m.SourceModule, m.SourceReferenceId, m.IsPosted, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.Lines {
		_, err := tx.Exec(ctx, `INSERT INTO fin_journal_lines (journal_id, account_id, debit, credit, memo) VALUES ($1, $2, $3, $4, $5)`, newID, item.AccountId, item.Debit, item.Credit, item.Memo)
		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing journal_entry
func (r *JournalEntryRepository) Update(ctx context.Context, id string, m *model.JournalEntry) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE fin_journal_entries SET entry_number = $1, entry_date = $2, description = $3, source_module = $4::journal_source_type, source_reference_id = $5, is_posted = $6, updated_by = $7, updated_at = CURRENT_TIMESTAMP WHERE journal_id = $8`
	_, err = tx.Exec(ctx, updateQuery, m.EntryNumber, m.EntryDate, m.Description, m.SourceModule, m.SourceReferenceId, m.IsPosted, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	if m.Lines != nil {
		_, _ = tx.Exec(ctx, `DELETE FROM fin_journal_lines WHERE journal_id = $1`, id)
		for _, item := range m.Lines {
			_, err := tx.Exec(ctx, `INSERT INTO fin_journal_lines (journal_id, account_id, debit, credit, memo) VALUES ($1, $2, $3, $4, $5)`, id, item.AccountId, item.Debit, item.Credit, item.Memo)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes journal_entry
func (r *JournalEntryRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE fin_journal_entries SET deleted_at = CURRENT_TIMESTAMP WHERE journal_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
