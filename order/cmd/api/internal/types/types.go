// Code generated by goctl. DO NOT EDIT.
package types

type Reply struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BuyBookReq struct {
	BookID int `json:"book_id"`
}

type SearchBookReq struct {
	BookID int `json:"book_id"`
}
