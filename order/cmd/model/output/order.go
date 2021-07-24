package output

import (
	"time"
)

type Order struct {
	Id       uint      `json:"id"`
	CreateAt int64     `json:"create_at"`
	BookId   int       `json:"book_id"`
	UserId   int       `json:"user_id"`
	BuyDate  time.Time `json:"buy_date"`
	BookName string    `json:"book_name"`
	UserName string    `json:"user_name"`
}
