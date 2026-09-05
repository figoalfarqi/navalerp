package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PlatformTcoRepository struct {
	DB *pgxpool.Pool
}

func NewPlatformTcoRepository(db *pgxpool.Pool) *PlatformTcoRepository {
	return &PlatformTcoRepository{DB: db}
}

// Get retrieves a single platform_tco by tco_id
func (r *PlatformTcoRepository) Get(ctx context.Context, id string) (*model.PlatformTco, error) {
	query := `SELECT tco_id, ship_id, fiscal_year, acquisition_amortization, fuel_lube_cost, mro_spareparts_cost, docking_services_cost, crew_payroll_allowances, modernization_upgrades_cost, total_annual_operating_cost, operating_hours_sea, cost_per_operating_hour, remarks, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM fin_platform_tco_summaries WHERE tco_id = $1 AND deleted_at IS NULL`

	var m model.PlatformTco
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.TcoId, &m.ShipId, &m.FiscalYear, &m.AcquisitionAmortization, &m.FuelLubeCost, &m.MroSparepartsCost, &m.DockingServicesCost, &m.CrewPayrollAllowances, &m.ModernizationUpgradesCost, &m.TotalAnnualOperatingCost, &m.OperatingHoursSea, &m.CostPerOperatingHour, &m.Remarks, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated platform_tco records
func (r *PlatformTcoRepository) List(ctx context.Context, opts model.ListOptions) ([]model.PlatformTco, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(remarks ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM fin_platform_tco_summaries WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT tco_id, ship_id, fiscal_year, acquisition_amortization, fuel_lube_cost, mro_spareparts_cost, docking_services_cost, crew_payroll_allowances, modernization_upgrades_cost, total_annual_operating_cost, operating_hours_sea, cost_per_operating_hour, remarks, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM fin_platform_tco_summaries WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.PlatformTco
	for rows.Next() {
		var m model.PlatformTco
		if err := rows.Scan(&m.TcoId, &m.ShipId, &m.FiscalYear, &m.AcquisitionAmortization, &m.FuelLubeCost, &m.MroSparepartsCost, &m.DockingServicesCost, &m.CrewPayrollAllowances, &m.ModernizationUpgradesCost, &m.TotalAnnualOperatingCost, &m.OperatingHoursSea, &m.CostPerOperatingHour, &m.Remarks, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new platform_tco with optional child items
func (r *PlatformTcoRepository) Create(ctx context.Context, m *model.PlatformTco) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO fin_platform_tco_summaries (ship_id, fiscal_year, acquisition_amortization, fuel_lube_cost, mro_spareparts_cost, docking_services_cost, crew_payroll_allowances, modernization_upgrades_cost, operating_hours_sea, cost_per_operating_hour, remarks, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING tco_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.ShipId, m.FiscalYear, m.AcquisitionAmortization, m.FuelLubeCost, m.MroSparepartsCost, m.DockingServicesCost, m.CrewPayrollAllowances, m.ModernizationUpgradesCost, m.OperatingHoursSea, m.CostPerOperatingHour, m.Remarks, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing platform_tco
func (r *PlatformTcoRepository) Update(ctx context.Context, id string, m *model.PlatformTco) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE fin_platform_tco_summaries SET ship_id = $1, fiscal_year = $2, acquisition_amortization = $3, fuel_lube_cost = $4, mro_spareparts_cost = $5, docking_services_cost = $6, crew_payroll_allowances = $7, modernization_upgrades_cost = $8, operating_hours_sea = $9, cost_per_operating_hour = $10, remarks = $11, updated_by = $12, updated_at = CURRENT_TIMESTAMP WHERE tco_id = $13`
	_, err = tx.Exec(ctx, updateQuery, m.ShipId, m.FiscalYear, m.AcquisitionAmortization, m.FuelLubeCost, m.MroSparepartsCost, m.DockingServicesCost, m.CrewPayrollAllowances, m.ModernizationUpgradesCost, m.OperatingHoursSea, m.CostPerOperatingHour, m.Remarks, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes platform_tco
func (r *PlatformTcoRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE fin_platform_tco_summaries SET deleted_at = CURRENT_TIMESTAMP WHERE tco_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
