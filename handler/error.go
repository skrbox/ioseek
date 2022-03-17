package handler

type ApiErr struct {
	Message string
	Code    int
}

func (a ApiErr) Error() string {
	return a.Message
}

func NewApiError(code int, msg string) ApiErr {
	return ApiErr{
		Message: msg,
		Code:    code,
	}
}

var (
	HTTP404      = NewApiError(404, "资源不存在或已删除")
	HTTP405      = NewApiError(405, "不合法的请求方式")
	NotCompleted = NewApiError(1000, "功能未完成")
)
