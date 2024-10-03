package entity

import (
	"github.com/volatiletech/sqlboiler/v4/types"
	"time"
)

type User struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	IP          string    `json:"ip"`
	IsActivated bool      `json:"is_activated"`
	IsAdmin     bool      `json:"is_admin"`
	StreamKey   string    `json:"stream_key"`
	Picture     string    `json:"picture"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserStats struct {
	ID           int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID       int64     `json:"user_id"`
	GamesPlayed  int64     `json:"games_played"`
	RatingElo    int       `json:"rating_elo"`
	HighestElo   int       `json:"highest_elo"`
	FavoriteGame string    `json:"favorite_game"`
	Wins         int64     `json:"wins"`
	Losses       int64     `json:"losses"`
	Draws        int64     `json:"draws"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type UserEconomy struct {
	ID        int64         `json:"id"`
	UserID    int64         `json:"user_id"`
	Balance   types.Decimal `json:"balance"`
	Lightning int64         `json:"lightning"`
	IsFreeze  bool          `json:"is_freeze"`
	UpdatedAt time.Time     `json:"updated_at"`
	CreatedAt time.Time     `json:"created_at"`
}

type FullProfile struct {
	User        *User        `json:"user"`
	UserStats   *UserStats   `json:"user_stats"`
	UserEconomy *UserEconomy `json:"user_economy"`
}

type UserModifyELO struct {
	UserID      int64 `json:"user_id"`
	OpponentELO int   `json:"opponent_elo"`
}
