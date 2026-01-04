package pkg

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type PlayerClaim struct {
	UserId string
	jwt.RegisteredClaims
}

func GenerateToken(userId string, secret string) (string, error) {
	claims := PlayerClaim{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "grpc_game_server",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("Failed to sign token")
	}

	return signedToken, nil
}

func ValidateToken(token string, secret string) (*PlayerClaim, error) {
	parsedToken, err := jwt.ParseWithClaims(
		token,
		&PlayerClaim{},
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("unexpected signing method")
			}

			return secret, nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("Authentication failed, err: %v", err)
	}

	claims, ok := parsedToken.Claims.(*PlayerClaim)
	if !ok || parsedToken.Valid {
		return nil, errors.New("Invalid token")
	}

	return claims, nil
}
