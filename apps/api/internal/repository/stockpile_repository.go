package repository

import (
	"context"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StockpileRepository struct{ DB *pgxpool.Pool }

func NewStockpileRepository(db *pgxpool.Pool) *StockpileRepository {
	return &StockpileRepository{DB: db}
}
func (r *StockpileRepository) Create(ctx context.Context, item *model.Stockpile) (int, error) {
	var id int
	err := r.DB.QueryRow(ctx, `
		INSERT INTO stockpile (
			stockpile_name,city_id,stockpile_address,stockpile_latitude,
			stockpile_longitude,stockpile_map_url,is_active,created_by,updated_by
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$8) RETURNING stockpile_id`,
		item.StockpileName, item.CityID, item.StockpileAddress, item.StockpileLatitude,
		item.StockpileLongitude, item.StockpileMapUrl, item.IsActive, item.CreatedBy).Scan(&id)
	return id, err
}
func (r *StockpileRepository) Update(ctx context.Context, id int, item *model.Stockpile) error {
	_, err := r.DB.Exec(ctx, `
		UPDATE stockpile SET stockpile_name=$1,city_id=$2,stockpile_address=$3,
			stockpile_latitude=$4,stockpile_longitude=$5,stockpile_map_url=$6,
			is_active=COALESCE($7,is_active),updated_by=$8,updated_at=CURRENT_TIMESTAMP
		WHERE stockpile_id=$9 AND deleted_at IS NULL`,
		item.StockpileName, item.CityID, item.StockpileAddress, item.StockpileLatitude,
		item.StockpileLongitude, item.StockpileMapUrl, nullableActive(item.IsActive), item.UpdatedBy, id)
	return err
}
func (r *StockpileRepository) SoftDelete(ctx context.Context, userID, id int) error {
	_, err := r.DB.Exec(ctx, `UPDATE stockpile SET is_active=0,deleted_by=$1,deleted_at=CURRENT_TIMESTAMP WHERE stockpile_id=$2 AND deleted_at IS NULL`, userID, id)
	return err
}

const stockpileJSON = `to_jsonb(s) || jsonb_build_object('city',CASE WHEN c.city_id IS NULL THEN NULL ELSE to_jsonb(c) END)`

func (r *StockpileRepository) GetByID(ctx context.Context, id int) (*model.Stockpile, error) {
	return scanJSONRow[model.Stockpile](r.DB.QueryRow(ctx, `
		SELECT `+stockpileJSON+` FROM stockpile s
		LEFT JOIN city c ON c.city_id=s.city_id AND c.deleted_at IS NULL
		WHERE s.stockpile_id=$1 AND s.deleted_at IS NULL`, id))
}
func (r *StockpileRepository) List(ctx context.Context, _ interface{}, cursorKey *int, limit int, filters map[string]string, _, _ string) ([]model.Stockpile, error) {
	rows, err := r.DB.Query(ctx, `
		SELECT `+stockpileJSON+` FROM stockpile s
		LEFT JOIN city c ON c.city_id=s.city_id AND c.deleted_at IS NULL
		WHERE s.deleted_at IS NULL
		  AND ($1::int IS NULL OR s.stockpile_id < $1)
		  AND ($2='' OR s.stockpile_name ILIKE '%' || $2 || '%')
		  AND ($3::int IS NULL OR s.city_id=$3)
		  AND ($4::smallint IS NULL OR s.is_active=$4)
		ORDER BY s.stockpile_id DESC LIMIT $5`,
		cursorKey, filters["stockpile_name"], positiveFilter(filters["city_id"]),
		binaryFilter(filters["is_active"]), limit)
	if err != nil {
		return nil, err
	}
	return scanJSONRows[model.Stockpile](rows)
}
