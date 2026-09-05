package service

import (
	"context"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type ReportService struct {
	Repo *repository.ReportRepository
}

func NewReportService(repo *repository.ReportRepository) *ReportService {
	return &ReportService{Repo: repo}
}

func (s *ReportService) Get(ctx context.Context, projectID *int, month string) (*model.ReportResponse, error) {
	return s.Repo.GetReportData(ctx)
}
