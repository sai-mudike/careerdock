package services

import (
	"context"
	"errors"
	"net/http"

	"github.com/sai-mudike/careerdock.git/internals/customErr"
	"github.com/sai-mudike/careerdock.git/internals/models"
	"github.com/sai-mudike/careerdock.git/internals/repositories"
	"github.com/sai-mudike/careerdock.git/internals/utils"
)

func CreateUser(ctx context.Context, user models.User) error {

	hashPass, err := utils.GenerateHashPass(user.PassWord)

	if err != nil {
		return err
	}

	user.PassWord = hashPass

	err = repositories.CreateUser(ctx, user)

	if err != nil {
		return err
	}

	return nil

}

func UserLogin(ctx context.Context, user models.User) (string, error) {
	userFromDb, err := repositories.GetUser(ctx, user)
	if err != nil {
		var apperror *customErr.CustomErr
		if errors.As(err, &apperror) && apperror.Code == "USER_NOT_FOUND" {
			return "", customErr.New(
				customErr.CodeInvalidCredentials,
				"invalid email or password",
				http.StatusUnauthorized,
				err)
		}
		return "", err
	}

	isPassValid := utils.ComapreHashAndPass(userFromDb.PassWord, user.PassWord)

	if !isPassValid {
		return "", customErr.New(
			customErr.CodeInvalidCredentials,
			"invalid email or password",
			http.StatusUnauthorized,
			err)
	}

	token, err := utils.GenerateToken(userFromDb.UserName, userFromDb.Id)

	return token, err
}
