package models

import (
	"github.com/exPriceD/WBTech-L0_Microservice/internal/entities"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Delivery struct {
	DeliveryID uint64 `json:"delivery_id"`
	OrderUID   string `json:"order_uid"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Zip        string `json:"zip"`
	City       string `json:"city"`
	Address    string `json:"address"`
	Region     string `json:"region"`
	Email      string `json:"email"`
}

func (d Delivery) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Name, validation.Required),
		validation.Field(&d.Phone, validation.Required),
		validation.Field(&d.Zip, validation.Required),
		validation.Field(&d.City, validation.Required),
		validation.Field(&d.Address, validation.Required),
		validation.Field(&d.Region, validation.Required),
		validation.Field(&d.Email, validation.Required, is.Email),
	)
}

func (d Delivery) ToEntity() entities.Delivery {
	return entities.Delivery{
		Name:    d.Name,
		Phone:   d.Phone,
		Zip:     d.Zip,
		City:    d.City,
		Address: d.Address,
		Region:  d.Region,
		Email:   d.Email,
	}
}
