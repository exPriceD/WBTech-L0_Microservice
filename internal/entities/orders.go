package entities

import "time"

type OrderWithDetails struct {
	OrderUID            string    `db:"order_uid" json:"order_uid"`
	TrackNumber         string    `db:"track_number" json:"track_number"`
	DeliveryName        string    `db:"delivery_name" json:"delivery_name"`
	DeliveryPhone       string    `db:"delivery_phone" json:"delivery_phone"`
	DeliveryZip         string    `db:"delivery_zip" json:"delivery_zip"`
	DeliveryCity        string    `db:"delivery_city" json:"delivery_city"`
	DeliveryAddress     string    `db:"delivery_address" json:"delivery_address"`
	DeliveryRegion      string    `db:"delivery_region" json:"delivery_region"`
	DeliveryEmail       string    `db:"delivery_email" json:"delivery_email"`
	PaymentTransaction  string    `db:"payment_transaction" json:"payment_transaction"`
	PaymentRequestID    string    `db:"payment_request_id" json:"payment_request_id"`
	PaymentCurrency     string    `db:"payment_currency" json:"payment_currency"`
	PaymentProvider     string    `db:"payment_provider" json:"payment_provider"`
	PaymentAmount       int       `db:"payment_amount" json:"payment_amount"`
	PaymentDT           int       `db:"payment_dt" json:"payment_dt"`
	PaymentBank         string    `db:"payment_bank" json:"payment_bank"`
	PaymentDeliveryCost int       `db:"payment_delivery_cost" json:"payment_delivery_cost"`
	PaymentGoodsTotal   int       `db:"payment_goods_total" json:"payment_goods_total"`
	PaymentCustomFee    int       `db:"payment_custom_fee" json:"payment_custom_fee"`
	Items               []Items   `json:"items"`
	Entry               string    `db:"entry" json:"entry"`
	Locale              string    `db:"locale" json:"locale"`
	InternalSignature   string    `db:"internal_signature" json:"internal_signature"`
	CustomerID          string    `db:"customer_id" json:"customer_id"`
	DeliveryService     string    `db:"delivery_service" json:"delivery_service"`
	Shardkey            string    `db:"shardkey" json:"shardkey"`
	SmID                int       `db:"sm_id" json:"sm_id"`
	DateCreated         time.Time `db:"date_created" json:"date_created"`
	OofShard            string    `db:"oof_shard" json:"oof_shard"`
}

type Items struct {
	ChrtID      int    `db:"chrt_id" json:"chrt_id"`
	TrackNumber string `db:"track_number" json:"item_track_number"`
	Price       int    `db:"price" json:"item_price"`
	Rid         string `db:"rid" json:"item_rid"`
	Name        string `db:"name" json:"item_name"`
	Sale        int    `db:"sale" json:"item_sale"`
	Size        string `db:"size" json:"item_size"`
	TotalPrice  int    `db:"total_price" json:"item_total_price"`
	NmID        int    `db:"nm_id" json:"item_nm_id"`
	Brand       string `db:"brand" json:"item_brand"`
	Status      int    `db:"status" json:"item_status"`
}
