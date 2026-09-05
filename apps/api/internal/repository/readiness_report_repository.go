package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ReadinessReportRepository struct {
	DB *pgxpool.Pool
}

func NewReadinessReportRepository(db *pgxpool.Pool) *ReadinessReportRepository {
	return &ReadinessReportRepository{DB: db}
}

// Get retrieves a single readiness_report by snapshot_id
func (r *ReadinessReportRepository) Get(ctx context.Context, id string) (*model.ReadinessReport, error) {
	query := `SELECT snapshot_id, ship_id, snapshot_timestamp, readiness_category, mro_readiness_score, personnel_manning_score, logistics_supply_score, composite_readiness_index, remarks, created_at FROM ops_ship_readiness_snapshots WHERE snapshot_id = $1`

	var m model.ReadinessReport
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.SnapshotId, &m.ShipId, &m.SnapshotTimestamp, &m.ReadinessCategory, &m.MroReadinessScore, &m.PersonnelManningScore, &m.LogisticsSupplyScore, &m.CompositeReadinessIndex, &m.Remarks, &m.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated readiness_report records
func (r *ReadinessReportRepository) List(ctx context.Context, opts model.ListOptions) ([]model.ReadinessReport, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(remarks ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM ops_ship_readiness_snapshots WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT snapshot_id, ship_id, snapshot_timestamp, readiness_category, mro_readiness_score, personnel_manning_score, logistics_supply_score, composite_readiness_index, remarks, created_at FROM ops_ship_readiness_snapshots WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.ReadinessReport
	for rows.Next() {
		var m model.ReadinessReport
		if err := rows.Scan(&m.SnapshotId, &m.ShipId, &m.SnapshotTimestamp, &m.ReadinessCategory, &m.MroReadinessScore, &m.PersonnelManningScore, &m.LogisticsSupplyScore, &m.CompositeReadinessIndex, &m.Remarks, &m.CreatedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new readiness_report with optional child items
func (r *ReadinessReportRepository) Create(ctx context.Context, m *model.ReadinessReport) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO ops_ship_readiness_snapshots (ship_id, snapshot_timestamp, readiness_category, mro_readiness_score, personnel_manning_score, logistics_supply_score, remarks) VALUES ($1, $2, $3::readiness_category_type, $4, $5, $6, $7) RETURNING snapshot_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.ShipId, m.SnapshotTimestamp, m.ReadinessCategory, m.MroReadinessScore, m.PersonnelManningScore, m.LogisticsSupplyScore, m.Remarks).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing readiness_report
func (r *ReadinessReportRepository) Update(ctx context.Context, id string, m *model.ReadinessReport) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE ops_ship_readiness_snapshots SET ship_id = $1, snapshot_timestamp = $2, readiness_category = $3::readiness_category_type, mro_readiness_score = $4, personnel_manning_score = $5, logistics_supply_score = $6, remarks = $7 WHERE snapshot_id = $8`
	_, err = tx.Exec(ctx, updateQuery, m.ShipId, m.SnapshotTimestamp, m.ReadinessCategory, m.MroReadinessScore, m.PersonnelManningScore, m.LogisticsSupplyScore, m.Remarks, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes readiness_report
func (r *ReadinessReportRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM ops_ship_readiness_snapshots WHERE snapshot_id = $1`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
