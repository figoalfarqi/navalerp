package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CrewAssignmentRepository struct {
	DB *pgxpool.Pool
}

func NewCrewAssignmentRepository(db *pgxpool.Pool) *CrewAssignmentRepository {
	return &CrewAssignmentRepository{DB: db}
}

// Get retrieves a single crew_assignment by assignment_id
func (r *CrewAssignmentRepository) Get(ctx context.Context, id string) (*model.CrewAssignment, error) {
	query := `SELECT t.assignment_id, t.ship_id, COALESCE(j_ship.ship_name, ''), t.personnel_id, COALESCE(j_prs.full_name, ''), t.crew_role, t.department, t.watch_bill_duty, t.assigned_date, t.relieved_date, t.is_active, t.created_by, t.created_at, t.updated_at
	FROM hcm_crew_assignments t
	LEFT JOIN mro_ships j_ship ON j_ship.ship_id = t.ship_id
	LEFT JOIN hcm_personnel j_prs ON j_prs.personnel_id = t.personnel_id
	WHERE t.assignment_id = $1`

	var m model.CrewAssignment
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.AssignmentId, &m.ShipId, &m.ShipName, &m.PersonnelId, &m.PersonnelName, &m.CrewRole, &m.Department, &m.WatchBillDuty, &m.AssignedDate, &m.RelievedDate, &m.IsActive, &m.CreatedBy, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// Load Allowances
	childRowsAllowances, err := r.DB.Query(ctx, `SELECT allowance_id, personnel_id, ship_id, mission_name, start_date, end_date, days_at_sea, daily_allowance_rate, total_allowance, payment_status, payment_reference_no, created_at FROM hcm_sea_duty_allowances WHERE assignment_id = $1`, id)
	if err == nil {
		defer childRowsAllowances.Close()
		for childRowsAllowances.Next() {
			var item model.SeaDutyAllowances
			if err := childRowsAllowances.Scan(&item.AllowanceId, &item.PersonnelId, &item.ShipId, &item.MissionName, &item.StartDate, &item.EndDate, &item.DaysAtSea, &item.DailyAllowanceRate, &item.TotalAllowance, &item.PaymentStatus, &item.PaymentReferenceNo, &item.CreatedAt); err == nil {
				m.Allowances = append(m.Allowances, item)
			}
		}
	}

	return &m, nil
}

// List retrieves paginated crew_assignment records
func (r *CrewAssignmentRepository) List(ctx context.Context, opts model.ListOptions) ([]model.CrewAssignment, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(crew_role ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM hcm_crew_assignments t WHERE %s", whereSql)
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

	listQuery := fmt.Sprintf(`SELECT t.assignment_id, t.ship_id, COALESCE(j_ship.ship_name, ''), t.personnel_id, COALESCE(j_prs.full_name, ''), t.crew_role, t.department, t.watch_bill_duty, t.assigned_date, t.relieved_date, t.is_active, t.created_by, t.created_at, t.updated_at
	FROM hcm_crew_assignments t
	LEFT JOIN mro_ships j_ship ON j_ship.ship_id = t.ship_id
	LEFT JOIN hcm_personnel j_prs ON j_prs.personnel_id = t.personnel_id
	WHERE %s ORDER BY t.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.CrewAssignment
	for rows.Next() {
		var m model.CrewAssignment
		if err := rows.Scan(&m.AssignmentId, &m.ShipId, &m.ShipName, &m.PersonnelId, &m.PersonnelName, &m.CrewRole, &m.Department, &m.WatchBillDuty, &m.AssignedDate, &m.RelievedDate, &m.IsActive, &m.CreatedBy, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new crew_assignment with optional child items
func (r *CrewAssignmentRepository) Create(ctx context.Context, m *model.CrewAssignment) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO hcm_crew_assignments (ship_id, personnel_id, crew_role, department, watch_bill_duty, assigned_date, relieved_date, is_active, created_by) VALUES ($1, $2, $3, $4::crew_department_type, $5::watch_bill_duty_type, $6, $7, $8, $9) RETURNING assignment_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.ShipId, m.PersonnelId, m.CrewRole, m.Department, m.WatchBillDuty, m.AssignedDate, m.RelievedDate, m.IsActive, m.CreatedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.Allowances {
		_, err := tx.Exec(ctx, `INSERT INTO hcm_sea_duty_allowances (personnel_id, ship_id, mission_name, start_date, end_date, days_at_sea, daily_allowance_rate, payment_status, payment_reference_no) VALUES ($1, $2, $3, $4, $5, $6, $7, $8::allowance_payment_status_type, $9)`, item.PersonnelId, item.ShipId, item.MissionName, item.StartDate, item.EndDate, item.DaysAtSea, item.DailyAllowanceRate, item.PaymentStatus, item.PaymentReferenceNo)
		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing crew_assignment
func (r *CrewAssignmentRepository) Update(ctx context.Context, id string, m *model.CrewAssignment) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE hcm_crew_assignments SET ship_id = $1, personnel_id = $2, crew_role = $3, department = $4::crew_department_type, watch_bill_duty = $5::watch_bill_duty_type, assigned_date = $6, relieved_date = $7, is_active = $8, updated_at = CURRENT_TIMESTAMP WHERE assignment_id = $9`
	_, err = tx.Exec(ctx, updateQuery, m.ShipId, m.PersonnelId, m.CrewRole, m.Department, m.WatchBillDuty, m.AssignedDate, m.RelievedDate, m.IsActive, id)
	if err != nil {
		return err
	}

	if len(m.Allowances) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM hcm_sea_duty_allowances WHERE assignment_id = $1`, id)
		for _, item := range m.Allowances {
			_, err := tx.Exec(ctx, `INSERT INTO hcm_sea_duty_allowances (personnel_id, ship_id, mission_name, start_date, end_date, days_at_sea, daily_allowance_rate, payment_status, payment_reference_no) VALUES ($1, $2, $3, $4, $5, $6, $7, $8::allowance_payment_status_type, $9)`, item.PersonnelId, item.ShipId, item.MissionName, item.StartDate, item.EndDate, item.DaysAtSea, item.DailyAllowanceRate, item.PaymentStatus, item.PaymentReferenceNo)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes crew_assignment
func (r *CrewAssignmentRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM hcm_crew_assignments t WHERE assignment_id = $1`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
