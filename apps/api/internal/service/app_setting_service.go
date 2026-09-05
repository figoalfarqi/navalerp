package service

import (
	"context"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type AppSettingService struct {
	Repo *repository.AppSettingRepository
}

func NewAppSettingService(repo *repository.AppSettingRepository) *AppSettingService {
	return &AppSettingService{Repo: repo}
}

func (s *AppSettingService) Create(ctx context.Context, userID int, req *model.AppSettingRequest) (*model.AppSetting, error, map[string]string) {
	if err, fields := validateRequest(req); err != nil {
		return nil, err, fields
	}
	id, err := s.Repo.Create(ctx, req, userID)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id)
	return item, err, nil
}

func (s *AppSettingService) Update(ctx context.Context, userID, id int, req *model.AppSettingRequest) (*model.AppSetting, error, map[string]string) {
	if err, fields := validateRequest(req); err != nil {
		return nil, err, fields
	}
	if err := s.Repo.Update(ctx, id, req, userID); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id)
	return item, err, nil
}

func (s *AppSettingService) Delete(ctx context.Context, userID, id int) error {
	return s.Repo.SoftDelete(ctx, id, userID)
}

func (s *AppSettingService) GetByID(ctx context.Context, id int) (*model.AppSetting, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *AppSettingService) List(ctx context.Context, opts model.ListOptions, checkerOnly bool) ([]model.AppSetting, error) {
	return s.Repo.List(ctx, normalizeListOptions(opts), checkerOnly)
}
