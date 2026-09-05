package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/figoalfarqi/apipml/internal/helper"
	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CityRepository struct {
	DB *pgxpool.Pool
}

func NewCityRepository(db *pgxpool.Pool) *CityRepository {
	return &CityRepository{DB: db}
}

// ==================================================
// Create City
// ==================================================
func (r *CityRepository) Create(ctx context.Context, c *model.City) (int, error) {
	query := `
		INSERT INTO city (
			province_id,
			city_name,
			is_active,
			created_by,
			updated_by
		) VALUES ($1,$2,$3,$4,$5)
		RETURNING city_id
	`

	var id int
	err := r.DB.QueryRow(ctx, query,
		c.ProvinceID,
		c.CityName,
		c.IsActive,
		c.CreatedBy,
		c.UpdatedBy,
	).Scan(&id)

	return id, err
}

// ==================================================
// Update City (Conditional is_active like Client)
// ==================================================
func (r *CityRepository) Update(ctx context.Context, id int, c *model.City) error {
	setIsActive := ""
	args := []interface{}{
		c.ProvinceID, // $1
		c.CityName,   // $2
	}
	argPos := 3

	// If IsActive != -1 → update is_active
	if c.IsActive != -1 {
		setIsActive = fmt.Sprintf(", is_active = $%d", argPos)
		args = append(args, c.IsActive)
		argPos++
	}

	args = append(args,
		c.UpdatedBy, // $argPos
		c.UpdatedAt, // $argPos+1
		id,          // $argPos+2
	)

	query := fmt.Sprintf(`
		UPDATE city
		SET province_id = $1,
			city_name = $2
			%s,
			updated_by = $%d,
			updated_at = $%d
		WHERE city_id = $%d AND deleted_at IS NULL
	`,
		setIsActive,
		argPos,   // updated_by
		argPos+1, // updated_at
		argPos+2, // WHERE city_id
	)

	_, err := r.DB.Exec(ctx, query, args...)
	return err
}

// ==================================================
// Soft Delete City (same pattern as Client)
// ==================================================
func (r *CityRepository) SoftDelete(ctx context.Context, deletedBy, id int) error {
	query := `
		UPDATE city
		SET deleted_at = $1,
			deleted_by = $2,
			is_active = 0,
			updated_at = $1,
			updated_by = $2
		WHERE city_id = $3 AND deleted_at IS NULL
	`
	_, err := r.DB.Exec(ctx, query, time.Now(), deletedBy, id)
	return err
}

// ==================================================
// Get By ID (JOIN Province)
// ==================================================
func (r *CityRepository) GetByID(ctx context.Context, id int) (*model.City, error) {
	query := `
		SELECT 
			c.city_id,
			c.province_id,
			c.city_name,
			c.is_active,
			c.created_by,
			c.updated_by,
			c.deleted_by,
			c.created_at,
			c.updated_at,
			c.deleted_at,

			p.province_id,
			p.province_name,
			p.province_real_name,
			p.created_at,
			p.updated_at,
			p.deleted_at
		FROM city c
		LEFT JOIN province p ON c.province_id = p.province_id AND p.deleted_at IS NULL
		WHERE c.city_id = $1 AND c.deleted_at IS NULL
	`

	var city model.City
	var province model.Province
	var provinceDeletedAt *time.Time

	err := r.DB.QueryRow(ctx, query, id).Scan(
		&city.CityID,
		&city.ProvinceID,
		&city.CityName,
		&city.IsActive,
		&city.CreatedBy,
		&city.UpdatedBy,
		&city.DeletedBy,
		&city.CreatedAt,
		&city.UpdatedAt,
		&city.DeletedAt,

		&province.ProvinceID,
		&province.ProvinceName,
		&province.ProvinceRealName,
		&province.CreatedAt,
		&province.UpdatedAt,
		&provinceDeletedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	if province.ProvinceID != 0 {
		province.DeletedAt = provinceDeletedAt
		city.Province = &province
	}

	return &city, nil
}

// ==================================================
// List (Cursor Pagination + Filters)
// ==================================================
func (r *CityRepository) List(ctx context.Context,
	cursorValue interface{}, cursorKey *int, limit int, filters map[string]string,
	orderBy, sort string) ([]model.City, error) {
	tableKey := "c.city_id"
	parsedOrderBy := "c." + orderBy
	baseQuery := `
		SELECT 
			c.city_id,
			c.province_id,
			c.city_name,
			c.is_active,
			c.created_by,
			c.updated_by,
			c.deleted_by,
			c.created_at,
			c.updated_at,
			c.deleted_at,

			p.province_id,
			p.province_name,
			p.province_real_name,
			p.created_at,
			p.updated_at,
			p.deleted_at
		FROM city c
		LEFT JOIN province p ON c.province_id = p.province_id AND p.deleted_at IS NULL
		WHERE c.deleted_at IS NULL
	`

	args := []interface{}{}
	argPos := 1

	cursorClause, cursorArgs, argPos := helper.BuildCursorCondition(parsedOrderBy, tableKey, sort, cursorValue, cursorKey, argPos)
	baseQuery += cursorClause
	args = append(args, cursorArgs...)

	// LIKE filter
	if v, ok := filters["city_name"]; ok && v != "" {
		baseQuery += fmt.Sprintf(" AND c.city_name ILIKE $%d", argPos)
		args = append(args, "%"+v+"%")
		argPos++
	}

	// exact filters
	exactFilters := []string{"province_id", "created_by", "updated_by", "is_active"}
	for _, key := range exactFilters {
		if v, ok := filters[key]; ok && v != "" {
			baseQuery += fmt.Sprintf(" AND c.%s = $%d", key, argPos)
			args = append(args, v)
			argPos++
		}
	}

	// time filters
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

	// Order & limit
	baseQuery += fmt.Sprintf(" ORDER BY %s %s, %s %s LIMIT $%d", parsedOrderBy, sort, tableKey, sort, argPos)
	args = append(args, limit)
	rows, err := r.DB.Query(ctx, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.City

	for rows.Next() {
		var city model.City
		var province model.Province
		var provDeletedAt *time.Time

		err := rows.Scan(
			&city.CityID,
			&city.ProvinceID,
			&city.CityName,
			&city.IsActive,
			&city.CreatedBy,
			&city.UpdatedBy,
			&city.DeletedBy,
			&city.CreatedAt,
			&city.UpdatedAt,
			&city.DeletedAt,

			&province.ProvinceID,
			&province.ProvinceName,
			&province.ProvinceRealName,
			&province.CreatedAt,
			&province.UpdatedAt,
			&provDeletedAt,
		)
		if err != nil {
			return nil, err
		}
		if province.ProvinceID != 0 {
			province.DeletedAt = provDeletedAt
			city.Province = &province
		}
		list = append(list, city)
	}
	return list, rows.Err()
}
