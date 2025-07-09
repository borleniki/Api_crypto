package service

import (
	"gormapp/dao"
	"gormapp/model"
)

// GetOrders fetches all orders
func GetOrders() ([]model.Order, error) {
	return dao.GetAllOrders()
}

// GetOrder fetches a single order by ID
func GetOrder(id int) (model.Order, error) {
	return dao.GetOrderByID(id)
}

// AddOrder creates a new order
func AddOrder(order model.Order) error {
	return dao.CreateOrder(order)
}

// EditOrder updates an existing order
func EditOrder(order model.Order) error {
	return dao.UpdateOrder(order)
}

// RemoveOrder deletes an order by ID
func RemoveOrder(id int) error {
	return dao.DeleteOrder(id)
}
