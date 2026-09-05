package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TheaterRepository struct {
	DB *pgxpool.Pool
}

func NewTheaterRepository(db *pgxpool.Pool) *TheaterRepository {
	return &TheaterRepository{DB: db}
}

// Get retrieves a single theater by theater_id
func (r *TheaterRepository) Get(ctx context.Context, id string) (*model.Theater, error) {
	query := `SELECT theater_id, theater_code, theater_name, responsible_command_unit_id, threat_level, description, created_at FROM ops_theaters WHERE theater_id = $1`

	var m model.Theater
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.TheaterId, &m.TheaterCode, &m.TheaterName, &m.ResponsibleCommandUnitId, &m.ThreatLevel, &m.Description, &m.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated theater records
func (r *TheaterRepository) List(ctx context.Context, opts model.ListOptions) ([]model.Theater, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(theater_code ILIKE $%[1]d OR theater_name ILIKE $%[1]d OR description ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM ops_theaters WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT theater_id, theater_code, theater_name, responsible_command_unit_id, threat_level, description, created_at FROM ops_theaters WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.Theater
	for rows.Next() {
		var m model.Theater
		if err := rows.Scan(&m.TheaterId, &m.TheaterCode, &m.TheaterName, &m.ResponsibleCommandUnitId, &m.ThreatLevel, &m.Description, &m.CreatedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new theater with optional child items
func (r *TheaterRepository) Create(ctx context.Context, m *model.Theater) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO ops_theaters (theater_code, theater_name, responsible_command_unit_id, threat_level, description) VALUES ($1, $2, $3, $4::ops_defcon_level_type, $5) RETURNING theater_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.TheaterCode, m.TheaterName, m.ResponsibleCommandUnitId, m.ThreatLevel, m.Description).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing theater
func (r *TheaterRepository) Update(ctx context.Context, id string, m *model.Theater) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE ops_theaters SET theater_code = $1, theater_name = $2, responsible_command_unit_id = $3, threat_level = $4::ops_defcon_level_type, description = $5 WHERE theater_id = $6`
	_, err = tx.Exec(ctx, updateQuery, m.TheaterCode, m.TheaterName, m.ResponsibleCommandUnitId, m.ThreatLevel, m.Description, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes theater
func (r *TheaterRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM ops_theaters WHERE theater_id = $1`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
