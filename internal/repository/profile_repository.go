package repository

import (
	"ProfileService/internal/app/postgre"
	"ProfileService/internal/app/redis"
	"ProfileService/internal/entity"
	"ProfileService/internal/models"
	"ProfileService/internal/pkg/elo"
	"ProfileService/internal/pkg/streamkey"
	"context"
	"database/sql"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/rs/zerolog"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"strconv"
)

type profileRepository struct {
	db    *postgre.PostgreSQLService
	redis *redis.RedisService

	logger zerolog.Logger
}

func NewProfileRepository(db *postgre.PostgreSQLService, redis *redis.RedisService, logger zerolog.Logger) ProfileRepository {
	return &profileRepository{db: db, redis: redis, logger: logger}
}

func (pr *profileRepository) GetProfile(ctx context.Context, id string) (*entity.FullProfile, error) {
	dbReplica := pr.db.GetReplica()

	pr.logger.Info().Msg("GetProfile: отсылаем запрос в бд")

	user, err := models.Users(qm.Where("id = ?", id)).One(ctx, dbReplica)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.logger.Error().Err(err).Msg("GetProfile: Ошибка, пользователь не найден")
			return nil, fmt.Errorf("404")
		}
		pr.logger.Error().Err(err).Msg("GetProfile: Ошибка с базой данных")
		return nil, fmt.Errorf("500")
	}

	userStats, err := models.UsersStats(qm.Where("user_id = ?", user.ID)).One(ctx, dbReplica)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.logger.Error().Err(err).Msg("GetProfile: Ошибка, статистика пользователя не найдена")
			return nil, fmt.Errorf("404")
		}
		pr.logger.Error().Err(err).Msg("GetProfile: Ошибка с базой данных")
		return nil, fmt.Errorf("500")
	}

	userEconomy, err := models.UsersEconomies(qm.Where("user_id = ?", user.ID)).One(ctx, dbReplica)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.logger.Error().Err(err).Msg("GetProfile: Ошибка, баланс пользователя не найден")
			return nil, fmt.Errorf("404")
		}
		pr.logger.Error().Err(err).Msg("GetProfile: Ошибка с базой данных")
		return nil, fmt.Errorf("500")
	}

	userFull := &entity.FullProfile{
		User: &entity.User{
			ID:          strconv.Itoa(int(user.ID)),
			Username:    user.Username,
			Email:       user.Email,
			IP:          user.IP,
			IsActivated: user.IsActivated,
			IsAdmin:     user.IsAdmin,
			StreamKey:   user.StreamKey,
			Picture:     user.Picture,
			UpdatedAt:   user.UpdatedAt,
			CreatedAt:   user.CreatedAt,
		},
		UserStats: &entity.UserStats{
			ID:           userStats.ID,
			UserID:       userStats.UserID,
			GamesPlayed:  userStats.GamesPlayed,
			RatingElo:    userStats.RatingElo,
			HighestElo:   userStats.HighestElo,
			FavoriteGame: userStats.FavoriteGame,
			Wins:         userStats.Wins,
			Losses:       userStats.Losses,
			Draws:        userStats.Draws,
			UpdatedAt:    userStats.UpdatedAt,
			CreatedAt:    userStats.CreatedAt,
		},
		UserEconomy: &entity.UserEconomy{
			ID:        userEconomy.ID,
			UserID:    userEconomy.UserID,
			Balance:   userEconomy.Balance,
			Lightning: userEconomy.Lightning,
			IsFreeze:  userEconomy.IsFreeze,
			UpdatedAt: userEconomy.UpdatedAt,
			CreatedAt: userEconomy.CreatedAt,
		},
	}

	return userFull, nil
}

func (pr *profileRepository) GetStreamKey(ctx context.Context, id string) (string, error) {
	dbReplica := pr.db.GetReplica()

	pr.logger.Info().Msg("GetStreamKey: отсылаем запрос в бд")

	user, err := models.Users(qm.Where("id = ?", id)).One(ctx, dbReplica)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.logger.Error().Err(err).Msg("GetStreamKey: Ошибка, пользователь не найден")
			return "", fmt.Errorf("404")
		}
		pr.logger.Error().Err(err).Msg("GetStreamKey: Ошибка с базой данных")
		return "", fmt.Errorf("500")
	}

	return user.StreamKey, nil
}

func (pr *profileRepository) GetAvatarProfile(ctx context.Context, id string) (string, error) {
	dbReplica := pr.db.GetReplica()

	pr.logger.Info().Msg("GetAvatarProfile: отсылаем запрос в бд")

	user, err := models.Users(qm.Where("id = ?", id)).One(ctx, dbReplica)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.logger.Error().Err(err).Msg("GetAvatarProfile: Ошибка, пользователь не найден")
			return "", fmt.Errorf("404")
		}
		pr.logger.Error().Err(err).Msg("GetAvatarProfile: Ошибка с базой данных")
		return "", fmt.Errorf("500")
	}

	return user.Picture, nil
}

func (pr *profileRepository) GetProfileStats(ctx context.Context, id string) (*entity.UserStats, error) {
	dbReplica := pr.db.GetReplica()

	pr.logger.Info().Msg("GetProfileStats: отсылаем запрос в бд")

	user, err := models.Users(qm.Where("id = ?", id)).One(ctx, dbReplica)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.logger.Error().Err(err).Msg("GetProfileStats: Ошибка, пользователь не найден")
			return nil, fmt.Errorf("404")
		}
		pr.logger.Error().Err(err).Msg("GetProfileStats: Ошибка с базой данных")
		return nil, fmt.Errorf("500")
	}

	userStats, err := models.UsersStats(qm.Where("user_id = ?", user.ID)).One(ctx, dbReplica)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.logger.Error().Err(err).Msg("GetProfileStats: Ошибка, статистика пользователя не найден")
			return nil, fmt.Errorf("404")
		}
		pr.logger.Error().Err(err).Msg("GetProfileStats: Ошибка с базой данных")
		return nil, fmt.Errorf("500")
	}

	userStatsResult := &entity.UserStats{
		ID:           userStats.ID,
		UserID:       userStats.UserID,
		GamesPlayed:  userStats.GamesPlayed,
		RatingElo:    userStats.RatingElo,
		HighestElo:   userStats.HighestElo,
		FavoriteGame: userStats.FavoriteGame,
		Wins:         userStats.Wins,
		Losses:       userStats.Losses,
		Draws:        userStats.Draws,
		UpdatedAt:    userStats.UpdatedAt,
		CreatedAt:    userStats.CreatedAt,
	}

	return userStatsResult, nil
}

func (pr *profileRepository) SetAdmin(ctx context.Context, id string) error {
	dbMaster := pr.db.GetMaster()

	pr.logger.Info().Msg("SetAdmin: отсылаем запрос в бд")

	user, err := models.Users(qm.Where("id = ?", id)).One(ctx, dbMaster)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.logger.Error().Err(err).Msg("SetAdmin: Ошибка, пользователь не найден")
			return fmt.Errorf("404")
		}
		pr.logger.Error().Err(err).Msg("SetAdmin: Ошибка с базой данных")
		return fmt.Errorf("500")
	}

	userUpd := &models.User{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Password:    user.Password,
		IP:          user.IP,
		IsActivated: user.IsActivated,
		IsAdmin:     true,
		StreamKey:   user.StreamKey,
		Picture:     user.Picture,
	}

	_, err = userUpd.Update(ctx, dbMaster, boil.Infer())

	if err != nil {
		pr.logger.Error().Err(err).Msg("SetAdmin: Ошибка с базой данных")
		return fmt.Errorf("500")
	}

	return nil
}

func (pr *profileRepository) UnsetAdmin(ctx context.Context, id string) error {
	dbMaster := pr.db.GetMaster()

	pr.logger.Info().Msg("UnsetAdmin: отсылаем запрос в бд")

	user, err := models.Users(qm.Where("id = ?", id)).One(ctx, dbMaster)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.logger.Error().Err(err).Msg("UnsetAdmin: Ошибка, пользователь не найден")
			return fmt.Errorf("404")
		}
		pr.logger.Error().Err(err).Msg("UnsetAdmin: Ошибка с базой данных")
		return fmt.Errorf("500")
	}

	userUpd := &models.User{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Password:    user.Password,
		IP:          user.IP,
		IsActivated: user.IsActivated,
		IsAdmin:     false,
		StreamKey:   user.StreamKey,
		Picture:     user.Picture,
	}

	_, err = userUpd.Update(ctx, dbMaster, boil.Infer())

	if err != nil {
		pr.logger.Error().Err(err).Msg("UnsetAdmin: Ошибка с базой данных")
		return fmt.Errorf("500")
	}

	return nil
}

func (pr *profileRepository) RandomStreamKey(ctx context.Context, id string) (string, error) {
	dbMaster := pr.db.GetMaster()

	pr.logger.Info().Msg("RandomStreamKey: отсылаем запрос в бд")

	user, err := models.Users(qm.Where("id = ?", id)).One(ctx, dbMaster)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.logger.Error().Err(err).Msg("RandomStreamKey: Ошибка, пользователь не найден")
			return "", fmt.Errorf("404")
		}
		pr.logger.Error().Err(err).Msg("RandomStreamKey: Ошибка с базой данных")
		return "", fmt.Errorf("500")
	}

	pr.logger.Info().Msg("RandomStreamKey: генерация стрим ключа")

	streamKey, err := streamkey.GenerateStreamKey()
	if err != nil {
		pr.logger.Error().Err(err).Msg("RandomStreamKey: Ошибка генерации ключа")
		return "", fmt.Errorf("500")
	}

	userUpd := &models.User{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Password:    user.Password,
		IP:          user.IP,
		IsActivated: user.IsActivated,
		IsAdmin:     user.IsAdmin,
		StreamKey:   streamKey,
		Picture:     user.Picture,
	}

	_, err = userUpd.Update(ctx, dbMaster, boil.Infer())

	if err != nil {
		pr.logger.Error().Err(err).Msg("RandomStreamKey: Ошибка с базой данных")
		return "", fmt.Errorf("500")
	}

	return streamKey, nil
}

func (pr *profileRepository) EditAvatarProfile(ctx context.Context, user *entity.User) (string, error) {
	dbMaster := pr.db.GetMaster()

	pr.logger.Info().Msg("EditAvatarProfile: отсылаем запрос в бд")

	userFind, err := models.Users(qm.Where("id = ?", user.ID)).One(ctx, dbMaster)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.logger.Error().Err(err).Msg("EditAvatarProfile: Ошибка, пользователь не найден")
			return "", fmt.Errorf("404")
		}
		pr.logger.Error().Err(err).Msg("EditAvatarProfile: Ошибка с базой данных")
		return "", fmt.Errorf("500")
	}

	pr.logger.Info().Msg("EditAvatarProfile: меняем ссылку на аватарку")

	userUpd := &models.User{
		ID:          userFind.ID,
		Username:    userFind.Username,
		Email:       userFind.Email,
		Password:    userFind.Password,
		IP:          userFind.IP,
		IsActivated: userFind.IsActivated,
		IsAdmin:     userFind.IsAdmin,
		StreamKey:   userFind.StreamKey,
		Picture:     user.Picture,
	}

	_, err = userUpd.Update(ctx, dbMaster, boil.Infer())

	if err != nil {
		pr.logger.Error().Err(err).Msg("EditAvatarProfile: Ошибка с базой данных")
		return "", fmt.Errorf("500")
	}

	return user.Picture, nil
}

func (pr *profileRepository) EditProfile(ctx context.Context, profile *entity.FullProfile) error {
	dbMaster := pr.db.GetMaster()

	pr.logger.Info().Msg("EditProfile: отсылаем запрос в бд")

	userFind, err := models.Users(qm.Where("id = ?", profile.User.ID)).One(ctx, dbMaster)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.logger.Error().Err(err).Msg("EditProfile: Ошибка, пользователь не найден")
			return fmt.Errorf("404")
		}
		pr.logger.Error().Err(err).Msg("EditProfile: Ошибка с базой данных")
		return fmt.Errorf("500")
	}

	pr.logger.Info().Msg("EditProfile: меняем данные")

	userStats := &models.UsersStat{
		ID:           userFind.ID,
		UserID:       userFind.ID,
		GamesPlayed:  profile.UserStats.GamesPlayed,
		RatingElo:    profile.UserStats.RatingElo,
		HighestElo:   profile.UserStats.HighestElo,
		FavoriteGame: profile.UserStats.FavoriteGame,
		Wins:         profile.UserStats.Wins,
		Losses:       profile.UserStats.Losses,
		Draws:        profile.UserStats.Draws,
	}

	_, err = userStats.Update(ctx, dbMaster, boil.Infer())

	if err != nil {
		pr.logger.Error().Err(err).Msg("EditProfile: Ошибка с базой данных")
		return fmt.Errorf("500")
	}

	return nil
}

func (pr *profileRepository) ModifyElo(ctx context.Context, elomodify *entity.UserModifyELO) (int, error) {
	dbMaster := pr.db.GetMaster()

	pr.logger.Info().Msg("ModifyElo: отсылаем запрос в бд")

	userStatsFind, err := models.UsersStats(qm.Where("user_id = ?", elomodify.UserID)).One(ctx, dbMaster)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.logger.Error().Err(err).Msg("ModifyElo: Ошибка, статистика пользователя не найдена")
			return 0, fmt.Errorf("404")
		}
		pr.logger.Error().Err(err).Msg("ModifyElo: Ошибка с базой данных")
		return 0, fmt.Errorf("500")
	}

	pr.logger.Info().Msg("ModifyElo: обновляем ELO")
	resultElo := elo.CalculatorELO(int(userStatsFind.GamesPlayed), int(userStatsFind.Wins), int(userStatsFind.Draws),
		int(userStatsFind.Losses), userStatsFind.RatingElo, elomodify.OpponentELO)

	highestElo := userStatsFind.HighestElo
	if resultElo > highestElo {
		highestElo = resultElo
	}

	userStats := &models.UsersStat{
		ID:           userStatsFind.ID,
		UserID:       userStatsFind.ID,
		GamesPlayed:  userStatsFind.GamesPlayed,
		RatingElo:    resultElo,
		HighestElo:   highestElo,
		FavoriteGame: userStatsFind.FavoriteGame,
		Wins:         userStatsFind.Wins,
		Losses:       userStatsFind.Losses,
		Draws:        userStatsFind.Draws,
		CreatedAt:    userStatsFind.CreatedAt,
	}

	_, err = userStats.Update(ctx, dbMaster, boil.Infer())

	if err != nil {
		pr.logger.Error().Err(err).Msg("ModifyElo: Ошибка с базой данных")
		return 0, fmt.Errorf("500")
	}

	return resultElo, nil
}
