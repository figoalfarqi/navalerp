package repository

import (
	"context"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MineRepository struct{ DB *pgxpool.Pool }

func NewMineRepository(db *pgxpool.Pool) *MineRepository { return &MineRepository{DB: db} }

func (r *MineRepository) Create(ctx context.Context, item *model.Mine) (int, error) {
	var id int
	err := r.DB.QueryRow(ctx, `
		INSERT INTO mine (
			mine_name,city_id,mine_address,mine_latitude,mine_longitude,mine_map_url,
			is_active,created_by,updated_by
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$8) RETURNING mine_id`,
		item.MineName, item.CityID, item.MineAddress, item.MineLatitude, item.MineLongitude,
		item.MineMapUrl, item.IsActive, item.CreatedBy).Scan(&id)
	return id, err
}
func (r *MineRepository) Update(ctx context.Context, id int, item *model.Mine) error {
	_, err := r.DB.Exec(ctx, `
		UPDATE mine SET mine_name=$1,city_id=$2,mine_address=$3,mine_latitude=$4,
			mine_longitude=$5,mine_map_url=$6,is_active=COALESCE($7,is_active),
			updated_by=$8,updated_at=CURRENT_TIMESTAMP
		WHERE mine_id=$9 AND deleted_at IS NULL`,
		item.MineName, item.CityID, item.MineAddress, item.MineLatitude, item.MineLongitude,
		item.MineMapUrl, nullableActive(item.IsActive), item.UpdatedBy, id)
	return err
}
func (r *MineRepository) SoftDelete(ctx context.Context, userID, id int) error {
	_, err := r.DB.Exec(ctx, `UPDATE mine SET is_active=0,deleted_by=$1,deleted_at=CURRENT_TIMESTAMP WHERE mine_id=$2 AND deleted_at IS NULL`, userID, id)
	return err
}

const mineJSON = `to_jsonb(m) || jsonb_build_object('city',CASE WHEN c.city_id IS NULL THEN NULL ELSE to_jsonb(c) END)`

func (r *MineRepository) GetByID(ctx context.Context, id int) (*model.Mine, error) {
	return scanJSONRow[model.Mine](r.DB.QueryRow(ctx, `
		SELECT `+mineJSON+` FROM mine m
		LEFT JOIN city c ON c.city_id=m.city_id AND c.deleted_at IS NULL
		WHERE m.mine_id=$1 AND m.deleted_at IS NULL`, id))
}
func (r *MineRepository) List(ctx context.Context, _ interface{}, cursorKey *int, limit int, filters map[string]string, _, _ string) ([]model.Mine, error) {
	rows, err := r.DB.Query(ctx, `
		SELECT `+mineJSON+` FROM mine m
		LEFT JOIN city c ON c.city_id=m.city_id AND c.deleted_at IS NULL
		WHERE m.deleted_at IS NULL
		  AND ($1::int IS NULL OR m.mine_id < $1)
		  AND ($2='' OR m.mine_name ILIKE '%' || $2 || '%')
		  AND ($3::int IS NULL OR m.city_id=$3)
		  AND ($4::smallint IS NULL OR m.is_active=$4)
		ORDER BY m.mine_id DESC LIMIT $5`,
		cursorKey, filters["mine_name"], positiveFilter(filters["city_id"]),
		binaryFilter(filters["is_active"]), limit)
	if err != nil {
		return nil, err
	}
	return scanJSONRows[model.Mine](rows)
}

func nullableActive(value int) *int {
	if value != 0 && value != 1 {
		return nil
	}
	return &value
}
