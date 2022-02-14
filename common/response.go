package common

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// ResSucModel ...
type ResponseInfos struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// todo 改为 ResponseItems
type Items struct {
	Items interface{} `json:"items"`
}

// todo 改为 ResponseItems
type PageItems struct {
	Items interface{} `json:"items"`
	Pages Pages       `json:"pages"`
}

// Pages
type Pages struct {
	Total    int64 `json:"total"`
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

// Response 接口统一返回
func Response(w http.ResponseWriter, data interface{}, err error) {
	v := ResponseInfos{}
	items := Items{Items: data}
	if err != nil {
		// 判断是否是token验证失败问题
		if err.Error() == AuthenticationFailed {
			httpx.WriteJson(w, http.StatusUnauthorized, v)
		} else {
			v.Code = FALSE_CODE
			v.Msg = err.Error()
			v.Data = items
			httpx.OkJson(w, v)
		}
	} else {
		v.Code = SUCCESS_CODE
		v.Msg = "ok"
		v.Data = items
		httpx.OkJson(w, v)
	}
}

// ResponsePage 接口统一返回
func ResponsePage(w http.ResponseWriter, data interface{}, err error) {
	v := ResponseInfos{}
	if err != nil {
		// 判断是否是token验证失败问题
		if err.Error() == AuthenticationFailed {
			httpx.WriteJson(w, http.StatusUnauthorized, v)
		} else {
			v.Code = FALSE_CODE
			v.Msg = err.Error()
			v.Data = data
			httpx.OkJson(w, v)
		}
	} else {
		v.Code = SUCCESS_CODE
		v.Msg = "ok"
		v.Data = data
		httpx.OkJson(w, v)
	}
}
