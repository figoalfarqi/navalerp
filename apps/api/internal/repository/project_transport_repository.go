package repository

import (
	"context"
	"time"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectTransportRepository struct{ DB *pgxpool.Pool }

func NewProjectTransportRepository(db *pgxpool.Pool) *ProjectTransportRepository {
	return &ProjectTransportRepository{DB: db}
}

func (r *ProjectTransportRepository) Create(ctx context.Context, req *model.ProjectTransportRequest, transportNumber string, userID int) (int, error) {
	transportedAt := time.Now()
	if req.TransportedAt != nil {
		transportedAt = *req.TransportedAt
	}
	completed := 0
	if req.IsCompleted != nil {
		completed = *req.IsCompleted
	}
	var id int
	err := r.DB.QueryRow(ctx, `
		INSERT INTO project_transport (
			project_id,project_route_id,transport_number,delivery_note_number,transported_at,
			truck_id,project_truck_assignment_id,driver_id,transport_vendor_id,
			loaded_volume_cubic,loaded_weight_ton,delivered_volume_cubic,delivered_weight_ton,
			purchase_volume_cubic,purchase_weight_ton,sale_volume_cubic,sale_weight_ton,
			transport_service_volume_cubic,transport_service_weight_ton,
			transport_cost_volume_cubic,transport_cost_weight_ton,volume_to_weight_conversion,
			material_purchase_unit,material_buy_price_per_cubic,material_buy_price_per_ton,
			material_sale_unit,material_sell_price_per_cubic,material_sell_price_per_ton,
			transport_service_unit,transport_service_price_per_cubic,
			transport_service_price_per_ton,transport_service_price_per_transport,
			transport_cost_unit,transport_cost_per_cubic,transport_cost_per_ton,
			transport_cost_per_transport,road_money_amount,loading_cost_amount,
			unloading_cost_amount,fuel_cost_amount,toll_cost_amount,other_income_amount,
			other_expense_amount,is_completed,transport_note,created_by,updated_by
		)
		SELECT
			p.project_id,pr.project_route_id,$3,$4,$5,$6,$7,
			COALESCE($8,t.driver_id),COALESCE($9,t.vendor_id),
			$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,
			p.volume_to_weight_conversion,p.material_purchase_unit,
			p.material_buy_price_per_cubic,p.material_buy_price_per_ton,
			p.material_sale_unit,p.material_sell_price_per_cubic,p.material_sell_price_per_ton,
			pr.transport_service_unit,pr.transport_service_price_per_cubic,
			pr.transport_service_price_per_ton,pr.transport_service_price_per_transport,
			pr.transport_cost_unit,pr.transport_cost_per_cubic,pr.transport_cost_per_ton,
			pr.transport_cost_per_transport,
			COALESCE($22,pr.road_money_per_transport),
			COALESCE($23,pr.loading_cost_per_transport),
			COALESCE($24,pr.unloading_cost_per_transport),
			COALESCE($25,pr.fuel_cost_per_transport),
			COALESCE($26,pr.toll_cost_per_transport),
			COALESCE($27,pr.other_income_per_transport),
			COALESCE($28,pr.other_expense_per_transport),
			$29,$30,$31,$31
		FROM project p
		JOIN project_route pr ON pr.project_route_id=$2 AND pr.project_id=p.project_id
			AND pr.deleted_at IS NULL AND pr.is_active=1
		LEFT JOIN truck t ON t.truck_id=$6 AND t.deleted_at IS NULL
		WHERE p.project_id=$1 AND p.deleted_at IS NULL AND p.is_active=1
		RETURNING project_transport_id`,
		req.ProjectID, req.ProjectRouteID, transportNumber, req.DeliveryNoteNumber,
		transportedAt, req.TruckID, req.ProjectTruckAssignmentID, req.DriverID,
		req.TransportVendorID, req.LoadedVolumeCubic, req.LoadedWeightTon,
		req.DeliveredVolumeCubic, req.DeliveredWeightTon, req.PurchaseVolumeCubic,
		req.PurchaseWeightTon, req.SaleVolumeCubic, req.SaleWeightTon,
		req.TransportServiceVolumeCubic, req.TransportServiceWeightTon,
		req.TransportCostVolumeCubic, req.TransportCostWeightTon,
		req.RoadMoneyAmount, req.LoadingCostAmount, req.UnloadingCostAmount,
		req.FuelCostAmount, req.TollCostAmount, req.OtherIncomeAmount,
		req.OtherExpenseAmount, completed, req.TransportNote, userID,
	).Scan(&id)
	return id, err
}

func (r *ProjectTransportRepository) Update(ctx context.Context, id int, req *model.ProjectTransportRequest, userID int) error {
	_, err := r.DB.Exec(ctx, `
		UPDATE project_transport SET
			project_id=$1,project_route_id=$2,
			delivery_note_number=COALESCE($3,delivery_note_number),
			transported_at=COALESCE($4,transported_at),
			truck_id=$5,project_truck_assignment_id=$6,
			driver_id=COALESCE($7,driver_id),transport_vendor_id=COALESCE($8,transport_vendor_id),
			loaded_volume_cubic=COALESCE($9,loaded_volume_cubic),
			loaded_weight_ton=COALESCE($10,loaded_weight_ton),
			delivered_volume_cubic=COALESCE($11,delivered_volume_cubic),
			delivered_weight_ton=COALESCE($12,delivered_weight_ton),
			purchase_volume_cubic=COALESCE($13,purchase_volume_cubic),
			purchase_weight_ton=COALESCE($14,purchase_weight_ton),
			sale_volume_cubic=COALESCE($15,sale_volume_cubic),
			sale_weight_ton=COALESCE($16,sale_weight_ton),
			transport_service_volume_cubic=COALESCE($17,transport_service_volume_cubic),
			transport_service_weight_ton=COALESCE($18,transport_service_weight_ton),
			transport_cost_volume_cubic=COALESCE($19,transport_cost_volume_cubic),
			transport_cost_weight_ton=COALESCE($20,transport_cost_weight_ton),
			road_money_amount=COALESCE($21,road_money_amount),
			loading_cost_amount=COALESCE($22,loading_cost_amount),
			unloading_cost_amount=COALESCE($23,unloading_cost_amount),
			fuel_cost_amount=COALESCE($24,fuel_cost_amount),
			toll_cost_amount=COALESCE($25,toll_cost_amount),
			other_income_amount=COALESCE($26,other_income_amount),
			other_expense_amount=COALESCE($27,other_expense_amount),
			is_completed=COALESCE($28,is_completed),transport_note=COALESCE($29,transport_note),
			updated_by=$30,updated_at=CURRENT_TIMESTAMP
		WHERE project_transport_id=$31 AND deleted_at IS NULL`,
		req.ProjectID, req.ProjectRouteID, req.DeliveryNoteNumber, req.TransportedAt,
		req.TruckID, req.ProjectTruckAssignmentID, req.DriverID, req.TransportVendorID,
		req.LoadedVolumeCubic, req.LoadedWeightTon, req.DeliveredVolumeCubic,
		req.DeliveredWeightTon, req.PurchaseVolumeCubic, req.PurchaseWeightTon,
		req.SaleVolumeCubic, req.SaleWeightTon, req.TransportServiceVolumeCubic,
		req.TransportServiceWeightTon, req.TransportCostVolumeCubic,
		req.TransportCostWeightTon, req.RoadMoneyAmount, req.LoadingCostAmount,
		req.UnloadingCostAmount, req.FuelCostAmount, req.TollCostAmount,
		req.OtherIncomeAmount, req.OtherExpenseAmount, req.IsCompleted,
		req.TransportNote, userID, id)
	return err
}

func (r *ProjectTransportRepository) SoftDelete(ctx context.Context, id, userID int) error {
	_, err := r.DB.Exec(ctx, `
		UPDATE project_transport
		SET deleted_by=$1,deleted_at=CURRENT_TIMESTAMP,updated_by=$1,updated_at=CURRENT_TIMESTAMP
		WHERE project_transport_id=$2 AND deleted_at IS NULL`, userID, id)
	return err
}

const projectTransportProjectJSON = `jsonb_build_object(
	'project_id',p.project_id,
	'project_code',p.project_code,
	'project_name',p.project_name
)`

const projectTransportRouteJSON = `jsonb_build_object(
	'project_route_id',pr.project_route_id,
	'route_name',pr.route_name,
	'route_type',pr.route_type
)`

const projectTransportTruckJSON = `CASE
	WHEN t.truck_id IS NULL THEN NULL
	ELSE jsonb_build_object(
		'truck_id',t.truck_id,
		'license_plate',t.license_plate
	)
END`

const projectTransportDriverJSON = `CASE
	WHEN d.app_user_id IS NULL THEN NULL
	ELSE jsonb_build_object(
		'app_user_id',d.app_user_id,
		'app_user_name',d.app_user_name,
		'username',d.username
	)
END`

const projectTransportVendorJSON = `CASE
	WHEN tv.vendor_id IS NULL THEN NULL
	ELSE jsonb_build_object(
		'vendor_id',tv.vendor_id,
		'vendor_name',tv.vendor_name
	)
END`

const projectTransportReferenceJSON = `jsonb_build_object(
	'project_transport_id',pt.project_transport_id,
	'transport_number',pt.transport_number,
	'project',` + projectTransportProjectJSON + `,
	'project_route',` + projectTransportRouteJSON + `,
	'truck',` + projectTransportTruckJSON + `,
	'driver',` + projectTransportDriverJSON + `,
	'transport_vendor',` + projectTransportVendorJSON + `
)`

const fullTransportJSON = `
	to_jsonb(pt) || jsonb_build_object(
		'project',` + projectTransportProjectJSON + `,
		'project_route',` + projectTransportRouteJSON + `,
		'truck',` + projectTransportTruckJSON + `,
		'driver',` + projectTransportDriverJSON + `,
		'transport_vendor',` + projectTransportVendorJSON + `,
		'latest_status',CASE
			WHEN latest_status.project_transport_status_type_id IS NULL THEN NULL
			ELSE jsonb_build_object(
				'project_transport_status_type_id',latest_status.project_transport_status_type_id,
				'status_time',latest_status.status_time
			)
		END,
		'status_count',(
			SELECT COUNT(*)
			FROM project_transport_status status_count
			WHERE status_count.project_transport_id=pt.project_transport_id
			  AND status_count.deleted_at IS NULL
		),
		'photo_count',(
			SELECT COUNT(*)
			FROM project_transport_photo photo_count
			JOIN project_transport_status photo_status
			  ON photo_status.project_transport_status_id=photo_count.project_transport_status_id
			WHERE photo_status.project_transport_id=pt.project_transport_id
			  AND photo_count.deleted_at IS NULL
			  AND photo_status.deleted_at IS NULL
		)
	)`

const fullTransportDetailJSON = fullTransportJSON + ` || jsonb_build_object(
	'statuses',COALESCE((
		SELECT jsonb_agg(
			to_jsonb(s) || jsonb_build_object(
				'photos',COALESCE((
					SELECT jsonb_agg(
						to_jsonb(ph)
						ORDER BY ph.created_at,ph.project_transport_photo_id
					)
					FROM project_transport_photo ph
					WHERE ph.project_transport_status_id=s.project_transport_status_id
					  AND ph.deleted_at IS NULL
				),'[]'::jsonb)
			)
			ORDER BY s.status_time,s.project_transport_status_id
		)
		FROM project_transport_status s
		WHERE s.project_transport_id=pt.project_transport_id
		  AND s.deleted_at IS NULL
	),'[]'::jsonb)
)`

func (r *ProjectTransportRepository) GetByID(ctx context.Context, id int, checkerID *int) (*model.ProjectTransport, error) {
	return scanJSONRow[model.ProjectTransport](r.DB.QueryRow(ctx, `
		SELECT `+fullTransportDetailJSON+`
		FROM project_transport pt
		JOIN project p ON p.project_id=pt.project_id
		JOIN project_route pr ON pr.project_route_id=pt.project_route_id
		LEFT JOIN truck t ON t.truck_id=pt.truck_id
		LEFT JOIN app_user d ON d.app_user_id=pt.driver_id AND d.deleted_at IS NULL
		LEFT JOIN vendor tv ON tv.vendor_id=pt.transport_vendor_id AND tv.deleted_at IS NULL
		LEFT JOIN LATERAL (
			SELECT s.project_transport_status_type_id,s.status_time
			FROM project_transport_status s
			WHERE s.project_transport_id=pt.project_transport_id AND s.deleted_at IS NULL
			ORDER BY s.status_time DESC,s.project_transport_status_id DESC
			LIMIT 1
		) latest_status ON TRUE
		WHERE pt.project_transport_id=$1 AND pt.deleted_at IS NULL
		  AND (
			$2::int IS NULL OR EXISTS (
				SELECT 1 FROM project_checker_assignment ca
				WHERE ca.project_id=pt.project_id AND ca.checker_id=$2
				  AND ca.is_active=1 AND ca.deleted_at IS NULL
				  AND ca.access_started_at <= CURRENT_TIMESTAMP
				  AND (ca.access_ended_at IS NULL OR ca.access_ended_at >= CURRENT_TIMESTAMP)
			)
		  )`, id, checkerID))
}

func (r *ProjectTransportRepository) List(ctx context.Context, opts model.ListOptions) ([]model.ProjectTransport, error) {
	rows, err := r.DB.Query(ctx, `
		SELECT `+fullTransportJSON+`
		FROM project_transport pt
		JOIN project p ON p.project_id=pt.project_id
		JOIN project_route pr ON pr.project_route_id=pt.project_route_id
		LEFT JOIN truck t ON t.truck_id=pt.truck_id
		LEFT JOIN app_user d ON d.app_user_id=pt.driver_id AND d.deleted_at IS NULL
		LEFT JOIN vendor tv ON tv.vendor_id=pt.transport_vendor_id AND tv.deleted_at IS NULL
		LEFT JOIN LATERAL (
			SELECT s.project_transport_status_type_id,s.status_time
			FROM project_transport_status s
			WHERE s.project_transport_id=pt.project_transport_id AND s.deleted_at IS NULL
			ORDER BY s.status_time DESC,s.project_transport_status_id DESC
			LIMIT 1
		) latest_status ON TRUE
		WHERE pt.deleted_at IS NULL
		  AND ($1::int IS NULL OR pt.project_transport_id < $1)
		  AND ($2::int IS NULL OR pt.project_id=$2)
		  AND ($3::int IS NULL OR pt.driver_id=$3)
		  AND ($4::timestamptz IS NULL OR pt.transported_at >= $4)
		  AND ($5::timestamptz IS NULL OR pt.transported_at < $5)
		  AND ($6::int IS NULL OR pt.is_completed=$6)
		  AND (
			$7::int IS NULL OR EXISTS (
				SELECT 1 FROM project_checker_assignment ca
				WHERE ca.project_id=pt.project_id AND ca.checker_id=$7
				  AND ca.is_active=1 AND ca.deleted_at IS NULL
				  AND ca.access_started_at <= CURRENT_TIMESTAMP
				  AND (ca.access_ended_at IS NULL OR ca.access_ended_at >= CURRENT_TIMESTAMP)
			)
		  )
		  AND ($8::int IS NULL OR pt.project_transport_id=$8)
		  AND ($11::int IS NULL OR pt.truck_id=$11)
		  AND ($12='' OR pt.transport_number ILIKE '%' || $12 || '%')
		  AND ($13='' OR pr.route_type=$13)
		ORDER BY pt.project_transport_id DESC LIMIT $9 OFFSET $10`,
		opts.CursorKey, opts.ProjectID, opts.DriverID, opts.DateFrom, opts.DateTo,
		opts.IsCompleted, opts.CheckerID, opts.TransportID, opts.Limit, opts.Offset,
		opts.TruckID, opts.Filters["transport_number"], opts.RouteType)
	if err != nil {
		return nil, err
	}
	return scanJSONRows[model.ProjectTransport](rows)
}

func (r *ProjectTransportRepository) ListOperational(ctx context.Context, opts model.ListOptions) ([]model.DriverProjectTransport, error) {
	rows, err := r.DB.Query(ctx, `
		SELECT jsonb_build_object(
			'project_transport_id',pt.project_transport_id,
			'project_id',pt.project_id,
			'project_code',p.project_code,
			'project_name',p.project_name,
			'project_route_id',pt.project_route_id,
			'route_name',pr.route_name,
			'route_type',pr.route_type,
			'origin_name',CASE
				WHEN pr.route_type='STOCKPILE_TO_CLIENT' THEN COALESCE(sp.stockpile_name,'Stockpile')
				ELSE COALESCE(m.mine_name,NULLIF(CONCAT_WS(' - ',v.vessel_name,po.port_name),''),'Lokasi asal')
			END,
			'destination_name',CASE
				WHEN pr.route_type='SOURCE_TO_STOCKPILE' THEN COALESCE(sp.stockpile_name,'Stockpile')
				ELSE COALESCE(cd.client_destination_name,'Lokasi tujuan')
			END,
			'distance_km',pr.distance_km,
			'transport_number',pt.transport_number,
			'transported_at',pt.transported_at,
			'truck_id',pt.truck_id,
			'license_plate',t.license_plate,
			'loaded_volume_cubic',pt.loaded_volume_cubic,
			'loaded_weight_ton',pt.loaded_weight_ton,
			'delivered_volume_cubic',pt.delivered_volume_cubic,
			'delivered_weight_ton',pt.delivered_weight_ton,
			'is_completed',pt.is_completed,
			'statuses',COALESCE((
				SELECT jsonb_agg(to_jsonb(s) ORDER BY s.status_time,s.project_transport_status_id)
				FROM project_transport_status s
				WHERE s.project_transport_id=pt.project_transport_id AND s.deleted_at IS NULL
			),'[]'::jsonb)
		)
		FROM project_transport pt
		JOIN project p ON p.project_id=pt.project_id
		JOIN project_route pr ON pr.project_route_id=pt.project_route_id
		LEFT JOIN truck t ON t.truck_id=pt.truck_id
		LEFT JOIN mine m ON m.mine_id=p.mine_id
		LEFT JOIN vessel_cargo vc ON vc.vessel_cargo_id=p.vessel_cargo_id
		LEFT JOIN vessel v ON v.vessel_id=vc.vessel_id
		LEFT JOIN port po ON po.port_id=vc.port_id
		LEFT JOIN stockpile_cargo sc ON sc.stockpile_cargo_id=p.stockpile_cargo_id
		LEFT JOIN stockpile sp ON sp.stockpile_id=sc.stockpile_id
		LEFT JOIN client_destination cd ON cd.client_destination_id=p.client_destination_id
		WHERE pt.deleted_at IS NULL
		  AND ($1::int IS NULL OR pt.project_transport_id < $1)
		  AND ($2::int IS NULL OR pt.project_id=$2)
		  AND ($3::int IS NULL OR pt.driver_id=$3)
		  AND ($4::timestamptz IS NULL OR pt.transported_at >= $4)
		  AND ($5::timestamptz IS NULL OR pt.transported_at < $5)
		  AND ($6::int IS NULL OR pt.is_completed=$6)
		  AND (
			$7::int IS NULL OR EXISTS (
				SELECT 1 FROM project_checker_assignment ca
				WHERE ca.project_id=pt.project_id AND ca.checker_id=$7
				  AND ca.is_active=1 AND ca.deleted_at IS NULL
				  AND ca.access_started_at <= CURRENT_TIMESTAMP
				  AND (ca.access_ended_at IS NULL OR ca.access_ended_at >= CURRENT_TIMESTAMP)
			)
		  )
		  AND ($8::int IS NULL OR pt.project_transport_id=$8)
		  AND ($11::int IS NULL OR pt.truck_id=$11)
		  AND ($12='' OR pt.transport_number ILIKE '%' || $12 || '%')
		  AND ($13='' OR pr.route_type=$13)
		ORDER BY pt.project_transport_id DESC LIMIT $9 OFFSET $10`,
		opts.CursorKey, opts.ProjectID, opts.DriverID, opts.DateFrom, opts.DateTo,
		opts.IsCompleted, opts.CheckerID, opts.TransportID, opts.Limit, opts.Offset,
		opts.TruckID, opts.Filters["transport_number"], opts.RouteType)
	if err != nil {
		return nil, err
	}
	return scanJSONRows[model.DriverProjectTransport](rows)
}

func (r *ProjectTransportRepository) ResolveProjectID(ctx context.Context, transportID int) (int, error) {
	var projectID int
	err := r.DB.QueryRow(ctx, `
		SELECT project_id FROM project_transport
		WHERE project_transport_id=$1 AND deleted_at IS NULL`, transportID).Scan(&projectID)
	return projectID, err
}
