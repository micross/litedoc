package models

import (
	"time"
)

type Token struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Token     string    `json:"token"`
	Revoked   bool      `json:"revoked"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

type RefreshToken struct {
	Id        int       `json:"id"`
	TokenId   int       `json:"token_id"`
	Revoked   bool      `json:"revoked"`
	ExpiresAt time.Time `json:"expires_at"`
}
