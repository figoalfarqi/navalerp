package app

import (
	"net/http"
	"slices"
	"strings"

	"github.com/figoalfarqi/apipml/config"
	"github.com/figoalfarqi/apipml/internal/app/middleware"
	"github.com/figoalfarqi/apipml/internal/helper"
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

func registerCreate(mux *http.ServeMux, cfg *config.Config, path string, roles []int, fn http.HandlerFunc) {
	mux.Handle("POST "+path, authenticated(cfg, roles, fn))
}

func registerMutations(mux *http.ServeMux, cfg *config.Config, path string, roles []int, h crudHandler) {
	registerCreate(mux, cfg, path, roles, h.Create)
	mux.Handle("PUT "+path+"/{id}", authenticated(cfg, roles, h.Update))
	mux.Handle("PATCH "+path+"/{id}", authenticated(cfg, roles, h.Update))
	mux.Handle("DELETE "+path+"/{id}", authenticated(cfg, roles, h.Delete))
}

func registerCRUD(mux *http.ServeMux, cfg *config.Config, path string, roles []int, h crudHandler) {
	registerRead(mux, cfg, path, roles, h.Get)
	registerMutations(mux, cfg, path, roles, h)
}

func registerApplicationRoutes(mux *http.ServeMux, cfg *config.Config, h *routeHandlers) {
	operational := helper.GetAppRoleIDsByRoleTypeName("admin_operational")
	technical := helper.GetAppRoleIDsByRoleTypeName("admin_technical")
	dashboardRoles := helper.GetAppRoleIDsByRoleTypeName("admin_dashboard")
	reportRoles := helper.GetAppRoleIDsByRoleTypeName("admin_report")
	checkerRoles := helper.GetAppRoleIDsByRoleTypeName("checker")
	driverRoles := helper.GetAppRoleIDsByRoleTypeName("driver")
	allRoles := helper.GetAppRoleIDsByRoleTypeName("allrole")
	userManagementReads := []int{
		helper.AppRoleITDev,
		helper.AppRoleSuperAdmin,
		helper.AppRoleAdmin,
	}

	mux.HandleFunc("POST /api/v1/admin/login", h.appUser.Login)
	mux.HandleFunc("POST /api/v1/checker/login", h.appUser.Login)
	mux.HandleFunc("POST /api/v1/driver/login", h.appUser.Login)

	registerRead(mux, cfg, "/api/v1/admin/app_role", userManagementReads, h.appRole.Get)
	registerMutations(mux, cfg, "/api/v1/admin/app_role", technical, h.appRole)
	registerCRUD(mux, cfg, "/api/v1/admin/app_setting", technical, h.appSetting)

	registerCRUD(mux, cfg, "/api/v1/admin/app_user", technical, h.appUser)
	registerCRUD(mux, cfg, "/api/v1/admin/admin", []int{
		helper.AppRoleITDev,
		helper.AppRoleSuperAdmin,
		helper.AppRoleAdmin,
	}, h.appUser)
	registerCRUD(mux, cfg, "/api/v1/admin/driver", operational, h.appUser)
	registerCRUD(mux, cfg, "/api/v1/admin/checker", operational, h.appUser)

	registerRead(mux, cfg, "/api/v1/admin/bank_merk", userManagementReads, h.bankMerk.Get)
	registerMutations(mux, cfg, "/api/v1/admin/bank_merk", operational, h.bankMerk)
	registerCRUD(mux, cfg, "/api/v1/admin/province", operational, h.province)
	registerRead(mux, cfg, "/api/v1/admin/city", userManagementReads, h.city.Get)
	registerMutations(mux, cfg, "/api/v1/admin/city", operational, h.city)
	registerCRUD(mux, cfg, "/api/v1/admin/vendor_type", operational, h.vendorType)
	registerCRUD(mux, cfg, "/api/v1/admin/vendor", operational, h.vendor)
	registerCRUD(mux, cfg, "/api/v1/admin/truck_merk", operational, h.truckMerk)
	registerCRUD(mux, cfg, "/api/v1/admin/truck_type", operational, h.truckType)
	registerCRUD(mux, cfg, "/api/v1/admin/truck", operational, h.truck)
	registerRead(mux, cfg, "/api/v1/admin/client", userManagementReads, h.client.Get)
	registerMutations(mux, cfg, "/api/v1/admin/client", operational, h.client)
	registerCRUD(mux, cfg, "/api/v1/admin/client_destination", operational, h.clientDestination)
	registerCRUD(mux, cfg, "/api/v1/admin/cargo_type", operational, h.cargoType)
	registerCRUD(mux, cfg, "/api/v1/admin/mine", operational, h.mine)
	registerCRUD(mux, cfg, "/api/v1/admin/stockpile", operational, h.stockpile)
	registerCRUD(mux, cfg, "/api/v1/admin/stockpile_cargo", operational, h.stockpileCargo)
	registerCRUD(mux, cfg, "/api/v1/admin/stockpile_adjustment", operational, h.stockpileAdjustment)
	registerRead(mux, cfg, "/api/v1/admin/stockpile_ledger", operational, h.stockpileLedger.Get)
	registerCRUD(mux, cfg, "/api/v1/admin/port", operational, h.port)
	registerCRUD(mux, cfg, "/api/v1/admin/vessel", operational, h.vessel)
	registerCRUD(mux, cfg, "/api/v1/admin/vessel_cargo", operational, h.vesselCargo)

	registerRead(mux, cfg, "/api/v1/admin/project", dashboardRoles, h.project.Get)
	registerMutations(mux, cfg, "/api/v1/admin/project", operational, h.project)
	registerCRUD(mux, cfg, "/api/v1/admin/project_checker_assignment", operational, h.projectCheckerAssignment)
	registerCRUD(mux, cfg, "/api/v1/admin/project_truck_assignment", operational, h.projectTruckAssignment)
	registerCRUD(mux, cfg, "/api/v1/admin/project_route", operational, h.projectRoute)
	registerCRUD(mux, cfg, "/api/v1/admin/project_transport", operational, h.projectTransport)
	registerCRUD(mux, cfg, "/api/v1/admin/project_transport_status", operational, h.projectTransportStatus)
	registerCRUD(mux, cfg, "/api/v1/admin/project_transport_photo", operational, h.projectTransportPhoto)
	registerCRUD(mux, cfg, "/api/v1/admin/project_financial_transaction", operational, h.projectFinancialTransaction)

	mux.Handle("GET /api/v1/admin/dashboard", authenticated(cfg, dashboardRoles, h.dashboard.Get))
	mux.Handle("GET /api/v1/admin/report", authenticated(cfg, reportRoles, h.report.Get))

	mux.Handle("GET /api/v1/checker/app_setting", authenticated(cfg, checkerRoles, h.appSetting.CheckerGet))
	registerRead(mux, cfg, "/api/v1/checker/project", checkerRoles, h.project.CheckerGet)
	mux.Handle("GET /api/v1/checker/project_checker_assignment/default", authenticated(cfg, checkerRoles, h.projectCheckerAssignment.CheckerDefault))
	registerRead(mux, cfg, "/api/v1/checker/project_route", checkerRoles, h.projectRoute.CheckerGet)
	registerRead(mux, cfg, "/api/v1/checker/project_truck_assignment", checkerRoles, h.projectTruckAssignment.CheckerGet)
	registerRead(mux, cfg, "/api/v1/checker/project_transport", checkerRoles, h.projectTransport.CheckerGet)
	registerCreate(mux, cfg, "/api/v1/checker/project_transport", checkerRoles, h.projectTransport.CheckerCreate)
	registerRead(mux, cfg, "/api/v1/checker/project_transport_status", checkerRoles, h.projectTransportStatus.CheckerGet)
	registerCreate(mux, cfg, "/api/v1/checker/project_transport_status", checkerRoles, h.projectTransportStatus.CheckerCreate)
	registerRead(mux, cfg, "/api/v1/checker/project_transport_photo", checkerRoles, h.projectTransportPhoto.CheckerGet)
	registerCreate(mux, cfg, "/api/v1/checker/project_transport_photo", checkerRoles, h.projectTransportPhoto.CheckerCreate)

	registerRead(mux, cfg, "/api/v1/driver/project_transport", driverRoles, h.projectTransport.DriverGet)

	registerRead(mux, cfg, "/api/v1/driver/driver", driverRoles, h.appUser.Get)
	mux.Handle("PUT /api/v1/driver/driver/{id}", authenticated(cfg, driverRoles, h.appUser.Update))
	mux.Handle("PATCH /api/v1/driver/driver/{id}", authenticated(cfg, driverRoles, h.appUser.Update))
	registerRead(mux, cfg, "/api/v1/checker/checker", checkerRoles, h.appUser.Get)
	mux.Handle("PUT /api/v1/checker/checker/{id}", authenticated(cfg, checkerRoles, h.appUser.Update))
	mux.Handle("PATCH /api/v1/checker/checker/{id}", authenticated(cfg, checkerRoles, h.appUser.Update))

	registerRead(mux, cfg, "/api/v1/driver/province", driverRoles, h.province.Get)
	registerRead(mux, cfg, "/api/v1/driver/city", driverRoles, h.city.Get)
	registerRead(mux, cfg, "/api/v1/checker/province", checkerRoles, h.province.Get)
	registerRead(mux, cfg, "/api/v1/checker/city", checkerRoles, h.city.Get)

	mux.Handle("PUT /api/v1/driver/change_password", authenticated(cfg, driverRoles, h.appUser.ChangePassword))
	mux.Handle("PATCH /api/v1/driver/change_password", authenticated(cfg, driverRoles, h.appUser.ChangePassword))
	mux.Handle("PUT /api/v1/checker/change_password", authenticated(cfg, checkerRoles, h.appUser.ChangePassword))
	mux.Handle("PATCH /api/v1/checker/change_password", authenticated(cfg, checkerRoles, h.appUser.ChangePassword))
	mux.Handle("PUT /api/v1/admin/change_password", authenticated(cfg, dashboardRoles, h.appUser.ChangePassword))
	mux.Handle("PATCH /api/v1/admin/change_password", authenticated(cfg, dashboardRoles, h.appUser.ChangePassword))
	mux.Handle("POST /api/v1/logout", authenticated(cfg, allRoles, h.appUser.Logout))
	mux.Handle("POST /api/v1/allrole/presigned", authenticated(cfg, allRoles, h.fileUpload.GetPresignedURL))

}
