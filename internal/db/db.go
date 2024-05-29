package db

import (
	"WBTech_L0/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func InitDB(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DatabaseName)
	log.Println("Connection string:", connStr)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := DB.Ping(); err != nil {
		return nil, err
	}
	return DB, nil
}
