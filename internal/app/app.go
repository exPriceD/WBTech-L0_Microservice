package app

import (
	"database/sql"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/config"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/db"
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

	defer func(db *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db.DB)

	srv := NewServer(cfg, DB)

	if err := srv.Start(); err != nil {
		return err
	}
	return nil
}
