package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zhuxinlei/micro_book/book/cmd/model/common"
	"github.com/zhuxinlei/micro_book/book/cmd/rpc/bookclient"
	common_local "github.com/zhuxinlei/micro_order/order/cmd/pkg/common"
	"strings"

	"github.com/zhuxinlei/micro_order/order/cmd/api/internal/svc"
	"github.com/zhuxinlei/micro_order/order/cmd/api/internal/types"

	"github.com/zhuxinlei/micro_order/order/cmd/model/input"

	"github.com/tal-tech/go-zero/core/logx"
)

type Search_bookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearch_bookLogic(ctx context.Context, svcCtx *svc.ServiceContext) Search_bookLogic {
	return Search_bookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Search_bookLogic) Search_book() (*types.Reply, error) {
	// todo: add your logic here and delete this line
	userId := common_local.ToInt(l.ctx.Value("userId"))
	where := input.Order{
		UserId: userId,
	}
	order, err := l.svcCtx.Order.GetOrders(where)
	if err != nil {
		return nil, errors.Wrapf(common_local.GetOrderError, common.ErrCodeMap[common_local.GetOrderErrorCode])
	}
	var temp []int
	for _, k := range order {
		temp = append(temp, k.BookId)
	}
	ids := strings.Replace(strings.Trim(fmt.Sprint(temp), "[]"), " ", ",", -1)

	data, err := l.svcCtx.BookRpc.GetBooks(l.ctx, &bookclient.IdsReq{Ids: ids})
	if err != nil {
		return nil, errors.Wrapf(common_local.GetBookError, common.ErrCodeMap[common_local.GetBookErrorCode])
	}

	for i,k := range order{
		for _,v := range data.Data{
			if k.BookId == int(v.Id){
				order[i].BookName = v.BookName
				order[i].UserName = "因为懒，这里就不走rpc 服务接口获取数据了"
				continue
			}
		}
	}
	return &types.Reply{
		Code:    200,
		Message: "",
		Data:    order,
	}, nil
}
