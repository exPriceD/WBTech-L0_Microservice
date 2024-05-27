package app

import (
	"WBTech_L0/internal/config"
	"net/http"
)

func StartServer() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}
	// Соединение с DB

	srv := NewServer()

	return http.ListenAndServe(cfg.Server.Port, srv)
}
