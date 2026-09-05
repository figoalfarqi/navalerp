package app

import (
	"github.com/figoalfarqi/navalerp/config"
	"github.com/figoalfarqi/navalerp/internal/handler"
	"github.com/figoalfarqi/navalerp/internal/repository"
	"github.com/figoalfarqi/navalerp/internal/service"
	"github.com/figoalfarqi/navalerp/pkg/database"
)

type routeHandlers struct {
	fileUpload                  *handler.FileUploadHandler
	appRole                     *handler.AppRoleHandler
	appSetting                  *handler.AppSettingHandler
	appUser                     *handler.AppUserHandler
	bankMerk                    *handler.BankMerkHandler
	province                    *handler.ProvinceHandler
	city                        *handler.CityHandler
	vendorType                  *handler.VendorTypeHandler
	vendor                      *handler.VendorHandler
	truckMerk                   *handler.TruckMerkHandler
	truckType                   *handler.TruckTypeHandler
	truck                       *handler.TruckHandler
	client                      *handler.ClientHandler
	clientDestination           *handler.ClientDestinationHandler
	cargoType                   *handler.CargoTypeHandler
	mine                        *handler.MineHandler
	stockpile                   *handler.StockpileHandler
	stockpileCargo              *handler.StockpileCargoHandler
	stockpileAdjustment         *handler.StockpileAdjustmentHandler
	stockpileLedger             *handler.StockpileLedgerHandler
	port                        *handler.PortHandler
	vessel                      *handler.VesselHandler
	vesselCargo                 *handler.VesselCargoHandler
	project                     *handler.ProjectHandler
	projectCheckerAssignment    *handler.ProjectCheckerAssignmentHandler
	projectTruckAssignment      *handler.ProjectTruckAssignmentHandler
	projectRoute                *handler.ProjectRouteHandler
	projectTransport            *handler.ProjectTransportHandler
	projectTransportStatus      *handler.ProjectTransportStatusHandler
	projectTransportPhoto       *handler.ProjectTransportPhotoHandler
	projectFinancialTransaction *handler.ProjectFinancialTransactionHandler
	dashboard                   *handler.DashboardHandler
	report                      *handler.ReportHandler
}

func buildRouteHandlers(cfg *config.Config) *routeHandlers {
	db := database.GetPool()
	fileService := service.NewFileUploadService(cfg)

	appRoleRepository := repository.NewAppRoleRepository(db)
	appSettingRepository := repository.NewAppSettingRepository(db)
	appUserRepository := repository.NewAppUserRepository(db)
	bankMerkRepository := repository.NewBankMerkRepository(db)
	provinceRepository := repository.NewProvinceRepository(db)
	cityRepository := repository.NewCityRepository(db)
	vendorTypeRepository := repository.NewVendorTypeRepository(db)
	vendorRepository := repository.NewVendorRepository(db)
	truckMerkRepository := repository.NewTruckMerkRepository(db)
	truckTypeRepository := repository.NewTruckTypeRepository(db)
	truckRepository := repository.NewTruckRepository(db)
	clientRepository := repository.NewClientRepository(db)
	clientDestinationRepository := repository.NewClientDestinationRepository(db)
	cargoTypeRepository := repository.NewCargoTypeRepository(db)
	mineRepository := repository.NewMineRepository(db)
	stockpileRepository := repository.NewStockpileRepository(db)
	stockpileCargoRepository := repository.NewStockpileCargoRepository(db)
	stockpileAdjustmentRepository := repository.NewStockpileAdjustmentRepository(db)
	stockpileLedgerRepository := repository.NewStockpileLedgerRepository(db)

	projectRepository := repository.NewProjectRepository(db)
	projectCheckerAssignmentRepository := repository.NewProjectCheckerAssignmentRepository(db)
	projectTruckAssignmentRepository := repository.NewProjectTruckAssignmentRepository(db)
	projectRouteRepository := repository.NewProjectRouteRepository(db)
	projectTransportRepository := repository.NewProjectTransportRepository(db)
	projectTransportStatusRepository := repository.NewProjectTransportStatusRepository(db)
	projectTransportPhotoRepository := repository.NewProjectTransportPhotoRepository(db)
	projectFinancialRepository := repository.NewProjectFinancialTransactionRepository(db)
	dashboardRepository := repository.NewDashboardRepository(db)

	stockpileLedgerService := service.NewStockpileLedgerService(stockpileLedgerRepository)

	return &routeHandlers{
		fileUpload:                  handler.NewFileUploadHandler(fileService, cfg),
		appRole:                     handler.NewAppRoleHandler(service.NewAppRoleService(appRoleRepository), cfg),
		appSetting:                  handler.NewAppSettingHandler(service.NewAppSettingService(appSettingRepository)),
		appUser:                     handler.NewAppUserHandler(service.NewAppUserService(appUserRepository, fileService), cfg),
		bankMerk:                    handler.NewBankMerkHandler(service.NewBankMerkService(bankMerkRepository), cfg),
		province:                    handler.NewProvinceHandler(service.NewProvinceService(provinceRepository), cfg),
		city:                        handler.NewCityHandler(service.NewCityService(cityRepository), cfg),
		vendorType:                  handler.NewVendorTypeHandler(service.NewVendorTypeService(vendorTypeRepository), cfg),
		vendor:                      handler.NewVendorHandler(service.NewVendorService(vendorRepository), cfg),
		truckMerk:                   handler.NewTruckMerkHandler(service.NewTruckMerkService(truckMerkRepository), cfg),
		truckType:                   handler.NewTruckTypeHandler(service.NewTruckTypeService(truckTypeRepository), cfg),
		truck:                       handler.NewTruckHandler(service.NewTruckService(truckRepository), cfg),
		client:                      handler.NewClientHandler(service.NewClientService(clientRepository), cfg),
		clientDestination:           handler.NewClientDestinationHandler(service.NewClientDestinationService(clientDestinationRepository), cfg),
		cargoType:                   handler.NewCargoTypeHandler(service.NewCargoTypeService(cargoTypeRepository), cfg),
		mine:                        handler.NewMineHandler(service.NewMineService(mineRepository), cfg),
		stockpile:                   handler.NewStockpileHandler(service.NewStockpileService(stockpileRepository), cfg),
		stockpileCargo:              handler.NewStockpileCargoHandler(service.NewStockpileCargoService(stockpileCargoRepository), cfg),
		stockpileAdjustment:         handler.NewStockpileAdjustmentHandler(service.NewStockpileAdjustmentService(stockpileAdjustmentRepository, stockpileCargoRepository, stockpileLedgerRepository, stockpileLedgerService), cfg),
		stockpileLedger:             handler.NewStockpileLedgerHandler(stockpileLedgerService, cfg),
		port:                        handler.NewPortHandler(service.NewPortService(repository.NewPortRepository(db))),
		vessel:                      handler.NewVesselHandler(service.NewVesselService(repository.NewVesselRepository(db))),
		vesselCargo:                 handler.NewVesselCargoHandler(service.NewVesselCargoService(repository.NewVesselCargoRepository(db))),
		project:                     handler.NewProjectHandler(service.NewProjectService(projectRepository)),
		projectCheckerAssignment:    handler.NewProjectCheckerAssignmentHandler(service.NewProjectCheckerAssignmentService(projectCheckerAssignmentRepository)),
		projectTruckAssignment:      handler.NewProjectTruckAssignmentHandler(service.NewProjectTruckAssignmentService(projectTruckAssignmentRepository)),
		projectRoute:                handler.NewProjectRouteHandler(service.NewProjectRouteService(projectRouteRepository)),
		projectTransport:            handler.NewProjectTransportHandler(service.NewProjectTransportService(projectTransportRepository, projectRepository, projectTruckAssignmentRepository)),
		projectTransportStatus:      handler.NewProjectTransportStatusHandler(service.NewProjectTransportStatusService(projectTransportStatusRepository, projectTransportRepository, projectRepository)),
		projectTransportPhoto:       handler.NewProjectTransportPhotoHandler(service.NewProjectTransportPhotoService(projectTransportPhotoRepository, projectRepository, fileService)),
		projectFinancialTransaction: handler.NewProjectFinancialTransactionHandler(service.NewProjectFinancialTransactionService(projectFinancialRepository)),
		dashboard:                   handler.NewDashboardHandler(service.NewDashboardService(dashboardRepository)),
		report: handler.NewReportHandler(service.NewReportService(repository.NewReportRepository(
			dashboardRepository,
			projectTransportRepository,
			projectFinancialRepository,
		))),
	}
}
