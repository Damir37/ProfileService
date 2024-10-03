package usecases

import (
	"ProfileService/internal/entity"
	"ProfileService/internal/repository"
	"context"
	"fmt"
	"github.com/rs/zerolog"
)

type ProfileServiceImpl struct {
	profileService ProfileService

	profileRepo repository.ProfileRepository
	logger      zerolog.Logger
}

func NewProfileService(profileRepo repository.ProfileRepository, logger zerolog.Logger) ProfileService {
	return &ProfileServiceImpl{
		profileRepo: profileRepo,
		logger:      logger,
	}
}

func (ps *ProfileServiceImpl) GetProfile(ctx context.Context, id string) (*entity.FullProfile, error) {
	if id == "" {
		return nil, fmt.Errorf("422")
	}
	return ps.profileRepo.GetProfile(ctx, id)
}

func (ps *ProfileServiceImpl) GetStreamKey(ctx context.Context, id string) (string, error) {
	if id == "" {
		return "", fmt.Errorf("422")
	}
	return ps.profileRepo.GetStreamKey(ctx, id)
}

func (ps *ProfileServiceImpl) GetAvatarProfile(ctx context.Context, id string) (string, error) {
	if id == "" {
		return "", fmt.Errorf("422")
	}
	return ps.profileRepo.GetAvatarProfile(ctx, id)
}

func (ps *ProfileServiceImpl) GetProfileStats(ctx context.Context, id string) (*entity.UserStats, error) {
	if id == "" {
		return nil, fmt.Errorf("422")
	}
	return ps.profileRepo.GetProfileStats(ctx, id)
}

func (ps *ProfileServiceImpl) SetAdmin(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("422")
	}
	return ps.profileRepo.SetAdmin(ctx, id)
}

func (ps *ProfileServiceImpl) UnsetAdmin(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("422")
	}
	return ps.profileRepo.UnsetAdmin(ctx, id)
}

func (ps *ProfileServiceImpl) RandomStreamKey(ctx context.Context, id string) (string, error) {
	if id == "" {
		return "", fmt.Errorf("422")
	}
	return ps.profileRepo.RandomStreamKey(ctx, id)
}

func (ps *ProfileServiceImpl) EditAvatarProfile(ctx context.Context, user *entity.User) (string, error) {
	return ps.profileRepo.EditAvatarProfile(ctx, user)
}

func (ps *ProfileServiceImpl) EditProfile(ctx context.Context, profile *entity.FullProfile) error {
	return ps.profileRepo.EditProfile(ctx, profile)
}

func (ps *ProfileServiceImpl) ModifyElo(ctx context.Context, elomodify *entity.UserModifyELO) (int, error) {
	return ps.profileRepo.ModifyElo(ctx, elomodify)
}
