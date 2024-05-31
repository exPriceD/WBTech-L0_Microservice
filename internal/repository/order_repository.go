package repository

import (
	"database/sql"
	"github.com/exPriceD/WBTech-L0_Microservice/internal/models"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) GetByID(orderUID string) (*models.Order, error) {
	query := `
			SELECT * 
			FROM orders o 
			JOIN delivery d ON d.order_uid = o.order_uid
			JOIN payment p ON o.order_uid = p.transaction
			JOIN items i ON o.order_uid = i.order_uid WHERE o.order_uid = $1;
	`
	rows, err := r.db.Query(query, orderUID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	// Проверка на наличие строк
	if !rows.Next() {
		return nil, sql.ErrNoRows
	}

	var order models.Order
	var items []models.Item
	var unused1, unused2, unused3 interface{}

	for rows.Next() {
		var item models.Item
		if err := rows.Scan(
			&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature,
			&order.CustomerID, &order.DeliveryService, &order.Shardkey, &order.SmID, &order.DateCreated, &order.OofShard,
			&unused1, &unused2, &order.Delivery.Name, &order.Delivery.Phone, &order.Delivery.Zip, &order.Delivery.City, &order.Delivery.Address, &order.Delivery.Region, &order.Delivery.Email,
			&order.Payment.Transaction, &order.Payment.RequestID, &order.Payment.Currency, &order.Payment.Provider, &order.Payment.Amount, &order.Payment.PaymentDT, &order.Payment.Bank, &order.Payment.DeliveryCost, &order.Payment.GoodsTotal, &order.Payment.CustomFee,
			&item.ChrtID, &unused3, &item.TrackNumber, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmID, &item.Brand, &item.Status,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	order.Items = items

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &order, nil
}
