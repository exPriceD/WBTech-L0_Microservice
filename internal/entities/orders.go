package entities

import "time"

type OrderWithDetails struct {
	OrderUID          string    `db:"order_uid" json:"order_uid"`
	TrackNumber       string    `db:"track_number" json:"track_number"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Items   `json:"items"`
	Entry             string    `db:"entry" json:"entry"`
	Locale            string    `db:"locale" json:"locale"`
	InternalSignature string    `db:"internal_signature" json:"internal_signature"`
	CustomerID        string    `db:"customer_id" json:"customer_id"`
	DeliveryService   string    `db:"delivery_service" json:"delivery_service"`
	Shardkey          string    `db:"shardkey" json:"shardkey"`
	SmID              int       `db:"sm_id" json:"sm_id"`
	DateCreated       time.Time `db:"date_created" json:"date_created"`
	OofShard          string    `db:"oof_shard" json:"oof_shard"`
}

type Delivery struct {
	Name    string `db:"delivery_name" json:"name"`
	Phone   string `db:"delivery_phone" json:"phone"`
	Zip     string `db:"delivery_zip" json:"zip"`
	City    string `db:"delivery_city" json:"city"`
	Address string `db:"delivery_address" json:"address"`
	Region  string `db:"delivery_region" json:"region"`
	Email   string `db:"delivery_email" json:"email"`
}

type Payment struct {
	Transaction  string `db:"payment_transaction" json:"transaction"`
	RequestID    string `db:"payment_request_id" json:"request_id"`
	Currency     string `db:"payment_currency" json:"currency"`
	Provider     string `db:"payment_provider" json:"provider"`
	Amount       int    `db:"payment_amount" json:"amount"`
	PaymentDT    int    `db:"payment_dt" json:"payment_dt"`
	Bank         string `db:"payment_bank" json:"bank"`
	DeliveryCost int    `db:"payment_delivery_cost" json:"delivery_cost"`
	GoodsTotal   int    `db:"payment_goods_total" json:"goods_total"`
	CustomFee    int    `db:"payment_custom_fee" json:"custom_fee"`
}

type Items struct {
	ChrtID      int    `db:"chrt_id" json:"chrt_id"`
	TrackNumber string `db:"track_number" json:"track_number"`
	Price       int    `db:"price" json:"price"`
	Rid         string `db:"rid" json:"rid"`
	Name        string `db:"name" json:"name"`
	Sale        int    `db:"sale" json:"sale"`
	Size        string `db:"size" json:"size"`
	TotalPrice  int    `db:"total_price" json:"total_price"`
	NmID        int    `db:"nm_id" json:"nm_id"`
	Brand       string `db:"brand" json:"brand"`
	Status      int    `db:"status" json:"status"`
}
