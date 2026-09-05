package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type EquipmentService struct {
	Repo *repository.EquipmentRepository
}

func NewEquipmentService(repo *repository.EquipmentRepository) *EquipmentService {
	return &EquipmentService{Repo: repo}
}

func (s *EquipmentService) GetByID(ctx context.Context, id string) (*model.Equipment, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *EquipmentService) List(ctx context.Context, opts model.ListOptions) ([]model.Equipment, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *EquipmentService) Create(ctx context.Context, loginID string, req *model.EquipmentRequest) (*model.Equipment, error, map[string]string) {
	m := &model.Equipment{}
	if req.SystemId != nil { m.SystemId = *req.SystemId }
	if req.SerialNumber != nil { m.SerialNumber = *req.SerialNumber }
	m.EquipmentTag = req.EquipmentTag
	if req.EquipmentName != nil { m.EquipmentName = *req.EquipmentName }
	m.Manufacturer = req.Manufacturer
	m.ModelNumber = req.ModelNumber
	m.CountryOfOrigin = req.CountryOfOrigin
	m.InstallationDate = req.InstallationDate
	m.TotalOperatingHours = req.TotalOperatingHours
	m.DesignLifeHours = req.DesignLifeHours
	m.CriticalityLevel = req.CriticalityLevel
	m.HealthStatus = req.HealthStatus
	m.Parameters = req.Parameters
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *EquipmentService) Update(ctx context.Context, loginID string, id string, req *model.EquipmentRequest) (*model.Equipment, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.SystemId != nil { m.SystemId = *req.SystemId }
	if req.SerialNumber != nil { m.SerialNumber = *req.SerialNumber }
	if req.EquipmentTag != nil { m.EquipmentTag = req.EquipmentTag }
	if req.EquipmentName != nil { m.EquipmentName = *req.EquipmentName }
	if req.Manufacturer != nil { m.Manufacturer = req.Manufacturer }
	if req.ModelNumber != nil { m.ModelNumber = req.ModelNumber }
	if req.CountryOfOrigin != nil { m.CountryOfOrigin = req.CountryOfOrigin }
	if req.InstallationDate != nil { m.InstallationDate = req.InstallationDate }
	if req.TotalOperatingHours != nil { m.TotalOperatingHours = req.TotalOperatingHours }
	if req.DesignLifeHours != nil { m.DesignLifeHours = req.DesignLifeHours }
	if req.CriticalityLevel != nil { m.CriticalityLevel = req.CriticalityLevel }
	if req.HealthStatus != nil { m.HealthStatus = req.HealthStatus }
	if req.Parameters != nil { m.Parameters = req.Parameters }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *EquipmentService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
