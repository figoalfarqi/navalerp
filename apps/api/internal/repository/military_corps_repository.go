package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MilitaryCorpsRepository struct {
	DB *pgxpool.Pool
}

func NewMilitaryCorpsRepository(db *pgxpool.Pool) *MilitaryCorpsRepository {
	return &MilitaryCorpsRepository{DB: db}
}

// Get retrieves a single military_corps by corps_id
func (r *MilitaryCorpsRepository) Get(ctx context.Context, id string) (*model.MilitaryCorps, error) {
	query := `SELECT corps_id, corps_code, corps_name, description, created_at FROM hcm_corps WHERE corps_id = $1`

	var m model.MilitaryCorps
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.CorpsId, &m.CorpsCode, &m.CorpsName, &m.Description, &m.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated military_corps records
func (r *MilitaryCorpsRepository) List(ctx context.Context, opts model.ListOptions) ([]model.MilitaryCorps, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(corps_name ILIKE $%[1]d OR description ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM hcm_corps WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT corps_id, corps_code, corps_name, description, created_at FROM hcm_corps WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.MilitaryCorps
	for rows.Next() {
		var m model.MilitaryCorps
		if err := rows.Scan(&m.CorpsId, &m.CorpsCode, &m.CorpsName, &m.Description, &m.CreatedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new military_corps with optional child items
func (r *MilitaryCorpsRepository) Create(ctx context.Context, m *model.MilitaryCorps) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO hcm_corps (corps_code, corps_name, description) VALUES ($1::corps_code_type, $2, $3) RETURNING corps_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.CorpsCode, m.CorpsName, m.Description).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing military_corps
func (r *MilitaryCorpsRepository) Update(ctx context.Context, id string, m *model.MilitaryCorps) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE hcm_corps SET corps_code = $1::corps_code_type, corps_name = $2, description = $3 WHERE corps_id = $4`
	_, err = tx.Exec(ctx, updateQuery, m.CorpsCode, m.CorpsName, m.Description, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes military_corps
func (r *MilitaryCorpsRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM hcm_corps WHERE corps_id = $1`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
