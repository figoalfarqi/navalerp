package app

import (
	"net/http"
	"slices"
	"strings"

	"github.com/figoalfarqi/navalerp/config"
	"github.com/figoalfarqi/navalerp/internal/app/middleware"
	"github.com/figoalfarqi/navalerp/internal/helper"
)

type crudHandler interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

func authenticated(cfg *config.Config, roles []int, fn http.HandlerFunc) http.Handler {
	return middleware.AuthJWT(cfg, roles, fn)
}

func registerRead(mux *http.ServeMux, cfg *config.Config, path string, roles []int, fn http.HandlerFunc) {
	if strings.HasPrefix(path, "/api/v1/admin/") {
		for _, roleID := range []int{helper.AppRoleOwner, helper.AppRoleITDev} {
			if !slices.Contains(roles, roleID) {
				roles = append(roles, roleID)
			}
		}
	}
	mux.Handle("GET "+path, authenticated(cfg, roles, fn))
	mux.Handle("GET "+path+"/{id}", authenticated(cfg, roles, fn))
}

func registerCRUD(mux *http.ServeMux, cfg *config.Config, path string, roles []int, h crudHandler) {
	registerRead(mux, cfg, path, roles, h.Get)
	mux.Handle("POST "+path, authenticated(cfg, roles, h.Create))
	mux.Handle("PUT "+path+"/{id}", authenticated(cfg, roles, h.Update))
	mux.Handle("PATCH "+path+"/{id}", authenticated(cfg, roles, h.Update))
	mux.Handle("DELETE "+path+"/{id}", authenticated(cfg, roles, h.Delete))
}

func registerApplicationRoutes(mux *http.ServeMux, cfg *config.Config, h *routeHandlers) {
	dashboardRoles := helper.GetAppRoleIDsByRoleTypeName("admin_dashboard")
	reportRoles := helper.GetAppRoleIDsByRoleTypeName("admin_report")
	allRoles := helper.GetAppRoleIDsByRoleTypeName("allrole")

	mux.HandleFunc("POST /api/v1/admin/login", h.navalAuth.Login)
	mux.HandleFunc("POST /api/v1/auth/login", h.navalAuth.Login)
	mux.HandleFunc("POST /api/v1/auth/logout", h.navalAuth.Logout)
	mux.Handle("GET /api/v1/auth/me", authenticated(cfg, allRoles, h.navalAuth.Me))

	mux.Handle("GET /api/v1/admin/dashboard", authenticated(cfg, dashboardRoles, h.dashboard.Get))
	mux.Handle("GET /api/v1/admin/report", authenticated(cfg, reportRoles, h.report.Get))
	mux.Handle("GET /api/v1/admin/generate_number", authenticated(cfg, allRoles, h.generator.GenerateNumber))

	// 48 Modular NavalERP CRUD Routes
	registerCRUD(mux, cfg, "/api/v1/admin/org_unit", allRoles, h.orgUnit)
	registerCRUD(mux, cfg, "/api/v1/admin/sys_user", allRoles, h.sysUser)
	registerCRUD(mux, cfg, "/api/v1/admin/audit_log", allRoles, h.auditLog)
	registerCRUD(mux, cfg, "/api/v1/admin/ship_class", allRoles, h.shipClass)
	registerCRUD(mux, cfg, "/api/v1/admin/ship", allRoles, h.ship)
	registerCRUD(mux, cfg, "/api/v1/admin/ship_system", allRoles, h.shipSystem)
	registerCRUD(mux, cfg, "/api/v1/admin/equipment", allRoles, h.equipment)
	registerCRUD(mux, cfg, "/api/v1/admin/pm_schedule", allRoles, h.pmSchedule)
	registerCRUD(mux, cfg, "/api/v1/admin/failure_report", allRoles, h.failureReport)
	registerCRUD(mux, cfg, "/api/v1/admin/work_order", allRoles, h.workOrder)
	registerCRUD(mux, cfg, "/api/v1/admin/docking_record", allRoles, h.dockingRecord)
	registerCRUD(mux, cfg, "/api/v1/admin/warehouse", allRoles, h.warehouse)
	registerCRUD(mux, cfg, "/api/v1/admin/material", allRoles, h.material)
	registerCRUD(mux, cfg, "/api/v1/admin/stock_balance", allRoles, h.stockBalance)
	registerCRUD(mux, cfg, "/api/v1/admin/item_instance", allRoles, h.itemInstance)
	registerCRUD(mux, cfg, "/api/v1/admin/stock_transfer", allRoles, h.stockTransfer)
	registerCRUD(mux, cfg, "/api/v1/admin/stock_adjustment", allRoles, h.stockAdjustment)
	registerCRUD(mux, cfg, "/api/v1/admin/vendor", allRoles, h.vendor)
	registerCRUD(mux, cfg, "/api/v1/admin/requisition", allRoles, h.requisition)
	registerCRUD(mux, cfg, "/api/v1/admin/tender", allRoles, h.tender)
	registerCRUD(mux, cfg, "/api/v1/admin/contract", allRoles, h.contract)
	registerCRUD(mux, cfg, "/api/v1/admin/purchase_order", allRoles, h.purchaseOrder)
	registerCRUD(mux, cfg, "/api/v1/admin/goods_receipt", allRoles, h.goodsReceipt)
	registerCRUD(mux, cfg, "/api/v1/admin/chart_of_account", allRoles, h.chartOfAccount)
	registerCRUD(mux, cfg, "/api/v1/admin/budget_program", allRoles, h.budgetProgram)
	registerCRUD(mux, cfg, "/api/v1/admin/budget_commitment", allRoles, h.budgetCommitment)
	registerCRUD(mux, cfg, "/api/v1/admin/invoice", allRoles, h.invoice)
	registerCRUD(mux, cfg, "/api/v1/admin/payment", allRoles, h.payment)
	registerCRUD(mux, cfg, "/api/v1/admin/journal_entry", allRoles, h.journalEntry)
	registerCRUD(mux, cfg, "/api/v1/admin/platform_tco", allRoles, h.platformTco)
	registerCRUD(mux, cfg, "/api/v1/admin/military_rank", allRoles, h.militaryRank)
	registerCRUD(mux, cfg, "/api/v1/admin/military_corps", allRoles, h.militaryCorps)
	registerCRUD(mux, cfg, "/api/v1/admin/qualification", allRoles, h.qualification)
	registerCRUD(mux, cfg, "/api/v1/admin/personnel", allRoles, h.personnel)
	registerCRUD(mux, cfg, "/api/v1/admin/crew_assignment", allRoles, h.crewAssignment)
	registerCRUD(mux, cfg, "/api/v1/admin/base_facility", allRoles, h.baseFacility)
	registerCRUD(mux, cfg, "/api/v1/admin/berth_booking", allRoles, h.berthBooking)
	registerCRUD(mux, cfg, "/api/v1/admin/fuel_bunker", allRoles, h.fuelBunker)
	registerCRUD(mux, cfg, "/api/v1/admin/transport_unit", allRoles, h.transportUnit)
	registerCRUD(mux, cfg, "/api/v1/admin/route", allRoles, h.route)
	registerCRUD(mux, cfg, "/api/v1/admin/shipment", allRoles, h.shipment)
	registerCRUD(mux, cfg, "/api/v1/admin/document_category", allRoles, h.documentCategory)
	registerCRUD(mux, cfg, "/api/v1/admin/document", allRoles, h.document)
	registerCRUD(mux, cfg, "/api/v1/admin/theater", allRoles, h.theater)
	registerCRUD(mux, cfg, "/api/v1/admin/mission", allRoles, h.mission)
	registerCRUD(mux, cfg, "/api/v1/admin/daily_log", allRoles, h.dailyLog)
	registerCRUD(mux, cfg, "/api/v1/admin/readiness_report", allRoles, h.readinessReport)
	registerCRUD(mux, cfg, "/api/v1/admin/readiness_alert", allRoles, h.readinessAlert)

	// Critical Underwater Infrastructure (CUI) Module
	registerCRUD(mux, cfg, "/api/v1/admin/cui_asset", allRoles, h.cuiAsset)
	registerCRUD(mux, cfg, "/api/v1/admin/cui_monitoring_log", allRoles, h.cuiMonitoringLog)
	registerCRUD(mux, cfg, "/api/v1/admin/cui_alert", allRoles, h.cuiAlert)
	registerCRUD(mux, cfg, "/api/v1/admin/cui_inspection", allRoles, h.cuiInspection)
	mux.Handle("GET /api/v1/admin/cui/overview", authenticated(cfg, allRoles, h.cuiOverview.Get))

	// Digital Approval Workflow
	mux.Handle("POST /api/v1/admin/approval", authenticated(cfg, allRoles, h.approval.Process))
	mux.Handle("GET /api/v1/admin/approval/pending", authenticated(cfg, allRoles, h.approval.ListPending))
}
