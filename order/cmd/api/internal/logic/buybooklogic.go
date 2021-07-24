package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zhuxinlei/micro_book/book/cmd/model/common"
	"github.com/zhuxinlei/micro_book/book/cmd/rpc/bookclient"
	"github.com/zhuxinlei/micro_order/order/cmd/api/internal/types"
	"github.com/zhuxinlei/micro_order/order/cmd/model/input"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/zhuxinlei/micro_order/order/cmd/api/internal/svc"
	common_local "github.com/zhuxinlei/micro_order/order/cmd/pkg/common"
)

type Buy_bookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBuy_bookLogic(ctx context.Context, svcCtx *svc.ServiceContext) Buy_bookLogic {
	return Buy_bookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Buy_bookLogic) Buy_book(req types.BuyBookReq) (*types.Reply, error) {
	// todo: add your logic here and delete this line

	//验证book_id是否真的存在，这里要调用book rpc
	bookInfo, err := l.svcCtx.BookRpc.GetBook(l.ctx,
		&bookclient.IdReq{
			Id: int64(req.BookID),
		},
	)

	if err != nil {
		fmt.Println("book rpc server获取book info error")
		return nil, errors.Wrapf(common.GetBookError, common.ErrCodeMap[common.GetBookErrorCode])
	}
	if bookInfo.Id == 0 {
		fmt.Println("书籍不存在")
		return nil, errors.Wrapf(common_local.GetBookError, common.ErrCodeMap[common_local.GetBookErrorCode])
	}

	userId := common_local.ToInt(l.ctx.Value("userId"))
	order := input.Order{
		BookId:  int(bookInfo.Id),
		UserId:  userId,
		BuyDate: time.Now(),
	}
	fmt.Println(2323)
	err = l.svcCtx.Order.Insert(order)
	if err != nil {
		fmt.Println("插入order error")
		return nil, errors.Wrapf(common_local.InsertOrderError, common.ErrCodeMap[common_local.InsertOrderErrorCode])
	}
	return &types.Reply{
		Code:    200,
		Message: "buy success",
	}, nil
}
