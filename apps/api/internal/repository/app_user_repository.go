package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/figoalfarqi/navalerp/internal/helper"
	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AppUserRepository struct {
	DB *pgxpool.Pool
}

func NewAppUserRepository(db *pgxpool.Pool) *AppUserRepository {
	return &AppUserRepository{DB: db}
}

func (r *AppUserRepository) Create(ctx context.Context, u *model.AppUser) (int, error) {
	query := `
		INSERT INTO app_user (
			app_role_id, client_id, bank_merk_id, username, password,
			app_user_status_id, app_user_name, app_user_preferred_name,
			app_user_phone, city_id, app_user_address,
			app_user_photo_url, id_card_photo_url, id_card_number,
			family_card_photo_url, family_card_number,
			driver_license_b_photo_url, driver_license_b_number, driver_license_b_expiry,
			tax_id_photo_url, tax_id_number,
			bpjs_photo_url, bpjs_number,
			bank_account_number, bank_account_name, salary_percentage,
			created_by, updated_by
		) VALUES (
			$1,$2,$3,$4,$5,
			$6,$7,$8,
			$9,$10,$11,$12,
			$13,$14,$15,
			$16,$17,
			$18,$19,$20,
			$21,$22,
			$23,$24,
			$25,$26,$27,
			$28
		)
		RETURNING app_user_id
	`

	// fmt.Println(helper.DebugSQL(query, []interface{}{
	// 	u.AppUserID,
	// 	u.AppRoleID, u.BankMerkID, u.Username, u.Password,
	// 	u.AppUserStatusID, u.AppUserName, u.AppUserPreferredName,
	// 	u.AppUserPhone, u.CityID, u.AppUserAddress,
	// 	u.AppUserPhotoURL, u.IDCardPhotoURL, u.IDCardNumber,
	// 	u.FamilyCardPhotoURL, u.FamilyCardNumber,
	// 	u.DriverLicenseBPhotoURL, u.DriverLicenseBNumber, u.DriverLicenseBExpiry,
	// 	u.TaxIDPhotoURL, u.TaxIDNumber,
	// 	u.BpjsPhotoURL, u.BpjsNumber,
	// 	u.BankAccountNumber, u.BankAccountName, u.SalaryPercentage,
	// 	u.CreatedBy, u.UpdatedBy}))
	var id int
	err := r.DB.QueryRow(ctx, query,
		u.AppRoleID, u.ClientID, u.BankMerkID, u.Username, u.Password,
		u.AppUserStatusID, u.AppUserName, u.AppUserPreferredName,
		u.AppUserPhone, u.CityID, u.AppUserAddress,
		u.AppUserPhotoURL, u.IDCardPhotoURL, u.IDCardNumber,
		u.FamilyCardPhotoURL, u.FamilyCardNumber,
		u.DriverLicenseBPhotoURL, u.DriverLicenseBNumber, u.DriverLicenseBExpiry,
		u.TaxIDPhotoURL, u.TaxIDNumber,
		u.BpjsPhotoURL, u.BpjsNumber,
		u.BankAccountNumber, u.BankAccountName, u.SalaryPercentage,
		u.CreatedBy, u.UpdatedBy,
	).Scan(&id)

	return id, err
}

func (r *AppUserRepository) CreateClientPicTx(ctx context.Context, tx pgx.Tx, u *model.AppUser) (int, error) {
	query := `
		INSERT INTO app_user (
			app_role_id, client_id, bank_merk_id, username, password,
			app_user_status_id, app_user_name, app_user_preferred_name,
			app_user_phone, city_id, app_user_address,
			app_user_photo_url,
			created_by, updated_by
		) VALUES (
			$1,$2,$3,$4,$5,
			$6,$7,$8,
			$9,$10,$11,
			$12,
			$13,$14
		)
		RETURNING app_user_id
	`

	var id int
	err := tx.QueryRow(ctx, query,
		u.AppRoleID, u.ClientID, u.BankMerkID, u.Username, u.Password,
		u.AppUserStatusID, u.AppUserName, u.AppUserPreferredName,
		u.AppUserPhone, u.CityID, u.AppUserAddress,
		u.AppUserPhotoURL,
		u.CreatedBy, u.UpdatedBy,
	).Scan(&id)

	return id, err
}

func (r *AppUserRepository) Update(ctx context.Context, id int, u *model.AppUser) error {

	setQuery := `
		client_id = $1,
		bank_merk_id = $2,
		app_user_name = $3,
		app_user_preferred_name = $4,
		app_user_phone = $5,
		city_id = $6,
		app_user_address = $7,
		id_card_number = $8,
		family_card_number = $9,
		driver_license_b_number = $10,
		driver_license_b_expiry = $11,
		tax_id_number = $12,
		bpjs_number = $13,
		bank_account_number = $14,
		bank_account_name = $15,
		salary_percentage = $16,
		app_user_photo_url = $17,
		id_card_photo_url = $18,
		family_card_photo_url = $19,
		driver_license_b_photo_url = $20,
		tax_id_photo_url = $21,
		bpjs_photo_url = $22
	`

	args := []interface{}{
		u.ClientID, u.BankMerkID, u.AppUserName,
		u.AppUserPreferredName, u.AppUserPhone,
		u.CityID, u.AppUserAddress, u.IDCardNumber,
		u.FamilyCardNumber,
		u.DriverLicenseBNumber, u.DriverLicenseBExpiry,
		u.TaxIDNumber,
		u.BpjsNumber,
		u.BankAccountNumber, u.BankAccountName, u.SalaryPercentage,
		u.AppUserPhotoURL, u.IDCardPhotoURL, u.FamilyCardPhotoURL,
		u.DriverLicenseBPhotoURL, u.TaxIDPhotoURL, u.BpjsPhotoURL,
	}

	argPos := len(args) + 1
	invalidateAuth := false

	// is_active conditional update

	// if u.AppUserPhotoURL != nil && *u.AppUserPhotoURL != "-1" {
	// 	setQuery += fmt.Sprintf(", app_user_photo_url = $%d", argPos)
	// 	args = append(args, *u.AppUserPhotoURL)
	// 	argPos++
	// }

	// if u.IDCardPhotoURL != nil && *u.IDCardPhotoURL != "-1" {
	// 	setQuery += fmt.Sprintf(", id_card_photo_url = $%d", argPos)
	// 	args = append(args, *u.IDCardPhotoURL)
	// 	argPos++
	// }

	// if u.FamilyCardPhotoURL != nil && *u.FamilyCardPhotoURL != "-1" {
	// 	setQuery += fmt.Sprintf(", family_card_photo_url = $%d", argPos)
	// 	args = append(args, *u.FamilyCardPhotoURL)
	// 	argPos++
	// }

	// if u.DriverLicenseBPhotoURL != nil && *u.DriverLicenseBPhotoURL != "-1" {
	// 	setQuery += fmt.Sprintf(", driver_license_b_photo_url = $%d", argPos)
	// 	args = append(args, *u.DriverLicenseBPhotoURL)
	// 	argPos++
	// }

	// if u.TaxIDPhotoURL != nil && *u.TaxIDPhotoURL != "-1" {
	// 	setQuery += fmt.Sprintf(", tax_id_photo_url = $%d", argPos)
	// 	args = append(args, *u.TaxIDPhotoURL)
	// 	argPos++
	// }

	// if u.BpjsPhotoURL != nil && *u.BpjsPhotoURL != "-1" {
	// 	setQuery += fmt.Sprintf(", bpjs_photo_url = $%d", argPos)
	// 	args = append(args, *u.BpjsPhotoURL)
	// 	argPos++
	// }

	if u.AppRoleID != -1 {
		setQuery += fmt.Sprintf(", app_role_id = $%d", argPos)
		args = append(args, u.AppRoleID)
		argPos++
		invalidateAuth = true
	}

	if u.AppUserStatusID != -1 {
		setQuery += fmt.Sprintf(", app_user_status_id = $%d", argPos)
		args = append(args, u.AppUserStatusID)
		argPos++
		invalidateAuth = true
	}
	if u.Username != "" {
		setQuery += fmt.Sprintf(", username = $%d", argPos)
		args = append(args, u.Username)
		argPos++
		invalidateAuth = true
	}
	if u.Password != "" {
		setQuery += fmt.Sprintf(", password = $%d", argPos)
		args = append(args, u.Password)
		argPos++
		invalidateAuth = true
	}
	if invalidateAuth {
		setQuery += ", auth_version = auth_version + 1"
	}

	args = append(args, u.UpdatedBy, u.UpdatedAt, id)

	query := fmt.Sprintf(`
		UPDATE app_user
		SET %s,
			updated_by = $%d,
			updated_at = $%d
		WHERE app_user_id = $%d AND deleted_at IS NULL
	`,
		setQuery,
		argPos,   // updated_by
		argPos+1, // updated_at
		argPos+2, // where id
	)

	_, err := r.DB.Exec(ctx, query, args...)
	return err
}

func (r *AppUserRepository) UpdateClientPicTx(ctx context.Context, tx pgx.Tx, id int, u *model.AppUser) error {

	setQuery := `
		bank_merk_id = $1,
		app_user_name = $2,
		app_user_preferred_name = $3,
		app_user_phone = $4,
		city_id = $5,
		app_user_address = $6,
		app_user_photo_url = $7
	`

	args := []interface{}{
		u.BankMerkID, u.AppUserName,
		u.AppUserPreferredName, u.AppUserPhone,
		u.CityID, u.AppUserAddress,
		u.AppUserPhotoURL,
	}

	argPos := len(args) + 1

	// Conditional status update
	if u.AppRoleID != -1 {
		setQuery += fmt.Sprintf(", app_role_id = $%d", argPos)
		args = append(args, u.AppRoleID)
		argPos++
	}

	if u.AppUserStatusID != -1 {
		setQuery += fmt.Sprintf(", app_user_status_id = $%d", argPos)
		args = append(args, u.AppUserStatusID)
		argPos++
	}

	args = append(args, u.UpdatedBy, u.UpdatedAt, id)

	query := fmt.Sprintf(`
		UPDATE app_user
		SET %s,
			updated_by = $%d,
			updated_at = $%d
		WHERE app_user_id = $%d AND deleted_at IS NULL
	`,
		setQuery,
		argPos,
		argPos+1,
		argPos+2,
	)

	_, err := tx.Exec(ctx, query, args...)
	return err
}

func (r *AppUserRepository) SoftDelete(ctx context.Context, deletedBy, id int) error {
	query := `
		UPDATE app_user
		SET deleted_at = $1,
			deleted_by = $2,
			app_user_status_id = 0,
			auth_version = auth_version + 1,
			updated_at = $1,
			updated_by = $2
		WHERE app_user_id = $3 AND deleted_at IS NULL
	`
	_, err := r.DB.Exec(ctx, query, time.Now(), deletedBy, id)
	return err
}

func (r *AppUserRepository) BatchSoftDeleteTx(ctx context.Context, tx pgx.Tx, deletedBy int, ids []int) error {
	if len(ids) == 0 {
		return nil
	}

	now := time.Now()

	query := `
		UPDATE app_user
		SET deleted_at = $1,
			deleted_by = $2,
			app_user_status_id = 0,
			updated_at = $1,
			updated_by = $2
		WHERE app_user_id = ANY($3) AND deleted_at IS NULL
	`

	_, err := tx.Exec(ctx, query, now, deletedBy, ids)
	return err
}

func (r *AppUserRepository) GetPasswordByID(ctx context.Context, id int) (*string, error) {
	query := `
	SELECT u.password
	FROM app_user u
	WHERE u.app_user_id = $1 AND u.deleted_at IS NULL
	`
	var password string
	err := r.DB.QueryRow(ctx, query, id).Scan(
		&password,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}
	return &password, nil
}

func (r *AppUserRepository) GetUserNameByID(ctx context.Context, id int) (*string, error) {

	query := `
	SELECT u.app_user_name
	FROM app_user u
	WHERE u.app_user_id = $1 AND u.deleted_at IS NULL
	`
	var userName string
	err := r.DB.QueryRow(ctx, query, id).Scan(
		&userName,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}
	return &userName, nil
}

func (r *AppUserRepository) GetByID(ctx context.Context, id int) (*model.AppUser, error) {

	query := `
	SELECT 
		u.app_user_id, u.app_role_id, u.client_id, u.bank_merk_id,
		u.username, u.app_user_status_id,
		u.app_user_name, u.app_user_preferred_name,
		u.app_user_phone, u.city_id,
		u.app_user_address, u.app_user_photo_url,
		u.id_card_photo_url, u.id_card_number,
		u.family_card_photo_url, u.family_card_number,
		u.driver_license_b_photo_url, u.driver_license_b_number, u.driver_license_b_expiry,
		u.tax_id_photo_url, u.tax_id_number,
		u.bpjs_photo_url, u.bpjs_number,
		u.bank_account_number, u.bank_account_name, u.salary_percentage,
		u.created_by, u.updated_by, u.deleted_by,
		u.created_at, u.updated_at, u.deleted_at,

		-- City
		c.city_id, c.city_name, c.province_id, c.is_active, c.created_at, c.updated_at, c.deleted_at,

		-- Role
		r2.app_role_id, r2.app_role_type_id, r2.app_role_name, r2.app_role_description,
		r2.created_at, r2.updated_at, r2.deleted_at,

		-- Bank Merk
		b.bank_merk_id, b.bank_merk_name, b.bank_merk_description, b.is_active,
		b.created_at, b.updated_at, b.deleted_at
	FROM app_user u
	LEFT JOIN city c ON u.city_id = c.city_id AND c.deleted_at IS NULL
	LEFT JOIN app_role r2 ON u.app_role_id = r2.app_role_id AND r2.deleted_at IS NULL
	LEFT JOIN bank_merk b ON u.bank_merk_id = b.bank_merk_id AND b.deleted_at IS NULL
	WHERE u.app_user_id = $1 AND u.deleted_at IS NULL
	`

	var (
		user model.AppUser
		city model.CityNullable
		role model.AppRole
		bank model.BankMerkNullable

		userCreatedBy, userUpdatedBy *int
		cityDel, roleDel, bankDel    *time.Time
	)

	err := r.DB.QueryRow(ctx, query, id).Scan(
		&user.AppUserID, &user.AppRoleID, &user.ClientID, &user.BankMerkID,
		&user.Username, &user.AppUserStatusID,
		&user.AppUserName, &user.AppUserPreferredName,
		&user.AppUserPhone, &user.CityID,
		&user.AppUserAddress, &user.AppUserPhotoURL,
		&user.IDCardPhotoURL, &user.IDCardNumber,
		&user.FamilyCardPhotoURL, &user.FamilyCardNumber,
		&user.DriverLicenseBPhotoURL, &user.DriverLicenseBNumber, &user.DriverLicenseBExpiry,
		&user.TaxIDPhotoURL, &user.TaxIDNumber,
		&user.BpjsPhotoURL, &user.BpjsNumber,
		&user.BankAccountNumber, &user.BankAccountName, &user.SalaryPercentage,
		&userCreatedBy, &userUpdatedBy, &user.DeletedBy,
		&user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,

		&city.CityID, &city.CityName, &city.ProvinceID, &city.IsActive, &city.CreatedAt, &city.UpdatedAt, &cityDel,

		&role.AppRoleID, &role.AppRoleTypeID, &role.AppRoleName, &role.AppRoleDescription, &role.CreatedAt, &role.UpdatedAt, &roleDel,

		&bank.BankMerkID, &bank.BankMerkName, &bank.BankMerkDescription, &bank.IsActive,

		&bank.CreatedAt, &bank.UpdatedAt, &bankDel,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}
	if userCreatedBy != nil {
		user.CreatedBy = *userCreatedBy
	}
	if userUpdatedBy != nil {
		user.UpdatedBy = *userUpdatedBy
	}

	if city.CityID != nil {
		city.DeletedAt = cityDel
		user.City = city.ToNotNullable()
	}

	if role.AppRoleID != 0 {
		role.DeletedAt = roleDel
		user.AppRole = &role
	}

	if bank.BankMerkID != nil {
		bank.DeletedAt = bankDel
		user.BankMerk = bank.ToNotNullable()
	}

	return &user, nil
}

// List returns list of AppUser with JOIN to city, app_role and bank_merk.
// Supports cursor pagination, LIKE filters, exact-match filters and time-range filters.
func (r *AppUserRepository) List(ctx context.Context,
	cursorValue interface{}, cursorKey *int, limit int, filters map[string]string,
	orderBy, sort string) ([]model.AppUser, error) {
	tableKey := "u.app_user_id"
	parsedOrderBy := "u." + orderBy

	baseQuery := `
		SELECT
			u.app_user_id,
			u.app_role_id,
			u.client_id,
			u.bank_merk_id,
			u.username,
			u.app_user_status_id,
			u.app_user_name,
			u.app_user_preferred_name,
			u.app_user_phone,
			u.city_id,
			u.app_user_address,
			u.app_user_photo_url,
			u.id_card_photo_url,
			u.id_card_number,
			u.family_card_photo_url,
			u.family_card_number,
			u.driver_license_b_photo_url,
			u.driver_license_b_number,
			u.driver_license_b_expiry,
			u.tax_id_photo_url,
			u.tax_id_number,
			u.bpjs_photo_url,
			u.bpjs_number,
			u.bank_account_number,
			u.bank_account_name,
			u.salary_percentage,
			u.created_by,
			u.updated_by,
			u.deleted_by,
			u.created_at,
			u.updated_at,
			u.deleted_at,

			-- city
			c.city_id,
			c.city_name,
			c.province_id,
			c.is_active,
			c.created_at,
			c.updated_at,
			c.deleted_at,

			-- province
			prov.province_id,
			prov.province_name,
			prov.province_real_name,
			prov.is_active,
			prov.created_at,
			prov.updated_at,
			prov.deleted_at,

			-- app_role
			r2.app_role_id,
			r2.app_role_type_id,
			r2.app_role_name,
			r2.app_role_description,
			r2.is_active,
			r2.created_at,
			r2.updated_at,
			r2.deleted_at,

			-- bank_merk
			b.bank_merk_id,
			b.bank_merk_name,
			b.bank_merk_description,
			b.is_active,
			b.created_at,
			b.updated_at,
			b.deleted_at
		FROM app_user u
		LEFT JOIN city c ON u.city_id = c.city_id AND c.deleted_at IS NULL
		LEFT JOIN province prov ON c.province_id = prov.province_id AND prov.deleted_at IS NULL
		LEFT JOIN app_role r2 ON u.app_role_id = r2.app_role_id AND r2.deleted_at IS NULL
		LEFT JOIN bank_merk b ON u.bank_merk_id = b.bank_merk_id AND b.deleted_at IS NULL
		WHERE u.deleted_at IS NULL
	`

	args := []interface{}{}
	argPos := 1

	// Cursor pagination
	cursorClause, cursorArgs, argPos := helper.BuildCursorCondition(parsedOrderBy, tableKey, sort, cursorValue, cursorKey, argPos)
	baseQuery += cursorClause
	args = append(args, cursorArgs...)
	// LIKE filters (text)
	likeFilters := map[string]string{
		"username":                "u.username ILIKE $%d",
		"app_user_name":           "u.app_user_name ILIKE $%d",
		"app_user_preferred_name": "u.app_user_preferred_name ILIKE $%d",
		"app_user_phone":          "u.app_user_phone ILIKE $%d",
		"app_user_address":        "u.app_user_address ILIKE $%d",
		"id_card_number":          "u.id_card_number ILIKE $%d",
		"driver_license_b_number": "u.driver_license_b_number ILIKE $%d",
		"tax_id_number":           "u.tax_id_number ILIKE $%d",
		"bpjs_number":             "u.bpjs_number ILIKE $%d",
		"bank_account_number":     "u.bank_account_number ILIKE $%d",
		"bank_account_name":       "u.bank_account_name ILIKE $%d",
	}
	for key, clause := range likeFilters {
		if v, ok := filters[key]; ok && v != "" {
			baseQuery += " AND " + fmt.Sprintf(clause, argPos)
			args = append(args, "%"+v+"%")
			argPos++
		}
	}

	// Exact match filters (IDs, numeric, status)
	exactFilters := []string{
		"app_role_id",
		"app_role_type_id",
		"bank_merk_id",
		"city_id",
		"client_id",
		"app_user_status_id",
		"created_by",
		"updated_by",
		"salary_percentage", // treat as exact match (provided by user)
		"usernameEXACT",
	}
	for _, key := range exactFilters {
		if v, ok := filters[key]; ok && v != "" {
			// For salary_percentage we reference u.salary_percentage, else u.<key>
			col := key
			if key == "salary_percentage" {
				col = "u.salary_percentage"
			} else if key == "app_role_type_id" {
				col = "r2.app_role_type_id"
			} else if key == "usernameEXACT" {
				col = "u.username"
			} else {
				col = "u." + key
			}
			baseQuery += fmt.Sprintf(" AND %s = $%d", col, argPos)
			args = append(args, v)
			argPos++
		}
	}

	// Exact match filters (IDs, numeric, status)
	exactNotFilters := []string{
		"app_user_idNOT",
	}
	for _, key := range exactNotFilters {
		if v, ok := filters[key]; ok && v != "" {
			// For salary_percentage we reference u.salary_percentage, else u.<key>
			col := key
			if key == "app_user_idNOT" {
				col = "u.app_user_id"
			} else {
				col = "u." + key
			}
			baseQuery += fmt.Sprintf(" AND %s != $%d", col, argPos)
			args = append(args, v)
			argPos++
		}
	}

	// Time-range filters
	timeFilters := map[string]string{
		"created_at_after":  "u.created_at >= $%d",
		"created_at_before": "u.created_at <= $%d",
		"updated_at_after":  "u.updated_at >= $%d",
		"updated_at_before": "u.updated_at <= $%d",
	}
	for key, clause := range timeFilters {
		if v, ok := filters[key]; ok && v != "" {
			baseQuery += " AND " + fmt.Sprintf(clause, argPos)
			args = append(args, v)
			argPos++
		}
	}

	// Final order + limit
	baseQuery += fmt.Sprintf(" ORDER BY %s %s, %s %s LIMIT $%d", parsedOrderBy, sort, tableKey, sort, argPos)
	args = append(args, limit)
	rows, err := r.DB.Query(ctx, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.AppUser

	for rows.Next() {
		var u model.AppUser

		var prov model.ProvinceNullable
		var city model.CityNullable
		var role model.AppRole
		var bank model.BankMerkNullable

		var userCreatedBy, userUpdatedBy *int
		var provDel, cityDel, roleDel, bankDel *time.Time

		if err := rows.Scan(
			&u.AppUserID,
			&u.AppRoleID,
			&u.ClientID,
			&u.BankMerkID,
			&u.Username,
			&u.AppUserStatusID,
			&u.AppUserName,
			&u.AppUserPreferredName,
			&u.AppUserPhone,
			&u.CityID,
			&u.AppUserAddress,
			&u.AppUserPhotoURL,
			&u.IDCardPhotoURL,
			&u.IDCardNumber,
			&u.FamilyCardPhotoURL,
			&u.FamilyCardNumber,
			&u.DriverLicenseBPhotoURL,
			&u.DriverLicenseBNumber,
			&u.DriverLicenseBExpiry,
			&u.TaxIDPhotoURL,
			&u.TaxIDNumber,
			&u.BpjsPhotoURL,
			&u.BpjsNumber,
			&u.BankAccountNumber,
			&u.BankAccountName,
			&u.SalaryPercentage,
			&userCreatedBy,
			&userUpdatedBy,
			&u.DeletedBy,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.DeletedAt,

			&city.CityID,
			&city.CityName,
			&city.ProvinceID,
			&city.IsActive,
			&city.CreatedAt,
			&city.UpdatedAt,
			&cityDel,

			&prov.ProvinceID,
			&prov.ProvinceName,
			&prov.ProvinceRealName,
			&prov.IsActive,
			&prov.CreatedAt,
			&prov.UpdatedAt,
			&provDel,

			&role.AppRoleID,
			&role.AppRoleTypeID,
			&role.AppRoleName,
			&role.AppRoleDescription,
			&role.IsActive,
			&role.CreatedAt,
			&role.UpdatedAt,
			&roleDel,

			&bank.BankMerkID,
			&bank.BankMerkName,
			&bank.BankMerkDescription,
			&bank.IsActive,
			&bank.CreatedAt,
			&bank.UpdatedAt,
			&bankDel,
		); err != nil {
			return nil, err
		}
		if userCreatedBy != nil {
			u.CreatedBy = *userCreatedBy
		}
		if userUpdatedBy != nil {
			u.UpdatedBy = *userUpdatedBy
		}

		// Attach join objects only if present (id != 0)
		if city.CityID != nil {
			city.DeletedAt = cityDel
			if prov.ProvinceID != nil {
				prov.DeletedAt = provDel
				city.Province = &prov
			}
			u.City = city.ToNotNullable()
		}
		if role.AppRoleID != 0 {
			role.DeletedAt = roleDel
			u.AppRole = &role
		}
		if bank.BankMerkID != nil {
			bank.DeletedAt = bankDel
			u.BankMerk = bank.ToNotNullable()
		}

		list = append(list, u)
	}

	return list, rows.Err()
}

func (r *AppUserRepository) GetByUsername(ctx context.Context, uname string) (*model.AppUser, error) {
	row := r.DB.QueryRow(ctx,
		`SELECT app_user_id, app_user_name, app_user_phone, app_user_photo_url, city_id, app_user_address,
		        app_user_status_id, u.app_role_id, username, password, u.created_at, u.updated_at, u.deleted_at,
		        r.app_role_name
		   FROM app_user u
		   JOIN app_role r ON r.app_role_id = u.app_role_id AND r.deleted_at IS NULL
		  WHERE u.deleted_at IS NULL AND username = $1`, uname)

	var u model.AppUser
	err := row.Scan(&u.AppUserID, &u.AppUserName, &u.AppUserPhone, &u.AppUserPhotoURL, &u.CityID, &u.AppUserAddress,
		&u.AppUserStatusID, &u.AppRoleID, &u.Username, &u.Password, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt,
		&u.LoginRoleName)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *AppUserRepository) UpdatePassword(ctx context.Context, id int, hashedPassword string) error {
	_, err := r.DB.Exec(ctx,
		`UPDATE app_user
		 SET password=$1,auth_version=auth_version+1,updated_at=$2
		 WHERE app_user_id=$3 AND deleted_at IS NULL`,
		hashedPassword, time.Now(), id,
	)
	return err
}

func (r *AppUserRepository) GetAuthVersion(ctx context.Context, id int) (int, error) {
	var version int
	err := r.DB.QueryRow(ctx, `
		SELECT auth_version FROM app_user
		WHERE app_user_id=$1 AND deleted_at IS NULL`, id).Scan(&version)
	return version, err
}

func (r *AppUserRepository) InvalidateAuth(ctx context.Context, id int) error {
	_, err := r.DB.Exec(ctx, `
		UPDATE app_user
		SET auth_version=auth_version+1,updated_at=CURRENT_TIMESTAMP
		WHERE app_user_id=$1 AND deleted_at IS NULL`, id)
	return err
}
