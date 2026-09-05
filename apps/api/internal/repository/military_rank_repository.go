package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MilitaryRankRepository struct {
	DB *pgxpool.Pool
}

func NewMilitaryRankRepository(db *pgxpool.Pool) *MilitaryRankRepository {
	return &MilitaryRankRepository{DB: db}
}

// Get retrieves a single military_rank by rank_id
func (r *MilitaryRankRepository) Get(ctx context.Context, id string) (*model.MilitaryRank, error) {
	query := `SELECT rank_id, rank_code, rank_name, rank_category, nato_rank_code, seniority_order, is_active, created_at FROM hcm_ranks WHERE rank_id = $1`

	var m model.MilitaryRank
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.RankId, &m.RankCode, &m.RankName, &m.RankCategory, &m.NatoRankCode, &m.SeniorityOrder, &m.IsActive, &m.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated military_rank records
func (r *MilitaryRankRepository) List(ctx context.Context, opts model.ListOptions) ([]model.MilitaryRank, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(rank_code ILIKE $%[1]d OR rank_name ILIKE $%[1]d OR nato_rank_code ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM hcm_ranks WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT rank_id, rank_code, rank_name, rank_category, nato_rank_code, seniority_order, is_active, created_at FROM hcm_ranks WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.MilitaryRank
	for rows.Next() {
		var m model.MilitaryRank
		if err := rows.Scan(&m.RankId, &m.RankCode, &m.RankName, &m.RankCategory, &m.NatoRankCode, &m.SeniorityOrder, &m.IsActive, &m.CreatedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new military_rank with optional child items
func (r *MilitaryRankRepository) Create(ctx context.Context, m *model.MilitaryRank) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO hcm_ranks (rank_code, rank_name, rank_category, nato_rank_code, seniority_order, is_active) VALUES ($1, $2, $3::rank_category_type, $4, $5, $6) RETURNING rank_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.RankCode, m.RankName, m.RankCategory, m.NatoRankCode, m.SeniorityOrder, m.IsActive).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing military_rank
func (r *MilitaryRankRepository) Update(ctx context.Context, id string, m *model.MilitaryRank) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE hcm_ranks SET rank_code = $1, rank_name = $2, rank_category = $3::rank_category_type, nato_rank_code = $4, seniority_order = $5, is_active = $6 WHERE rank_id = $7`
	_, err = tx.Exec(ctx, updateQuery, m.RankCode, m.RankName, m.RankCategory, m.NatoRankCode, m.SeniorityOrder, m.IsActive, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes military_rank
func (r *MilitaryRankRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM hcm_ranks WHERE rank_id = $1`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
