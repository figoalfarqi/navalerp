package repository

import (
	"context"
	"time"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DashboardRepository struct{ DB *pgxpool.Pool }

func NewDashboardRepository(db *pgxpool.Pool) *DashboardRepository {
	return &DashboardRepository{DB: db}
}

const dashboardTransportTotals = `
	WITH selected_projects AS (
		SELECT p.*
		FROM project p
		WHERE p.deleted_at IS NULL
		  AND ($1::int IS NULL OR p.project_id=$1)
	), transport_totals AS (
		SELECT pt.project_id,
			COUNT(*)::int AS transport_count,
			COUNT(*) FILTER (WHERE pt.is_completed=1)::int AS completed_transport_count,
			COALESCE(SUM(COALESCE(pt.delivered_volume_cubic,pt.loaded_volume_cubic,0)),0) AS volume_cubic,
			COALESCE(SUM(COALESCE(pt.delivered_weight_ton,pt.loaded_weight_ton,0)),0) AS weight_ton,
			COALESCE(SUM(pt.material_sale_amount + pt.transport_service_income_amount + pt.other_income_amount),0) AS total_income,
			COALESCE(SUM(pt.material_purchase_amount + pt.transport_expense_amount
				+ pt.road_money_amount + pt.loading_cost_amount + pt.unloading_cost_amount
				+ pt.fuel_cost_amount + pt.toll_cost_amount + pt.other_expense_amount),0) AS total_expense
		FROM project_transport pt
		JOIN selected_projects sp ON sp.project_id=pt.project_id
		WHERE pt.deleted_at IS NULL AND pt.transported_at >= $2 AND pt.transported_at < $3
		GROUP BY pt.project_id
	)
`

func (r *DashboardRepository) Summary(ctx context.Context, projectID *int, from, to time.Time) (*model.DashboardSummary, error) {
	return scanJSONRow[model.DashboardSummary](r.DB.QueryRow(ctx, dashboardTransportTotals+`
		SELECT jsonb_build_object(
			'project_count',COUNT(sp.project_id)::int,
			'transport_count',COALESCE(SUM(tt.transport_count),0)::int,
			'completed_transport_count',COALESCE(SUM(tt.completed_transport_count),0)::int,
			'volume_cubic',COALESCE(SUM(tt.volume_cubic),0),
			'weight_ton',COALESCE(SUM(tt.weight_ton),0),
			'total_income',COALESCE(SUM(COALESCE(tt.total_income,0)+sp.fixed_other_income),0),
			'total_expense',COALESCE(SUM(COALESCE(tt.total_expense,0)+sp.fixed_other_expense),0),
			'net_profit',COALESCE(SUM(
				COALESCE(tt.total_income,0)+sp.fixed_other_income
				-COALESCE(tt.total_expense,0)-sp.fixed_other_expense
			),0)
		)
		FROM selected_projects sp
		LEFT JOIN transport_totals tt ON tt.project_id=sp.project_id`,
		projectID, from, to))
}

func (r *DashboardRepository) Projects(ctx context.Context, projectID *int, from, to time.Time) ([]model.DashboardProject, error) {
	rows, err := r.DB.Query(ctx, dashboardTransportTotals+`
		SELECT jsonb_build_object(
			'project_id',sp.project_id,
			'project_code',sp.project_code,
			'project_name',sp.project_name,
			'project_count',1,
			'transport_count',COALESCE(tt.transport_count,0),
			'completed_transport_count',COALESCE(tt.completed_transport_count,0),
			'volume_cubic',COALESCE(tt.volume_cubic,0),
			'weight_ton',COALESCE(tt.weight_ton,0),
			'total_income',COALESCE(tt.total_income,0)+sp.fixed_other_income,
			'total_expense',COALESCE(tt.total_expense,0)+sp.fixed_other_expense,
			'net_profit',COALESCE(tt.total_income,0)+sp.fixed_other_income
				-COALESCE(tt.total_expense,0)-sp.fixed_other_expense
		)
		FROM selected_projects sp
		LEFT JOIN transport_totals tt ON tt.project_id=sp.project_id
		ORDER BY sp.project_name`, projectID, from, to)
	if err != nil {
		return nil, err
	}
	return scanJSONRows[model.DashboardProject](rows)
}

func (r *DashboardRepository) Daily(ctx context.Context, projectID *int, from, to time.Time) ([]model.DashboardDaily, error) {
	rows, err := r.DB.Query(ctx, `
		SELECT jsonb_build_object(
			'date',TO_CHAR(d.day,'YYYY-MM-DD'),
			'project_count',COALESCE(a.project_count,0),
			'transport_count',COALESCE(a.transport_count,0),
			'completed_transport_count',COALESCE(a.completed_transport_count,0),
			'volume_cubic',COALESCE(a.volume_cubic,0),
			'weight_ton',COALESCE(a.weight_ton,0),
			'total_income',COALESCE(a.total_income,0),
			'total_expense',COALESCE(a.total_expense,0),
			'net_profit',COALESCE(a.total_income,0)-COALESCE(a.total_expense,0)
		)
		FROM generate_series($2::date,($3::date-INTERVAL '1 day')::date,INTERVAL '1 day') d(day)
		LEFT JOIN (
			SELECT (pt.transported_at AT TIME ZONE 'Asia/Jakarta')::date AS day,
				COUNT(DISTINCT pt.project_id)::int AS project_count,
				COUNT(*)::int AS transport_count,
				COUNT(*) FILTER (WHERE pt.is_completed=1)::int AS completed_transport_count,
				SUM(COALESCE(pt.delivered_volume_cubic,pt.loaded_volume_cubic,0)) AS volume_cubic,
				SUM(COALESCE(pt.delivered_weight_ton,pt.loaded_weight_ton,0)) AS weight_ton,
				SUM(pt.material_sale_amount+pt.transport_service_income_amount+pt.other_income_amount) AS total_income,
				SUM(pt.material_purchase_amount+pt.transport_expense_amount+pt.road_money_amount
					+pt.loading_cost_amount+pt.unloading_cost_amount+pt.fuel_cost_amount
					+pt.toll_cost_amount+pt.other_expense_amount) AS total_expense
			FROM project_transport pt
			JOIN project p ON p.project_id=pt.project_id AND p.deleted_at IS NULL
			WHERE pt.deleted_at IS NULL AND pt.transported_at >= $2 AND pt.transported_at < $3
			  AND ($1::int IS NULL OR pt.project_id=$1)
			GROUP BY (pt.transported_at AT TIME ZONE 'Asia/Jakarta')::date
		) a ON a.day=d.day::date
		ORDER BY d.day`, projectID, from, to)
	if err != nil {
		return nil, err
	}
	return scanJSONRows[model.DashboardDaily](rows)
}
