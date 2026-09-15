package services

import (
	"context"

	"github.com/sai-mudike/careerdock.git/internals/customErr"
	"github.com/sai-mudike/careerdock.git/internals/models"
	"github.com/sai-mudike/careerdock.git/internals/repositories"
)

func CreateUser(ctx context.Context, user models.User) error {

	err := repositories.CreateUser(ctx, user)

	if err != nil {
		return err
	}

	return nil

}

func UserLogin(ctx context.Context, user models.User) error {
	userFromDb, err := repositories.GetUser(ctx, user)
	if err != nil {
		return err
	}

	if userFromDb.PassWord != user.PassWord {
		return customErr.ErrInvalidCredentials
	}

	return nil
}
