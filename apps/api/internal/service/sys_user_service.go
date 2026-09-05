package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/auth"
	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type SysUserService struct {
	Repo *repository.SysUserRepository
}

func NewSysUserService(repo *repository.SysUserRepository) *SysUserService {
	return &SysUserService{Repo: repo}
}

func (s *SysUserService) GetByID(ctx context.Context, id string) (*model.SysUser, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *SysUserService) List(ctx context.Context, opts model.ListOptions) ([]model.SysUser, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *SysUserService) Create(ctx context.Context, loginID string, req *model.SysUserRequest) (*model.SysUser, error, map[string]string) {
	m := &model.SysUser{}
	if req.UnitId != nil { m.UnitId = *req.UnitId }
	if req.Username != nil { m.Username = *req.Username }
	if req.PasswordHash != nil { m.PasswordHash = *req.PasswordHash }
	if req.FullName != nil { m.FullName = *req.FullName }
	m.Email = req.Email
	m.Phone = req.Phone
	m.MilitaryId = req.MilitaryId
	m.RankTitle = req.RankTitle
	m.Department = req.Department
	if req.Role != nil { m.Role = *req.Role }
	m.IsActive = req.IsActive
	m.LastLoginAt = req.LastLoginAt
	m.FailedLoginAttempts = req.FailedLoginAttempts
	if req.AuthVersion != nil { m.AuthVersion = *req.AuthVersion }
	if req.PasswordHash != nil && *req.PasswordHash != "" {
		hashed, err := auth.HashPassword(*req.PasswordHash)
		if err == nil { m.PasswordHash = hashed }
	}
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *SysUserService) Update(ctx context.Context, loginID string, id string, req *model.SysUserRequest) (*model.SysUser, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.UnitId != nil { m.UnitId = *req.UnitId }
	if req.Username != nil { m.Username = *req.Username }
	if req.PasswordHash != nil { m.PasswordHash = *req.PasswordHash }
	if req.FullName != nil { m.FullName = *req.FullName }
	if req.Email != nil { m.Email = req.Email }
	if req.Phone != nil { m.Phone = req.Phone }
	if req.MilitaryId != nil { m.MilitaryId = req.MilitaryId }
	if req.RankTitle != nil { m.RankTitle = req.RankTitle }
	if req.Department != nil { m.Department = req.Department }
	if req.Role != nil { m.Role = *req.Role }
	if req.IsActive != nil { m.IsActive = req.IsActive }
	if req.LastLoginAt != nil { m.LastLoginAt = req.LastLoginAt }
	if req.FailedLoginAttempts != nil { m.FailedLoginAttempts = req.FailedLoginAttempts }
	if req.AuthVersion != nil { m.AuthVersion = *req.AuthVersion }
	if req.PasswordHash != nil && *req.PasswordHash != "" {
		hashed, err := auth.HashPassword(*req.PasswordHash)
		if err == nil { m.PasswordHash = hashed }
	}
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *SysUserService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
