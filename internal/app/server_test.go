package app

import (
	"github.com/exPriceD/WBTech-L0_Microservice/internal/cache"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/config"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/db"
	nats_streming "github.com/exPriceD/WBTech-L0_Microservice/internal/nats_streaming"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestGetOrders(t *testing.T) {
	projectRoot, err := filepath.Abs("../..")
	if err != nil {
		t.Fatalf("Failed to determine project root: %v", err)
	}

	// Установите рабочую директорию на projectRoot
	if err := os.Chdir(projectRoot); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	sc, err := nats_streming.InitNATS(cfg)
	if err != nil {
		t.Fatalf("Failed to initialize NATS: %v", err)
	}

	DB, err := db.InitDB(cfg)
	if err != nil {
		t.Fatalf("Failed to initialize DB: %v", err)
	}

	defer func(db *sqlx.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db.DB)

	redisClient := cache.InitRedis(cfg)

	s := NewServer(cfg, DB, redisClient, sc)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/orders", nil)
	s.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
