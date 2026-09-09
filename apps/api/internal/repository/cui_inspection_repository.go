package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CuiInspectionRepository struct {
	DB *pgxpool.Pool
}

func NewCuiInspectionRepository(db *pgxpool.Pool) *CuiInspectionRepository {
	return &CuiInspectionRepository{DB: db}
}

// Get retrieves a single cui_inspection by ID
// Get retrieves a single cui_inspection by ID
func (r *CuiInspectionRepository) Get(ctx context.Context, id string) (*model.CuiInspection, error) {
	query := `SELECT 
		i.inspection_id, i.inspection_number, i.cui_asset_id, i.ship_id, i.inspection_date, 
		i.inspector_officer_id, i.method, i.condition_rating, i.findings, i.remedial_action_required, 
		i.next_inspection_date, i.created_by, i.updated_by, i.deleted_by, i.created_at, i.updated_at, i.deleted_at,
		a.asset_name, a.asset_code, s.ship_name, s.hull_number, p.full_name, p.nrp
	FROM cui_inspections i
	LEFT JOIN cui_assets a ON a.cui_asset_id = i.cui_asset_id
	LEFT JOIN mro_ships s ON s.ship_id = i.ship_id
	LEFT JOIN hcm_personnel p ON p.personnel_id = i.inspector_officer_id
	WHERE i.inspection_id = $1 AND i.deleted_at IS NULL`

	row := r.DB.QueryRow(ctx, query, id)
	var m model.CuiInspection
	err := row.Scan(&m.InspectionId, &m.InspectionNumber, &m.CuiAssetId, &m.ShipId, &m.InspectionDate, &m.InspectorOfficerId, &m.Method, &m.ConditionRating, &m.Findings, &m.RemedialActionRequired, &m.NextInspectionDate, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt, &m.AssetName, &m.AssetCode, &m.ShipName, &m.HullNumber, &m.InspectorName, &m.InspectorNrp)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves cui_inspection records with search, filter, and pagination
func (r *CuiInspectionRepository) List(ctx context.Context, opts model.ListOptions) ([]model.CuiInspection, int, error) {
	whereQuery := `FROM cui_inspections i
	LEFT JOIN cui_assets a ON a.cui_asset_id = i.cui_asset_id
	LEFT JOIN mro_ships s ON s.ship_id = i.ship_id
	LEFT JOIN hcm_personnel p ON p.personnel_id = i.inspector_officer_id
	WHERE 1=1 AND i.deleted_at IS NULL`
	var args []interface{}
	argIndex := 1

	if opts.Search != "" {
		whereQuery += fmt.Sprintf(" AND (i.inspection_number ILIKE $%d OR i.findings ILIKE $%d OR a.asset_name ILIKE $%d OR a.asset_code ILIKE $%d OR s.ship_name ILIKE $%d OR p.full_name ILIKE $%d)", argIndex, argIndex, argIndex, argIndex, argIndex, argIndex)
		args = append(args, "%"+opts.Search+"%")
		argIndex++
	}

	allowedCols := map[string]bool{
		"inspection_id":            true,
		"inspection_number":        true,
		"cui_asset_id":             true,
		"ship_id":                  true,
		"inspection_date":          true,
		"inspector_officer_id":     true,
		"method":                   true,
		"condition_rating":         true,
		"findings":                 true,
		"remedial_action_required": true,
		"next_inspection_date":     true,
		"created_at":               true,
		"updated_at":               true,
	}

	// Generic filter support with column whitelisting
	for k, v := range opts.Filters {
		if v != "" && allowedCols[k] {
			whereQuery += fmt.Sprintf(" AND i.%s = $%d", k, argIndex)
			args = append(args, v)
			argIndex++
		}
	}

	// Count total
	var total int
	countQuery := "SELECT COUNT(*) " + whereQuery
	err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Sorting and Pagination
	sortCol := "created_at"
	if opts.SortBy != "" && allowedCols[opts.SortBy] {
		sortCol = opts.SortBy
	}
	order := "DESC"
	if strings.ToUpper(opts.Order) == "ASC" {
		order = "ASC"
	}
	whereQuery += fmt.Sprintf(" ORDER BY i.%s %s", sortCol, order)

	if opts.Limit > 0 {
		whereQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
		args = append(args, opts.Limit, opts.Offset)
	}

	selectQuery := `SELECT 
		i.inspection_id, i.inspection_number, i.cui_asset_id, i.ship_id, i.inspection_date, 
		i.inspector_officer_id, i.method, i.condition_rating, i.findings, i.remedial_action_required, 
		i.next_inspection_date, i.created_by, i.updated_by, i.deleted_by, i.created_at, i.updated_at, i.deleted_at,
		a.asset_name, a.asset_code, s.ship_name, s.hull_number, p.full_name, p.nrp ` + whereQuery
	rows, err := r.DB.Query(ctx, selectQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.CuiInspection
	for rows.Next() {
		var m model.CuiInspection
		if err := rows.Scan(&m.InspectionId, &m.InspectionNumber, &m.CuiAssetId, &m.ShipId, &m.InspectionDate, &m.InspectorOfficerId, &m.Method, &m.ConditionRating, &m.Findings, &m.RemedialActionRequired, &m.NextInspectionDate, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt, &m.AssetName, &m.AssetCode, &m.ShipName, &m.HullNumber, &m.InspectorName, &m.InspectorNrp); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new cui_inspection
func (r *CuiInspectionRepository) Create(ctx context.Context, m *model.CuiInspection) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO cui_inspections (inspection_number, cui_asset_id, ship_id, inspection_date, inspector_officer_id, method, condition_rating, findings, remedial_action_required, next_inspection_date, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6::cui_inspection_method_enum, $7::cui_condition_enum, $8, $9, $10, $11, $12, $13) RETURNING inspection_id`
	var newID string
	err = tx.QueryRow(ctx, query, m.InspectionNumber, m.CuiAssetId, m.ShipId, m.InspectionDate, m.InspectorOfficerId, m.Method, m.ConditionRating, m.Findings, m.RemedialActionRequired, m.NextInspectionDate, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	return newID, tx.Commit(ctx)
}

// Update modifies an existing cui_inspection
func (r *CuiInspectionRepository) Update(ctx context.Context, id string, m *model.CuiInspection) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `UPDATE cui_inspections SET inspection_number = $1, cui_asset_id = $2, ship_id = $3, inspection_date = $4, inspector_officer_id = $5, method = $6::cui_inspection_method_enum, condition_rating = $7::cui_condition_enum, findings = $8, remedial_action_required = $9, next_inspection_date = $10, updated_by = $11, updated_at = CURRENT_TIMESTAMP WHERE inspection_id = $12`
	_, err = tx.Exec(ctx, query, m.InspectionNumber, m.CuiAssetId, m.ShipId, m.InspectionDate, m.InspectorOfficerId, m.Method, m.ConditionRating, m.Findings, m.RemedialActionRequired, m.NextInspectionDate, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes cui_inspection
func (r *CuiInspectionRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE cui_inspections SET deleted_at = CURRENT_TIMESTAMP WHERE inspection_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
