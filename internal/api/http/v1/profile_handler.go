package v1

import (
	"ProfileService/internal/api/http/v1/model"
	"ProfileService/internal/entity"
	"ProfileService/internal/usecases"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
)

type ProfileHandler struct {
	profileUsecase usecases.ProfileService
	logger         zerolog.Logger
}

func NewProfileHandler(profileUsecase usecases.ProfileService, logger zerolog.Logger) *ProfileHandler {
	return &ProfileHandler{profileUsecase: profileUsecase, logger: logger}
}

func (h *ProfileHandler) GetProfile(ctx *gin.Context) {
	h.logger.Info().Msg("HTTP-HANDLER GetProfile вызов функции")

	h.logger.Info().Msg("HTTP-HANDLER GetProfile получаем id")
	id := ctx.Param("id")

	profile, err := h.profileUsecase.GetProfile(ctx, id)
	if err != nil {
		switch {
		case err.Error() == "422":
			h.handleErrorResponse(ctx, http.StatusUnprocessableEntity, "ID пользователя пустой", err)
		case err.Error() == "404":
			h.handleErrorResponse(ctx, http.StatusNotFound, "Пользователь не найден", err)
		case err.Error() == "500":
			h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка на стороне сервера", err)
		}
		return
	}

	ctx.JSON(http.StatusOK, profile)
}

func (h *ProfileHandler) GetStreamKey(ctx *gin.Context) {
	h.logger.Info().Msg("HTTP-HANDLER GetStreamKey вызов функции")

	h.logger.Info().Msg("HTTP-HANDLER GetStreamKey получаем id")
	id := ctx.Param("id")

	streamKey, err := h.profileUsecase.GetStreamKey(ctx, id)
	if err != nil {
		switch {
		case err.Error() == "422":
			h.handleErrorResponse(ctx, http.StatusUnprocessableEntity, "ID пользователя пустой", err)
		case err.Error() == "404":
			h.handleErrorResponse(ctx, http.StatusNotFound, "Пользователь не найден", err)
		case err.Error() == "500":
			h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка на стороне сервера", err)
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"streamKey": streamKey})
}

func (h *ProfileHandler) GetAvatarProfile(ctx *gin.Context) {
	h.logger.Info().Msg("HTTP-HANDLER GetAvatarProfile вызов функции")

	h.logger.Info().Msg("HTTP-HANDLER GetAvatarProfile получаем id")
	id := ctx.Param("id")

	avatar, err := h.profileUsecase.GetAvatarProfile(ctx, id)
	if err != nil {
		switch {
		case err.Error() == "422":
			h.handleErrorResponse(ctx, http.StatusUnprocessableEntity, "ID пользователя пустой", err)
		case err.Error() == "404":
			h.handleErrorResponse(ctx, http.StatusNotFound, "Пользователь не найден", err)
		case err.Error() == "500":
			h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка на стороне сервера", err)
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"picture": avatar})
}

func (h *ProfileHandler) GetProfileStats(ctx *gin.Context) {
	h.logger.Info().Msg("HTTP-HANDLER GetProfileStats вызов функции")

	h.logger.Info().Msg("HTTP-HANDLER GetProfileStats получаем id")
	id := ctx.Param("id")

	userStats, err := h.profileUsecase.GetProfileStats(ctx, id)
	if err != nil {
		switch {
		case err.Error() == "422":
			h.handleErrorResponse(ctx, http.StatusUnprocessableEntity, "ID пользователя пустой", err)
		case err.Error() == "404":
			h.handleErrorResponse(ctx, http.StatusNotFound, "Пользователь не найден", err)
		case err.Error() == "500":
			h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка на стороне сервера", err)
		}
		return
	}

	ctx.JSON(http.StatusOK, userStats)
}

func (h *ProfileHandler) SetAdmin(ctx *gin.Context) {
	h.logger.Info().Msg("HTTP-HANDLER SetAdmin вызов функции")

	h.logger.Info().Msg("HTTP-HANDLER SetAdmin получаем id")
	id := ctx.Param("id")

	err := h.profileUsecase.SetAdmin(ctx, id)
	if err != nil {
		switch {
		case err.Error() == "422":
			h.handleErrorResponse(ctx, http.StatusUnprocessableEntity, "ID пользователя пустой", err)
		case err.Error() == "404":
			h.handleErrorResponse(ctx, http.StatusNotFound, "Пользователь не найден", err)
		case err.Error() == "500":
			h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка на стороне сервера", err)
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Админка выдана"})
}

func (h *ProfileHandler) UnsetAdmin(ctx *gin.Context) {
	h.logger.Info().Msg("HTTP-HANDLER UnsetAdmin вызов функции")

	h.logger.Info().Msg("HTTP-HANDLER UnsetAdmin получаем id")
	id := ctx.Param("id")

	err := h.profileUsecase.UnsetAdmin(ctx, id)
	if err != nil {
		switch {
		case err.Error() == "422":
			h.handleErrorResponse(ctx, http.StatusUnprocessableEntity, "ID пользователя пустой", err)
		case err.Error() == "404":
			h.handleErrorResponse(ctx, http.StatusNotFound, "Пользователь не найден", err)
		case err.Error() == "500":
			h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка на стороне сервера", err)
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Админка снята"})
}

func (h *ProfileHandler) RandomStreamKey(ctx *gin.Context) {
	h.logger.Info().Msg("HTTP-HANDLER RandomStreamKey вызов функции")

	h.logger.Info().Msg("HTTP-HANDLER RandomStreamKey получаем id")
	id := ctx.Param("id")

	streamKey, err := h.profileUsecase.RandomStreamKey(ctx, id)
	if err != nil {
		switch {
		case err.Error() == "422":
			h.handleErrorResponse(ctx, http.StatusUnprocessableEntity, "ID пользователя пустой", err)
		case err.Error() == "404":
			h.handleErrorResponse(ctx, http.StatusNotFound, "Пользователь не найден", err)
		case err.Error() == "500":
			h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка на стороне сервера", err)
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Ключ сгенерирован", "streamKey": streamKey})
}

func (h *ProfileHandler) EditAvatar(ctx *gin.Context) {
	h.logger.Info().Msg("HTTP-HANDLER EditAvatar вызов функции")

	h.logger.Info().Msg("HTTP-HANDLER EditAvatar получаем данные")
	var avatar model.EditAvatarRequest
	if err := ctx.ShouldBindJSON(&avatar); err != nil {
		h.handleErrorResponse(ctx, http.StatusBadRequest, "Запрос отсуствует", err)
	}

	user := &entity.User{
		ID:      avatar.UserId,
		Picture: avatar.Url,
	}

	avaUrl, err := h.profileUsecase.EditAvatarProfile(ctx, user)
	if err != nil {
		switch {
		case err.Error() == "422":
			h.handleErrorResponse(ctx, http.StatusUnprocessableEntity, "ID пользователя пустой", err)
		case err.Error() == "404":
			h.handleErrorResponse(ctx, http.StatusNotFound, "Пользователь не найден", err)
		case err.Error() == "500":
			h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка на стороне сервера", err)
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"avatar": avaUrl})
}

func (h *ProfileHandler) EditProfile(ctx *gin.Context) {
	h.logger.Info().Msg("HTTP-HANDLER EditProfile вызов функции")

	h.logger.Info().Msg("HTTP-HANDLER EditProfile получаем данные")
	var profile model.EditProfileRequest
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		h.handleErrorResponse(ctx, http.StatusBadRequest, "Запрос отсуствует", err)
	}

	user := &entity.FullProfile{
		User:      profile.User,
		UserStats: profile.UserStats,
	}

	err := h.profileUsecase.EditProfile(ctx, user)
	if err != nil {
		switch {
		case err.Error() == "404":
			h.handleErrorResponse(ctx, http.StatusNotFound, "Пользователь не найден", err)
		case err.Error() == "500":
			h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка на стороне сервера", err)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Статистика обнавлена"})
}

func (h *ProfileHandler) ModifyElo(ctx *gin.Context) {
	h.logger.Info().Msg("HTTP-HANDLER ModifyElo вызов функции")

	h.logger.Info().Msg("HTTP-HANDLER ModifyElo получаем данные")
	var modifyElo model.ModifyELO
	if err := ctx.ShouldBindJSON(&modifyElo); err != nil {
		h.handleErrorResponse(ctx, http.StatusBadRequest, "Запрос отсуствует", err)
	}

	userId, err := strconv.Atoi(modifyElo.UserId)
	if err != nil {
		h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка на стороне сервера", err)
	}

	opponentElo, err := strconv.Atoi(modifyElo.OpponentElo)
	if err != nil {
		h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка на стороне сервера", err)
	}

	editElo := &entity.UserModifyELO{
		UserID:      int64(userId),
		OpponentELO: int(opponentElo),
	}

	elo, err := h.profileUsecase.ModifyElo(ctx, editElo)
	if err != nil {
		switch {
		case err.Error() == "404":
			h.handleErrorResponse(ctx, http.StatusNotFound, "Пользователь не найден", err)
		case err.Error() == "500":
			h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка на стороне сервера", err)
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"elo": elo})
}

func (h *ProfileHandler) handleErrorResponse(ctx *gin.Context, statusCode int, message string, err error) {
	h.logger.Error().Err(err).Msgf("HTTP-Handler: %s", message)
	ctx.JSON(statusCode, model.ErrorResponse{
		Status:  statusCode,
		Message: message,
		Error:   err.Error(),
	})
}
