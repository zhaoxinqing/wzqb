package common

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS_CODE       = 200             //成功的状态码
	FAIL_CODE          = 30000           //失败的状态码
	MD5_PREFIX         = "yijingzhilian" //MD5加密前缀字符串，防止彩虹表暴力破解
	TOKEN_KEY          = "Authorization" //页面token键名
	CONTENT_TYPE_KET   = "Content-Type"
	CONTENT_TYPE_VALUE = "application/json"
	USER_ID_Key        = "UserID"   //页面用户ID键名
	USER_UUID_Key      = "X-UUID"   //页面UUID键名
	USER_TENANT_ID_Key = "TenantID" // 租户id
	USER_GROUP_ID_Key  = "GroupID"  // 组id
	SUPER_ADMIN_ID     = 1          // 超级管理员账号ID

)

var (
	SucAction = &ControllerError{90000, "ok.ActionSuccess", "操作成功", ""}
	SucUpdate = &ControllerError{90001, "ok.UpdateDone", "更新成功", ""}
	SucDelete = &ControllerError{90002, "ok.DeleteDone", "删除成功", ""}

	ErrLogin  = &ControllerError{9999, "err.ErrInstall", "安装失败", ""}
	ErrAction = &ControllerError{9998, "err.ErrAction", "操作失败", ""}

	Err404 = &ControllerError{404, "err.Err404", "页面没有找到", ""}
	Err403 = &ControllerError{403, "err.Err403", "您没有相关权限", ""}

	ErrInstall             = &ControllerError{10000, "err.ErrInstall", "安装失败", ""}
	ErrInputData           = &ControllerError{10001, "err.ErrInputData", "数据输入错误", ""}
	ErrDatabase            = &ControllerError{10002, "err.ErrDatabase", "数据库错误", ""}
	ErrDupUser             = &ControllerError{10003, "err.ErrDupUser", "用户信息已存在", ""}
	ErrNoUser              = &ControllerError{10004, "err.ErrNoUser", "用户信息不存在", ""}
	ErrPass                = &ControllerError{10005, "err.ErrPass", "用户信息不存在或密码不正确", ""}
	ErrNoUserPass          = &ControllerError{10006, "err.ErrNoUserPass", "用户信息不存在或密码不正确", ""}
	ErrNoUserChange        = &ControllerError{10007, "err.ErrNoUserChange", "用户信息不存在或数据未改变", ""}
	ErrInvalidUser         = &ControllerError{10008, "err.ErrInvalidUser", "用户信息不正确", ""}
	ErrOpenFile            = &ControllerError{10009, "err.ErrOpenFile", "打开文件错误", ""}
	ErrWriteFile           = &ControllerError{10010, "err.ErrWriteFile", "写文件出错", ""}
	ErrSystem              = &ControllerError{10011, "err.ErrSystem", "操作系统错误", ""}
	ErrExpired             = &ControllerError{10012, "err.ErrExpired", "登录已过期", ""}
	ErrPermission          = &ControllerError{10013, "err.ErrPermission", "没有权限", ""}
	ErrGenJwt              = &ControllerError{10014, "err.ErrGenJwt", "获取令牌失败", ""}
	ErrChkJwt              = &ControllerError{10012, "err.ErrChkJwt", "无效的令牌", ""}
	ErrIDData              = &ControllerError{10016, "err.ErrIdData", "此ID无数据记录", ""}
	ErrDifferentPasswords  = &ControllerError{10017, "err.ErrDifferentPasswords", "输入的密码不一致", ""}
	ErrSamePasswords       = &ControllerError{10018, "err.ErrSamePasswords", "请输入与旧密码不一样的密码", ""}
	ErrUserOrPassEmpty     = &ControllerError{10019, "err.ErrUserOrPassEmpty", "用户名或密码不能为空", ""}
	ErrUserHasBeenDisabled = &ControllerError{10020, "err.ErrUserHasBeenDisabled", "用户已被禁用", ""}
	ErrOldPass             = &ControllerError{10021, "err.ErrOldPass", "原密码输入不正确", ""}
	ErrJwtExp              = &ControllerError{10022, "err.ErrJwtExp", "令牌过期", ""}

	ErrAddFail       = &ControllerError{11000, "err.ErrAddFail", "创建失败", ""}
	ErrEditFail      = &ControllerError{11001, "err.ErrEditFail", "更新失败", ""}
	ErrDelFail       = &ControllerError{11002, "err.ErrDelFail", "删除失败", ""}
	ErrInvalidParams = &ControllerError{11003, "err.ErrInvalidParams", "验证参数失败", ""}

	ErrRoleAssignFail      = &ControllerError{12000, "err.ErrRoleAssignFail", "权限分配失败", ""}
	ErrMenuData            = &ControllerError{12001, "err.ErrMenuData", "请传递菜单ids", ""}
	ErrRoleExists          = &ControllerError{12002, "err.ErrRoleExists", "角色已存在", ""}
	ErrRoleNameOrKeyExists = &ControllerError{12003, "err.ErrRoleNameOrKeyExists", "已存在同名或同关键字角色", ""}

	ErrCaptchaEmpty      = &ControllerError{13001, "err.ErrCaptchaEmpty", "验证码不能为空", ""}
	ErrCaptcha           = &ControllerError{13002, "err.ErrCaptcha", "验证码错误", ""}
	ErrDeptDel           = &ControllerError{13003, "err.ErrDeptDel", "部门无法删除", "部门内仍有成员,请先行转移到其它部门"}
	ErrDeptHasMember     = &ControllerError{13004, "err.ErrDeptHasMember", "部门不可删除", "部门内仍有成员"}
	ErrDupRecord         = &ControllerError{13005, "err.ErrDupRecord", "记录已存在", ""}
	ErrWrongRefreshToken = &ControllerError{13006, "err.ErrWrongRefreshToken", "无效的refresh令牌", ""}
	ErrBindDingtalk      = &ControllerError{13007, "err.ErrBindDingtalk", "", ""}
	ErrUnBindDingtalk    = &ControllerError{13008, "err.ErrUnBindDingtalk", "", ""}
	ErrGoogleBindCode    = &ControllerError{13009, "err.ErrGoogleBindCode", "", ""}
	ErrSendMail          = &ControllerError{13010, "err.ErrSendMail", "发送邮件失败", ""}
	ErrValidation        = &ControllerError{13011, "err.ErrValidate", "请求参数验证失败", ""}
	ErrNoRecord          = &ControllerError{13012, "err.ErrNoRecord", "记录不存在", ""}
	ErrHasSubRecord      = &ControllerError{13013, "err.ErrHasSubRecord", "子节点不为空", ""}
	ErrUploadAvatar      = &ControllerError{13014, "err.ErrUploadAvatar", "上传头像失败", ""}
	ErrSmsSendCode       = &ControllerError{13015, "err.ErrSendCode", "验证码发送失败", ""}
)

// ControllerError ...
type ControllerError struct {
	Code     int    `json:"code"`
	Langkey  string `json:"langkey"`
	Message  string `json:"msg"`
	Moreinfo string `json:"moreinfo"`
}

// ResponseModel ...
type ResponseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ResponseModelBase ...
type ResponseModelBase struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ResSuccess 响应成功
func ResSuccess(c *gin.Context, v interface{}) {
	ret := ResponseModel{Code: SUCCESS_CODE, Message: "ok", Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

// ResSuccessMsg 响应成功
func ResSuccessMsg(c *gin.Context) {
	ret := ResponseModelBase{Code: SUCCESS_CODE, Message: "ok"}
	ResJSON(c, http.StatusOK, &ret)
}

// ResFail 响应失败
func ResFail(c *gin.Context, msg string) {
	ret := ResponseModelBase{Code: FAIL_CODE, Message: msg}
	ResJSON(c, http.StatusOK, &ret)
}

// ResFailCode 响应失败
func ResFailCode(c *gin.Context, msg string, code int) {
	ret := ResponseModelBase{Code: code, Message: msg}
	ResJSON(c, http.StatusOK, &ret)
}

// ResJSON 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	c.JSON(status, v)
	c.Abort()
}

// ResErrSrv 响应错误-服务端故障
func ResErrSrv(c *gin.Context, err error) {
	ret := ResponseModelBase{Code: FAIL_CODE, Message: err.Error()}
	ResJSON(c, http.StatusOK, &ret)
}

// ResErrCli 响应错误-用户端故障
func ResErrCli(c *gin.Context, err error) {
	ret := ResponseModelBase{Code: FAIL_CODE, Message: err.Error()}
	ResJSON(c, http.StatusOK, &ret)
}

// ResponsePageData ...
type ResponsePageData struct {
	Total int64       `json:"total"`
	Items interface{} `json:"items"`
}

// ResponsePage ...
type ResponsePage struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    ResponsePageData `json:"data"`
}

// ResSuccessPage 响应成功-分页数据
func ResSuccessPage(c *gin.Context, total int64, list interface{}) {
	ret := ResponsePage{Code: SUCCESS_CODE, Message: "ok", Data: ResponsePageData{Total: total, Items: list}}
	ResJSON(c, http.StatusOK, &ret)
}

// // GetPageIndex 获取页码
// func GetPageIndex(c *gin.Context) int64 {
// 	return GetQueryToInt64(c, "page", 1)
// }

// // GetPageSize 获取每页记录数
// func GetPageSize(c *gin.Context) int64 {
// 	limit := GetQueryToInt64(c, "pageSize", 10)
// 	if limit > 500 {
// 		limit = 20
// 	}
// 	return limit
// }

// GetPageSort 获取排序信息
func GetPageSort(c *gin.Context) string {
	return GetQueryToStr(c, "sort")
}

// GetPageKey 获取搜索关键词信息
func GetPageKey(c *gin.Context) string {
	return GetQueryToStr(c, "key")
}

// GetQueryToStrE ...
func GetQueryToStrE(c *gin.Context, key string) (string, error) {
	str, ok := c.GetQuery(key)
	if !ok {
		return "", errors.New("没有这个值传入")
	}
	return str, nil
}

// GetQueryToStr ...
func GetQueryToStr(c *gin.Context, key string, defaultValues ...string) string {
	var defaultValue string
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	str, err := GetQueryToStrE(c, key)
	if str == "" || err != nil {
		return defaultValue
	}
	return str
}
