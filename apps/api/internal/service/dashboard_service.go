package service

import (
	"context"
	"errors"
	"time"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/repository"
)

type DashboardService struct {
	Repo *repository.DashboardRepository
}

func NewDashboardService(repo *repository.DashboardRepository) *DashboardService {
	return &DashboardService{Repo: repo}
}

func (s *DashboardService) Get(ctx context.Context, projectID *int, month string) (*model.DashboardResponse, error) {
	jakarta := time.FixedZone("Asia/Jakarta", 7*60*60)
	if month == "" {
		month = time.Now().In(jakarta).Format("2006-01")
	}
	from, err := time.ParseInLocation("2006-01", month, jakarta)
	if err != nil {
		return nil, errors.New("month must use YYYY-MM")
	}
	to := from.AddDate(0, 1, 0)
	summary, err := s.Repo.Summary(ctx, projectID, from, to)
	if err != nil {
		return nil, err
	}
	projects, err := s.Repo.Projects(ctx, projectID, from, to)
	if err != nil {
		return nil, err
	}
	daily, err := s.Repo.Daily(ctx, projectID, from, to)
	if err != nil {
		return nil, err
	}
	return &model.DashboardResponse{Summary: *summary, Projects: projects, Daily: daily}, nil
}
