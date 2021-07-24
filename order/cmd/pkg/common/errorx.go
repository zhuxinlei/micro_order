package common

// 全局的 错误号 类型，用于API调用之间传递
type ServerErrorxCode int

// 全局的 错误号 的具体定义
const (
	InsertOrderErrorCode = 3000 //插入Order出错
	GetOrderErrorCode    = 3001 //获取用户信息出错
	GetBookErrorCode     = 3002 //按正规的业务逻辑来看，不应该将book相关的错误码放在order中，book相关的错误码应该
	放在book中返回
)

// 内部的错误map，用来对应 错误号和错误信息
var ErrCodeMap = map[ServerErrorxCode]string{

	InsertOrderErrorCode: "insert into Orders error",
	GetOrderErrorCode:    "get Order info error",
	GetBookErrorCode:     "get book info error",
}

// Sentinel Error： 即全局定义的Static错误变量
// 注意，这里的全局error是没有保存堆栈信息的，所以需要在初始调用处使用 errors.Wrap
var (
	InsertOrderError = NewServerErrorx(InsertOrderErrorCode)
	GetOrderError    = NewServerErrorx(GetOrderErrorCode)
	GetBookError     = NewServerErrorx(GetBookErrorCode)
)

func NewServerErrorx(code ServerErrorxCode) *ServerErrorx {
	return &ServerErrorx{
		Code:    code,
		Message: ErrCodeMap[code],
	}
}

// error的具体实现
type ServerErrorx struct {
	// 对外使用 - 错误码
	Code ServerErrorxCode
	// 对外使用 - 错误信息
	Message string
}

func (e *ServerErrorx) Error() string {
	return e.Message
}
