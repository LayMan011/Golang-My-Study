package domain

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	accessTTL  = 15 * time.Minute
	refreshTTL = 7 * 24 * time.Hour
	ctx        = context.Background()
	rdb        = redis.NewClient(&redis.Options{Addr: "localhost:6379"})
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewTokenPair(accessToken string, refreshToken string) *TokenPair {
	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

type Claims struct {
	Login string `json:"login"`
	jwt.RegisteredClaims
}

func NewClaims(login string, nameJTI string, TTL time.Duration) Claims {
	return Claims{
		Login: login,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   login,
			ID:        nameJTI,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}

func GenerateTokenPair(login string) (*TokenPair, error) {
	accessJTI, err := randomID()
	if err != nil {
		return nil, err
	}
	refreshJTI, err := randomID()
	if err != nil {
		return nil, err
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secret := []byte(os.Getenv("SECTER_KEY_FOR_JWT"))

	accessClaims := NewClaims(login, accessJTI, accessTTL)
	refreshClaims := NewClaims(login, refreshJTI, refreshTTL)

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(secret)
	if err != nil {
		return nil, err
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(secret)
	if err != nil {
		return nil, err
	}

	if err := rdb.Set(ctx, "refresh:"+refreshJTI, login, refreshTTL).Err(); err != nil {
		return nil, err
	}

	return NewTokenPair(accessToken, refreshToken), nil
}

func ParseToken(tokenString string) (*Claims, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secret := []byte(os.Getenv("SECTER_KEY_FOR_JWT"))

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
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

func randomID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
