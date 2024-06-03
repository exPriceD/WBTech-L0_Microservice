package items

import (
	"github.com/exPriceD/WBTech-L0_Microservice/internal/entities"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewItemsRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetItems(orderUID string) ([]entities.Items, error) {
	var items []entities.Items
	itemsQuery := `SELECT chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status FROM items WHERE order_uid = $1;`
	if err := r.db.Select(&items, itemsQuery, orderUID); err != nil {
		return nil, err
	}

	return items, nil
}
