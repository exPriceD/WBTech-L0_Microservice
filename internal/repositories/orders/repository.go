package orders

import (
	"github.com/exPriceD/WBTech-L0_Microservice/internal/entities"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetByID(orderUID string) (*entities.OrderWithDetails, error) {
	var order entities.OrderWithDetails
	orderQuery := `
		SELECT
		    o.order_uid, o.track_number, o.entry, o.locale, o.internal_signature, o.customer_id, o.delivery_service, o.shardkey, o.sm_id, o.date_created, o.oof_shard,
			d.name AS delivery_name, d.phone AS delivery_phone, d.zip AS delivery_zip, d.city AS delivery_city,
			d.address AS delivery_address, d.region AS delivery_region, d.email AS delivery_email,
			p.transaction AS payment_transaction, p.request_id AS payment_request_id, p.currency AS payment_currency, p.provider AS payment_provider, 
			p.amount AS payment_amount, p.payment_dt AS payment_dt, p.bank AS payment_bank,
			p.delivery_cost AS payment_delivery_cost, p.goods_total AS payment_goods_total, p.custom_fee AS payment_custom_fee
		FROM orders o
				 JOIN delivery d ON o.order_uid = d.order_uid
				 JOIN payment p ON o.order_uid = p.transaction
		WHERE o.order_uid = $1;
`
	if err := r.db.QueryRow(orderQuery, orderUID).Scan(
		&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature,
		&order.CustomerID, &order.DeliveryService, &order.Shardkey, &order.SmID, &order.DateCreated, &order.OofShard,
		&order.Delivery.Name, &order.Delivery.Phone, &order.Delivery.Zip, &order.Delivery.City, &order.Delivery.Address, &order.Delivery.Region, &order.Delivery.Email,
		&order.Payment.Transaction, &order.Payment.RequestID, &order.Payment.Currency, &order.Payment.Provider, &order.Payment.Amount, &order.Payment.PaymentDT, &order.Payment.Bank, &order.Payment.DeliveryCost, &order.Payment.GoodsTotal, &order.Payment.CustomFee,
	); err != nil {
		return nil, err
	}

	items := make([]entities.Items, 0)
	itemsQuery := `SELECT chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status FROM items WHERE order_uid = $1;`
	err := r.db.Select(&items, itemsQuery, orderUID)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		order.Items = append(order.Items, item)
	}

	return &order, nil
}
