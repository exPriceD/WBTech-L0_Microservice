package orders

import (
	"github.com/exPriceD/WBTech-L0_Microservice/internal/entities"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/repositories/items"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db        *sqlx.DB
	itemsRepo *items.Repository
}

func NewOrderRepository(db *sqlx.DB, itemsRepo *items.Repository) *Repository {
	return &Repository{
		db:        db,
		itemsRepo: itemsRepo,
	}
}

func (r *Repository) GetByID(orderUID string) (*entities.OrderWithDetails, error) {
	order, err := r.getOrderWithPaymentAndDelivery(orderUID)
	if err != nil {
		return nil, err
	}

	receivedItems, err := r.itemsRepo.GetItems(orderUID)
	if err != nil {
		return nil, err
	}

	for _, item := range receivedItems {
		order.Items = append(order.Items, item)
	}

	return order, nil
}

func (r *Repository) getOrderWithPaymentAndDelivery(orderUID string) (*entities.OrderWithDetails, error) {
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

	return &order, nil
}

func (r *Repository) Insert(order *entities.OrderWithDetails) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	orderQuery := `INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`
	if _, err := tx.Exec(orderQuery, order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService, order.Shardkey, order.SmID, order.DateCreated, order.OofShard); err != nil {
		_ = tx.Rollback()
		return err
	}

	deliveryQuery := `INSERT INTO delivery (order_uid, name, phone, zip, city, address, region, email) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`
	if _, err := tx.Exec(deliveryQuery, order.OrderUID, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email); err != nil {
		_ = tx.Rollback()
		return err
	}

	paymentQuery := `INSERT INTO payment (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`
	if _, err := tx.Exec(paymentQuery, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDT, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee); err != nil {
		_ = tx.Rollback()
		return err
	}

	for _, item := range order.Items {
		itemQuery := `INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`
		if _, err := tx.Exec(itemQuery, item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status); err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	_ = tx.Commit()
	return nil
}

func (r *Repository) GetAll() ([]entities.OrderWithoutDetails, error) {
	var orders []entities.OrderWithoutDetails
	query := `SELECT order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard FROM orders;`
	if err := r.db.Select(&orders, query); err != nil {
		return nil, err
	}
	return orders, nil
}
