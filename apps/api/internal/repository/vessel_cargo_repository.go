package repository

import (
	"context"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type VesselCargoRepository struct{ DB *pgxpool.Pool }

const vesselCargoJSON = `
	to_jsonb(vc) || jsonb_build_object(
		'vessel',to_jsonb(v) || jsonb_build_object(
			'vendor',CASE WHEN vendor.vendor_id IS NULL THEN NULL ELSE to_jsonb(vendor) END
		),
		'port',to_jsonb(po) || jsonb_build_object(
			'city',CASE WHEN city.city_id IS NULL THEN NULL ELSE to_jsonb(city) END
		),
		'cargo_type',to_jsonb(ct)
	)`

func NewVesselCargoRepository(db *pgxpool.Pool) *VesselCargoRepository {
	return &VesselCargoRepository{DB: db}
}

func (r *VesselCargoRepository) Create(ctx context.Context, req *model.VesselCargoRequest, userID int) (int, error) {
	active := 1
	if req.IsActive != nil {
		active = *req.IsActive
	}
	var id int
	err := r.DB.QueryRow(ctx, `
		INSERT INTO vessel_cargo (
			vessel_id,port_id,cargo_type_id,voyage_number,bill_of_lading_number,
			arrival_at,unloading_started_at,unloading_completed_at,
			manifest_volume_cubic,manifest_weight_ton,is_active,created_by,updated_by
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$12)
		RETURNING vessel_cargo_id`,
		req.VesselID, req.PortID, req.CargoTypeID, req.VoyageNumber,
		req.BillOfLadingNumber, req.ArrivalAt, req.UnloadingStartedAt,
		req.UnloadingCompletedAt, req.ManifestVolumeCubic, req.ManifestWeightTon,
		active, userID,
	).Scan(&id)
	return id, err
}

func (r *VesselCargoRepository) Update(ctx context.Context, id int, req *model.VesselCargoRequest, userID int) error {
	active := 1
	if req.IsActive != nil {
		active = *req.IsActive
	}
	_, err := r.DB.Exec(ctx, `
		UPDATE vessel_cargo SET vessel_id=$1,port_id=$2,cargo_type_id=$3,voyage_number=$4,
			bill_of_lading_number=$5,arrival_at=$6,unloading_started_at=$7,
			unloading_completed_at=$8,manifest_volume_cubic=$9,manifest_weight_ton=$10,
			is_active=$11,updated_by=$12,updated_at=CURRENT_TIMESTAMP
		WHERE vessel_cargo_id=$13 AND deleted_at IS NULL`,
		req.VesselID, req.PortID, req.CargoTypeID, req.VoyageNumber,
		req.BillOfLadingNumber, req.ArrivalAt, req.UnloadingStartedAt,
		req.UnloadingCompletedAt, req.ManifestVolumeCubic, req.ManifestWeightTon,
		active, userID, id)
	return err
}

func (r *VesselCargoRepository) SoftDelete(ctx context.Context, id, userID int) error {
	_, err := r.DB.Exec(ctx, `UPDATE vessel_cargo SET deleted_by=$1,deleted_at=CURRENT_TIMESTAMP WHERE vessel_cargo_id=$2 AND deleted_at IS NULL`, userID, id)
	return err
}

func (r *VesselCargoRepository) GetByID(ctx context.Context, id int) (*model.VesselCargo, error) {
	return scanJSONRow[model.VesselCargo](r.DB.QueryRow(ctx, `
		SELECT `+vesselCargoJSON+`
		FROM vessel_cargo vc
		JOIN vessel v ON v.vessel_id=vc.vessel_id AND v.deleted_at IS NULL
		LEFT JOIN vendor vendor ON vendor.vendor_id=v.vendor_id AND vendor.deleted_at IS NULL
		JOIN port po ON po.port_id=vc.port_id AND po.deleted_at IS NULL
		LEFT JOIN city city ON city.city_id=po.city_id AND city.deleted_at IS NULL
		JOIN cargo_type ct ON ct.cargo_type_id=vc.cargo_type_id AND ct.deleted_at IS NULL
		WHERE vc.vessel_cargo_id=$1 AND vc.deleted_at IS NULL`, id))
}

func (r *VesselCargoRepository) List(ctx context.Context, opts model.ListOptions) ([]model.VesselCargo, error) {
	rows, err := r.DB.Query(ctx, `
		SELECT `+vesselCargoJSON+`
		FROM vessel_cargo vc
		JOIN vessel v ON v.vessel_id=vc.vessel_id AND v.deleted_at IS NULL
		LEFT JOIN vendor vendor ON vendor.vendor_id=v.vendor_id AND vendor.deleted_at IS NULL
		JOIN port po ON po.port_id=vc.port_id AND po.deleted_at IS NULL
		LEFT JOIN city city ON city.city_id=po.city_id AND city.deleted_at IS NULL
		JOIN cargo_type ct ON ct.cargo_type_id=vc.cargo_type_id AND ct.deleted_at IS NULL
		WHERE vc.deleted_at IS NULL
		  AND ($1::int IS NULL OR vc.vessel_cargo_id < $1)
		  AND ($2::int IS NULL OR vc.vessel_id=$2)
		  AND ($3::int IS NULL OR vc.port_id=$3)
		  AND ($4::int IS NULL OR vc.cargo_type_id=$4)
		  AND ($5::smallint IS NULL OR vc.is_active=$5)
		  AND ($6='' OR COALESCE(vc.voyage_number,'') ILIKE '%' || $6 || '%'
			OR COALESCE(vc.bill_of_lading_number,'') ILIKE '%' || $6 || '%')
		ORDER BY vc.vessel_cargo_id DESC LIMIT $7 OFFSET $8`,
		opts.CursorKey, opts.VesselID, opts.PortID, opts.CargoTypeID, opts.IsActive,
		opts.Filters["voyage_number"], opts.Limit, opts.Offset)
	if err != nil {
		return nil, err
	}
	return scanJSONRows[model.VesselCargo](rows)
}
