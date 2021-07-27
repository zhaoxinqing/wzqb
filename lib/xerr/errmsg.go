package xerr

var message map[int]string

func init() {
	message = make(map[int]string)
	message[OK] = "SUCCESS"
	message[ErrCodeBadReq] = "服务器繁忙,请稍后再试"
	message[ErrCodeParams] = "参数错误"
	message[ErrCodeUserNotFound] = "用户不存在"
}

func MapErrMsg(errcode int) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器繁忙,请稍后再试"
	}
}
