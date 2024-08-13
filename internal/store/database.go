package store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/anvidev/nit/config"
	_ "github.com/lib/pq"
)

func OpenDB(cfg *config.Config) *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error opening database", err)
	}
	return db
}
