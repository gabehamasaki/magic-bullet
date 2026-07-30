package helper

import (
	"fmt"
	"magic-bullet/backend/internal/config"

	"github.com/golang-jwt/jwt/v5"
)


func CreateJWT(config *config.Config, claims jwt.MapClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) 

	tokenString, err := token.SignedString([]byte(config.APP_SECRET));
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func ParseToken(config *config.Config, tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.APP_SECRET), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims) 
	if !ok || !token.Valid{
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
