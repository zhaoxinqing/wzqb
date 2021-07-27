package xerr

//成功返回
const OK = 200

const (
	ErrCodeApiServer     = 10001 //api服务异常
	ErrCodeRpcServer     = 10002 //rpc服务异常
	ErrCodeDataBase      = 10003 //数据库异常
	ErrCodeBadReq        = 10004 //错误的请求
	ErrCodeToken         = 10005 //token错误
	ErrCodeTokenInvalid  = 10006 //token失效
	ErrCodeParams        = 10007 //参数错误
	ErrCodeParamsInvalid = 10008 //参数非法
	ErrCodeRedis         = 10009 //连接异常
	ErrCodeDataNotFound  = 10010 //数据不存在
	ErrCodeDataRepeat    = 10011 //数据重复
)
const (
	ErrMsgRpcServer    = "grpc服务异常"
	ErrMsgApiServer    = "api服务异常"
	ErrMsgDataBase     = "数据库异常"
	ErrMsgToken        = "token错误"
	ErrMsgTokeInvalid  = "token失效"
	ErrMsgDataRepeat   = "数据已存在"
	ErrMsgDataNotFound = "数据不存在"
	ErrMsgPwdInvalid   = "密码无效"

	/**************用户*****************/
	ErrMsgUserNotFount = "用户不存在"
	/**************合同****************/
	ErrMsgContractDel = "合同已删除"
)

//用户模块 2开头
const (
	ErrCodeUserNotFound = 20001
)

// 数据台账 8开头
const (
	ErrCodeContractDel = 80001
)
