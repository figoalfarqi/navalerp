package repository

import (
	"context"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PortRepository struct{ DB *pgxpool.Pool }

const portJSON = `
	to_jsonb(po) || jsonb_build_object(
		'city',CASE
			WHEN c.city_id IS NULL THEN NULL
			ELSE to_jsonb(c) || jsonb_build_object(
				'province',CASE WHEN province.province_id IS NULL THEN NULL ELSE to_jsonb(province) END
			)
		END
	)`

func NewPortRepository(db *pgxpool.Pool) *PortRepository { return &PortRepository{DB: db} }

func (r *PortRepository) Create(ctx context.Context, req *model.PortRequest, userID int) (int, error) {
	active := 1
	if req.IsActive != nil {
		active = *req.IsActive
	}
	var id int
	err := r.DB.QueryRow(ctx, `
		INSERT INTO port (
			port_name,city_id,port_address,port_latitude,port_longitude,port_map_url,
			is_active,created_by,updated_by
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$8)
		RETURNING port_id`,
		req.PortName, req.CityID, req.PortAddress, req.PortLatitude, req.PortLongitude,
		req.PortMapURL, active, userID,
	).Scan(&id)
	return id, err
}

func (r *PortRepository) Update(ctx context.Context, id int, req *model.PortRequest, userID int) error {
	active := 1
	if req.IsActive != nil {
		active = *req.IsActive
	}
	_, err := r.DB.Exec(ctx, `
		UPDATE port SET port_name=$1,city_id=$2,port_address=$3,port_latitude=$4,
			port_longitude=$5,port_map_url=$6,is_active=$7,updated_by=$8,
			updated_at=CURRENT_TIMESTAMP
		WHERE port_id=$9 AND deleted_at IS NULL`,
		req.PortName, req.CityID, req.PortAddress, req.PortLatitude, req.PortLongitude,
		req.PortMapURL, active, userID, id,
	)
	return err
}

func (r *PortRepository) SoftDelete(ctx context.Context, id, userID int) error {
	_, err := r.DB.Exec(ctx, `UPDATE port SET deleted_by=$1,deleted_at=CURRENT_TIMESTAMP WHERE port_id=$2 AND deleted_at IS NULL`, userID, id)
	return err
}

func (r *PortRepository) GetByID(ctx context.Context, id int) (*model.Port, error) {
	return scanJSONRow[model.Port](r.DB.QueryRow(ctx, `
		SELECT `+portJSON+`
		FROM port po
		LEFT JOIN city c ON c.city_id=po.city_id AND c.deleted_at IS NULL
		LEFT JOIN province province ON province.province_id=c.province_id AND province.deleted_at IS NULL
		WHERE po.port_id=$1 AND po.deleted_at IS NULL`, id))
}

func (r *PortRepository) List(ctx context.Context, opts model.ListOptions) ([]model.Port, error) {
	rows, err := r.DB.Query(ctx, `
		SELECT `+portJSON+`
		FROM port po
		LEFT JOIN city c ON c.city_id=po.city_id AND c.deleted_at IS NULL
		LEFT JOIN province province ON province.province_id=c.province_id AND province.deleted_at IS NULL
		WHERE po.deleted_at IS NULL
		  AND ($1::int IS NULL OR po.port_id < $1)
		  AND ($2='' OR po.port_name ILIKE '%' || $2 || '%')
		  AND ($3::int IS NULL OR po.city_id=$3)
		  AND ($4::smallint IS NULL OR po.is_active=$4)
		ORDER BY po.port_id DESC LIMIT $5 OFFSET $6`,
		opts.CursorKey, opts.Search, positiveFilter(opts.Filters["city_id"]), opts.IsActive, opts.Limit, opts.Offset)
	if err != nil {
		return nil, err
	}
	return scanJSONRows[model.Port](rows)
}
