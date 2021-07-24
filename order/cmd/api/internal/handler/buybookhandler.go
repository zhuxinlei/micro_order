package handler

import (
	"github.com/zhuxinlei/micro_order/order/cmd/api/internal/logic"
	"github.com/zhuxinlei/micro_order/order/cmd/api/internal/svc"
	"github.com/zhuxinlei/micro_order/order/cmd/api/internal/types"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"github.com/zhuxinlei/micro_order/order/cmd/pkg/common"
)

func buy_bookHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BuyBookReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewBuy_bookLogic(r.Context(), ctx)
		resp, err := l.Buy_book(req)
		if err != nil {
			//httpx.Error(w, err)
			common.ServerError(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
