package response

type ResCode int64

// 定义状态码
const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeInvalidPassword
	CodeUserExist
	CodeUserNotExist
	CodeServerBusy

	CodeInvalidToken
	CodeNeedLogin
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeInvalidPassword: "用户名或密码错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeServerBusy:      "服务繁忙",
	CodeInvalidToken:    "无效token",
	CodeNeedLogin:       "需要登录",
}

func (c ResCode) Msg() string {
	msg := codeMsgMap[c]
	return msg
}
