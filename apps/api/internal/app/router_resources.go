package app

import (
	"github.com/figoalfarqi/navalerp/config"
	"github.com/figoalfarqi/navalerp/internal/handler"
	"github.com/figoalfarqi/navalerp/internal/repository"
	"github.com/figoalfarqi/navalerp/internal/service"
	"github.com/figoalfarqi/navalerp/pkg/database"
)

type routeHandlers struct {
	fileUpload *handler.FileUploadHandler
	dashboard  *handler.DashboardHandler
	report     *handler.ReportHandler
	navalAuth  *handler.NavalAuthHandler
	orgUnit *handler.OrgUnitHandler
	sysUser *handler.SysUserHandler
	auditLog *handler.AuditLogHandler
	shipClass *handler.ShipClassHandler
	ship *handler.ShipHandler
	shipSystem *handler.ShipSystemHandler
	equipment *handler.EquipmentHandler
	pmSchedule *handler.PmScheduleHandler
	failureReport *handler.FailureReportHandler
	workOrder *handler.WorkOrderHandler
	dockingRecord *handler.DockingRecordHandler
	warehouse *handler.WarehouseHandler
	material *handler.MaterialHandler
	stockBalance *handler.StockBalanceHandler
	itemInstance *handler.ItemInstanceHandler
	stockTransfer *handler.StockTransferHandler
	stockAdjustment *handler.StockAdjustmentHandler
	vendor *handler.VendorHandler
	requisition *handler.RequisitionHandler
	tender *handler.TenderHandler
	contract *handler.ContractHandler
	purchaseOrder *handler.PurchaseOrderHandler
	goodsReceipt *handler.GoodsReceiptHandler
	chartOfAccount *handler.ChartOfAccountHandler
	budgetProgram *handler.BudgetProgramHandler
	budgetCommitment *handler.BudgetCommitmentHandler
	invoice *handler.InvoiceHandler
	payment *handler.PaymentHandler
	journalEntry *handler.JournalEntryHandler
	platformTco *handler.PlatformTcoHandler
	militaryRank *handler.MilitaryRankHandler
	militaryCorps *handler.MilitaryCorpsHandler
	qualification *handler.QualificationHandler
	personnel *handler.PersonnelHandler
	crewAssignment *handler.CrewAssignmentHandler
	baseFacility *handler.BaseFacilityHandler
	berthBooking *handler.BerthBookingHandler
	fuelBunker *handler.FuelBunkerHandler
	transportUnit *handler.TransportUnitHandler
	route *handler.RouteHandler
	shipment *handler.ShipmentHandler
	documentCategory *handler.DocumentCategoryHandler
	document *handler.DocumentHandler
	theater *handler.TheaterHandler
	mission *handler.MissionHandler
	dailyLog *handler.DailyLogHandler
	readinessReport *handler.ReadinessReportHandler
	readinessAlert *handler.ReadinessAlertHandler
}

func buildRouteHandlers(cfg *config.Config) *routeHandlers {
	db := database.GetPool()
	fileService := service.NewFileUploadService(cfg)
	dashboardRepo := repository.NewDashboardRepository(db)
	dashboardService := service.NewDashboardService(dashboardRepo)
	reportRepo := repository.NewReportRepository(db)
	reportService := service.NewReportService(reportRepo)

	orgUnitRepo := repository.NewOrgUnitRepository(db)
	orgUnitService := service.NewOrgUnitService(orgUnitRepo)
	sysUserRepo := repository.NewSysUserRepository(db)
	sysUserService := service.NewSysUserService(sysUserRepo)
	auditLogRepo := repository.NewAuditLogRepository(db)
	auditLogService := service.NewAuditLogService(auditLogRepo)
	shipClassRepo := repository.NewShipClassRepository(db)
	shipClassService := service.NewShipClassService(shipClassRepo)
	shipRepo := repository.NewShipRepository(db)
	shipService := service.NewShipService(shipRepo)
	shipSystemRepo := repository.NewShipSystemRepository(db)
	shipSystemService := service.NewShipSystemService(shipSystemRepo)
	equipmentRepo := repository.NewEquipmentRepository(db)
	equipmentService := service.NewEquipmentService(equipmentRepo)
	pmScheduleRepo := repository.NewPmScheduleRepository(db)
	pmScheduleService := service.NewPmScheduleService(pmScheduleRepo)
	failureReportRepo := repository.NewFailureReportRepository(db)
	failureReportService := service.NewFailureReportService(failureReportRepo)
	workOrderRepo := repository.NewWorkOrderRepository(db)
	workOrderService := service.NewWorkOrderService(workOrderRepo)
	dockingRecordRepo := repository.NewDockingRecordRepository(db)
	dockingRecordService := service.NewDockingRecordService(dockingRecordRepo)
	warehouseRepo := repository.NewWarehouseRepository(db)
	warehouseService := service.NewWarehouseService(warehouseRepo)
	materialRepo := repository.NewMaterialRepository(db)
	materialService := service.NewMaterialService(materialRepo)
	stockBalanceRepo := repository.NewStockBalanceRepository(db)
	stockBalanceService := service.NewStockBalanceService(stockBalanceRepo)
	itemInstanceRepo := repository.NewItemInstanceRepository(db)
	itemInstanceService := service.NewItemInstanceService(itemInstanceRepo)
	stockTransferRepo := repository.NewStockTransferRepository(db)
	stockTransferService := service.NewStockTransferService(stockTransferRepo)
	stockAdjustmentRepo := repository.NewStockAdjustmentRepository(db)
	stockAdjustmentService := service.NewStockAdjustmentService(stockAdjustmentRepo)
	vendorRepo := repository.NewVendorRepository(db)
	vendorService := service.NewVendorService(vendorRepo)
	requisitionRepo := repository.NewRequisitionRepository(db)
	requisitionService := service.NewRequisitionService(requisitionRepo)
	tenderRepo := repository.NewTenderRepository(db)
	tenderService := service.NewTenderService(tenderRepo)
	contractRepo := repository.NewContractRepository(db)
	contractService := service.NewContractService(contractRepo)
	purchaseOrderRepo := repository.NewPurchaseOrderRepository(db)
	purchaseOrderService := service.NewPurchaseOrderService(purchaseOrderRepo)
	goodsReceiptRepo := repository.NewGoodsReceiptRepository(db)
	goodsReceiptService := service.NewGoodsReceiptService(goodsReceiptRepo)
	chartOfAccountRepo := repository.NewChartOfAccountRepository(db)
	chartOfAccountService := service.NewChartOfAccountService(chartOfAccountRepo)
	budgetProgramRepo := repository.NewBudgetProgramRepository(db)
	budgetProgramService := service.NewBudgetProgramService(budgetProgramRepo)
	budgetCommitmentRepo := repository.NewBudgetCommitmentRepository(db)
	budgetCommitmentService := service.NewBudgetCommitmentService(budgetCommitmentRepo)
	invoiceRepo := repository.NewInvoiceRepository(db)
	invoiceService := service.NewInvoiceService(invoiceRepo)
	paymentRepo := repository.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepo)
	journalEntryRepo := repository.NewJournalEntryRepository(db)
	journalEntryService := service.NewJournalEntryService(journalEntryRepo)
	platformTcoRepo := repository.NewPlatformTcoRepository(db)
	platformTcoService := service.NewPlatformTcoService(platformTcoRepo)
	militaryRankRepo := repository.NewMilitaryRankRepository(db)
	militaryRankService := service.NewMilitaryRankService(militaryRankRepo)
	militaryCorpsRepo := repository.NewMilitaryCorpsRepository(db)
	militaryCorpsService := service.NewMilitaryCorpsService(militaryCorpsRepo)
	qualificationRepo := repository.NewQualificationRepository(db)
	qualificationService := service.NewQualificationService(qualificationRepo)
	personnelRepo := repository.NewPersonnelRepository(db)
	personnelService := service.NewPersonnelService(personnelRepo)
	crewAssignmentRepo := repository.NewCrewAssignmentRepository(db)
	crewAssignmentService := service.NewCrewAssignmentService(crewAssignmentRepo)
	baseFacilityRepo := repository.NewBaseFacilityRepository(db)
	baseFacilityService := service.NewBaseFacilityService(baseFacilityRepo)
	berthBookingRepo := repository.NewBerthBookingRepository(db)
	berthBookingService := service.NewBerthBookingService(berthBookingRepo)
	fuelBunkerRepo := repository.NewFuelBunkerRepository(db)
	fuelBunkerService := service.NewFuelBunkerService(fuelBunkerRepo)
	transportUnitRepo := repository.NewTransportUnitRepository(db)
	transportUnitService := service.NewTransportUnitService(transportUnitRepo)
	routeRepo := repository.NewRouteRepository(db)
	routeService := service.NewRouteService(routeRepo)
	shipmentRepo := repository.NewShipmentRepository(db)
	shipmentService := service.NewShipmentService(shipmentRepo)
	documentCategoryRepo := repository.NewDocumentCategoryRepository(db)
	documentCategoryService := service.NewDocumentCategoryService(documentCategoryRepo)
	documentRepo := repository.NewDocumentRepository(db)
	documentService := service.NewDocumentService(documentRepo)
	theaterRepo := repository.NewTheaterRepository(db)
	theaterService := service.NewTheaterService(theaterRepo)
	missionRepo := repository.NewMissionRepository(db)
	missionService := service.NewMissionService(missionRepo)
	dailyLogRepo := repository.NewDailyLogRepository(db)
	dailyLogService := service.NewDailyLogService(dailyLogRepo)
	readinessReportRepo := repository.NewReadinessReportRepository(db)
	readinessReportService := service.NewReadinessReportService(readinessReportRepo)
	readinessAlertRepo := repository.NewReadinessAlertRepository(db)
	readinessAlertService := service.NewReadinessAlertService(readinessAlertRepo)

	return &routeHandlers{
		fileUpload: handler.NewFileUploadHandler(fileService, cfg),
		dashboard:  handler.NewDashboardHandler(dashboardService, cfg),
		report:     handler.NewReportHandler(reportService, cfg),
		navalAuth:  handler.NewNavalAuthHandler(db, cfg),
		orgUnit: handler.NewOrgUnitHandler(orgUnitService, cfg),
		sysUser: handler.NewSysUserHandler(sysUserService, cfg),
		auditLog: handler.NewAuditLogHandler(auditLogService, cfg),
		shipClass: handler.NewShipClassHandler(shipClassService, cfg),
		ship: handler.NewShipHandler(shipService, cfg),
		shipSystem: handler.NewShipSystemHandler(shipSystemService, cfg),
		equipment: handler.NewEquipmentHandler(equipmentService, cfg),
		pmSchedule: handler.NewPmScheduleHandler(pmScheduleService, cfg),
		failureReport: handler.NewFailureReportHandler(failureReportService, cfg),
		workOrder: handler.NewWorkOrderHandler(workOrderService, cfg),
		dockingRecord: handler.NewDockingRecordHandler(dockingRecordService, cfg),
		warehouse: handler.NewWarehouseHandler(warehouseService, cfg),
		material: handler.NewMaterialHandler(materialService, cfg),
		stockBalance: handler.NewStockBalanceHandler(stockBalanceService, cfg),
		itemInstance: handler.NewItemInstanceHandler(itemInstanceService, cfg),
		stockTransfer: handler.NewStockTransferHandler(stockTransferService, cfg),
		stockAdjustment: handler.NewStockAdjustmentHandler(stockAdjustmentService, cfg),
		vendor: handler.NewVendorHandler(vendorService, cfg),
		requisition: handler.NewRequisitionHandler(requisitionService, cfg),
		tender: handler.NewTenderHandler(tenderService, cfg),
		contract: handler.NewContractHandler(contractService, cfg),
		purchaseOrder: handler.NewPurchaseOrderHandler(purchaseOrderService, cfg),
		goodsReceipt: handler.NewGoodsReceiptHandler(goodsReceiptService, cfg),
		chartOfAccount: handler.NewChartOfAccountHandler(chartOfAccountService, cfg),
		budgetProgram: handler.NewBudgetProgramHandler(budgetProgramService, cfg),
		budgetCommitment: handler.NewBudgetCommitmentHandler(budgetCommitmentService, cfg),
		invoice: handler.NewInvoiceHandler(invoiceService, cfg),
		payment: handler.NewPaymentHandler(paymentService, cfg),
		journalEntry: handler.NewJournalEntryHandler(journalEntryService, cfg),
		platformTco: handler.NewPlatformTcoHandler(platformTcoService, cfg),
		militaryRank: handler.NewMilitaryRankHandler(militaryRankService, cfg),
		militaryCorps: handler.NewMilitaryCorpsHandler(militaryCorpsService, cfg),
		qualification: handler.NewQualificationHandler(qualificationService, cfg),
		personnel: handler.NewPersonnelHandler(personnelService, cfg),
		crewAssignment: handler.NewCrewAssignmentHandler(crewAssignmentService, cfg),
		baseFacility: handler.NewBaseFacilityHandler(baseFacilityService, cfg),
		berthBooking: handler.NewBerthBookingHandler(berthBookingService, cfg),
		fuelBunker: handler.NewFuelBunkerHandler(fuelBunkerService, cfg),
		transportUnit: handler.NewTransportUnitHandler(transportUnitService, cfg),
		route: handler.NewRouteHandler(routeService, cfg),
		shipment: handler.NewShipmentHandler(shipmentService, cfg),
		documentCategory: handler.NewDocumentCategoryHandler(documentCategoryService, cfg),
		document: handler.NewDocumentHandler(documentService, cfg),
		theater: handler.NewTheaterHandler(theaterService, cfg),
		mission: handler.NewMissionHandler(missionService, cfg),
		dailyLog: handler.NewDailyLogHandler(dailyLogService, cfg),
		readinessReport: handler.NewReadinessReportHandler(readinessReportService, cfg),
		readinessAlert: handler.NewReadinessAlertHandler(readinessAlertService, cfg),
	}
}
