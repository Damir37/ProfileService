package model

import "ProfileService/internal/entity"

type EditAvatarRequest struct {
	UserId string `json:"user_id"`
	Url    string `json:"url"`
}

type EditProfileRequest struct {
	User      *entity.User      `json:"user"`
	UserStats *entity.UserStats `json:"user_stats"`
}

type ModifyELO struct {
	UserId      string `json:"user_id"`
	OpponentElo string `json:"opponent_elo"`
}
