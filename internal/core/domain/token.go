package domain

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenPair struct {
	AccessToken  string `json:"access_token" redis:"access_token"`
	RefreshToken string `json:"refresh_token" redis:"refresh_token"`
}

func NewTokenPair(accessToken string, refreshToken string) *TokenPair {
	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

type Claims struct {
	jwt.RegisteredClaims
	UserID int `json:"user_id,omitempty"`
}

func NewClaims(login string, nameJTI string, userID int, TTL time.Duration) Claims {
	return Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   login,
			ID:        nameJTI,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID: userID,
	}
}

func ParseToken(tokenString string) (*Claims, error) {
	secret := []byte(os.Getenv("SECTER_KEY_FOR_JWT"))
	if len(secret) == 0 {
		return nil, fmt.Errorf("SECTER_KEY_FOR_JWT is not set")
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
