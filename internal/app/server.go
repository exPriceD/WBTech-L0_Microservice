package app

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/config"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/middleware"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/models"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/repositories/items"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/repositories/orders"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"time"
)

// Server struct represents the server with its dependencies.
type Server struct {
	addr        string
	router      *mux.Router
	orderRepo   *orders.Repository
	redisClient *redis.Client
	stanConn    stan.Conn
}

// NewServer creates a new server with the given configuration, database connection, Redis client, and STAN connection.
func NewServer(cfg *config.Config, db *sqlx.DB, redisClient *redis.Client, stanConn stan.Conn) *Server {
	s := &Server{
		addr:        cfg.Server.Port,
		router:      mux.NewRouter(),
		redisClient: redisClient,
		stanConn:    stanConn,
	}

	s.configureRouter()

	itemsRepo := items.NewItemsRepository(db)
	s.orderRepo = orders.NewOrderRepository(db, itemsRepo)

	return s
}

// configureRouter sets up the routes for the server.
func (s *Server) configureRouter() {
	s.router.Use(middleware.LoggingMiddleware)
	s.router.HandleFunc("/order/{id}", s.getOrderHandler).Methods("GET")
}

// subscribeToOrders subscribes to the "orders" topic on the STAN connection.
func (s *Server) subscribeToOrders() {
	sub, err := s.stanConn.Subscribe("orders", func(msg *stan.Msg) {
		var order models.Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Printf("Error unmarshaling order: %v", err)
			err := msg.Ack()
			if err != nil {
				log.Printf("Error msg.Ack(): %v", err)
			}
		}
		if err := models.Validate(order); err != nil {
			log.Printf("Error validating order: %v", err)
		}

		if err := models.Validate(order.Payment); err != nil {
			log.Printf("Error validating payment: %v", err)
		}

		if err := models.Validate(order.Delivery); err != nil {
			log.Printf("Error validating delivery: %v", err)
		}

		for _, item := range order.Items {
			if err := models.Validate(item); err != nil {
				log.Printf("Error validating item: %v", err)
			}
		}

		err = msg.Ack()
		if err != nil {
			log.Printf("Error msg.Ack(): %v", err)
		}

		orderEntity := order.ToEntity()
		err = s.orderRepo.Insert(&orderEntity)
		if err != nil {
			log.Printf("Error inserting order: %v", err)
		}
	}, stan.StartAt(pb.StartPosition_NewOnly))
	if err != nil {
		log.Fatal(err)
	}
	defer func(sub stan.Subscription) {
		err := sub.Unsubscribe()
		if err != nil {
			log.Fatal(err)
		}
	}(sub)

	select {}
}

// Start starts the server and begins listening for requests.
func (s *Server) Start() error {
	go s.subscribeToOrders()
	log.Printf("Starting server on %s", s.addr)
	return http.ListenAndServe(s.addr, s.router)
}

// ServeHTTP delegates the HTTP request to the router.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// getOrderHandler handles the GET request for an order by its ID.
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
