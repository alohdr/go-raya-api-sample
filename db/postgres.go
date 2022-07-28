package db

import (
	"database/sql"
	"fmt"
	"go-rest-api/config"
	"log"

	_ "github.com/lib/pq"
)

const (
	POSTGRES = "postgres"
)

func Postgres() *sql.DB {

	cfg := config.Get()

	conn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUsername,
		cfg.DbPassword,
		cfg.DbName)

	db, err := sql.Open(POSTGRES, conn)
	if err != nil {
		log.Fatal("DB: connection error")
	}

	return db
}
