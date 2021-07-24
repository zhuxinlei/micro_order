package input

import (
	"github.com/jinzhu/gorm"
	"github.com/zhuxinlei/micro_order/order/cmd/model/output"
	"time"
)

type Order struct {
	gorm.Model
	BookId  int       `json:"book_id"`
	UserId  int       `json:"user_id"`
	BuyDate time.Time `json:"buy_date"`
}

func (o *Order) ToOutput() output.Order {

	return output.Order{
		Id:       o.ID,
		CreateAt: o.CreatedAt.Unix(),
		BookId:   o.BookId,
		UserId:   o.UserId,
		BuyDate:  o.BuyDate,
	}
}
