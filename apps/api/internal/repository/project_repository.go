package repository

import (
	"context"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectRepository struct{ DB *pgxpool.Pool }

const projectRelationsJSON = `
	to_jsonb(p) || jsonb_build_object(
		'mine',CASE WHEN mine.mine_id IS NULL THEN NULL ELSE to_jsonb(mine) END,
		'vessel_cargo',CASE
			WHEN vc.vessel_cargo_id IS NULL THEN NULL
			ELSE to_jsonb(vc) || jsonb_build_object(
				'vessel',CASE WHEN vessel.vessel_id IS NULL THEN NULL ELSE to_jsonb(vessel) END,
				'port',CASE WHEN port.port_id IS NULL THEN NULL ELSE to_jsonb(port) END,
				'cargo_type',CASE WHEN vessel_cargo_type.cargo_type_id IS NULL THEN NULL ELSE to_jsonb(vessel_cargo_type) END
			)
		END,
		'stockpile_cargo',CASE
			WHEN sc.stockpile_cargo_id IS NULL THEN NULL
			ELSE to_jsonb(sc) || jsonb_build_object(
				'stockpile',CASE WHEN stockpile.stockpile_id IS NULL THEN NULL ELSE to_jsonb(stockpile) END,
				'cargo_type',CASE WHEN stockpile_cargo_type.cargo_type_id IS NULL THEN NULL ELSE to_jsonb(stockpile_cargo_type) END
			)
		END,
		'client_destination',to_jsonb(cd) || jsonb_build_object(
			'client',CASE WHEN client.client_id IS NULL THEN NULL ELSE to_jsonb(client) END,
			'city',CASE WHEN destination_city.city_id IS NULL THEN NULL ELSE to_jsonb(destination_city) END
		),
		'cargo_type',to_jsonb(project_cargo_type)
	)`

func NewProjectRepository(db *pgxpool.Pool) *ProjectRepository {
	return &ProjectRepository{DB: db}
}

func (r *ProjectRepository) Create(ctx context.Context, req *model.ProjectRequest, userID int) (int, error) {
	active := 1
	if req.IsActive != nil {
		active = *req.IsActive
	}
	fixedIncome, fixedExpense := 0.0, 0.0
	if req.FixedOtherIncome != nil {
		fixedIncome = *req.FixedOtherIncome
	}
	if req.FixedOtherExpense != nil {
		fixedExpense = *req.FixedOtherExpense
	}
	var id int
	err := r.DB.QueryRow(ctx, `
		INSERT INTO project (
			project_code,project_name,route_type,mine_id,vessel_cargo_id,stockpile_cargo_id,
			client_destination_id,cargo_type_id,project_status,start_date,end_date,
			planned_volume_cubic,planned_weight_ton,volume_to_weight_conversion,
			material_purchase_unit,material_buy_price_per_cubic,material_buy_price_per_ton,
			material_sale_unit,material_sell_price_per_cubic,material_sell_price_per_ton,
			fixed_other_income,fixed_other_expense,project_note,is_active,created_by,updated_by
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10::date,$11::date,$12,$13,$14,
			$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$25
		) RETURNING project_id`,
		req.ProjectCode, req.ProjectName, req.RouteType, req.MineID, req.VesselCargoID,
		req.StockpileCargoID, req.ClientDestinationID, req.CargoTypeID, req.ProjectStatus,
		req.StartDate, req.EndDate, req.PlannedVolumeCubic, req.PlannedWeightTon,
		req.VolumeToWeightConversion, req.MaterialPurchaseUnit, req.MaterialBuyPricePerCubic,
		req.MaterialBuyPricePerTon, req.MaterialSaleUnit, req.MaterialSellPricePerCubic,
		req.MaterialSellPricePerTon, fixedIncome, fixedExpense, req.ProjectNote, active, userID,
	).Scan(&id)
	return id, err
}

func (r *ProjectRepository) Update(ctx context.Context, id int, req *model.ProjectRequest, userID int) error {
	active := 1
	if req.IsActive != nil {
		active = *req.IsActive
	}
	fixedIncome, fixedExpense := 0.0, 0.0
	if req.FixedOtherIncome != nil {
		fixedIncome = *req.FixedOtherIncome
	}
	if req.FixedOtherExpense != nil {
		fixedExpense = *req.FixedOtherExpense
	}
	_, err := r.DB.Exec(ctx, `
		UPDATE project SET
			project_code=$1,project_name=$2,route_type=$3,mine_id=$4,vessel_cargo_id=$5,
			stockpile_cargo_id=$6,client_destination_id=$7,cargo_type_id=$8,
			project_status=$9,start_date=$10::date,end_date=$11::date,
			planned_volume_cubic=$12,planned_weight_ton=$13,volume_to_weight_conversion=$14,
			material_purchase_unit=$15,material_buy_price_per_cubic=$16,
			material_buy_price_per_ton=$17,material_sale_unit=$18,
			material_sell_price_per_cubic=$19,material_sell_price_per_ton=$20,
			fixed_other_income=$21,fixed_other_expense=$22,project_note=$23,is_active=$24,
			updated_by=$25,updated_at=CURRENT_TIMESTAMP
		WHERE project_id=$26 AND deleted_at IS NULL`,
		req.ProjectCode, req.ProjectName, req.RouteType, req.MineID, req.VesselCargoID,
		req.StockpileCargoID, req.ClientDestinationID, req.CargoTypeID, req.ProjectStatus,
		req.StartDate, req.EndDate, req.PlannedVolumeCubic, req.PlannedWeightTon,
		req.VolumeToWeightConversion, req.MaterialPurchaseUnit, req.MaterialBuyPricePerCubic,
		req.MaterialBuyPricePerTon, req.MaterialSaleUnit, req.MaterialSellPricePerCubic,
		req.MaterialSellPricePerTon, fixedIncome, fixedExpense, req.ProjectNote, active, userID, id)
	return err
}

func (r *ProjectRepository) SoftDelete(ctx context.Context, id, userID int) error {
	_, err := r.DB.Exec(ctx, `
		UPDATE project SET deleted_by=$1,deleted_at=CURRENT_TIMESTAMP,updated_by=$1,updated_at=CURRENT_TIMESTAMP
		WHERE project_id=$2 AND deleted_at IS NULL`, userID, id)
	return err
}

func (r *ProjectRepository) GetByID(ctx context.Context, id int, checkerID *int) (*model.Project, error) {
	return scanJSONRow[model.Project](r.DB.QueryRow(ctx, `
		SELECT `+projectRelationsJSON+` || jsonb_build_object(
			'is_default',
			COALESCE((
				SELECT a.is_default
				FROM project_checker_assignment a
				WHERE a.project_id=p.project_id AND a.checker_id=$2
				  AND a.is_active=1 AND a.deleted_at IS NULL
				  AND a.access_started_at <= CURRENT_TIMESTAMP
				  AND (a.access_ended_at IS NULL OR a.access_ended_at >= CURRENT_TIMESTAMP)
				ORDER BY a.project_checker_assignment_id DESC
				LIMIT 1
			),0)
		)
		FROM project p
		LEFT JOIN mine mine ON mine.mine_id=p.mine_id AND mine.deleted_at IS NULL
		LEFT JOIN vessel_cargo vc ON vc.vessel_cargo_id=p.vessel_cargo_id AND vc.deleted_at IS NULL
		LEFT JOIN vessel vessel ON vessel.vessel_id=vc.vessel_id AND vessel.deleted_at IS NULL
		LEFT JOIN port port ON port.port_id=vc.port_id AND port.deleted_at IS NULL
		LEFT JOIN cargo_type vessel_cargo_type ON vessel_cargo_type.cargo_type_id=vc.cargo_type_id AND vessel_cargo_type.deleted_at IS NULL
		LEFT JOIN stockpile_cargo sc ON sc.stockpile_cargo_id=p.stockpile_cargo_id AND sc.deleted_at IS NULL
		LEFT JOIN stockpile stockpile ON stockpile.stockpile_id=sc.stockpile_id AND stockpile.deleted_at IS NULL
		LEFT JOIN cargo_type stockpile_cargo_type ON stockpile_cargo_type.cargo_type_id=sc.cargo_type_id AND stockpile_cargo_type.deleted_at IS NULL
		JOIN client_destination cd ON cd.client_destination_id=p.client_destination_id AND cd.deleted_at IS NULL
		LEFT JOIN client client ON client.client_id=cd.client_id AND client.deleted_at IS NULL
		LEFT JOIN city destination_city ON destination_city.city_id=cd.city_id AND destination_city.deleted_at IS NULL
		JOIN cargo_type project_cargo_type ON project_cargo_type.cargo_type_id=p.cargo_type_id AND project_cargo_type.deleted_at IS NULL
		WHERE p.project_id=$1 AND p.deleted_at IS NULL
		  AND (
			$2::int IS NULL OR EXISTS (
				SELECT 1 FROM project_checker_assignment a
				WHERE a.project_id=p.project_id AND a.checker_id=$2
				  AND a.is_active=1 AND a.deleted_at IS NULL
				  AND a.access_started_at <= CURRENT_TIMESTAMP
				  AND (a.access_ended_at IS NULL OR a.access_ended_at >= CURRENT_TIMESTAMP)
			)
		  )`, id, checkerID))
}

func (r *ProjectRepository) List(ctx context.Context, opts model.ListOptions) ([]model.Project, error) {
	rows, err := r.DB.Query(ctx, `
		SELECT `+projectRelationsJSON+` || jsonb_build_object(
			'is_default',
			COALESCE((
				SELECT a.is_default
				FROM project_checker_assignment a
				WHERE a.project_id=p.project_id AND a.checker_id=$4
				  AND a.is_active=1 AND a.deleted_at IS NULL
				  AND a.access_started_at <= CURRENT_TIMESTAMP
				  AND (a.access_ended_at IS NULL OR a.access_ended_at >= CURRENT_TIMESTAMP)
				ORDER BY a.project_checker_assignment_id DESC
				LIMIT 1
			),0)
		)
		FROM project p
		LEFT JOIN mine mine ON mine.mine_id=p.mine_id AND mine.deleted_at IS NULL
		LEFT JOIN vessel_cargo vc ON vc.vessel_cargo_id=p.vessel_cargo_id AND vc.deleted_at IS NULL
		LEFT JOIN vessel vessel ON vessel.vessel_id=vc.vessel_id AND vessel.deleted_at IS NULL
		LEFT JOIN port port ON port.port_id=vc.port_id AND port.deleted_at IS NULL
		LEFT JOIN cargo_type vessel_cargo_type ON vessel_cargo_type.cargo_type_id=vc.cargo_type_id AND vessel_cargo_type.deleted_at IS NULL
		LEFT JOIN stockpile_cargo sc ON sc.stockpile_cargo_id=p.stockpile_cargo_id AND sc.deleted_at IS NULL
		LEFT JOIN stockpile stockpile ON stockpile.stockpile_id=sc.stockpile_id AND stockpile.deleted_at IS NULL
		LEFT JOIN cargo_type stockpile_cargo_type ON stockpile_cargo_type.cargo_type_id=sc.cargo_type_id AND stockpile_cargo_type.deleted_at IS NULL
		JOIN client_destination cd ON cd.client_destination_id=p.client_destination_id AND cd.deleted_at IS NULL
		LEFT JOIN client client ON client.client_id=cd.client_id AND client.deleted_at IS NULL
		LEFT JOIN city destination_city ON destination_city.city_id=cd.city_id AND destination_city.deleted_at IS NULL
		JOIN cargo_type project_cargo_type ON project_cargo_type.cargo_type_id=p.cargo_type_id AND project_cargo_type.deleted_at IS NULL
		WHERE p.deleted_at IS NULL
		  AND ($1::int IS NULL OR p.project_id < $1)
		  AND ($2::int IS NULL OR p.project_id=$2)
		  AND ($3='' OR p.project_name ILIKE '%' || $3 || '%' OR p.project_code ILIKE '%' || $3 || '%')
		  AND (
			$4::int IS NULL OR EXISTS (
				SELECT 1 FROM project_checker_assignment a
				WHERE a.project_id=p.project_id AND a.checker_id=$4
				  AND a.is_active=1 AND a.deleted_at IS NULL
				  AND a.access_started_at <= CURRENT_TIMESTAMP
				  AND (a.access_ended_at IS NULL OR a.access_ended_at >= CURRENT_TIMESTAMP)
			)
		  )
		  AND ($5::smallint IS NULL OR p.is_active=$5)
		  AND ($6='' OR p.route_type=$6)
		  AND ($7='' OR p.project_status=$7)
		  AND ($8::int IS NULL OR p.cargo_type_id=$8)
		ORDER BY p.project_id DESC
		LIMIT $9 OFFSET $10`,
		opts.CursorKey, opts.ProjectID, opts.Search, opts.CheckerID, opts.IsActive,
		opts.RouteType, opts.Filters["project_status"], opts.CargoTypeID, opts.Limit, opts.Offset)
	if err != nil {
		return nil, err
	}
	return scanJSONRows[model.Project](rows)
}

func (r *ProjectRepository) CheckerCanAccess(ctx context.Context, projectID, checkerID int) (bool, error) {
	var allowed bool
	err := r.DB.QueryRow(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM project_checker_assignment
			WHERE project_id=$1 AND checker_id=$2 AND is_active=1 AND deleted_at IS NULL
			  AND access_started_at <= CURRENT_TIMESTAMP
			  AND (access_ended_at IS NULL OR access_ended_at >= CURRENT_TIMESTAMP)
		)`, projectID, checkerID).Scan(&allowed)
	return allowed, err
}
