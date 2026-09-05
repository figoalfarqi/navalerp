package repository

import (
	"context"
	"time"

	"github.com/figoalfarqi/navalerp/internal/model"
)

type ReportRepository struct {
	Dashboard *DashboardRepository
	Transport *ProjectTransportRepository
	Financial *ProjectFinancialTransactionRepository
}

func NewReportRepository(
	dashboard *DashboardRepository,
	transport *ProjectTransportRepository,
	financial *ProjectFinancialTransactionRepository,
) *ReportRepository {
	return &ReportRepository{Dashboard: dashboard, Transport: transport, Financial: financial}
}

func (r *ReportRepository) Build(
	ctx context.Context,
	projectID *int,
	from, to time.Time,
	period, periodLabel string,
) (*model.ReportResponse, error) {
	summary, err := r.Dashboard.Summary(ctx, projectID, from, to)
	if err != nil {
		return nil, err
	}
	dashboardProjects, err := r.Dashboard.Projects(ctx, projectID, from, to)
	if err != nil {
		return nil, err
	}
	projects := make([]model.ReportProject, 0, len(dashboardProjects))
	for _, p := range dashboardProjects {
		projects = append(projects, model.ReportProject{
			ProjectID:        p.ProjectID,
			ProjectCode:      p.ProjectCode,
			ProjectName:      p.ProjectName,
			PeriodLabel:      periodLabel,
			DashboardSummary: p.DashboardSummary,
		})
	}
	opts := model.ListOptions{
		Limit:     10000,
		ProjectID: projectID,
		DateFrom:  &from,
		DateTo:    &to,
	}
	transports, err := r.Transport.List(ctx, opts)
	if err != nil {
		return nil, err
	}
	financial, err := r.Financial.List(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &model.ReportResponse{
		Period:                period,
		PeriodLabel:           periodLabel,
		DateFrom:              from.Format("2006-01-02"),
		DateTo:                to.AddDate(0, 0, -1).Format("2006-01-02"),
		Summary:               *summary,
		Projects:              projects,
		Transports:            transports,
		FinancialTransactions: financial,
	}, nil
}
