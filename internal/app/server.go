package app

import (
	"database/sql"
	"encoding/json"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/config"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/middleware"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/repositories/orders"
	"github.com/gorilla/mux"
	"github.com/nats-io/stan.go"
	"log"
	"net/http"
)

type Server struct {
	addr          string
	router        *mux.Router
	natsConn      stan.Conn
	natsClusterID string
	natsClientID  string
	orderRepo     *orders.OrderRepository
}

func NewServer(cfg *config.Config, db *sql.DB) *Server {
	s := &Server{
		addr:          cfg.Server.Port,
		router:        mux.NewRouter(),
		natsClusterID: cfg.NATS.ClusterID,
		natsClientID:  cfg.NATS.ClientID,
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

	order, err := s.orderRepo.GetByID(orderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
