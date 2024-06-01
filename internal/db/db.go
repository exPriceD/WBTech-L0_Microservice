package db

import (
	"fmt"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var DB *sqlx.DB

func InitDB(cfg *config.Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DatabaseName)
	log.Println("Connection string:", connStr)

	var err error
	DB, err = sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := DB.Ping(); err != nil {
		return nil, err
	}
	return DB, nil
}
