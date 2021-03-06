package common

import (
	"encoding/json"
	"github.com/pkg/errors"

	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ServerError(w http.ResponseWriter, err error) {
	var serverError = new(ServerErrorx)
	if errors.As(err, &serverError) {

		resp := Response{
			Code: int(serverError.Code),
			Msg:  serverError.Message,
		}
		res, _ := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(res)

	}

}
