package domain

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
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
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func NewClaims(id string, nameJTI string, TTL time.Duration) Claims {
	return Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   id,
			ID:        nameJTI,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}

func ParseToken(ID string) (*Claims, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secret := []byte(os.Getenv("SECTER_KEY_FOR_JWT"))

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(ID, claims, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secret, nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
