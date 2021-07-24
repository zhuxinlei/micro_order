package model

import (
	"github.com/zhuxinlei/micro_order/order/cmd/model/input"
	"github.com/zhuxinlei/micro_order/order/cmd/model/output"
)

type Order interface {
	Insert(order input.Order) error
	GetOrder(order input.Order) (output.Order, error)
	GetOrders(order input.Order) ([]output.Order, error)
}
