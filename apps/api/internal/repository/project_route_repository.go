package repository

import (
	"context"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectRouteRepository struct{ DB *pgxpool.Pool }

const projectRouteJSON = `
	to_jsonb(r) || jsonb_build_object(
		'project',jsonb_build_object(
			'project_id',p.project_id,
			'project_code',p.project_code,
			'project_name',p.project_name
		)
	)`

func NewProjectRouteRepository(db *pgxpool.Pool) *ProjectRouteRepository {
	return &ProjectRouteRepository{DB: db}
}

func routeAmount(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

func (r *ProjectRouteRepository) Create(ctx context.Context, req *model.ProjectRouteRequest, userID int) (int, error) {
	active := 1
	if req.IsActive != nil {
		active = *req.IsActive
	}
	var id int
	err := r.DB.QueryRow(ctx, `
		INSERT INTO project_route (
			project_id,project_pattern,route_sequence,route_type,route_name,distance_km,
			transport_service_unit,transport_service_price_per_cubic,transport_service_price_per_ton,
			transport_service_price_per_transport,transport_cost_unit,transport_cost_per_cubic,
			transport_cost_per_ton,transport_cost_per_transport,road_money_per_transport,
			loading_cost_per_transport,unloading_cost_per_transport,fuel_cost_per_transport,
			toll_cost_per_transport,other_income_per_transport,other_expense_per_transport,
			route_note,is_active,created_by,updated_by
		) VALUES (
			$1,COALESCE((
				SELECT p.route_type
				FROM project p
				WHERE p.project_id=$1 AND p.deleted_at IS NULL
			),$2),$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,
			$19,$20,$21,$22,$23,$24,$24
		) RETURNING project_route_id`,
		req.ProjectID, req.ProjectPattern, req.RouteSequence, req.RouteType, req.RouteName,
		req.DistanceKM, req.TransportServiceUnit, req.TransportServicePricePerCubic,
		req.TransportServicePricePerTon, req.TransportServicePricePerTransport,
		req.TransportCostUnit, req.TransportCostPerCubic, req.TransportCostPerTon,
		req.TransportCostPerTransport, routeAmount(req.RoadMoneyPerTransport),
		routeAmount(req.LoadingCostPerTransport), routeAmount(req.UnloadingCostPerTransport),
		routeAmount(req.FuelCostPerTransport), routeAmount(req.TollCostPerTransport),
		routeAmount(req.OtherIncomePerTransport), routeAmount(req.OtherExpensePerTransport),
		req.RouteNote, active, userID,
	).Scan(&id)
	return id, err
}

func (r *ProjectRouteRepository) Update(ctx context.Context, id int, req *model.ProjectRouteRequest, userID int) error {
	active := 1
	if req.IsActive != nil {
		active = *req.IsActive
	}
	_, err := r.DB.Exec(ctx, `
		UPDATE project_route SET
			project_id=$1,
			project_pattern=COALESCE((
				SELECT p.route_type
				FROM project p
				WHERE p.project_id=$1 AND p.deleted_at IS NULL
			),$2),
			route_sequence=$3,route_type=$4,route_name=$5,
			distance_km=$6,transport_service_unit=$7,transport_service_price_per_cubic=$8,
			transport_service_price_per_ton=$9,transport_service_price_per_transport=$10,
			transport_cost_unit=$11,transport_cost_per_cubic=$12,transport_cost_per_ton=$13,
			transport_cost_per_transport=$14,road_money_per_transport=$15,
			loading_cost_per_transport=$16,unloading_cost_per_transport=$17,
			fuel_cost_per_transport=$18,toll_cost_per_transport=$19,
			other_income_per_transport=$20,other_expense_per_transport=$21,route_note=$22,
			is_active=$23,updated_by=$24,updated_at=CURRENT_TIMESTAMP
		WHERE project_route_id=$25 AND deleted_at IS NULL`,
		req.ProjectID, req.ProjectPattern, req.RouteSequence, req.RouteType, req.RouteName,
		req.DistanceKM, req.TransportServiceUnit, req.TransportServicePricePerCubic,
		req.TransportServicePricePerTon, req.TransportServicePricePerTransport,
		req.TransportCostUnit, req.TransportCostPerCubic, req.TransportCostPerTon,
		req.TransportCostPerTransport, routeAmount(req.RoadMoneyPerTransport),
		routeAmount(req.LoadingCostPerTransport), routeAmount(req.UnloadingCostPerTransport),
		routeAmount(req.FuelCostPerTransport), routeAmount(req.TollCostPerTransport),
		routeAmount(req.OtherIncomePerTransport), routeAmount(req.OtherExpensePerTransport),
		req.RouteNote, active, userID, id)
	return err
}

func (r *ProjectRouteRepository) SoftDelete(ctx context.Context, id, userID int) error {
	_, err := r.DB.Exec(ctx, `UPDATE project_route SET deleted_by=$1,deleted_at=CURRENT_TIMESTAMP WHERE project_route_id=$2 AND deleted_at IS NULL`, userID, id)
	return err
}

func (r *ProjectRouteRepository) GetByID(ctx context.Context, id int, checkerID *int) (*model.ProjectRoute, error) {
	return scanJSONRow[model.ProjectRoute](r.DB.QueryRow(ctx, `
		SELECT `+projectRouteJSON+`
		FROM project_route r
		JOIN project p ON p.project_id=r.project_id AND p.deleted_at IS NULL
		WHERE r.project_route_id=$1 AND r.deleted_at IS NULL
		  AND (
			$2::int IS NULL OR EXISTS (
				SELECT 1 FROM project_checker_assignment a
				WHERE a.project_id=r.project_id AND a.checker_id=$2
				  AND a.is_active=1 AND a.deleted_at IS NULL
				  AND a.access_started_at <= CURRENT_TIMESTAMP
				  AND (a.access_ended_at IS NULL OR a.access_ended_at >= CURRENT_TIMESTAMP)
			)
		  )`, id, checkerID))
}

func (r *ProjectRouteRepository) List(ctx context.Context, opts model.ListOptions) ([]model.ProjectRoute, error) {
	rows, err := r.DB.Query(ctx, `
		SELECT `+projectRouteJSON+`
		FROM project_route r
		JOIN project p ON p.project_id=r.project_id AND p.deleted_at IS NULL
		WHERE r.deleted_at IS NULL
		  AND ($1::int IS NULL OR r.project_route_id < $1)
		  AND ($2::int IS NULL OR r.project_id=$2)
		  AND ($5='' OR r.route_type=$5)
		  AND ($6::smallint IS NULL OR r.is_active=$6)
		  AND (
			$3::int IS NULL OR EXISTS (
				SELECT 1 FROM project_checker_assignment a
				WHERE a.project_id=r.project_id AND a.checker_id=$3
				  AND a.is_active=1 AND a.deleted_at IS NULL
				  AND a.access_started_at <= CURRENT_TIMESTAMP
				  AND (a.access_ended_at IS NULL OR a.access_ended_at >= CURRENT_TIMESTAMP)
			)
		  )
		ORDER BY r.project_route_id DESC LIMIT $4 OFFSET $7`,
		opts.CursorKey, opts.ProjectID, opts.CheckerID, opts.Limit, opts.RouteType, opts.IsActive, opts.Offset)
	if err != nil {
		return nil, err
	}
	return scanJSONRows[model.ProjectRoute](rows)
}
