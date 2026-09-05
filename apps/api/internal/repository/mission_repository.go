package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionRepository struct {
	DB *pgxpool.Pool
}

func NewMissionRepository(db *pgxpool.Pool) *MissionRepository {
	return &MissionRepository{DB: db}
}

// Get retrieves a single mission by mission_id
func (r *MissionRepository) Get(ctx context.Context, id string) (*model.Mission, error) {
	query := `SELECT mission_id, theater_id, mission_code, mission_name, mission_type, start_date, end_date, commanding_officer_user_id, mission_status, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM ops_missions WHERE mission_id = $1 AND deleted_at IS NULL`

	var m model.Mission
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.MissionId, &m.TheaterId, &m.MissionCode, &m.MissionName, &m.MissionType, &m.StartDate, &m.EndDate, &m.CommandingOfficerUserId, &m.MissionStatus, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	// Load AssignedShips
	childRowsAssignedShips, err := r.DB.Query(ctx, `SELECT assignment_id, mission_id, ship_id, tactical_callsign, role_in_task_force, joined_date, released_date, status, created_at FROM ops_mission_ship_assignments WHERE mission_id = $1`, id)
	if err == nil {
		defer childRowsAssignedShips.Close()
		for childRowsAssignedShips.Next() {
			var item model.MissionShipAssignments
			if err := childRowsAssignedShips.Scan(&item.AssignmentId, &item.MissionId, &item.ShipId, &item.TacticalCallsign, &item.RoleInTaskForce, &item.JoinedDate, &item.ReleasedDate, &item.Status, &item.CreatedAt); err == nil {
				m.AssignedShips = append(m.AssignedShips, item)
			}
		}
	}

	return &m, nil
}

// List retrieves paginated mission records
func (r *MissionRepository) List(ctx context.Context, opts model.ListOptions) ([]model.Mission, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(mission_code ILIKE $%[1]d OR mission_name ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM ops_missions WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT mission_id, theater_id, mission_code, mission_name, mission_type, start_date, end_date, commanding_officer_user_id, mission_status, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM ops_missions WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.Mission
	for rows.Next() {
		var m model.Mission
		if err := rows.Scan(&m.MissionId, &m.TheaterId, &m.MissionCode, &m.MissionName, &m.MissionType, &m.StartDate, &m.EndDate, &m.CommandingOfficerUserId, &m.MissionStatus, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new mission with optional child items
func (r *MissionRepository) Create(ctx context.Context, m *model.Mission) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO ops_missions (theater_id, mission_code, mission_name, mission_type, start_date, end_date, commanding_officer_user_id, mission_status, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4::ops_mission_type_enum, $5, $6, $7, $8::ops_mission_status_type, $9, $10, $11) RETURNING mission_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.TheaterId, m.MissionCode, m.MissionName, m.MissionType, m.StartDate, m.EndDate, m.CommandingOfficerUserId, m.MissionStatus, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.AssignedShips {
		_, err := tx.Exec(ctx, `INSERT INTO ops_mission_ship_assignments (mission_id, ship_id, tactical_callsign, role_in_task_force, joined_date, released_date, status) VALUES ($1, $2, $3, $4::ops_task_force_role_type, $5, $6, $7::ops_ship_assignment_status_type)`, newID, item.ShipId, item.TacticalCallsign, item.RoleInTaskForce, item.JoinedDate, item.ReleasedDate, item.Status)
		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing mission
func (r *MissionRepository) Update(ctx context.Context, id string, m *model.Mission) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE ops_missions SET theater_id = $1, mission_code = $2, mission_name = $3, mission_type = $4::ops_mission_type_enum, start_date = $5, end_date = $6, commanding_officer_user_id = $7, mission_status = $8::ops_mission_status_type, updated_by = $9, updated_at = CURRENT_TIMESTAMP WHERE mission_id = $10`
	_, err = tx.Exec(ctx, updateQuery, m.TheaterId, m.MissionCode, m.MissionName, m.MissionType, m.StartDate, m.EndDate, m.CommandingOfficerUserId, m.MissionStatus, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	if len(m.AssignedShips) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM ops_mission_ship_assignments WHERE mission_id = $1`, id)
		for _, item := range m.AssignedShips {
			_, err := tx.Exec(ctx, `INSERT INTO ops_mission_ship_assignments (mission_id, ship_id, tactical_callsign, role_in_task_force, joined_date, released_date, status) VALUES ($1, $2, $3, $4::ops_task_force_role_type, $5, $6, $7::ops_ship_assignment_status_type)`, id, item.ShipId, item.TacticalCallsign, item.RoleInTaskForce, item.JoinedDate, item.ReleasedDate, item.Status)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes mission
func (r *MissionRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE ops_missions SET deleted_at = CURRENT_TIMESTAMP WHERE mission_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
