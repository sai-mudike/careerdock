package db

import (
	"fmt"
	"time"

	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Connect(dataBaseURL string) error {
	var err error
	DB, err = sql.Open("pgx", dataBaseURL)

	if err != nil {
		return fmt.Errorf("DB.connect: %v", err.Error())
	}

	DB.SetMaxOpenConns(15)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("DB.ping: %v", err.Error())
	}

	return nil
}
