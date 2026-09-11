package main

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/sai-mudike/careerdock.git/internals/config"
)

func main() {
	cfg := config.MustLoad()
	m, err := migrate.New(
		"file://migrations",
		cfg.DataBase_URL,
	)
	if err != nil {
		panic(fmt.Errorf("migrate.new: %v", err.Error()))
	}

	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "up":
			if err := m.Up(); err != nil {
				panic(err)
			}
			fmt.Println("migrated up")
		case "down":
			if err := m.Down(); err != nil {
				panic(err)
			}
			fmt.Println("migrated down")
		default:
			fmt.Println("only migrate <up|down> is allowed")
		}
	}
}
