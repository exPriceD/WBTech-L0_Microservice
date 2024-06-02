package app

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/config"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/middleware"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/repositories/orders"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/stan.go"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"time"
)

type Server struct {
	addr          string
	router        *mux.Router
	natsConn      stan.Conn
	natsClusterID string
	natsClientID  string
	orderRepo     *orders.Repository
	redisClient   *redis.Client
}

func NewServer(cfg *config.Config, db *sqlx.DB, redisClient *redis.Client) *Server {
	s := &Server{
		addr:          cfg.Server.Port,
		router:        mux.NewRouter(),
		natsClusterID: cfg.NATS.ClusterID,
		natsClientID:  cfg.NATS.ClientID,
		redisClient:   redisClient,
	}

	s.configureRouter()
	s.orderRepo = orders.NewOrderRepository(db)
	return s
}

func (s *Server) configureRouter() {
	s.router.Use(middleware.LoggingMiddleware)
	s.router.HandleFunc("/order/{id}", s.getOrderHandler).Methods("GET")
}

func (s *Server) configureNATS() error {
	natsConn, err := stan.Connect(s.natsClusterID, s.natsClientID)
	if err != nil {
		return err
	}
	s.natsConn = natsConn
	return nil
}

func (s *Server) Start() error {
	if err := s.configureNATS(); err != nil {
		return err
	}
	defer func(natsConn stan.Conn) {
		err := natsConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(s.natsConn)

	log.Printf("Starting server on %s", s.addr)
	return http.ListenAndServe(s.addr, s.router)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) getOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["id"]

	ctx := context.Background()
	cachedOrder, err := s.redisClient.Get(ctx, orderId).Result()

	if errors.Is(err, redis.Nil) {
		log.Printf("Cache miss for order ID %s", orderId)
	} else if err != nil {
		log.Printf("Error getting cache for order ID %s", orderId)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(cachedOrder))
		if err != nil {
			log.Printf("Error writing cache for order ID %s", orderId)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			return
		}
	}

	order, err := s.orderRepo.GetByID(orderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	serializedOrder, err := json.Marshal(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = s.redisClient.Set(ctx, orderId, serializedOrder, 3600*time.Second).Err()
	if err != nil {
		log.Printf("Error setting cache for order ID %s. Error: %s", orderId, err)
	}

	if order == nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	data, err := json.Marshal(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
