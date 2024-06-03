package nats_streaming

import (
	"github.com/exPriceD/WBTech-L0_Microservice/internal/config"
	"github.com/nats-io/stan.go"
)

func InitNATS(cfg *config.Config) (stan.Conn, error) {
	sc, err := stan.Connect(cfg.NATS.ClusterID, cfg.NATS.ClientID)
	if err != nil {
		return nil, err
	}
	return sc, nil
}
