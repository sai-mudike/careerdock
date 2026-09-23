package utils

import (
	"errors"
	"fmt"
	"net/http"
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
		return "", fmt.Errorf("token Generation: %w", err)
	}

	return parsedToken, nil
}

func VerifyToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, customErr.New(
				customErr.CodeInvalidToken,
				"invalid authentication token",
				http.StatusUnauthorized,
				nil,
			)
		}
		return []byte(cfg.JwtSecret), nil

	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return "", customErr.New(
				customErr.CodeExpiredToken,
				"authentication token has expired",
				http.StatusUnauthorized,
				err,
			)
		}

		return "", customErr.New(
			customErr.CodeInvalidToken,
			"invalid authentication token",
			http.StatusUnauthorized,
			err,
		)
	}

	isValid := parsedToken.Valid

	if !isValid {
		return "", customErr.New(
			customErr.CodeInvalidToken,
			"invalid authentication token",
			http.StatusUnauthorized,
			err,
		)
	}

	data, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return "", customErr.New(
			customErr.CodeInvalidToken,
			"invalid authentication token",
			http.StatusUnauthorized,
			err,
		)
	}

	userID, ok := data["userID"].(string)

	if !ok || userID == "" {
		return "", customErr.New(
			customErr.CodeInvalidToken,
			"invalid authentication token",
			http.StatusUnauthorized,
			err,
		)
	}

	return userID, nil

}
