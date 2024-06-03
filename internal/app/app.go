package app

import (
	"github.com/exPriceD/WBTech-L0_Microservice/internal/cache"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/config"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/db"
	nats_streming "github.com/exPriceD/WBTech-L0_Microservice/internal/nats_streaming"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/stan.go"
	"log"
)

func StartServer() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

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

	sc, err := nats_streming.InitNATS(cfg)
	if err != nil {
		return err
	}

	defer func(sc stan.Conn) {
		err := sc.Close()
		if err != nil {
		}
	}(sc)

	srv := NewServer(cfg, DB, redisClient, sc)

	if err := srv.Start(); err != nil {
		return err
	}
	return nil
}
