package repository

import (
	"context"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type VesselRepository struct{ DB *pgxpool.Pool }

const vesselJSON = `
	to_jsonb(v) || jsonb_build_object(
		'vendor',CASE
			WHEN vendor.vendor_id IS NULL THEN NULL
			ELSE to_jsonb(vendor) || jsonb_build_object(
				'vendor_type',CASE WHEN vt.vendor_type_id IS NULL THEN NULL ELSE to_jsonb(vt) END,
				'bank_merk',CASE WHEN bm.bank_merk_id IS NULL THEN NULL ELSE to_jsonb(bm) END,
				'city',CASE WHEN c.city_id IS NULL THEN NULL ELSE to_jsonb(c) END
			)
		END
	)`

func NewVesselRepository(db *pgxpool.Pool) *VesselRepository { return &VesselRepository{DB: db} }

func (r *VesselRepository) Create(ctx context.Context, req *model.VesselRequest, userID int) (int, error) {
	active := 1
	if req.IsActive != nil {
		active = *req.IsActive
	}
	var id int
	err := r.DB.QueryRow(ctx, `
		INSERT INTO vessel (
			vendor_id,vessel_name,imo_number,registration_number,is_active,created_by,updated_by
		) VALUES ($1,$2,$3,$4,$5,$6,$6)
		RETURNING vessel_id`,
		req.VendorID, req.VesselName, req.IMONumber, req.RegistrationNumber, active, userID,
	).Scan(&id)
	return id, err
}

func (r *VesselRepository) Update(ctx context.Context, id int, req *model.VesselRequest, userID int) error {
	active := 1
	if req.IsActive != nil {
		active = *req.IsActive
	}
	_, err := r.DB.Exec(ctx, `
		UPDATE vessel SET vendor_id=$1,vessel_name=$2,imo_number=$3,registration_number=$4,
			is_active=$5,updated_by=$6,updated_at=CURRENT_TIMESTAMP
		WHERE vessel_id=$7 AND deleted_at IS NULL`,
		req.VendorID, req.VesselName, req.IMONumber, req.RegistrationNumber, active, userID, id)
	return err
}

func (r *VesselRepository) SoftDelete(ctx context.Context, id, userID int) error {
	_, err := r.DB.Exec(ctx, `UPDATE vessel SET deleted_by=$1,deleted_at=CURRENT_TIMESTAMP WHERE vessel_id=$2 AND deleted_at IS NULL`, userID, id)
	return err
}

func (r *VesselRepository) GetByID(ctx context.Context, id int) (*model.Vessel, error) {
	return scanJSONRow[model.Vessel](r.DB.QueryRow(ctx, `
		SELECT `+vesselJSON+`
		FROM vessel v
		LEFT JOIN vendor vendor ON vendor.vendor_id=v.vendor_id AND vendor.deleted_at IS NULL
		LEFT JOIN vendor_type vt ON vt.vendor_type_id=vendor.vendor_type_id AND vt.deleted_at IS NULL
		LEFT JOIN bank_merk bm ON bm.bank_merk_id=vendor.bank_merk_id AND bm.deleted_at IS NULL
		LEFT JOIN city c ON c.city_id=vendor.city_id AND c.deleted_at IS NULL
		WHERE v.vessel_id=$1 AND v.deleted_at IS NULL`, id))
}

func (r *VesselRepository) List(ctx context.Context, opts model.ListOptions) ([]model.Vessel, error) {
	rows, err := r.DB.Query(ctx, `
		SELECT `+vesselJSON+`
		FROM vessel v
		LEFT JOIN vendor vendor ON vendor.vendor_id=v.vendor_id AND vendor.deleted_at IS NULL
		LEFT JOIN vendor_type vt ON vt.vendor_type_id=vendor.vendor_type_id AND vt.deleted_at IS NULL
		LEFT JOIN bank_merk bm ON bm.bank_merk_id=vendor.bank_merk_id AND bm.deleted_at IS NULL
		LEFT JOIN city c ON c.city_id=vendor.city_id AND c.deleted_at IS NULL
		WHERE v.deleted_at IS NULL
		  AND ($1::int IS NULL OR v.vessel_id < $1)
		  AND ($2='' OR v.vessel_name ILIKE '%' || $2 || '%'
			OR COALESCE(v.imo_number,'') ILIKE '%' || $2 || '%')
		  AND ($3::int IS NULL OR v.vendor_id=$3)
		  AND ($4::smallint IS NULL OR v.is_active=$4)
		ORDER BY v.vessel_id DESC LIMIT $5 OFFSET $6`,
		opts.CursorKey, opts.Search, positiveFilter(opts.Filters["vendor_id"]), opts.IsActive, opts.Limit, opts.Offset)
	if err != nil {
		return nil, err
	}
	return scanJSONRows[model.Vessel](rows)
}
