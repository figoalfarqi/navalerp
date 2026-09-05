package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type PersonnelService struct {
	Repo *repository.PersonnelRepository
}

func NewPersonnelService(repo *repository.PersonnelRepository) *PersonnelService {
	return &PersonnelService{Repo: repo}
}

func (s *PersonnelService) GetByID(ctx context.Context, id string) (*model.Personnel, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *PersonnelService) List(ctx context.Context, opts model.ListOptions) ([]model.Personnel, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *PersonnelService) Create(ctx context.Context, loginID string, req *model.PersonnelRequest) (*model.Personnel, error, map[string]string) {
	m := &model.Personnel{}
	if req.Nrp != nil { m.Nrp = *req.Nrp }
	if req.FullName != nil { m.FullName = *req.FullName }
	if req.RankId != nil { m.RankId = *req.RankId }
	if req.CorpsId != nil { m.CorpsId = *req.CorpsId }
	if req.CurrentUnitId != nil { m.CurrentUnitId = *req.CurrentUnitId }
	if req.CurrentPosition != nil { m.CurrentPosition = *req.CurrentPosition }
	m.BirthPlace = req.BirthPlace
	if req.BirthDate != nil { m.BirthDate = *req.BirthDate }
	m.Gender = req.Gender
	m.BloodType = req.BloodType
	m.Religion = req.Religion
	m.EducationLevel = req.EducationLevel
	if req.ServiceEntryDate != nil { m.ServiceEntryDate = *req.ServiceEntryDate }
	m.UserId = req.UserId
	m.Status = req.Status
	m.ServiceRecords = req.ServiceRecords
	m.Qualifications = req.Qualifications
	m.MedicalReadiness = req.MedicalReadiness
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *PersonnelService) Update(ctx context.Context, loginID string, id string, req *model.PersonnelRequest) (*model.Personnel, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.Nrp != nil { m.Nrp = *req.Nrp }
	if req.FullName != nil { m.FullName = *req.FullName }
	if req.RankId != nil { m.RankId = *req.RankId }
	if req.CorpsId != nil { m.CorpsId = *req.CorpsId }
	if req.CurrentUnitId != nil { m.CurrentUnitId = *req.CurrentUnitId }
	if req.CurrentPosition != nil { m.CurrentPosition = *req.CurrentPosition }
	if req.BirthPlace != nil { m.BirthPlace = req.BirthPlace }
	if req.BirthDate != nil { m.BirthDate = *req.BirthDate }
	if req.Gender != nil { m.Gender = req.Gender }
	if req.BloodType != nil { m.BloodType = req.BloodType }
	if req.Religion != nil { m.Religion = req.Religion }
	if req.EducationLevel != nil { m.EducationLevel = req.EducationLevel }
	if req.ServiceEntryDate != nil { m.ServiceEntryDate = *req.ServiceEntryDate }
	if req.UserId != nil { m.UserId = req.UserId }
	if req.Status != nil { m.Status = req.Status }
	if req.ServiceRecords != nil { m.ServiceRecords = req.ServiceRecords }
	if req.Qualifications != nil { m.Qualifications = req.Qualifications }
	if req.MedicalReadiness != nil { m.MedicalReadiness = req.MedicalReadiness }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *PersonnelService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
