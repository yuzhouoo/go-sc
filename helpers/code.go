package helpers

type Code struct {
}

// http状态码
const (
	HTTP_CODE_SUCCESS      = 200
	HTTP_CODE_SERVER_ERROR = 500
)

// 服务状态码
const (
	MSG_CODE_SUCCESS = 100
	MSG_DESC_SUCCESS = "成功"

	MSG_CODE_FAIL = 200
	MSG_DESC_FAIL = "失败"

	MSG_CODE_ACCIDENT = 300
	MSG_DESC_ACCIDENT = "异常"

	MSG_CODE_NOT_LOGIN = 101
	MSG_DESC_NOT_LOGIN = "未登录"
)
