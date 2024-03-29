package order

import "Back-End-Ecommers/entities"

type Order interface {
	Get(userId int) ([]entities.Order, error)
	GetById(orderId int) (entities.Order, error)
	Create(userId, paymentId int) (entities.Order, error)
	Update(orderId int, newOrder entities.Order) (entities.Order, error)
	Delete(orderId int) error
}
