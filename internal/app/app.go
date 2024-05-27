package app

import (
	"WBTech_L0/internal/config"
	"WBTech_L0/internal/db"
	"database/sql"
	"net/http"
)

func StartServer() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}
	// Соединение с DB
	if err := db.InitDB(cfg); err != nil {
		return err
	}
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
		}
	}(db.DB)

	srv := NewServer()

	return http.ListenAndServe(cfg.Server.Port, srv)
}
