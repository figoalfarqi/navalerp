package repository

import (
	"context"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DashboardRepository struct {
	DB *pgxpool.Pool
}

func NewDashboardRepository(db *pgxpool.Pool) *DashboardRepository {
	return &DashboardRepository{DB: db}
}

func (r *DashboardRepository) GetDashboardData(ctx context.Context) (*model.DashboardResponse, error) {
	var summary model.DashboardSummary

	// Total ships and readiness
	_ = r.DB.QueryRow(ctx, `
		SELECT 
			COUNT(*),
			COUNT(*) FILTER (WHERE status = 'ACTIVE'),
			COUNT(*) FILTER (WHERE current_readiness_status = 'FULLY_MISSION_CAPABLE')
		FROM mro_ships WHERE deleted_at IS NULL
	`).Scan(&summary.TotalShips, &summary.ActiveShips, &summary.FullyMissionCapable)

	// Personnel
	_ = r.DB.QueryRow(ctx, `SELECT COUNT(*) FROM hcm_personnel WHERE deleted_at IS NULL`).Scan(&summary.TotalPersonnel)

	// Warehouses
	_ = r.DB.QueryRow(ctx, `SELECT COUNT(*) FROM inv_warehouses WHERE deleted_at IS NULL`).Scan(&summary.TotalWarehouses)

	// Active missions
	_ = r.DB.QueryRow(ctx, `SELECT COUNT(*) FROM ops_missions WHERE mission_status = 'ACTIVE' AND deleted_at IS NULL`).Scan(&summary.ActiveMissions)

	// Open work orders
	_ = r.DB.QueryRow(ctx, `SELECT COUNT(*) FROM mro_work_orders WHERE status NOT IN ('COMPLETED', 'CANCELLED') AND deleted_at IS NULL`).Scan(&summary.OpenWorkOrders)

	if summary.TotalShips > 0 {
		summary.OverallReadinessScore = float64(summary.FullyMissionCapable) / float64(summary.TotalShips) * 100.0
	} else {
		summary.OverallReadinessScore = 100.0
	}

	// Ship readiness list
	rows, err := r.DB.Query(ctx, `
		SELECT 
			s.ship_id, s.ship_name, s.hull_number, s.current_readiness_status,
			COALESCE(rs.composite_readiness_index, 90.0)
		FROM mro_ships s
		LEFT JOIN LATERAL (
			SELECT composite_readiness_index 
			FROM ops_ship_readiness_snapshots 
			WHERE ship_id = s.ship_id 
			ORDER BY snapshot_timestamp DESC 
			LIMIT 1
		) rs ON true
		WHERE s.deleted_at IS NULL
		ORDER BY s.hull_number
	`)
	if err != nil {
		return &model.DashboardResponse{Summary: summary}, nil
	}
	defer rows.Close()

	var ships []model.DashboardShipReadiness
	for rows.Next() {
		var s model.DashboardShipReadiness
		if err := rows.Scan(&s.ShipID, &s.ShipName, &s.HullNumber, &s.ReadinessStatus, &s.CompositeScore); err == nil {
			ships = append(ships, s)
		}
	}

	return &model.DashboardResponse{
		Summary: summary,
		Ships:   ships,
	}, nil
}
