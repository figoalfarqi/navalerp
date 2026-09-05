package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MaterialRepository struct {
	DB *pgxpool.Pool
}

func NewMaterialRepository(db *pgxpool.Pool) *MaterialRepository {
	return &MaterialRepository{DB: db}
}

// Get retrieves a single material by material_id
func (r *MaterialRepository) Get(ctx context.Context, id string) (*model.Material, error) {
	query := `SELECT material_id, material_code, nsn, part_number, oem_name, material_name, category, uom, weight_kg, min_stock_level, max_stock_level, reorder_point, safety_stock, shelf_life_days, is_controlled_item, unit_price_idr, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM inv_materials WHERE material_id = $1 AND deleted_at IS NULL`

	var m model.Material
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.MaterialId, &m.MaterialCode, &m.Nsn, &m.PartNumber, &m.OemName, &m.MaterialName, &m.Category, &m.Uom, &m.WeightKg, &m.MinStockLevel, &m.MaxStockLevel, &m.ReorderPoint, &m.SafetyStock, &m.ShelfLifeDays, &m.IsControlledItem, &m.UnitPriceIdr, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	// Load EquipmentLinks
	childRowsEquipmentLinks, err := r.DB.Query(ctx, `SELECT link_id, material_id, system_id, equipment_id, is_mandatory_spare, interchangeability_code, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM inv_material_equipment_links WHERE material_id = $1`, id)
	if err == nil {
		defer childRowsEquipmentLinks.Close()
		for childRowsEquipmentLinks.Next() {
			var item model.MaterialEquipmentLinks
			if err := childRowsEquipmentLinks.Scan(&item.LinkId, &item.MaterialId, &item.SystemId, &item.EquipmentId, &item.IsMandatorySpare, &item.InterchangeabilityCode, &item.CreatedBy, &item.UpdatedBy, &item.DeletedBy, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt); err == nil {
				m.EquipmentLinks = append(m.EquipmentLinks, item)
			}
		}
	}

	return &m, nil
}

// List retrieves paginated material records
func (r *MaterialRepository) List(ctx context.Context, opts model.ListOptions) ([]model.Material, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(material_code ILIKE $%[1]d OR nsn ILIKE $%[1]d OR part_number ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM inv_materials WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT material_id, material_code, nsn, part_number, oem_name, material_name, category, uom, weight_kg, min_stock_level, max_stock_level, reorder_point, safety_stock, shelf_life_days, is_controlled_item, unit_price_idr, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM inv_materials WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.Material
	for rows.Next() {
		var m model.Material
		if err := rows.Scan(&m.MaterialId, &m.MaterialCode, &m.Nsn, &m.PartNumber, &m.OemName, &m.MaterialName, &m.Category, &m.Uom, &m.WeightKg, &m.MinStockLevel, &m.MaxStockLevel, &m.ReorderPoint, &m.SafetyStock, &m.ShelfLifeDays, &m.IsControlledItem, &m.UnitPriceIdr, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new material with optional child items
func (r *MaterialRepository) Create(ctx context.Context, m *model.Material) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO inv_materials (material_code, nsn, part_number, oem_name, material_name, category, uom, weight_kg, min_stock_level, max_stock_level, reorder_point, safety_stock, shelf_life_days, is_controlled_item, unit_price_idr, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6::material_category_type, $7::uom_type, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18) RETURNING material_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.MaterialCode, m.Nsn, m.PartNumber, m.OemName, m.MaterialName, m.Category, m.Uom, m.WeightKg, m.MinStockLevel, m.MaxStockLevel, m.ReorderPoint, m.SafetyStock, m.ShelfLifeDays, m.IsControlledItem, m.UnitPriceIdr, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.EquipmentLinks {
		_, err := tx.Exec(ctx, `INSERT INTO inv_material_equipment_links (material_id, system_id, equipment_id, is_mandatory_spare, interchangeability_code, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, newID, item.SystemId, item.EquipmentId, item.IsMandatorySpare, item.InterchangeabilityCode, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing material
func (r *MaterialRepository) Update(ctx context.Context, id string, m *model.Material) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE inv_materials SET material_code = $1, nsn = $2, part_number = $3, oem_name = $4, material_name = $5, category = $6::material_category_type, uom = $7::uom_type, weight_kg = $8, min_stock_level = $9, max_stock_level = $10, reorder_point = $11, safety_stock = $12, shelf_life_days = $13, is_controlled_item = $14, unit_price_idr = $15, updated_by = $16, updated_at = CURRENT_TIMESTAMP WHERE material_id = $17`
	_, err = tx.Exec(ctx, updateQuery, m.MaterialCode, m.Nsn, m.PartNumber, m.OemName, m.MaterialName, m.Category, m.Uom, m.WeightKg, m.MinStockLevel, m.MaxStockLevel, m.ReorderPoint, m.SafetyStock, m.ShelfLifeDays, m.IsControlledItem, m.UnitPriceIdr, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	if len(m.EquipmentLinks) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM inv_material_equipment_links WHERE material_id = $1`, id)
		for _, item := range m.EquipmentLinks {
			_, err := tx.Exec(ctx, `INSERT INTO inv_material_equipment_links (material_id, system_id, equipment_id, is_mandatory_spare, interchangeability_code, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, id, item.SystemId, item.EquipmentId, item.IsMandatorySpare, item.InterchangeabilityCode, item.CreatedBy, item.UpdatedBy, item.DeletedBy)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes material
func (r *MaterialRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE inv_materials SET deleted_at = CURRENT_TIMESTAMP WHERE material_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
