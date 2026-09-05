package repository

import (
	"context"
	"time"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectTruckAssignmentRepository struct{ DB *pgxpool.Pool }

const projectTruckAssignmentJSON = `
	to_jsonb(a) || jsonb_build_object(
		'project',jsonb_build_object(
			'project_id',p.project_id,
			'project_code',p.project_code,
			'project_name',p.project_name
		),
		'truck',jsonb_build_object(
			'truck_id',t.truck_id,
			'license_plate',t.license_plate,
			'driver',CASE
				WHEN driver.app_user_id IS NULL THEN NULL
				ELSE jsonb_build_object(
					'app_user_id',driver.app_user_id,
					'app_user_name',driver.app_user_name,
					'username',driver.username
				)
			END,
			'vendor',CASE
				WHEN vendor.vendor_id IS NULL THEN NULL
				ELSE jsonb_build_object(
					'vendor_id',vendor.vendor_id,
					'vendor_name',vendor.vendor_name
				)
			END,
			'truck_type',CASE
				WHEN tt.truck_type_id IS NULL THEN NULL
				ELSE jsonb_build_object(
					'truck_type_id',tt.truck_type_id,
					'truck_type_name',tt.truck_type_name
				)
			END,
			'truck_merk',CASE
				WHEN tm.truck_merk_id IS NULL THEN NULL
				ELSE jsonb_build_object(
					'truck_merk_id',tm.truck_merk_id,
					'truck_merk_name',tm.truck_merk_name
				)
			END
		)
	)`

func NewProjectTruckAssignmentRepository(db *pgxpool.Pool) *ProjectTruckAssignmentRepository {
	return &ProjectTruckAssignmentRepository{DB: db}
}

func (r *ProjectTruckAssignmentRepository) Create(ctx context.Context, req *model.ProjectTruckAssignmentRequest, userID int) (int, error) {
	active := 1
	if req.IsActive != nil {
		active = *req.IsActive
	}
	start := time.Now()
	if req.AssignmentStartedAt != nil {
		start = *req.AssignmentStartedAt
	}
	var id int
	err := r.DB.QueryRow(ctx, `
		INSERT INTO project_truck_assignment (
			project_id,truck_id,assignment_started_at,assignment_ended_at,assignment_note,
			is_active,created_by,updated_by
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$7)
		RETURNING project_truck_assignment_id`,
		req.ProjectID, req.TruckID, start, req.AssignmentEndedAt, req.AssignmentNote, active, userID,
	).Scan(&id)
	return id, err
}

func (r *ProjectTruckAssignmentRepository) Update(ctx context.Context, id int, req *model.ProjectTruckAssignmentRequest, userID int) error {
	active := 1
	if req.IsActive != nil {
		active = *req.IsActive
	}
	start := time.Now()
	if req.AssignmentStartedAt != nil {
		start = *req.AssignmentStartedAt
	}
	_, err := r.DB.Exec(ctx, `
		UPDATE project_truck_assignment SET project_id=$1,truck_id=$2,assignment_started_at=$3,
			assignment_ended_at=$4,assignment_note=$5,is_active=$6,updated_by=$7,updated_at=CURRENT_TIMESTAMP
		WHERE project_truck_assignment_id=$8 AND deleted_at IS NULL`,
		req.ProjectID, req.TruckID, start, req.AssignmentEndedAt, req.AssignmentNote, active, userID, id)
	return err
}

func (r *ProjectTruckAssignmentRepository) SoftDelete(ctx context.Context, id, userID int) error {
	_, err := r.DB.Exec(ctx, `
		UPDATE project_truck_assignment SET is_active=0,deleted_by=$1,deleted_at=CURRENT_TIMESTAMP,
			updated_by=$1,updated_at=CURRENT_TIMESTAMP
		WHERE project_truck_assignment_id=$2 AND deleted_at IS NULL`, userID, id)
	return err
}

func (r *ProjectTruckAssignmentRepository) GetByID(ctx context.Context, id int) (*model.ProjectTruckAssignment, error) {
	return scanJSONRow[model.ProjectTruckAssignment](r.DB.QueryRow(ctx, `
		SELECT `+projectTruckAssignmentJSON+`
		FROM project_truck_assignment a
		JOIN project p ON p.project_id=a.project_id AND p.deleted_at IS NULL
		JOIN truck t ON t.truck_id=a.truck_id AND t.deleted_at IS NULL
		LEFT JOIN app_user driver ON driver.app_user_id=t.driver_id AND driver.deleted_at IS NULL
		LEFT JOIN vendor vendor ON vendor.vendor_id=t.vendor_id AND vendor.deleted_at IS NULL
		LEFT JOIN truck_type tt ON tt.truck_type_id=t.truck_type_id AND tt.deleted_at IS NULL
		LEFT JOIN truck_merk tm ON tm.truck_merk_id=t.truck_merk_id AND tm.deleted_at IS NULL
		WHERE project_truck_assignment_id=$1 AND a.deleted_at IS NULL`, id))
}

func (r *ProjectTruckAssignmentRepository) List(ctx context.Context, opts model.ListOptions) ([]model.ProjectTruckAssignment, error) {
	rows, err := r.DB.Query(ctx, `
		SELECT `+projectTruckAssignmentJSON+`
		FROM project_truck_assignment a
		JOIN project p ON p.project_id=a.project_id AND p.deleted_at IS NULL
		JOIN truck t ON t.truck_id=a.truck_id AND t.deleted_at IS NULL
		LEFT JOIN app_user driver ON driver.app_user_id=t.driver_id AND driver.deleted_at IS NULL
		LEFT JOIN vendor vendor ON vendor.vendor_id=t.vendor_id AND vendor.deleted_at IS NULL
		LEFT JOIN truck_type tt ON tt.truck_type_id=t.truck_type_id AND tt.deleted_at IS NULL
		LEFT JOIN truck_merk tm ON tm.truck_merk_id=t.truck_merk_id AND tm.deleted_at IS NULL
		WHERE a.deleted_at IS NULL
		  AND ($1::int IS NULL OR a.project_truck_assignment_id < $1)
		  AND ($2::int IS NULL OR a.project_id=$2)
		  AND ($6::int IS NULL OR a.truck_id=$6)
		  AND ($7::smallint IS NULL OR a.is_active=$7)
		  AND (
			$3::int IS NULL OR EXISTS (
				SELECT 1 FROM project_checker_assignment ca
				WHERE ca.project_id=a.project_id AND ca.checker_id=$3
				  AND ca.is_active=1 AND ca.deleted_at IS NULL
				  AND ca.access_started_at <= CURRENT_TIMESTAMP
				  AND (ca.access_ended_at IS NULL OR ca.access_ended_at >= CURRENT_TIMESTAMP)
			)
		  )
		  AND (
			$3::int IS NULL OR (
				a.is_active=1 AND a.assignment_started_at <= CURRENT_TIMESTAMP
				AND (a.assignment_ended_at IS NULL OR a.assignment_ended_at >= CURRENT_TIMESTAMP)
			)
		  )
		ORDER BY a.project_truck_assignment_id DESC LIMIT $4 OFFSET $5`,
		opts.CursorKey, opts.ProjectID, opts.CheckerID, opts.Limit, opts.Offset,
		opts.TruckID, opts.IsActive)
	if err != nil {
		return nil, err
	}
	return scanJSONRows[model.ProjectTruckAssignment](rows)
}

func (r *ProjectTruckAssignmentRepository) IsValid(ctx context.Context, assignmentID, projectID, truckID int) (bool, error) {
	var valid bool
	err := r.DB.QueryRow(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM project_truck_assignment
			WHERE project_truck_assignment_id=$1 AND project_id=$2 AND truck_id=$3
			  AND is_active=1 AND deleted_at IS NULL
			  AND assignment_started_at <= CURRENT_TIMESTAMP
			  AND (assignment_ended_at IS NULL OR assignment_ended_at >= CURRENT_TIMESTAMP)
		)`, assignmentID, projectID, truckID).Scan(&valid)
	return valid, err
}
