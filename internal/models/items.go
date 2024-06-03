package models

import (
	"github.com/exPriceD/WBTech-L0_Microservice/internal/entities"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Items struct {
	ChrtID      int    `json:"chrt_id"`
	OrderUID    string `json:"order_uid"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

func (i Items) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.ChrtID, validation.Required),
		validation.Field(&i.TrackNumber, validation.Required),
		validation.Field(&i.Price, validation.Required),
		validation.Field(&i.Rid, validation.Required),
		validation.Field(&i.Name, validation.Required),
		validation.Field(&i.Sale, validation.Required),
		validation.Field(&i.Size, validation.Required),
		validation.Field(&i.TotalPrice, validation.Required),
		validation.Field(&i.NmID, validation.Required),
		validation.Field(&i.Brand, validation.Required),
		validation.Field(&i.Status, validation.Required),
	)
}

func (i Items) ToEntity() entities.Items {
	return entities.Items{
		ChrtID:      i.ChrtID,
		TrackNumber: i.TrackNumber,
		Price:       i.Price,
		Rid:         i.Rid,
		Name:        i.Name,
		Sale:        i.Sale,
		Size:        i.Size,
		TotalPrice:  i.TotalPrice,
		NmID:        i.NmID,
		Brand:       i.Brand,
		Status:      i.Status,
	}
}
