package database

import (
	"database/sql"

	"github.com/marcelobritu/go-expert-desafio-clean-architecture/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (o *OrderRepository) FindAll() ([]entity.Order, error) {
	rows, err := o.Db.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.Order
	for rows.Next() {
		var order entity.Order
		if err := rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
