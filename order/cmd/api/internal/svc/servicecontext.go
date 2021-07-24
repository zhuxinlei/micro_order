package svc

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/zhuxinlei/micro_book/book/cmd/rpc/bookclient"
	"github.com/zhuxinlei/micro_order/order/cmd/api/internal/config"
	"github.com/zhuxinlei/micro_order/order/cmd/model/table"
)

type ServiceContext struct {
	Config  config.Config
	BookRpc bookclient.Book
	Order   table.OrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	DB, err := gorm.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/micro_order?charset=utf8&parseTime=true&loc=UTC")
	if err != nil {
		logx.Errorf("连接数据库出错%s", err.Error())
		fmt.Println("连接数据库出错")
	}
	order := table.NewOrderModel(DB)
	return &ServiceContext{
		Config:  c,
		BookRpc: bookclient.NewBook(zrpc.MustNewClient(c.BookRpc)),
		Order:   order,
	}
}
