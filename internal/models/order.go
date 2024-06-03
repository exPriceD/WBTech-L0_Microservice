package models

import (
	"github.com/exPriceD/WBTech-L0_Microservice/internal/entities"
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type Order struct {
	OrderUID          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Items   `json:"items"`
	Entry             string    `json:"entry"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmID              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

func (o Order) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.OrderUID, validation.Required),
		validation.Field(&o.TrackNumber, validation.Required),
		validation.Field(&o.Delivery, validation.Required),
		validation.Field(&o.Payment, validation.Required),
		validation.Field(&o.Items, validation.Required),
		validation.Field(&o.Entry, validation.Required),
		validation.Field(&o.Locale, validation.Required),
		validation.Field(&o.CustomerID, validation.Required),
		validation.Field(&o.DeliveryService, validation.Required),
		validation.Field(&o.Shardkey, validation.Required),
		validation.Field(&o.SmID, validation.Required),
		validation.Field(&o.DateCreated, validation.Required),
		validation.Field(&o.OofShard, validation.Required),
	)
}

func (o Order) ToEntity() entities.OrderWithDetails {
	entityItems := make([]entities.Items, 0, len(o.Items))
	for _, item := range o.Items {
		entityItems = append(entityItems, item.ToEntity())
	}

	return entities.OrderWithDetails{
		OrderUID:          o.OrderUID,
		TrackNumber:       o.TrackNumber,
		Delivery:          o.Delivery.ToEntity(),
		Payment:           o.Payment.ToEntity(),
		Items:             entityItems,
		Entry:             o.Entry,
		Locale:            o.Locale,
		InternalSignature: o.InternalSignature,
		CustomerID:        o.CustomerID,
		DeliveryService:   o.DeliveryService,
		Shardkey:          o.Shardkey,
		SmID:              o.SmID,
		DateCreated:       o.DateCreated,
		OofShard:          o.OofShard,
	}
}
