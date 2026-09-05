package repository

import (
	"context"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StockpileCargoRepository struct{ DB *pgxpool.Pool }

func NewStockpileCargoRepository(db *pgxpool.Pool) *StockpileCargoRepository {
	return &StockpileCargoRepository{DB: db}
}

func cargoQuantity(value *float64) float64 {
	if value == nil {
		return 0
	}
	return *value
}

func (r *StockpileCargoRepository) Create(ctx context.Context, req *model.StockpileCargoRequest, userID int) (int, error) {
	active := 1
	if req.IsActive != nil {
		active = *req.IsActive
	}
	var id int
	err := r.DB.QueryRow(ctx, `
		INSERT INTO stockpile_cargo (
			stockpile_id,cargo_type_id,capacity_volume_cubic,capacity_weight_ton,
			current_volume_cubic,current_weight_ton,is_active,created_by,updated_by
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$8)
		RETURNING stockpile_cargo_id`,
		req.StockpileID, req.CargoTypeID, req.CapacityVolumeCubic, req.CapacityWeightTon,
		cargoQuantity(req.CurrentVolumeCubic), cargoQuantity(req.CurrentWeightTon), active, userID,
	).Scan(&id)
	return id, err
}

func (r *StockpileCargoRepository) Update(ctx context.Context, id int, req *model.StockpileCargoRequest, userID int) error {
	_, err := r.DB.Exec(ctx, `
		UPDATE stockpile_cargo SET
			stockpile_id=$1,cargo_type_id=$2,capacity_volume_cubic=$3,capacity_weight_ton=$4,
			current_volume_cubic=COALESCE($5,current_volume_cubic),
			current_weight_ton=COALESCE($6,current_weight_ton),
			is_active=COALESCE($7,is_active),updated_by=$8,updated_at=CURRENT_TIMESTAMP
		WHERE stockpile_cargo_id=$9 AND deleted_at IS NULL`,
		req.StockpileID, req.CargoTypeID, req.CapacityVolumeCubic, req.CapacityWeightTon,
		req.CurrentVolumeCubic, req.CurrentWeightTon, req.IsActive, userID, id)
	return err
}

func (r *StockpileCargoRepository) SoftDelete(ctx context.Context, userID, id int) error {
	_, err := r.DB.Exec(ctx, `
		UPDATE stockpile_cargo SET is_active=0,deleted_by=$1,deleted_at=CURRENT_TIMESTAMP,
			updated_by=$1,updated_at=CURRENT_TIMESTAMP
		WHERE stockpile_cargo_id=$2 AND deleted_at IS NULL`, userID, id)
	return err
}

const stockpileCargoJSON = `
	to_jsonb(sc) || jsonb_build_object(
		'stockpile',to_jsonb(s) || jsonb_build_object(
			'city',CASE WHEN city.city_id IS NULL THEN NULL ELSE to_jsonb(city) END
		),
		'cargo_type',to_jsonb(ct)
	)`

func (r *StockpileCargoRepository) GetByID(ctx context.Context, id int) (*model.StockpileCargo, error) {
	return scanJSONRow[model.StockpileCargo](r.DB.QueryRow(ctx, `
		SELECT `+stockpileCargoJSON+`
		FROM stockpile_cargo sc
		JOIN stockpile s ON s.stockpile_id=sc.stockpile_id AND s.deleted_at IS NULL
		LEFT JOIN city city ON city.city_id=s.city_id AND city.deleted_at IS NULL
		JOIN cargo_type ct ON ct.cargo_type_id=sc.cargo_type_id AND ct.deleted_at IS NULL
		WHERE sc.stockpile_cargo_id=$1 AND sc.deleted_at IS NULL`, id))
}

func (r *StockpileCargoRepository) List(ctx context.Context, opts model.ListOptions) ([]model.StockpileCargo, error) {
	rows, err := r.DB.Query(ctx, `
		SELECT `+stockpileCargoJSON+`
		FROM stockpile_cargo sc
		JOIN stockpile s ON s.stockpile_id=sc.stockpile_id AND s.deleted_at IS NULL
		LEFT JOIN city city ON city.city_id=s.city_id AND city.deleted_at IS NULL
		JOIN cargo_type ct ON ct.cargo_type_id=sc.cargo_type_id AND ct.deleted_at IS NULL
		WHERE sc.deleted_at IS NULL
		  AND ($1::int IS NULL OR sc.stockpile_cargo_id < $1)
		  AND ($2::int IS NULL OR sc.stockpile_id=$2)
		  AND ($3::int IS NULL OR sc.cargo_type_id=$3)
		  AND ($4::smallint IS NULL OR sc.is_active=$4)
		ORDER BY sc.stockpile_cargo_id DESC LIMIT $5 OFFSET $6`,
		opts.CursorKey, positiveFilter(opts.Filters["stockpile_id"]), opts.CargoTypeID,
		opts.IsActive, opts.Limit, opts.Offset)
	if err != nil {
		return nil, err
	}
	return scanJSONRows[model.StockpileCargo](rows)
}
