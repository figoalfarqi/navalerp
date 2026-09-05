package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DockingRecordRepository struct {
	DB *pgxpool.Pool
}

func NewDockingRecordRepository(db *pgxpool.Pool) *DockingRecordRepository {
	return &DockingRecordRepository{DB: db}
}

// Get retrieves a single docking_record by docking_id
func (r *DockingRecordRepository) Get(ctx context.Context, id string) (*model.DockingRecord, error) {
	query := `SELECT t.docking_id, t.ship_id, COALESCE(j_ship.ship_name, ''), t.shipyard_name, t.docking_type, t.entry_date, t.scheduled_exit_date, t.actual_exit_date, t.sea_trial_passed, t.classification_surveyor, t.certificate_number, t.total_docking_cost, t.docking_summary, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM mro_docking_records t
	LEFT JOIN mro_ships j_ship ON j_ship.ship_id = t.ship_id
	WHERE t.docking_id = $1 AND t.deleted_at IS NULL`

	var m model.DockingRecord
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.DockingId, &m.ShipId, &m.ShipName, &m.ShipyardName, &m.DockingType, &m.EntryDate, &m.ScheduledExitDate, &m.ActualExitDate, &m.SeaTrialPassed, &m.ClassificationSurveyor, &m.CertificateNumber, &m.TotalDockingCost, &m.DockingSummary, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated docking_record records
func (r *DockingRecordRepository) List(ctx context.Context, opts model.ListOptions) ([]model.DockingRecord, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "t.deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(shipyard_name ILIKE $%[1]d OR classification_surveyor ILIKE $%[1]d OR certificate_number ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM mro_docking_records t WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf(`SELECT t.docking_id, t.ship_id, COALESCE(j_ship.ship_name, ''), t.shipyard_name, t.docking_type, t.entry_date, t.scheduled_exit_date, t.actual_exit_date, t.sea_trial_passed, t.classification_surveyor, t.certificate_number, t.total_docking_cost, t.docking_summary, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM mro_docking_records t
	LEFT JOIN mro_ships j_ship ON j_ship.ship_id = t.ship_id
	WHERE %s ORDER BY t.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.DockingRecord
	for rows.Next() {
		var m model.DockingRecord
		if err := rows.Scan(&m.DockingId, &m.ShipId, &m.ShipName, &m.ShipyardName, &m.DockingType, &m.EntryDate, &m.ScheduledExitDate, &m.ActualExitDate, &m.SeaTrialPassed, &m.ClassificationSurveyor, &m.CertificateNumber, &m.TotalDockingCost, &m.DockingSummary, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new docking_record with optional child items
func (r *DockingRecordRepository) Create(ctx context.Context, m *model.DockingRecord) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO mro_docking_records (ship_id, shipyard_name, docking_type, entry_date, scheduled_exit_date, actual_exit_date, sea_trial_passed, classification_surveyor, certificate_number, total_docking_cost, docking_summary, created_by, updated_by, deleted_by) VALUES ($1, $2, $3::docking_type_enum, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING docking_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.ShipId, m.ShipyardName, m.DockingType, m.EntryDate, m.ScheduledExitDate, m.ActualExitDate, m.SeaTrialPassed, m.ClassificationSurveyor, m.CertificateNumber, m.TotalDockingCost, m.DockingSummary, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing docking_record
func (r *DockingRecordRepository) Update(ctx context.Context, id string, m *model.DockingRecord) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE mro_docking_records SET ship_id = $1, shipyard_name = $2, docking_type = $3::docking_type_enum, entry_date = $4, scheduled_exit_date = $5, actual_exit_date = $6, sea_trial_passed = $7, classification_surveyor = $8, certificate_number = $9, total_docking_cost = $10, docking_summary = $11, updated_by = $12, updated_at = CURRENT_TIMESTAMP WHERE docking_id = $13`
	_, err = tx.Exec(ctx, updateQuery, m.ShipId, m.ShipyardName, m.DockingType, m.EntryDate, m.ScheduledExitDate, m.ActualExitDate, m.SeaTrialPassed, m.ClassificationSurveyor, m.CertificateNumber, m.TotalDockingCost, m.DockingSummary, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes docking_record
func (r *DockingRecordRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE mro_docking_records SET deleted_at = CURRENT_TIMESTAMP WHERE docking_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
