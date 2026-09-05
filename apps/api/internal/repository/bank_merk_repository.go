package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/figoalfarqi/apipml/internal/helper"
	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BankMerkRepository struct {
	DB *pgxpool.Pool
}

func NewBankMerkRepository(db *pgxpool.Pool) *BankMerkRepository {
	return &BankMerkRepository{DB: db}
}

// ==================================================
// Create
// ==================================================
func (r *BankMerkRepository) Create(ctx context.Context, m *model.BankMerk) (int, error) {
	query := `
		INSERT INTO bank_merk (
			bank_merk_name,
			bank_merk_description,
			is_active,
			created_by,
			updated_by
		) VALUES ($1,$2,$3,$4,$5)
		RETURNING bank_merk_id
	`

	var id int
	err := r.DB.QueryRow(ctx, query,
		m.BankMerkName,
		m.BankMerkDescription,
		m.IsActive,
		m.CreatedBy,
		m.UpdatedBy,
	).Scan(&id)

	return id, err
}

// ==================================================
// Update (Dynamic like ClientRepository)
// ==================================================
func (r *BankMerkRepository) Update(ctx context.Context, id int, m *model.BankMerk) error {

	setIsActive := ""
	args := []interface{}{
		m.BankMerkName,        // $1
		m.BankMerkDescription, // $2
	}

	argPos := 3

	if m.IsActive != -1 { // gunakan -1 sebagai marker "jangan update"
		setIsActive = fmt.Sprintf(", is_active = $%d", argPos)
		args = append(args, m.IsActive)
		argPos++
	}

	args = append(args,
		m.UpdatedBy, // $argPos
		time.Now(),  // $argPos+1
		id,          // $argPos+2
	)

	query := fmt.Sprintf(`
		UPDATE bank_merk
		SET bank_merk_name = $1,
			bank_merk_description = $2
			%s,
			updated_by = $%d,
			updated_at = $%d
		WHERE bank_merk_id = $%d AND deleted_at IS NULL
	`,
		setIsActive,
		argPos,   // updated_by
		argPos+1, // updated_at
		argPos+2, // WHERE
	)

	_, err := r.DB.Exec(ctx, query, args...)
	return err
}

// ==================================================
// Soft Delete
// ==================================================
func (r *BankMerkRepository) SoftDelete(ctx context.Context, deletedBy, id int) error {
	query := `
		UPDATE bank_merk
		SET deleted_at = $1,
			deleted_by = $2,
			is_active = 0,
			updated_at = $1,
			updated_by = $2
		WHERE bank_merk_id = $3 AND deleted_at IS NULL
	`

	_, err := r.DB.Exec(ctx, query, time.Now(), deletedBy, id)
	return err
}

// ==================================================
// Get By ID
// ==================================================
func (r *BankMerkRepository) GetByID(ctx context.Context, id int) (*model.BankMerk, error) {
	query := `
		SELECT
			bank_merk_id,
			bank_merk_name,
			bank_merk_description,
			is_active,
			created_by,
			updated_by,
			deleted_by,
			created_at,
			updated_at,
			deleted_at
		FROM bank_merk
		WHERE bank_merk_id = $1 AND deleted_at IS NULL
	`

	var m model.BankMerk

	err := r.DB.QueryRow(ctx, query, id).Scan(
		&m.BankMerkID,
		&m.BankMerkName,
		&m.BankMerkDescription,
		&m.IsActive,
		&m.CreatedBy,
		&m.UpdatedBy,
		&m.DeletedBy,
		&m.CreatedAt,
		&m.UpdatedAt,
		&m.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return &m, nil
}

// ==================================================
// List (Cursor Pagination + Filters) — MATCH ClientRepository
// ==================================================
func (r *BankMerkRepository) List(ctx context.Context,
	cursorValue interface{}, cursorKey *int, limit int, filters map[string]string,
	orderBy, sort string) ([]model.BankMerk, error) {
	tableKey := "bank_merk_id"
	baseQuery := `
		SELECT 
			bank_merk_id,
			bank_merk_name,
			bank_merk_description,
			is_active,
			created_by,
			updated_by,
			deleted_by,
			created_at,
			updated_at,
			deleted_at
		FROM bank_merk
		WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	argPos := 1

	// cursor
	cursorClause, cursorArgs, argPos := helper.BuildCursorCondition(orderBy, tableKey, sort, cursorValue, cursorKey, argPos)
	baseQuery += cursorClause
	args = append(args, cursorArgs...)

	// LIKE filters
	likeFilters := map[string]string{
		"bank_merk_name":        "bank_merk_name ILIKE $%d",
		"bank_merk_description": "bank_merk_description ILIKE $%d",
	}

	for key, clause := range likeFilters {
		if v, ok := filters[key]; ok && v != "" {
			baseQuery += " AND " + fmt.Sprintf(clause, argPos)
			args = append(args, "%"+v+"%")
			argPos++
		}
	}

	// exact match
	exactFilters := []string{"is_active", "created_by", "updated_by"}

	for _, key := range exactFilters {
		if v, ok := filters[key]; ok && v != "" {
			baseQuery += fmt.Sprintf(" AND %s = $%d", key, argPos)
			args = append(args, v)
			argPos++
		}
	}

	// time range
	timeFilters := map[string]string{
		"created_at_after":  "created_at >= $%d",
		"created_at_before": "created_at <= $%d",
		"updated_at_after":  "updated_at >= $%d",
		"updated_at_before": "updated_at <= $%d",
	}
	for key, clause := range timeFilters {
		if v, ok := filters[key]; ok && v != "" {
			baseQuery += " AND " + fmt.Sprintf(clause, argPos)
			args = append(args, v)
			argPos++
		}
	}

	// final order + limit
	baseQuery += fmt.Sprintf(" ORDER BY %s %s, %s %s LIMIT $%d", orderBy, sort, tableKey, sort, argPos)
	args = append(args, limit)

	rows, err := r.DB.Query(ctx, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.BankMerk

	for rows.Next() {
		var m model.BankMerk

		if err := rows.Scan(
			&m.BankMerkID,
			&m.BankMerkName,
			&m.BankMerkDescription,
			&m.IsActive,
			&m.CreatedBy,
			&m.UpdatedBy,
			&m.DeletedBy,
			&m.CreatedAt,
			&m.UpdatedAt,
			&m.DeletedAt,
		); err != nil {
			return nil, err
		}

		list = append(list, m)
	}

	return list, rows.Err()
}
