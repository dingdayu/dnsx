package v1

import "fmt"

var (
	SuccessResponse = Response{Code: 1000, Message: "success"}

	ErrInternalServer = Response{Code: 1001, Message: "系统错误"}
	ErrMissParams     = Response{Code: 1002, Message: "缺少参数"}
	ErrFailParams     = Response{Code: 1003, Message: "参数格式错误"}
	ErrNotExist       = Response{Code: 1004, Message: "数据不存在"}
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (e Response) Error() string {
	return fmt.Sprintf(e.Message)
}

func NewErrResponse(code int, message string) Response {
	s := Response{}
	s.Code = code
	s.Message = message
	return s
}

// NewResponseWithData 返回携带data的自定义响应
func NewResponseWithData(response Response, data interface{}) Response {
	s := response
	s.Data = data
	return s
}

func NewSucResponse(data interface{}) Response {
	s := SuccessResponse
	s.Data = data
	return s
}
func NewErrParamsResponse(message string) Response {
	s := ErrMissParams
	s.Message = message
	return s
}
