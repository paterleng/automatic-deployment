package controller

type ResCode int64

const (
	CodeSuccess      ResCode = 200 + iota
	CodeInvalidParam ResCode = 1001 + iota
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeInvalidToken
	CodeNeedLogin
	CodeCreateError
	CodeNotNil
	CodeParamError
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeInvalidToken:    "无效token",
	CodeNeedLogin:       "需要登录",
	CodeCreateError:     "创建失败",
	CodeNotNil:          "选择不能为空",
	CodeParamError:      "参数错误",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
