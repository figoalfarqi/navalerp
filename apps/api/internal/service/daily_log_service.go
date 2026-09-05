package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type DailyLogService struct {
	Repo *repository.DailyLogRepository
}

func NewDailyLogService(repo *repository.DailyLogRepository) *DailyLogService {
	return &DailyLogService{Repo: repo}
}

func (s *DailyLogService) GetByID(ctx context.Context, id string) (*model.DailyLog, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *DailyLogService) List(ctx context.Context, opts model.ListOptions) ([]model.DailyLog, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *DailyLogService) Create(ctx context.Context, loginID string, req *model.DailyLogRequest) (*model.DailyLog, error, map[string]string) {
	m := &model.DailyLog{}
	if req.ShipId != nil { m.ShipId = *req.ShipId }
	if req.LogDate != nil { m.LogDate = *req.LogDate }
	if req.Latitude != nil { m.Latitude = *req.Latitude }
	if req.Longitude != nil { m.Longitude = *req.Longitude }
	m.HeadingDegrees = req.HeadingDegrees
	if req.SpeedKnots != nil { m.SpeedKnots = *req.SpeedKnots }
	m.SeaState = req.SeaState
	m.WeatherCondition = req.WeatherCondition
	if req.FuelRemainingLiters != nil { m.FuelRemainingLiters = *req.FuelRemainingLiters }
	if req.FreshWaterRemainingTons != nil { m.FreshWaterRemainingTons = *req.FreshWaterRemainingTons }
	m.TacticalSummary = req.TacticalSummary
	if req.LoggedByUserId != nil { m.LoggedByUserId = *req.LoggedByUserId }
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *DailyLogService) Update(ctx context.Context, loginID string, id string, req *model.DailyLogRequest) (*model.DailyLog, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.ShipId != nil { m.ShipId = *req.ShipId }
	if req.LogDate != nil { m.LogDate = *req.LogDate }
	if req.Latitude != nil { m.Latitude = *req.Latitude }
	if req.Longitude != nil { m.Longitude = *req.Longitude }
	if req.HeadingDegrees != nil { m.HeadingDegrees = req.HeadingDegrees }
	if req.SpeedKnots != nil { m.SpeedKnots = *req.SpeedKnots }
	if req.SeaState != nil { m.SeaState = req.SeaState }
	if req.WeatherCondition != nil { m.WeatherCondition = req.WeatherCondition }
	if req.FuelRemainingLiters != nil { m.FuelRemainingLiters = *req.FuelRemainingLiters }
	if req.FreshWaterRemainingTons != nil { m.FreshWaterRemainingTons = *req.FreshWaterRemainingTons }
	if req.TacticalSummary != nil { m.TacticalSummary = req.TacticalSummary }
	if req.LoggedByUserId != nil { m.LoggedByUserId = *req.LoggedByUserId }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *DailyLogService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
