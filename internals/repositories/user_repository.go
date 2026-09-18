package repositories

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sai-mudike/careerdock.git/internals/customErr"
	"github.com/sai-mudike/careerdock.git/internals/db"
	"github.com/sai-mudike/careerdock.git/internals/models"
)

func CreateUser(ctx context.Context, user models.User) error {

	query := `
	INSERT INTO users(username,password)
	VALUES ($1,$2);
	`

	_, err := db.DB.ExecContext(ctx, query, user.UserName, user.PassWord)
	var pgErr *pgconn.PgError
	if err != nil {
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return customErr.ErrEmailAlreadyExists
			}
		}
		return customErr.ErrUserNotCreated
	}
	return nil

}

func GetUser(ctx context.Context, user models.User) (*models.User, error) {

	query := `
	SELECT id,password FROM users WHERE username=$1;
	`

	row := db.DB.QueryRowContext(ctx, query, user.UserName)

	var userFromDB models.User

	err := row.Scan(&userFromDB.Id, &userFromDB.PassWord)

	if err != nil {
		return nil, customErr.ErrUserNotFound
	}

	userFromDB.UserName = user.UserName

	return &userFromDB, nil

}

// func UserExists(ctx context.Context, userID string) (bool, error) {
// 	query := `
// SELECT EXISTS(
// SELECT 1 FROM users WHERE id=$1
// );
// `

// 	var ifExists bool

// 	row := db.DB.QueryRowContext(ctx, query, userID)

// 	err := row.Scan(&ifExists)

// 	if err != nil {
// 		return false, customErr.ErrUnauthorized
// 	}

// 	return ifExists, nil

// }
