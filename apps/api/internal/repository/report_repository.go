package repository

import (
	"context"
	"time"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ReportRepository struct {
	DB *pgxpool.Pool
}

func NewReportRepository(db *pgxpool.Pool) *ReportRepository {
	return &ReportRepository{DB: db}
}

func (r *ReportRepository) GetReportData(ctx context.Context) (*model.ReportResponse, error) {
	dashRepo := NewDashboardRepository(r.DB)
	dash, _ := dashRepo.GetDashboardData(ctx)

	var items []model.NavalReportItem
	if dash != nil {
		items = append(items, model.NavalReportItem{
			Category: "Armada & KRI",
			Name:     "Kesiapan Tempur KRI (FMC)",
			Status:   "OPERATIONAL",
			Metric:   dash.Summary.OverallReadinessScore,
			Notes:    "Rasio kapal KRI yang siap tempur penuh",
		})
		items = append(items, model.NavalReportItem{
			Category: "SDM Militer",
			Name:     "Kekuatan Personel Aktif",
			Status:   "ACTIVE",
			Metric:   float64(dash.Summary.TotalPersonnel),
			Notes:    "Jumlah prajurit terdata",
		})
		items = append(items, model.NavalReportItem{
			Category: "Operasi & Misi",
			Name:     "Misi Tempur Berjalan",
			Status:   "DEPLOYED",
			Metric:   float64(dash.Summary.ActiveMissions),
			Notes:    "Jumlah misi operasi laut aktif",
		})
		items = append(items, model.NavalReportItem{
			Category: "Pemeliharaan (MRO)",
			Name:     "Work Order Terbuka",
			Status:   "IN_PROGRESS",
			Metric:   float64(dash.Summary.OpenWorkOrders),
			Notes:    "Perintah kerja perbaikan sedang diproses",
		})
	}

	return &model.ReportResponse{
		Period:    "Tahun Anggaran Berjalan",
		Generated: time.Now().Format("2006-01-02 15:04:05"),
		Summary:   dash.Summary,
		Items:     items,
	}, nil
}
