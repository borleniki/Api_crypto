package dao

import (
	"gormapp/config"
	"gormapp/model"
)

// GetAllOrders retrieves all orders from the database
func GetAllOrders() ([]model.Order, error) {
	var orders []model.Order
	err := config.DB.Find(&orders).Error
	return orders, err
}

// GetOrderByID retrieves a single order by its ID
func GetOrderByID(id int) (model.Order, error) {
	var order model.Order
	err := config.DB.First(&order, id).Error
	return order, err
}

// CreateOrder inserts a new order into the database
func CreateOrder(order model.Order) error {
	return config.DB.Create(&order).Error
}

// UpdateOrder updates an existing order record
func UpdateOrder(order model.Order) error {
	return config.DB.Save(&order).Error
}

// DeleteOrder deletes an order by ID
func DeleteOrder(id int) error {
	return config.DB.Delete(&model.Order{}, id).Error
}
