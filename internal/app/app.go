package app

import (
	"github.com/exPriceD/WBTech-L0_Microservice/internal/cache"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/config"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/db"
	"github.com/jmoiron/sqlx"
	"log"
)

func StartServer() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}
	// Соединение с DB
	DB, err := db.InitDB(cfg)
	if err != nil {
		return err
	}

	redisClient := cache.InitRedis(cfg)

	defer func(db *sqlx.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db.DB)

	srv := NewServer(cfg, DB, redisClient)

	if err := srv.Start(); err != nil {
		return err
	}
	return nil
}
