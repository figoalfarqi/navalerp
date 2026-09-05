package repository

import (
	"context"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClientDestinationRepository struct{ DB *pgxpool.Pool }

func NewClientDestinationRepository(db *pgxpool.Pool) *ClientDestinationRepository {
	return &ClientDestinationRepository{DB: db}
}
func (r *ClientDestinationRepository) Create(ctx context.Context, item *model.ClientDestination) (int, error) {
	var id int
	err := r.DB.QueryRow(ctx, `
		INSERT INTO client_destination (
			client_id,client_destination_name,city_id,client_destination_address,
			client_destination_latitude,client_destination_longitude,client_destination_map_url,
			operating_hours,is_active,created_by,updated_by
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$10)
		RETURNING client_destination_id`,
		item.ClientID, item.ClientDestinationName, item.CityID, item.ClientDestinationAddress,
		item.ClientDestinationLatitude, item.ClientDestinationLongitude, item.ClientDestinationMapUrl,
		item.OperatingHours, item.IsActive, item.CreatedBy).Scan(&id)
	return id, err
}
func (r *ClientDestinationRepository) Update(ctx context.Context, id int, item *model.ClientDestination) error {
	_, err := r.DB.Exec(ctx, `
		UPDATE client_destination SET client_id=$1,client_destination_name=$2,city_id=$3,
			client_destination_address=$4,client_destination_latitude=$5,
			client_destination_longitude=$6,client_destination_map_url=$7,operating_hours=$8,
			is_active=COALESCE($9,is_active),updated_by=$10,updated_at=CURRENT_TIMESTAMP
		WHERE client_destination_id=$11 AND deleted_at IS NULL`,
		item.ClientID, item.ClientDestinationName, item.CityID, item.ClientDestinationAddress,
		item.ClientDestinationLatitude, item.ClientDestinationLongitude, item.ClientDestinationMapUrl,
		item.OperatingHours, nullableActive(item.IsActive), item.UpdatedBy, id)
	return err
}
func (r *ClientDestinationRepository) SoftDelete(ctx context.Context, userID, id int) error {
	_, err := r.DB.Exec(ctx, `
		UPDATE client_destination SET is_active=0,deleted_by=$1,deleted_at=CURRENT_TIMESTAMP
		WHERE client_destination_id=$2 AND deleted_at IS NULL`, userID, id)
	return err
}

const clientDestinationJSON = `
	to_jsonb(cd) || jsonb_build_object(
		'client',CASE WHEN cl.client_id IS NULL THEN NULL ELSE to_jsonb(cl) END,
		'city',CASE WHEN c.city_id IS NULL THEN NULL ELSE to_jsonb(c) END
	)`

func (r *ClientDestinationRepository) GetByID(ctx context.Context, id int) (*model.ClientDestination, error) {
	return scanJSONRow[model.ClientDestination](r.DB.QueryRow(ctx, `
		SELECT `+clientDestinationJSON+` FROM client_destination cd
		LEFT JOIN client cl ON cl.client_id=cd.client_id AND cl.deleted_at IS NULL
		LEFT JOIN city c ON c.city_id=cd.city_id AND c.deleted_at IS NULL
		WHERE cd.client_destination_id=$1 AND cd.deleted_at IS NULL`, id))
}
func (r *ClientDestinationRepository) List(ctx context.Context, _ interface{}, cursorKey *int, limit int, filters map[string]string, _, _ string) ([]model.ClientDestination, error) {
	rows, err := r.DB.Query(ctx, `
		SELECT `+clientDestinationJSON+` FROM client_destination cd
		LEFT JOIN client cl ON cl.client_id=cd.client_id AND cl.deleted_at IS NULL
		LEFT JOIN city c ON c.city_id=cd.city_id AND c.deleted_at IS NULL
		WHERE cd.deleted_at IS NULL
		  AND ($1::int IS NULL OR cd.client_destination_id < $1)
		  AND ($2='' OR cd.client_destination_name ILIKE '%' || $2 || '%')
		  AND ($3::int IS NULL OR cd.client_id=$3)
		  AND ($4::int IS NULL OR cd.city_id=$4)
		  AND ($5::smallint IS NULL OR cd.is_active=$5)
		ORDER BY cd.client_destination_id DESC LIMIT $6`,
		cursorKey, filters["client_destination_name"], positiveFilter(filters["client_id"]),
		positiveFilter(filters["city_id"]), binaryFilter(filters["is_active"]), limit)
	if err != nil {
		return nil, err
	}
	return scanJSONRows[model.ClientDestination](rows)
}
