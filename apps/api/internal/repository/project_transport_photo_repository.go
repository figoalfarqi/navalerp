package repository

import (
	"context"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectTransportPhotoRepository struct{ DB *pgxpool.Pool }

const projectTransportPhotoJSON = `
	to_jsonb(ph) || jsonb_build_object(
		'project_transport_status',jsonb_build_object(
			'project_transport_status_id',s.project_transport_status_id,
			'project_transport_status_type_id',s.project_transport_status_type_id,
			'status_time',s.status_time,
			'project_transport',` + projectTransportReferenceJSON + `
		)
	)`

func NewProjectTransportPhotoRepository(db *pgxpool.Pool) *ProjectTransportPhotoRepository {
	return &ProjectTransportPhotoRepository{DB: db}
}

func (r *ProjectTransportPhotoRepository) Create(ctx context.Context, req *model.ProjectTransportPhotoRequest, userID int) (int, error) {
	var id int
	err := r.DB.QueryRow(ctx, `
		INSERT INTO project_transport_photo (
			project_transport_status_id,photo_url,photo_type_id,photo_description,created_by,updated_by
		) VALUES ($1,$2,$3,$4,$5,$5)
		RETURNING project_transport_photo_id`,
		req.ProjectTransportStatusID, req.PhotoURL, req.PhotoTypeID, req.PhotoDescription, userID,
	).Scan(&id)
	return id, err
}

func (r *ProjectTransportPhotoRepository) Update(ctx context.Context, id int, req *model.ProjectTransportPhotoRequest, userID int) error {
	_, err := r.DB.Exec(ctx, `
		UPDATE project_transport_photo SET project_transport_status_id=$1,photo_url=$2,
			photo_type_id=$3,photo_description=$4,updated_by=$5,updated_at=CURRENT_TIMESTAMP
		WHERE project_transport_photo_id=$6 AND deleted_at IS NULL`,
		req.ProjectTransportStatusID, req.PhotoURL, req.PhotoTypeID, req.PhotoDescription, userID, id)
	return err
}

func (r *ProjectTransportPhotoRepository) SoftDelete(ctx context.Context, id, userID int) error {
	_, err := r.DB.Exec(ctx, `
		UPDATE project_transport_photo
		SET deleted_by=$1,deleted_at=CURRENT_TIMESTAMP,updated_by=$1,updated_at=CURRENT_TIMESTAMP
		WHERE project_transport_photo_id=$2 AND deleted_at IS NULL`, userID, id)
	return err
}

func (r *ProjectTransportPhotoRepository) GetByID(ctx context.Context, id int, checkerID *int) (*model.ProjectTransportPhoto, error) {
	return scanJSONRow[model.ProjectTransportPhoto](r.DB.QueryRow(ctx, `
		SELECT `+projectTransportPhotoJSON+`
		FROM project_transport_photo ph
		JOIN project_transport_status s ON s.project_transport_status_id=ph.project_transport_status_id
		JOIN project_transport pt ON pt.project_transport_id=s.project_transport_id
		JOIN project p ON p.project_id=pt.project_id
		JOIN project_route pr ON pr.project_route_id=pt.project_route_id
		LEFT JOIN truck t ON t.truck_id=pt.truck_id
		LEFT JOIN app_user d ON d.app_user_id=pt.driver_id AND d.deleted_at IS NULL
		LEFT JOIN vendor tv ON tv.vendor_id=pt.transport_vendor_id AND tv.deleted_at IS NULL
		WHERE ph.project_transport_photo_id=$1 AND ph.deleted_at IS NULL
		  AND s.deleted_at IS NULL AND pt.deleted_at IS NULL
		  AND (
			$2::int IS NULL OR EXISTS (
				SELECT 1 FROM project_checker_assignment ca
				WHERE ca.project_id=pt.project_id AND ca.checker_id=$2
				  AND ca.is_active=1 AND ca.deleted_at IS NULL
				  AND ca.access_started_at <= CURRENT_TIMESTAMP
				  AND (ca.access_ended_at IS NULL OR ca.access_ended_at >= CURRENT_TIMESTAMP)
			)
		  )`, id, checkerID))
}

func (r *ProjectTransportPhotoRepository) List(ctx context.Context, opts model.ListOptions, statusID *int) ([]model.ProjectTransportPhoto, error) {
	rows, err := r.DB.Query(ctx, `
		SELECT `+projectTransportPhotoJSON+`
		FROM project_transport_photo ph
		JOIN project_transport_status s ON s.project_transport_status_id=ph.project_transport_status_id
		JOIN project_transport pt ON pt.project_transport_id=s.project_transport_id
		JOIN project p ON p.project_id=pt.project_id
		JOIN project_route pr ON pr.project_route_id=pt.project_route_id
		LEFT JOIN truck t ON t.truck_id=pt.truck_id
		LEFT JOIN app_user d ON d.app_user_id=pt.driver_id AND d.deleted_at IS NULL
		LEFT JOIN vendor tv ON tv.vendor_id=pt.transport_vendor_id AND tv.deleted_at IS NULL
		WHERE ph.deleted_at IS NULL AND s.deleted_at IS NULL AND pt.deleted_at IS NULL
		  AND ($1::int IS NULL OR ph.project_transport_photo_id < $1)
		  AND ($2::int IS NULL OR ph.project_transport_status_id=$2)
		  AND ($3::int IS NULL OR pt.project_id=$3)
		  AND ($7::smallint IS NULL OR ph.photo_type_id=$7)
		  AND ($8::timestamptz IS NULL OR ph.created_at >= $8)
		  AND ($9::timestamptz IS NULL OR ph.created_at < $9)
		  AND (
			$4::int IS NULL OR EXISTS (
				SELECT 1 FROM project_checker_assignment ca
				WHERE ca.project_id=pt.project_id AND ca.checker_id=$4
				  AND ca.is_active=1 AND ca.deleted_at IS NULL
				  AND ca.access_started_at <= CURRENT_TIMESTAMP
				  AND (ca.access_ended_at IS NULL OR ca.access_ended_at >= CURRENT_TIMESTAMP)
			)
		  )
		ORDER BY ph.project_transport_photo_id DESC LIMIT $5 OFFSET $6`,
		opts.CursorKey, statusID, opts.ProjectID, opts.CheckerID, opts.Limit, opts.Offset,
		positiveFilter(opts.Filters["photo_type_id"]), opts.DateFrom, opts.DateTo)
	if err != nil {
		return nil, err
	}
	return scanJSONRows[model.ProjectTransportPhoto](rows)
}

func (r *ProjectTransportPhotoRepository) ResolveProjectIDByStatus(ctx context.Context, statusID int) (int, error) {
	var projectID int
	err := r.DB.QueryRow(ctx, `
		SELECT pt.project_id
		FROM project_transport_status s
		JOIN project_transport pt ON pt.project_transport_id=s.project_transport_id
		WHERE s.project_transport_status_id=$1 AND s.deleted_at IS NULL AND pt.deleted_at IS NULL`,
		statusID).Scan(&projectID)
	return projectID, err
}
