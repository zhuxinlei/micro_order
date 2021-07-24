package handler

import (
	"github.com/zhuxinlei/micro_order/order/cmd/api/internal/logic"
	"github.com/zhuxinlei/micro_order/order/cmd/api/internal/svc"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func search_userHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewSearch_userLogic(r.Context(), ctx)
		resp, err := l.Search_user()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
