package handler

import (
	"errors"
	"net/http"

	"go-template/common"
	"go-template/service/account/api/internal/logic"
	"go-template/service/account/api/internal/svc"
	"go-template/service/account/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// LoginHandler 登陆
func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterParam
		//
		if err := httpx.Parse(r, &req); err != nil {
			common.Response(w, nil, errors.New(common.IncorrectParameteFormat))
			return
		}
		// 字段校验
		if req.Name == "" || req.Password == "" {
			common.Response(w, nil, errors.New(common.MissingRequiredParameter))
			return
		}
		//
		l := logic.NewLoginLogic(r.Context(), svcCtx)
		value, err := l.Login(req)
		//
		common.Response(w, value, err)
	}
}
