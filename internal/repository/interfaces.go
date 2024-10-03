package repository

import (
	"ProfileService/internal/entity"
	"context"
)

type (
	ProfileRepository interface {
		GetProfile(ctx context.Context, id string) (*entity.FullProfile, error)
		GetStreamKey(ctx context.Context, id string) (string, error)
		GetAvatarProfile(ctx context.Context, id string) (string, error)
		GetProfileStats(ctx context.Context, id string) (*entity.UserStats, error)

		SetAdmin(ctx context.Context, id string) error
		UnsetAdmin(ctx context.Context, id string) error
		RandomStreamKey(ctx context.Context, id string) (string, error)
		EditAvatarProfile(ctx context.Context, user *entity.User) (string, error)
		EditProfile(ctx context.Context, profile *entity.FullProfile) error
		ModifyElo(ctx context.Context, elomodify *entity.UserModifyELO) (int, error)
	}
)
