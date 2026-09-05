package repository

import (
	"context"
	"errors"
	"time"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectCheckerAssignmentRepository struct{ DB *pgxpool.Pool }

const projectCheckerAssignmentJSON = `
	to_jsonb(a) || jsonb_build_object(
		'project',jsonb_build_object(
			'project_id',p.project_id,
			'project_code',p.project_code,
			'project_name',p.project_name
		),
		'checker',jsonb_build_object(
			'app_user_id',checker.app_user_id,
			'app_user_name',checker.app_user_name,
			'username',checker.username
		)
	)`

func NewProjectCheckerAssignmentRepository(db *pgxpool.Pool) *ProjectCheckerAssignmentRepository {
	return &ProjectCheckerAssignmentRepository{DB: db}
}

func (r *ProjectCheckerAssignmentRepository) Create(ctx context.Context, req *model.ProjectCheckerAssignmentRequest, userID int) (int, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)
	var isChecker bool
	if err := tx.QueryRow(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM app_user u JOIN app_role ar ON ar.app_role_id=u.app_role_id
			WHERE u.app_user_id=$1 AND u.deleted_at IS NULL AND u.app_user_status_id=1
			  AND ar.app_role_type_id=2 AND ar.deleted_at IS NULL AND ar.is_active=1
		)`, req.CheckerID).Scan(&isChecker); err != nil {
		return 0, err
	}
	if !isChecker {
		return 0, errors.New("checker_id is not an active checker")
	}
	active, defaultValue := 1, 0
	if req.IsActive != nil {
		active = *req.IsActive
	}
	if req.IsDefault != nil {
		defaultValue = *req.IsDefault
	}
	if defaultValue == 1 {
		if _, err := tx.Exec(ctx, `
			UPDATE project_checker_assignment
			SET is_default=0,updated_by=$1,updated_at=CURRENT_TIMESTAMP
			WHERE checker_id=$2 AND deleted_at IS NULL AND is_active=1`, userID, req.CheckerID); err != nil {
			return 0, err
		}
	}
	start := time.Now()
	if req.AccessStartedAt != nil {
		start = *req.AccessStartedAt
	}
	var id int
	err = tx.QueryRow(ctx, `
		INSERT INTO project_checker_assignment (
			project_id,checker_id,access_started_at,access_ended_at,assignment_note,
			is_active,is_default,created_by,updated_by
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$8)
		RETURNING project_checker_assignment_id`,
		req.ProjectID, req.CheckerID, start, req.AccessEndedAt, req.AssignmentNote,
		active, defaultValue, userID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, tx.Commit(ctx)
}

func (r *ProjectCheckerAssignmentRepository) Update(ctx context.Context, id int, req *model.ProjectCheckerAssignmentRequest, userID int) error {
	active, defaultValue := 1, 0
	if req.IsActive != nil {
		active = *req.IsActive
	}
	if req.IsDefault != nil {
		defaultValue = *req.IsDefault
	}
	start := time.Now()
	if req.AccessStartedAt != nil {
		start = *req.AccessStartedAt
	}
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	if defaultValue == 1 {
		if _, err := tx.Exec(ctx, `
			UPDATE project_checker_assignment SET is_default=0,updated_by=$1,updated_at=CURRENT_TIMESTAMP
			WHERE checker_id=$2 AND project_checker_assignment_id<>$3
			  AND deleted_at IS NULL AND is_active=1`, userID, req.CheckerID, id); err != nil {
			return err
		}
	}
	_, err = tx.Exec(ctx, `
		UPDATE project_checker_assignment SET
			project_id=$1,checker_id=$2,access_started_at=$3,access_ended_at=$4,
			assignment_note=$5,is_active=$6,is_default=$7,updated_by=$8,updated_at=CURRENT_TIMESTAMP
		WHERE project_checker_assignment_id=$9 AND deleted_at IS NULL`,
		req.ProjectID, req.CheckerID, start, req.AccessEndedAt, req.AssignmentNote,
		active, defaultValue, userID, id)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (r *ProjectCheckerAssignmentRepository) SoftDelete(ctx context.Context, id, userID int) error {
	_, err := r.DB.Exec(ctx, `
		UPDATE project_checker_assignment
		SET is_active=0,is_default=0,deleted_by=$1,deleted_at=CURRENT_TIMESTAMP,updated_by=$1,updated_at=CURRENT_TIMESTAMP
		WHERE project_checker_assignment_id=$2 AND deleted_at IS NULL`, userID, id)
	return err
}

func (r *ProjectCheckerAssignmentRepository) GetByID(ctx context.Context, id int) (*model.ProjectCheckerAssignment, error) {
	return scanJSONRow[model.ProjectCheckerAssignment](r.DB.QueryRow(ctx, `
		SELECT `+projectCheckerAssignmentJSON+`
		FROM project_checker_assignment a
		JOIN project p ON p.project_id=a.project_id AND p.deleted_at IS NULL
		JOIN app_user checker ON checker.app_user_id=a.checker_id AND checker.deleted_at IS NULL
		WHERE a.project_checker_assignment_id=$1 AND a.deleted_at IS NULL`, id))
}

func (r *ProjectCheckerAssignmentRepository) List(ctx context.Context, opts model.ListOptions) ([]model.ProjectCheckerAssignment, error) {
	rows, err := r.DB.Query(ctx, `
		SELECT `+projectCheckerAssignmentJSON+`
		FROM project_checker_assignment a
		JOIN project p ON p.project_id=a.project_id AND p.deleted_at IS NULL
		JOIN app_user checker ON checker.app_user_id=a.checker_id AND checker.deleted_at IS NULL
		WHERE a.deleted_at IS NULL
		  AND ($1::int IS NULL OR a.project_checker_assignment_id < $1)
		  AND ($2::int IS NULL OR a.project_id=$2)
		  AND ($3::int IS NULL OR a.checker_id=$3)
		  AND ($4::smallint IS NULL OR a.is_active=$4)
		  AND ($5::smallint IS NULL OR a.is_default=$5)
		ORDER BY a.project_checker_assignment_id DESC LIMIT $6 OFFSET $7`,
		opts.CursorKey, opts.ProjectID, opts.CheckerID, opts.IsActive,
		binaryFilter(opts.Filters["is_default"]), opts.Limit, opts.Offset)
	if err != nil {
		return nil, err
	}
	return scanJSONRows[model.ProjectCheckerAssignment](rows)
}

func (r *ProjectCheckerAssignmentRepository) GetDefaultForChecker(ctx context.Context, checkerID int) (*model.ProjectCheckerAssignment, error) {
	item, err := scanJSONRow[model.ProjectCheckerAssignment](r.DB.QueryRow(ctx, `
		SELECT `+projectCheckerAssignmentJSON+`
		FROM project_checker_assignment a
		JOIN project p ON p.project_id=a.project_id AND p.deleted_at IS NULL
		JOIN app_user checker ON checker.app_user_id=a.checker_id AND checker.deleted_at IS NULL
		WHERE a.checker_id=$1 AND a.is_default=1 AND a.is_active=1 AND a.deleted_at IS NULL
		  AND a.access_started_at <= CURRENT_TIMESTAMP
		  AND (a.access_ended_at IS NULL OR a.access_ended_at >= CURRENT_TIMESTAMP)
		ORDER BY a.project_checker_assignment_id DESC LIMIT 1`, checkerID))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return item, err
}
