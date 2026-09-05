package service

import (
	"context"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type DashboardService struct {
	Repo *repository.DashboardRepository
}

func NewDashboardService(repo *repository.DashboardRepository) *DashboardService {
	return &DashboardService{Repo: repo}
}

func (s *DashboardService) Get(ctx context.Context, projectID *int, month string) (*model.DashboardResponse, error) {
	return s.Repo.GetDashboardData(ctx)
}
