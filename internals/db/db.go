package db

import (
	"fmt"
	"time"

	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(dataBaseURL string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dataBaseURL)

	if err != nil {
		return nil, fmt.Errorf("db.connect: %v", err.Error())
	}

	db.SetMaxOpenConns(15)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.ping: %v", err.Error())
	}

	return db, nil
}
