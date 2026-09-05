package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/figoalfarqi/navalerp/internal/helper"
	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClientRepository struct {
	DB *pgxpool.Pool
}

func NewClientRepository(db *pgxpool.Pool) *ClientRepository {
	return &ClientRepository{DB: db}
}

// ==================================================
// Create (TX)
// ==================================================
func (r *ClientRepository) Create(
	ctx context.Context,
	m *model.Client,
) (int, error) {

	query := `
        INSERT INTO client (
            client_name,
            client_email,
            client_tin,
            number_of_day_until_due,
            city_id,
            client_address,
            is_active,
            created_by,
            updated_by
        )
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
        RETURNING client_id
    `

	var id int
	err := r.DB.QueryRow(ctx, query,
		m.ClientName,
		m.ClientEmail,
		m.ClientTin,
		m.NumberOfDayUntilDue,
		m.CityID,
		m.ClientAddress,
		m.IsActive,
		m.CreatedBy,
		m.UpdatedBy,
	).Scan(&id)

	return id, err
}

// ==================================================
// Update (TX)
// ==================================================
func (r *ClientRepository) Update(
	ctx context.Context,
	id int,
	m *model.Client,
) error {

	// Field yang *selalu* di-update
	args := []interface{}{
		m.ClientName,          // $1
		m.ClientEmail,         // $2
		m.ClientTin,           // $3
		m.NumberOfDayUntilDue, // $4
		m.CityID,              // $5
		m.ClientAddress,       // $6
	}

	// Mulai parameter posisi
	argPos := len(args) + 1

	// Optional field: is_active
	setIsActive := ""
	if m.IsActive != -1 { // jika bukan -1, update
		setIsActive = fmt.Sprintf(", is_active = $%d", argPos)
		args = append(args, m.IsActive)
		argPos++
	}

	// updated_by dan updated_at
	args = append(args,
		m.UpdatedBy, // $argPos
		m.UpdatedAt, // $argPos+1
		id,          // $argPos+2
	)

	query := fmt.Sprintf(`
        UPDATE client
        SET
            client_name = $1,
            client_email = $2,
            client_tin = $3,
            number_of_day_until_due = $4,
            city_id = $5,
            client_address = $6
            %s,
            updated_by = $%d,
            updated_at = $%d
        WHERE client_id = $%d
          AND deleted_at IS NULL
        `,
		setIsActive,
		argPos,   // updated_by
		argPos+1, // updated_at
		argPos+2, // id
	)

	cmd, err := r.DB.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

// ==================================================
// Soft Delete
// ==================================================
func (r *ClientRepository) SoftDelete(ctx context.Context, deletedBy, id int) error {
	query := `
        UPDATE client
        SET deleted_at = $1,
            deleted_by = $2
        WHERE client_id = $3
          AND deleted_at IS NULL
    `

	_, err := r.DB.Exec(ctx, query, time.Now(), deletedBy, id)
	return err
}

// ==================================================
// Get By ID
// ==================================================
func (r *ClientRepository) GetByID(ctx context.Context, id int) (*model.Client, error) {

	query := `
        SELECT
            c.client_id,
            c.client_name,
            c.client_email,
            c.client_tin,
            c.number_of_day_until_due,
            c.city_id,
            c.client_address,
            c.is_active,
            c.created_by,
            c.updated_by,
            c.created_at,
            c.updated_at,
            c.deleted_at,

            ci.city_id,
            ci.province_id,
            ci.city_name,
            ci.is_active,
            ci.created_at,
            ci.updated_at,

            COALESCE(pics.data, '[]')

        FROM client c
        LEFT JOIN city ci ON ci.city_id = c.city_id

        LEFT JOIN LATERAL (
            SELECT json_agg(json_build_object(
                'app_user_id', u.app_user_id,
                '	', u.app_user_name,
                '	', u.app_user_phone,
				'	', u.username
            )) AS data
            FROM app_user u
            WHERE u.client_id = c.client_id
              AND u.deleted_at IS NULL
        ) pics ON TRUE

        WHERE c.client_id = $1
          AND c.deleted_at IS NULL
    `

	var m model.Client
	var picsJSON []byte

	var city model.CityNullable

	err := r.DB.QueryRow(ctx, query, id).Scan(
		&m.ClientID,
		&m.ClientName,
		&m.ClientEmail,
		&m.ClientTin,
		&m.NumberOfDayUntilDue,
		&m.CityID,
		&m.ClientAddress,
		&m.IsActive,
		&m.CreatedBy,
		&m.UpdatedBy,
		&m.CreatedAt,
		&m.UpdatedAt,
		&m.DeletedAt,

		&city.CityID,
		&city.ProvinceID,
		&city.CityName,
		&city.IsActive,
		&city.CreatedAt,
		&city.UpdatedAt,

		&picsJSON,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	_ = json.Unmarshal(picsJSON, &m.ClientPics)

	if city.CityID != nil {
		m.City = city.ToNotNullable()
	}
	return &m, nil
}

// ==================================================
// List
// ==================================================
func (r *ClientRepository) List(
	ctx context.Context,
	cursorValue interface{}, cursorKey *int, limit int, filters map[string]string,
	orderBy, sort string,
) ([]model.Client, error) {
	tableKey := "client_id"
	parsedOrderBy := "c." + orderBy
	baseQuery := `
        SELECT
            c.client_id,
            c.client_name,
            c.client_email,
            c.client_tin,
            c.number_of_day_until_due,
            c.city_id,
            c.client_address,
            c.is_active,
            c.created_by,
            c.updated_by,
            c.created_at,
            c.updated_at,
            c.deleted_at,

            ci.city_id,
            ci.province_id,
            ci.city_name,
            ci.is_active,
            ci.created_at,
            ci.updated_at,

            COALESCE(pics.data, '[]')

        FROM client c
        LEFT JOIN city ci ON ci.city_id = c.city_id

        LEFT JOIN LATERAL (
            SELECT json_agg(json_build_object(
                'app_user_id', u.app_user_id,
                'app_user_name', u.app_user_name,
				'app_user_phone', u.app_user_phone,
                'username', u.username
            )) AS data
            FROM app_user u
            WHERE u.client_id = c.client_id
              AND u.deleted_at IS NULL
        ) pics ON TRUE

        WHERE c.deleted_at IS NULL
    `

	args := []any{}
	argPos := 1

	cursorClause, cursorArgs, argPos := helper.BuildCursorCondition(parsedOrderBy, tableKey, sort, cursorValue, cursorKey, argPos)
	baseQuery += cursorClause
	args = append(args, cursorArgs...)
	// Exact match filters
	exactFilters := []string{
		"city_id",
		"is_active",
	}

	for _, key := range exactFilters {
		if v, ok := filters[key]; ok && v != "" {
			baseQuery += fmt.Sprintf(" AND c.%s = $%d", key, argPos)
			args = append(args, v)
			argPos++
		}
	}

	// Time filters
	timeFilters := map[string]string{
		"created_at_after":  "c.created_at >= $%d",
		"created_at_before": "c.created_at <= $%d",
		"updated_at_after":  "c.updated_at >= $%d",
		"updated_at_before": "c.updated_at <= $%d",
	}

	for key, clause := range timeFilters {
		if v, ok := filters[key]; ok && v != "" {
			baseQuery += " AND " + fmt.Sprintf(clause, argPos)
			args = append(args, v)
			argPos++
		}
	}

	// Search by name
	if v, ok := filters["client_name"]; ok && v != "" {
		baseQuery += fmt.Sprintf(" AND c.client_name ILIKE $%d", argPos)
		args = append(args, "%"+v+"%")
		argPos++
	}

	// Order + limit
	baseQuery += fmt.Sprintf(" ORDER BY %s %s, %s %s LIMIT $%d", parsedOrderBy, sort, tableKey, sort, argPos)

	args = append(args, limit)

	rows, err := r.DB.Query(ctx, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []model.Client{}

	for rows.Next() {
		var m model.Client
		var picsJSON []byte
		var city model.CityNullable

		err := rows.Scan(
			&m.ClientID,
			&m.ClientName,
			&m.ClientEmail,
			&m.ClientTin,
			&m.NumberOfDayUntilDue,
			&m.CityID,
			&m.ClientAddress,
			&m.IsActive,
			&m.CreatedBy,
			&m.UpdatedBy,
			&m.CreatedAt,
			&m.UpdatedAt,
			&m.DeletedAt,

			&city.CityID,
			&city.ProvinceID,
			&city.CityName,
			&city.IsActive,
			&city.CreatedAt,
			&city.UpdatedAt,

			&picsJSON,
		)
		if err != nil {
			return nil, err
		}

		_ = json.Unmarshal(picsJSON, &m.ClientPics)

		if city.CityID != nil {
			m.City = city.ToNotNullable()
		}
		result = append(result, m)
	}

	return result, nil
}
