package common

const (
	SUCCESS_CODE = 200   // 成功的状态码
	FALSE_CODE   = 10000 // 失败的状态码
)

var (
	ErrParam         = "传入参数有误"
	ErrAccountSystem = "账户系统响应失败"
	ErrFormat        = "格式不正确"
)

// // ControllerError ...
// type ControllerError struct {
// 	Code     int    `json:"code"`
// 	Langkey  string `json:"langkey"`
// 	Message  string `json:"msg"`
// 	Moreinfo string `json:"moreinfo"`
// }
