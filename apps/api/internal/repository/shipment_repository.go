package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ShipmentRepository struct {
	DB *pgxpool.Pool
}

func NewShipmentRepository(db *pgxpool.Pool) *ShipmentRepository {
	return &ShipmentRepository{DB: db}
}

// Get retrieves a single shipment by shipment_id
func (r *ShipmentRepository) Get(ctx context.Context, id string) (*model.Shipment, error) {
	query := `SELECT t.shipment_id, t.manifest_number, t.route_id, t.transport_unit_id, t.origin_warehouse_id, COALESCE(j_owh.warehouse_name, ''), t.destination_warehouse_id, COALESCE(j_dwh.warehouse_name, ''), t.departure_date, t.arrival_date, t.escort_security_level, t.status, t.authorized_by_user_id, t.remarks, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM log_shipments t
	LEFT JOIN inv_warehouses j_owh ON j_owh.warehouse_id = t.origin_warehouse_id
	LEFT JOIN inv_warehouses j_dwh ON j_dwh.warehouse_id = t.destination_warehouse_id
	WHERE t.shipment_id = $1 AND t.deleted_at IS NULL`

	var m model.Shipment
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.ShipmentId, &m.ManifestNumber, &m.RouteId, &m.TransportUnitId, &m.OriginWarehouseId, &m.OriginWarehouseName, &m.DestinationWarehouseId, &m.DestinationWarehouseName, &m.DepartureDate, &m.ArrivalDate, &m.EscortSecurityLevel, &m.Status, &m.AuthorizedByUserId, &m.Remarks, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	// Load Items
	childRowsItems, err := r.DB.Query(ctx, `SELECT shipment_item_id, shipment_id, material_id, quantity_dispatched, quantity_received, packaging_type, weight_kg, notes, created_at FROM log_shipment_items WHERE shipment_id = $1`, id)
	if err == nil {
		defer childRowsItems.Close()
		for childRowsItems.Next() {
			var item model.ShipmentItems
			if err := childRowsItems.Scan(&item.ShipmentItemId, &item.ShipmentId, &item.MaterialId, &item.QuantityDispatched, &item.QuantityReceived, &item.PackagingType, &item.WeightKg, &item.Notes, &item.CreatedAt); err == nil {
				m.Items = append(m.Items, item)
			}
		}
	}

	return &m, nil
}

// List retrieves paginated shipment records
func (r *ShipmentRepository) List(ctx context.Context, opts model.ListOptions) ([]model.Shipment, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "t.deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(manifest_number ILIKE $%[1]d OR remarks ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM log_shipments t WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf(`SELECT t.shipment_id, t.manifest_number, t.route_id, t.transport_unit_id, t.origin_warehouse_id, COALESCE(j_owh.warehouse_name, ''), t.destination_warehouse_id, COALESCE(j_dwh.warehouse_name, ''), t.departure_date, t.arrival_date, t.escort_security_level, t.status, t.authorized_by_user_id, t.remarks, t.created_by, t.updated_by, t.deleted_by, t.created_at, t.updated_at, t.deleted_at
	FROM log_shipments t
	LEFT JOIN inv_warehouses j_owh ON j_owh.warehouse_id = t.origin_warehouse_id
	LEFT JOIN inv_warehouses j_dwh ON j_dwh.warehouse_id = t.destination_warehouse_id
	WHERE %s ORDER BY t.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.Shipment
	for rows.Next() {
		var m model.Shipment
		if err := rows.Scan(&m.ShipmentId, &m.ManifestNumber, &m.RouteId, &m.TransportUnitId, &m.OriginWarehouseId, &m.OriginWarehouseName, &m.DestinationWarehouseId, &m.DestinationWarehouseName, &m.DepartureDate, &m.ArrivalDate, &m.EscortSecurityLevel, &m.Status, &m.AuthorizedByUserId, &m.Remarks, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new shipment with optional child items
func (r *ShipmentRepository) Create(ctx context.Context, m *model.Shipment) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO log_shipments (manifest_number, route_id, transport_unit_id, origin_warehouse_id, destination_warehouse_id, departure_date, arrival_date, escort_security_level, status, authorized_by_user_id, remarks, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8::escort_security_level_type, $9::shipment_status_type, $10, $11, $12, $13, $14) RETURNING shipment_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.ManifestNumber, m.RouteId, m.TransportUnitId, m.OriginWarehouseId, m.DestinationWarehouseId, m.DepartureDate, m.ArrivalDate, m.EscortSecurityLevel, m.Status, m.AuthorizedByUserId, m.Remarks, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.Items {
		_, err := tx.Exec(ctx, `INSERT INTO log_shipment_items (shipment_id, material_id, quantity_dispatched, quantity_received, packaging_type, weight_kg, notes) VALUES ($1, $2, $3, $4, $5::shipment_packaging_type, $6, $7)`, newID, item.MaterialId, item.QuantityDispatched, item.QuantityReceived, item.PackagingType, item.WeightKg, item.Notes)
		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing shipment
func (r *ShipmentRepository) Update(ctx context.Context, id string, m *model.Shipment) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE log_shipments SET manifest_number = $1, route_id = $2, transport_unit_id = $3, origin_warehouse_id = $4, destination_warehouse_id = $5, departure_date = $6, arrival_date = $7, escort_security_level = $8::escort_security_level_type, status = $9::shipment_status_type, authorized_by_user_id = $10, remarks = $11, updated_by = $12, updated_at = CURRENT_TIMESTAMP WHERE shipment_id = $13`
	_, err = tx.Exec(ctx, updateQuery, m.ManifestNumber, m.RouteId, m.TransportUnitId, m.OriginWarehouseId, m.DestinationWarehouseId, m.DepartureDate, m.ArrivalDate, m.EscortSecurityLevel, m.Status, m.AuthorizedByUserId, m.Remarks, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	if len(m.Items) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM log_shipment_items WHERE shipment_id = $1`, id)
		for _, item := range m.Items {
			_, err := tx.Exec(ctx, `INSERT INTO log_shipment_items (shipment_id, material_id, quantity_dispatched, quantity_received, packaging_type, weight_kg, notes) VALUES ($1, $2, $3, $4, $5::shipment_packaging_type, $6, $7)`, id, item.MaterialId, item.QuantityDispatched, item.QuantityReceived, item.PackagingType, item.WeightKg, item.Notes)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes shipment
func (r *ShipmentRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE log_shipments SET deleted_at = CURRENT_TIMESTAMP WHERE shipment_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
