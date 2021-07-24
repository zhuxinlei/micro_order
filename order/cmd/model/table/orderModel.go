package table

import (
	"github.com/jinzhu/gorm"
	"github.com/zhuxinlei/micro_order/order/cmd/model/input"
	"github.com/zhuxinlei/micro_order/order/cmd/model/output"
)

type OrderModel struct {
	db *gorm.DB
}

func NewOrderModel(db *gorm.DB) OrderModel {
	return OrderModel{
		db: db,
	}
}

func (u OrderModel) Insert(input input.Order) (err error) {
	err = u.db.Debug().Create(&input).Error
	return

}
func (u OrderModel) GetOrder(user input.Order) (output output.Order, err error) {
	var data input.Order
	err = u.db.Debug().Where(user).Take(&data).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}

	return data.ToOutput(), err
}

func (u OrderModel) GetOrders(where input.Order) ([]output.Order, error) {
	var data []input.Order
	err := u.db.Debug().Where(where).Find(&data).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}

	var res []output.Order
	for _, k := range data {
		res = append(res,k.ToOutput())
	}

	return res, err
}
