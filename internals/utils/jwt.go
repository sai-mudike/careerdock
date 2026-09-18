package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sai-mudike/careerdock.git/internals/config"
	"github.com/sai-mudike/careerdock.git/internals/customErr"
)

var cfg = config.MustLoad()

func GenerateToken(username, userID string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userName": username,
		"userID":   userID,
		"exp":      time.Now().Add(2 * time.Hour).Unix(),
	})

	parsedToken, err := token.SignedString([]byte(cfg.JwtSecret))

	if err != nil {
		return "", customErr.ErrTokenGeneration
	}

	return parsedToken, nil
}

func VerifyToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unverified Method used")
		}
		return []byte(cfg.JwtSecret), nil

	})

	if err != nil {
		return "", customErr.ErrUnauthorized
	}

	isValid := parsedToken.Valid

	if !isValid {
		return "", customErr.ErrInvalidToken
	}

	data, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return "", customErr.ErrTokenExpired
	}

	userID := data["userID"].(string)

	if userID == "" {
		return "", customErr.ErrInvalidToken
	}

	return userID, nil

}
