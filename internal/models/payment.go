package models

import (
	"github.com/exPriceD/WBTech-L0_Microservice/internal/entities"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDT    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

func (p Payment) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Transaction, validation.Required),
		validation.Field(&p.Currency, validation.Required),
		validation.Field(&p.Provider, validation.Required),
		validation.Field(&p.Amount, validation.Required),
		validation.Field(&p.PaymentDT, validation.Required),
		validation.Field(&p.Bank, validation.Required),
		validation.Field(&p.DeliveryCost, validation.Required),
		validation.Field(&p.GoodsTotal, validation.Required),
	)
}

func (p Payment) ToEntity() entities.Payment {
	return entities.Payment{
		Transaction:  p.Transaction,
		RequestID:    p.RequestID,
		Currency:     p.Currency,
		Provider:     p.Provider,
		Amount:       p.Amount,
		PaymentDT:    p.PaymentDT,
		Bank:         p.Bank,
		DeliveryCost: p.DeliveryCost,
		GoodsTotal:   p.GoodsTotal,
		CustomFee:    p.CustomFee,
	}
}
